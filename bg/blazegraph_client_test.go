package bg

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/tmcphillips/blazegraph-util/assert"
)

func TestBlazegraphClient_GetAllTriplesAsJSON_EmptyStore(t *testing.T) {
	bc := NewBlazegraphClient()
	bc.DeleteAllTriples()
	assert.JSONEquals(t, bc.SelectAllTriples(),
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
