package main

import (
	"os"
	"strings"
	"testing"

	"github.com/tmcphillips/blazegraph-util/assert"
)

func ExampleBlazegraphCmd_export_empty_store() {
	Main.OutWriter = os.Stdout
	Main.ErrWriter = os.Stdout
	run("blazegraph drop")
	run("blazegraph export")
	// Output:
	//
}

func ExampleBlazegraphCmd_drop_import_ttl_then_export_ttl() {
	Main.OutWriter = os.Stdout
	Main.ErrWriter = os.Stdout
	run("blazegraph drop")
	run("blazegraph import --file testdata/in.nt --format ttl")
	run("blazegraph export --format ttl")
	// Output:
	// @prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
	// @prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .
	// @prefix sesame: <http://www.openrdf.org/schema/sesame#> .
	// @prefix owl: <http://www.w3.org/2002/07/owl#> .
	// @prefix xsd: <http://www.w3.org/2001/XMLSchema#> .
	// @prefix fn: <http://www.w3.org/2005/xpath-functions#> .
	// @prefix foaf: <http://xmlns.com/foaf/0.1/> .
	// @prefix dc: <http://purl.org/dc/elements/1.1/> .
	// @prefix hint: <http://www.bigdata.com/queryHints#> .
	// @prefix bd: <http://www.bigdata.com/rdf#> .
	// @prefix bds: <http://www.bigdata.com/rdf/search#> .

	// <http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .

	// <http://tmcphill.net/tags#tag> a rdf:Property ;
	// 	rdfs:subPropertyOf <http://tmcphill.net/tags#tag> .

	// <http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
}

func ExampleBlazegraphCmd_drop_import_ttl_then_export_ntriples() {
	Main.OutWriter = os.Stdout
	Main.ErrWriter = os.Stdout
	run("blazegraph drop")
	run("blazegraph import --file testdata/in.nt --format ttl")
	run("blazegraph export --format nt")
	// Output:
	// <http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
	// <http://tmcphill.net/tags#tag> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/1999/02/22-rdf-syntax-ns#Property> .
	// <http://tmcphill.net/tags#tag> <http://www.w3.org/2000/01/rdf-schema#subPropertyOf> <http://tmcphill.net/tags#tag> .
	// <http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
}

func ExampleBlazegraphCmd_drop_import_ttl_then_export_xml() {
	Main.OutWriter = os.Stdout
	Main.ErrWriter = os.Stdout
	run("blazegraph drop")
	run("blazegraph import --file testdata/in.nt --format ttl")
	run("blazegraph export --format xml")
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <rdf:RDF
	// 	xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
	// 	xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#"
	// 	xmlns:sesame="http://www.openrdf.org/schema/sesame#"
	// 	xmlns:owl="http://www.w3.org/2002/07/owl#"
	// 	xmlns:xsd="http://www.w3.org/2001/XMLSchema#"
	// 	xmlns:fn="http://www.w3.org/2005/xpath-functions#"
	// 	xmlns:foaf="http://xmlns.com/foaf/0.1/"
	// 	xmlns:dc="http://purl.org/dc/elements/1.1/"
	// 	xmlns:hint="http://www.bigdata.com/queryHints#"
	// 	xmlns:bd="http://www.bigdata.com/rdf#"
	// 	xmlns:bds="http://www.bigdata.com/rdf/search#">
	//
	// <rdf:Description rdf:about="http://tmcphill.net/data#y">
	// 	<tag xmlns="http://tmcphill.net/tags#">eight</tag>
	// </rdf:Description>
	//
	// <rdf:Description rdf:about="http://tmcphill.net/tags#tag">
	// 	<rdf:type rdf:resource="http://www.w3.org/1999/02/22-rdf-syntax-ns#Property"/>
	// 	<rdfs:subPropertyOf rdf:resource="http://tmcphill.net/tags#tag"/>
	// </rdf:Description>
	//
	// <rdf:Description rdf:about="http://tmcphill.net/data#x">
	// 	<tag xmlns="http://tmcphill.net/tags#">seven</tag>
	// </rdf:Description>
	//
	// </rdf:RDF>
}

func TestBlazegraphCmd_drop_import_ttl_then_export_jsonld(t *testing.T) {
	var resultsBuffer strings.Builder
	Main.OutWriter = &resultsBuffer
	Main.ErrWriter = &resultsBuffer

	run("blazegraph drop")
	run("blazegraph import --file testdata/in.nt --format ttl")
	run("blazegraph export --format jsonld")

	assert.JSONEquals(t, resultsBuffer.String(), `
	[
		{
		  "@id": "http://tmcphill.net/data#x",
		  "http://tmcphill.net/tags#tag": [
			{
			  "@value": "seven"
			}
		  ]
		},
		{
		  "@id": "http://tmcphill.net/data#y",
		  "http://tmcphill.net/tags#tag": [
			{
			  "@value": "eight"
			}
		  ]
		},
		{
		  "@id": "http://tmcphill.net/tags#tag",
		  "@type": [
			"http://www.w3.org/1999/02/22-rdf-syntax-ns#Property"
		  ],
		  "http://www.w3.org/2000/01/rdf-schema#subPropertyOf": [
			{
			  "@id": "http://tmcphill.net/tags#tag"
			}
		  ]
		}
	  ]
	`)
}

func ExampleBlazegraphCmd_drop_import_jsonld_then_export_ntriples() {
	Main.OutWriter = os.Stdout
	Main.ErrWriter = os.Stdout
	run("blazegraph drop")
	run("blazegraph import --file testdata/address-book.jsonld --format jsonld")
	run("blazegraph export --format nt")
	// Output:
	// <http://learningsparql.com/ns/addressbook#email> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/1999/02/22-rdf-syntax-ns#Property> .
	// <http://learningsparql.com/ns/addressbook#email> <http://www.w3.org/2000/01/rdf-schema#subPropertyOf> <http://learningsparql.com/ns/addressbook#email> .
	// <http://learningsparql.com/ns/addressbook#firstname> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/1999/02/22-rdf-syntax-ns#Property> .
	// <http://learningsparql.com/ns/addressbook#firstname> <http://www.w3.org/2000/01/rdf-schema#subPropertyOf> <http://learningsparql.com/ns/addressbook#firstname> .
	// <http://learningsparql.com/ns/addressbook#homeTel> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/1999/02/22-rdf-syntax-ns#Property> .
	// <http://learningsparql.com/ns/addressbook#homeTel> <http://www.w3.org/2000/01/rdf-schema#subPropertyOf> <http://learningsparql.com/ns/addressbook#homeTel> .
	// <http://learningsparql.com/ns/addressbook#lastname> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/1999/02/22-rdf-syntax-ns#Property> .
	// <http://learningsparql.com/ns/addressbook#lastname> <http://www.w3.org/2000/01/rdf-schema#subPropertyOf> <http://learningsparql.com/ns/addressbook#lastname> .
	// <http://learningsparql.com/ns/addressbook#mobileTel> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/1999/02/22-rdf-syntax-ns#Property> .
	// <http://learningsparql.com/ns/addressbook#mobileTel> <http://www.w3.org/2000/01/rdf-schema#subPropertyOf> <http://learningsparql.com/ns/addressbook#mobileTel> .
	// <http://learningsparql.com/ns/addressbook#nickname> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/1999/02/22-rdf-syntax-ns#Property> .
	// <http://learningsparql.com/ns/addressbook#nickname> <http://www.w3.org/2000/01/rdf-schema#subPropertyOf> <http://learningsparql.com/ns/addressbook#nickname> .
	// <http://learningsparql.com/ns/data#i0432> <http://learningsparql.com/ns/addressbook#email> "richard49@hotmail.com"^^<http://www.w3.org/2001/XMLSchema#string> .
	// <http://learningsparql.com/ns/data#i0432> <http://learningsparql.com/ns/addressbook#firstname> "Richard"^^<http://www.w3.org/2001/XMLSchema#string> .
	// <http://learningsparql.com/ns/data#i0432> <http://learningsparql.com/ns/addressbook#homeTel> "(229) 276-5135"^^<http://www.w3.org/2001/XMLSchema#string> .
	// <http://learningsparql.com/ns/data#i0432> <http://learningsparql.com/ns/addressbook#lastname> "Mutt"^^<http://www.w3.org/2001/XMLSchema#string> .
	// <http://learningsparql.com/ns/data#i0432> <http://learningsparql.com/ns/addressbook#nickname> "Dick"^^<http://www.w3.org/2001/XMLSchema#string> .
	// <http://learningsparql.com/ns/data#i8301> <http://learningsparql.com/ns/addressbook#email> "c.ellis@usairwaysgroup.com"^^<http://www.w3.org/2001/XMLSchema#string> .
	// <http://learningsparql.com/ns/data#i8301> <http://learningsparql.com/ns/addressbook#email> "craigellis@yahoo.com"^^<http://www.w3.org/2001/XMLSchema#string> .
	// <http://learningsparql.com/ns/data#i8301> <http://learningsparql.com/ns/addressbook#firstname> "Craig"^^<http://www.w3.org/2001/XMLSchema#string> .
	// <http://learningsparql.com/ns/data#i8301> <http://learningsparql.com/ns/addressbook#homeTel> "(194) 966-1505"^^<http://www.w3.org/2001/XMLSchema#string> .
	// <http://learningsparql.com/ns/data#i8301> <http://learningsparql.com/ns/addressbook#lastname> "Ellis"^^<http://www.w3.org/2001/XMLSchema#string> .
	// <http://learningsparql.com/ns/data#i9771> <http://learningsparql.com/ns/addressbook#email> "cindym@gmail.com"^^<http://www.w3.org/2001/XMLSchema#string> .
	// <http://learningsparql.com/ns/data#i9771> <http://learningsparql.com/ns/addressbook#firstname> "Cindy"^^<http://www.w3.org/2001/XMLSchema#string> .
	// <http://learningsparql.com/ns/data#i9771> <http://learningsparql.com/ns/addressbook#homeTel> "(245) 646-5488"^^<http://www.w3.org/2001/XMLSchema#string> .
	// <http://learningsparql.com/ns/data#i9771> <http://learningsparql.com/ns/addressbook#lastname> "Marshall"^^<http://www.w3.org/2001/XMLSchema#string> .
	// <http://learningsparql.com/ns/data#i9771> <http://learningsparql.com/ns/addressbook#mobileTel> "(245) 732-8991"^^<http://www.w3.org/2001/XMLSchema#string> .
}

func TestBlazegraphCmd_drop_import_jsonld_then_export_jsonld(t *testing.T) {
	var resultsBuffer strings.Builder
	Main.OutWriter = &resultsBuffer
	Main.ErrWriter = &resultsBuffer

	run("blazegraph drop")
	run("blazegraph import --file testdata/address-book.jsonld --format jsonld")
	run("blazegraph export --format jsonld")

	assert.JSONEquals(t, resultsBuffer.String(), `
	[ {
	"@id" : "http://learningsparql.com/ns/addressbook#email",
	"@type" : [ "http://www.w3.org/1999/02/22-rdf-syntax-ns#Property" ],
	"http://www.w3.org/2000/01/rdf-schema#subPropertyOf" : [ {
		"@id" : "http://learningsparql.com/ns/addressbook#email"
	} ]
	}, {
	"@id" : "http://learningsparql.com/ns/addressbook#firstname",
	"@type" : [ "http://www.w3.org/1999/02/22-rdf-syntax-ns#Property" ],
	"http://www.w3.org/2000/01/rdf-schema#subPropertyOf" : [ {
		"@id" : "http://learningsparql.com/ns/addressbook#firstname"
	} ]
	}, {
	"@id" : "http://learningsparql.com/ns/addressbook#homeTel",
	"@type" : [ "http://www.w3.org/1999/02/22-rdf-syntax-ns#Property" ],
	"http://www.w3.org/2000/01/rdf-schema#subPropertyOf" : [ {
		"@id" : "http://learningsparql.com/ns/addressbook#homeTel"
	} ]
	}, {
	"@id" : "http://learningsparql.com/ns/addressbook#lastname",
	"@type" : [ "http://www.w3.org/1999/02/22-rdf-syntax-ns#Property" ],
	"http://www.w3.org/2000/01/rdf-schema#subPropertyOf" : [ {
		"@id" : "http://learningsparql.com/ns/addressbook#lastname"
	} ]
	}, {
	"@id" : "http://learningsparql.com/ns/addressbook#mobileTel",
	"@type" : [ "http://www.w3.org/1999/02/22-rdf-syntax-ns#Property" ],
	"http://www.w3.org/2000/01/rdf-schema#subPropertyOf" : [ {
		"@id" : "http://learningsparql.com/ns/addressbook#mobileTel"
	} ]
	}, {
	"@id" : "http://learningsparql.com/ns/addressbook#nickname",
	"@type" : [ "http://www.w3.org/1999/02/22-rdf-syntax-ns#Property" ],
	"http://www.w3.org/2000/01/rdf-schema#subPropertyOf" : [ {
		"@id" : "http://learningsparql.com/ns/addressbook#nickname"
	} ]
	}, {
	"@id" : "http://learningsparql.com/ns/data#i0432",
	"http://learningsparql.com/ns/addressbook#email" : [ {
		"@value" : "richard49@hotmail.com"
	} ],
	"http://learningsparql.com/ns/addressbook#firstname" : [ {
		"@value" : "Richard"
	} ],
	"http://learningsparql.com/ns/addressbook#homeTel" : [ {
		"@value" : "(229) 276-5135"
	} ],
	"http://learningsparql.com/ns/addressbook#lastname" : [ {
		"@value" : "Mutt"
	} ],
	"http://learningsparql.com/ns/addressbook#nickname" : [ {
		"@value" : "Dick"
	} ]
	}, {
	"@id" : "http://learningsparql.com/ns/data#i8301",
	"http://learningsparql.com/ns/addressbook#email" : [ {
		"@value" : "c.ellis@usairwaysgroup.com"
	}, {
		"@value" : "craigellis@yahoo.com"
	} ],
	"http://learningsparql.com/ns/addressbook#firstname" : [ {
		"@value" : "Craig"
	} ],
	"http://learningsparql.com/ns/addressbook#homeTel" : [ {
		"@value" : "(194) 966-1505"
	} ],
	"http://learningsparql.com/ns/addressbook#lastname" : [ {
		"@value" : "Ellis"
	} ]
	}, {
	"@id" : "http://learningsparql.com/ns/data#i9771",
	"http://learningsparql.com/ns/addressbook#email" : [ {
		"@value" : "cindym@gmail.com"
	} ],
	"http://learningsparql.com/ns/addressbook#firstname" : [ {
		"@value" : "Cindy"
	} ],
	"http://learningsparql.com/ns/addressbook#homeTel" : [ {
		"@value" : "(245) 646-5488"
	} ],
	"http://learningsparql.com/ns/addressbook#lastname" : [ {
		"@value" : "Marshall"
	} ],
	"http://learningsparql.com/ns/addressbook#mobileTel" : [ {
		"@value" : "(245) 732-8991"
	} ]
	} ]`)
}
