package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/cirss/geist"
	"github.com/cirss/geist/util"
)

func recreateDataset(bc *geist.BlazegraphClient) {
	bc.DestroyDataSet("kb")
	p := geist.NewDatasetProperties("kb")
	bc.CreateDataSet(p)
}

func TestBlazegraphClient_GetAllTriplesAsJSON_EmptyStore(t *testing.T) {
	bc := geist.NewBlazegraphClient(geist.DefaultUrl)
	recreateDataset(bc)
	triples, _ := bc.SelectAll()
	actual, _ := triples.JSONString()
	util.JSONEquals(t, actual,
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
	bc := geist.NewBlazegraphClient(geist.DefaultUrl)
	recreateDataset(bc)
	bc.PostData("application/x-turtle", []byte(`
	@prefix t: <http://tmcphill.net/tags#> .
	@prefix d: <http://tmcphill.net/data#> .
	d:y t:tag "seven" .
	`))

	resultSet, _ := bc.Select(
		`prefix t: <http://tmcphill.net/tags#>
		 SELECT ?s ?o
		 WHERE
		 { ?s t:tag ?o }
		 `)

	actual, _ := resultSet.JSONString()

	util.JSONEquals(t,
		actual,
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
	bc := geist.NewBlazegraphClient(geist.DefaultUrl)
	recreateDataset(bc)
	bc.PostData("application/x-turtle", []byte(`
		@prefix t: <http://tmcphill.net/tags#> .
		@prefix d: <http://tmcphill.net/data#> .

		d:x t:tag "seven" .
		d:y t:tag "eight" .
	`))

	resultSet, _ := bc.Select(
		`prefix ab: <http://tmcphill.net/tags#>
		 SELECT ?s ?o
		 WHERE
		 { ?s ab:tag ?o }
		 ORDER BY ?s ?o
		 `)

	actual, _ := resultSet.JSONString()

	util.JSONEquals(t,
		actual,
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
	bc := geist.NewBlazegraphClient(geist.DefaultUrl)
	recreateDataset(bc)
	bc.PostData("application/x-turtle", []byte(`
		@prefix t: <http://tmcphill.net/tags#> .
		@prefix d: <http://tmcphill.net/data#> .

		d:x t:tag "seven" .
		d:y t:tag "eight" .
	`))
	rs, _ := bc.Select(
		`prefix ab: <http://tmcphill.net/tags#>
		 SELECT ?s ?o
		 WHERE
		 { ?s ab:tag ?o }
		 ORDER BY ?s ?o
		`)

	util.StringEquals(t, strings.Join(rs.Variables(), ", "), "s, o")

	util.StringEquals(t, rs.Bindings()[0]["s"].Type, "uri")
	util.StringEquals(t, rs.Bindings()[0]["s"].Value, "http://tmcphill.net/data#x")
	util.StringEquals(t, rs.Bindings()[0]["o"].Type, "literal")
	util.StringEquals(t, rs.Bindings()[0]["o"].Value, "seven")

	util.StringEquals(t, rs.Bindings()[1]["s"].Type, "uri")
	util.StringEquals(t, rs.Bindings()[1]["s"].Value, "http://tmcphill.net/data#y")
	util.StringEquals(t, rs.Bindings()[1]["o"].Type, "literal")
	util.StringEquals(t, rs.Bindings()[1]["o"].Value, "eight")
}

func ExampleBlazegraphClient_DumpAsNTriples() {
	bc := geist.NewBlazegraphClient(geist.DefaultUrl)
	recreateDataset(bc)
	bc.PostData("application/x-turtle", []byte(`
		@prefix t: <http://tmcphill.net/tags#> .
		@prefix d: <http://tmcphill.net/data#> .

		d:x t:tag "seven" .
		d:y t:tag "eight" .
	`))
	triples, _ := bc.ConstructAll("text/plain", true)
	fmt.Println(triples)
	// Output:
	// <http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
	// <http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
}
