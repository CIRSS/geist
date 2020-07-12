package bg

import (
	"encoding/json"
	"strings"
	"testing"

	tu "github.com/tmcphillips/blazegraph-util/testutil"
)

func TestBlazegraphClient_GetAllTriplesAsJSON_EmptyStore(t *testing.T) {
	bc := NewBlazegraphClient()
	bc.DeleteAllTriples()
	tu.AssertJSONEquals(t, bc.SelectAllTriples(),
		`{
			"head" : {
				"vars" : [ "s", "p", "o" ]
			},
			"results" : {
				"bindings" : [ ]
			}
		}`)
}

func TestBlazegraphClient_InsertOneTriple(t *testing.T) {
	bc := NewBlazegraphClient()
	bc.DeleteAllTriples()
	bc.PostNewData(`
	@prefix t: <http://tmcphill.net/tags#> .
	@prefix d: <http://tmcphill.net/data#> .
	d:y t:tag "seven" .
	`)

	responseBody := bc.PostSparqlQuery(
		`prefix t: <http://tmcphill.net/tags#>
		 SELECT ?s ?o
		 WHERE
		 { ?s t:tag ?o }
		 `)

	var resultJSON interface{}
	json.Unmarshal(responseBody, &resultJSON)

	tu.AssertJSONEquals(t,
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

func TestBlazegraphClient_InsertTwoTriples(t *testing.T) {
	bc := NewBlazegraphClient()
	bc.DeleteAllTriples()
	bc.PostNewData(`
		@prefix t: <http://tmcphill.net/tags#> .
		@prefix d: <http://tmcphill.net/data#> .

		d:x t:tag "seven" .
		d:y t:tag "eight" .
	`)

	responseBody := bc.PostSparqlQuery(
		`prefix ab: <http://tmcphill.net/tags#>
		 SELECT ?s ?o
		 WHERE
		 { ?s ab:tag ?o }
		 `)

	var resultJSON interface{}
	json.Unmarshal(responseBody, &resultJSON)

	tu.AssertJSONEquals(t,
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

func TestBlazegraphClient_InsertTwoTriples_Struct(t *testing.T) {
	bc := NewBlazegraphClient()
	bc.DeleteAllTriples()
	bc.PostNewData(`
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

	tu.AssertStringEquals(t, strings.Join(sr.Vars(), ", "), "s, o")

	tu.AssertStringEquals(t, sr.Bindings()[0]["s"].Type, "uri")
	tu.AssertStringEquals(t, sr.Bindings()[0]["s"].Value, "http://tmcphill.net/data#x")
	tu.AssertStringEquals(t, sr.Bindings()[0]["o"].Type, "literal")
	tu.AssertStringEquals(t, sr.Bindings()[0]["o"].Value, "seven")

	tu.AssertStringEquals(t, sr.Bindings()[1]["s"].Type, "uri")
	tu.AssertStringEquals(t, sr.Bindings()[1]["s"].Value, "http://tmcphill.net/data#y")
	tu.AssertStringEquals(t, sr.Bindings()[1]["o"].Type, "literal")
	tu.AssertStringEquals(t, sr.Bindings()[1]["o"].Value, "eight")
}
