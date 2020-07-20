package main

import (
	"os"
	"strings"
	"testing"

	"github.com/tmcphillips/blazegraph-util/assert"
)

func runWithArgs(commandLine string) {
	os.Args = strings.Fields(commandLine)
	Main.Run()
}

func ExampleBlazegraphCmd_drop_then_dump() {
	Main.OutWriter = os.Stdout
	Main.ErrWriter = os.Stdout
	runWithArgs("blazegraph drop")
	runWithArgs("blazegraph dump")
	// Output:
	//
}

func ExampleBlazegraphCmd_drop_load_turtle_then_dump_default() {
	Main.OutWriter = os.Stdout
	Main.ErrWriter = os.Stdout
	runWithArgs("blazegraph drop")
	runWithArgs("blazegraph load --file testdata/in.nt --format turtle")
	runWithArgs("blazegraph dump")
	// Output:
	// <http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
	// <http://tmcphill.net/tags#tag> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/1999/02/22-rdf-syntax-ns#Property> .
	// <http://tmcphill.net/tags#tag> <http://www.w3.org/2000/01/rdf-schema#subPropertyOf> <http://tmcphill.net/tags#tag> .
	// <http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
}

func ExampleBlazegraphCmd_drop_load_turtle_then_dump_ntriples() {
	Main.OutWriter = os.Stdout
	Main.ErrWriter = os.Stdout
	runWithArgs("blazegraph drop")
	runWithArgs("blazegraph load --file testdata/in.nt --format turtle")
	runWithArgs("blazegraph dump --format n-triples")
	// Output:
	// <http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
	// <http://tmcphill.net/tags#tag> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/1999/02/22-rdf-syntax-ns#Property> .
	// <http://tmcphill.net/tags#tag> <http://www.w3.org/2000/01/rdf-schema#subPropertyOf> <http://tmcphill.net/tags#tag> .
	// <http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
}

func TestBlazegraphCmd_drop_load_turtle_then_dump_jsonld(t *testing.T) {
	var resultsBuffer strings.Builder
	Main.OutWriter = &resultsBuffer
	Main.ErrWriter = &resultsBuffer

	runWithArgs("blazegraph drop")
	runWithArgs("blazegraph load --file testdata/in.nt --format turtle")
	runWithArgs("blazegraph dump --format json-ld")

	actualJSON, _ := assert.CanonicalJSONFromString(resultsBuffer.String())
	expectedJSON, _ := assert.CanonicalJSONFromString(`
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
	assert.StringEquals(t, actualJSON, expectedJSON)
}

func ExampleBlazegraphCmd_drop_load_jsonld_then_dump_ntriples() {
	Main.OutWriter = os.Stdout
	Main.ErrWriter = os.Stdout
	runWithArgs("blazegraph drop")
	runWithArgs("blazegraph load --file testdata/address-book.jsonld --format json-ld")
	runWithArgs("blazegraph dump")
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

func TestBlazegraphCmd_drop_load_jsonld_then_dump_jsonld(t *testing.T) {
	var resultsBuffer strings.Builder
	Main.OutWriter = &resultsBuffer
	Main.ErrWriter = &resultsBuffer

	runWithArgs("blazegraph drop")
	runWithArgs("blazegraph load --file testdata/address-book.jsonld --format json-ld")
	runWithArgs("blazegraph dump --format json-ld")

	actualJSON, _ := assert.CanonicalJSONFromString(resultsBuffer.String())
	expectedJSON, _ := assert.CanonicalJSONFromString(`
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

	assert.StringEquals(t, actualJSON, expectedJSON)
}
