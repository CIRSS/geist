package bg

import (
	"strings"
	"testing"
)

func TestBlazegraphClient_GetAllTriplesAsJSON_EmptyStore(t *testing.T) {
	bc := NewBlazegraphClient()
	bc.DeleteAllTriples()
	AssertJSONEquals(t, bc.SelectAllTriples(),
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

	queryResult, err := bc.PostSparqlQuery(
		`prefix t: <http://tmcphill.net/tags#>
		 SELECT ?s ?o
		 WHERE
		 { ?s t:tag ?o }
		 `)

	t.Log(err)

	AssertJSONEquals(t,
		queryResult,
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
	queryResult, _ := bc.PostSparqlQuery(
		`prefix ab: <http://tmcphill.net/tags#>
		 SELECT ?s ?o
		 WHERE
		 { ?s ab:tag ?o }
		 `)

	AssertJSONEquals(t,
		queryResult,
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
	qr, _ := bc.SparqlQuery(
		`prefix ab: <http://tmcphill.net/tags#>
		 SELECT ?s ?o
		 WHERE
		 { ?s ab:tag ?o }
		 `)

	AssertStringEquals(t, strings.Join(qr.Head.Vars, ","), "s,o")
	AssertStringEquals(t, qr.Results.Bindings[0].S.Type, "uri")
	AssertStringEquals(t, qr.Results.Bindings[0].S.Value, "http://tmcphill.net/data#x")
	AssertStringEquals(t, qr.Results.Bindings[0].O.Type, "literal")
	AssertStringEquals(t, qr.Results.Bindings[0].O.Value, "seven")
	AssertStringEquals(t, qr.Results.Bindings[1].S.Type, "uri")
	AssertStringEquals(t, qr.Results.Bindings[1].S.Value, "http://tmcphill.net/data#y")
	AssertStringEquals(t, qr.Results.Bindings[1].O.Type, "literal")
	AssertStringEquals(t, qr.Results.Bindings[1].O.Value, "eight")
}
