package blazegraph

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/tmcphillips/blazegraph-util/assert"
)

func TestClient_GetAllTriplesAsJSON_EmptyStore(t *testing.T) {
	bc := NewClient()
	bc.DeleteAllTriples()
	triples, _ := bc.RequestAllTriplesAsJSON()
	assert.JSONEquals(t, triples,
		`{
			"head" : {
				"vars" : [ "s", "p", "o" ]
			},
			"results" : {
				"bindings" : [ ]
			}
		}`)
}

func TestClient_InsertOneTriple(t *testing.T) {
	bc := NewClient()
	bc.DeleteAllTriples()
	bc.PostTurtleString(`
	@prefix t: <http://tmcphill.net/tags#> .
	@prefix d: <http://tmcphill.net/data#> .
	d:y t:tag "seven" .
	`)

	responseBody, _ := bc.PostSparqlQuery(
		`prefix t: <http://tmcphill.net/tags#>
		 SELECT ?s ?o
		 WHERE
		 { ?s t:tag ?o }
		 `)

	var resultJSON interface{}
	json.Unmarshal(responseBody, &resultJSON)

	assert.JSONEquals(t,
		resultJSON,
		`{
			"head" : { "vars" : [ "s", "o" ] },
			"results" : {
			  "bindings" : [ {
				"s" : { "type" : "uri",     "value" : "http://tmcphill.net/data#y" },
				"o" : { "type" : "literal", "value" : "seven" }
			  } ]
			}
		  }`)
}

func TestClient_InsertTwoTriples(t *testing.T) {
	bc := NewClient()
	bc.DeleteAllTriples()
	bc.PostTurtleString(`
		@prefix t: <http://tmcphill.net/tags#> .
		@prefix d: <http://tmcphill.net/data#> .

		d:x t:tag "seven" .
		d:y t:tag "eight" .
	`)

	responseBody, _ := bc.PostSparqlQuery(
		`prefix ab: <http://tmcphill.net/tags#>
		 SELECT ?s ?o
		 WHERE
		 { ?s ab:tag ?o }
		 `)

	var resultJSON interface{}
	json.Unmarshal(responseBody, &resultJSON)

	assert.JSONEquals(t,
		resultJSON,
		`{
			"head" : { "vars" : [ "s", "o" ] },
			"results" : {
			  "bindings" : [ {
				"s" : { "type" : "uri",     "value" : "http://tmcphill.net/data#x" },
				"o" : { "type" : "literal", "value" : "seven" }
			}, {
				"s" : { "type" : "uri",     "value" : "http://tmcphill.net/data#y" },
				"o" : { "type" : "literal", "value" : "eight" }
			  } ]
			}
		  }`)
}

func TestClient_InsertTwoTriples_Struct(t *testing.T) {
	bc := NewClient()
	bc.DeleteAllTriples()
	bc.PostTurtleString(`
		@prefix t: <http://tmcphill.net/tags#> .
		@prefix d: <http://tmcphill.net/data#> .

		d:x t:tag "seven" .
		d:y t:tag "eight" .
	`)
	sr, _ := bc.SparqlQuery(
		`prefix ab: <http://tmcphill.net/tags#>
		 SELECT ?s ?o
		 WHERE
		 { ?s ab:tag ?o }
		 `)

	assert.StringEquals(t, strings.Join(sr.Vars(), ", "), "s, o")

	assert.StringEquals(t, sr.Bindings()[0]["s"].Type, "uri")
	assert.StringEquals(t, sr.Bindings()[0]["s"].Value, "http://tmcphill.net/data#x")
	assert.StringEquals(t, sr.Bindings()[0]["o"].Type, "literal")
	assert.StringEquals(t, sr.Bindings()[0]["o"].Value, "seven")

	assert.StringEquals(t, sr.Bindings()[1]["s"].Type, "uri")
	assert.StringEquals(t, sr.Bindings()[1]["s"].Value, "http://tmcphill.net/data#y")
	assert.StringEquals(t, sr.Bindings()[1]["o"].Type, "literal")
	assert.StringEquals(t, sr.Bindings()[1]["o"].Value, "eight")
}

func ExampleClient_DumpAsNTriples() {
	bc := NewClient()
	bc.DeleteAllTriples()
	bc.PostTurtleString(`
		@prefix t: <http://tmcphill.net/tags#> .
		@prefix d: <http://tmcphill.net/data#> .

		d:x t:tag "seven" .
		d:y t:tag "eight" .
	`)
	triples, _ := bc.DumpAsNTriples()
	fmt.Println(triples)
	// Output:
	// <http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
	// <http://tmcphill.net/tags#tag> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/1999/02/22-rdf-syntax-ns#Property> .
	// <http://tmcphill.net/tags#tag> <http://www.w3.org/2000/01/rdf-schema#subPropertyOf> <http://tmcphill.net/tags#tag> .
	// <http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
}
