package bg

import (
	"fmt"
	"testing"
)

func TestBlazegraphClient_GetAllTriplesAsJSON_EmptyStore(t *testing.T) {
	bc := NewBlazegraphClient()
	bc.deleteAllTriples()
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
	bc.deleteAllTriples()
	bc.PostNewData(`
	@prefix ab:    <http://learningsparql.com/ns/addressbook#> .
	@prefix d:     <http://learningsparql.com/ns/data#> .

	d:y ab:tag "seven" .
	`)

	queryResult := bc.PostSparqlQuery(
		`prefix ab: <http://learningsparql.com/ns/addressbook#>
		 SELECT ?s ?o
		 WHERE
		 { ?s ab:tag ?o }
		 `)

	AssertJSONEquals(t,
		queryResult,
		`{
			"head" : {
			  "vars" : [ "s", "o" ]
			},
			"results" : {
			  "bindings" : [ {
				"s" : {
				  "type" : "uri",
				  "value" : "http://learningsparql.com/ns/data#y"
				},
				"o" : {
				  "type" : "literal",
				  "value" : "seven"
				}
			  } ]
			}
		  }`)
}

func ExampleBlazegraph_Client_EmptyStore_PostTwoTriples() {
	bc := NewBlazegraphClient()
	bc.deleteAllTriples()
	result := bc.PostNewData(`
		@prefix ab:    <http://learningsparql.com/ns/addressbook#> .
		@prefix d:     <http://learningsparql.com/ns/data#> .

		d:y ab:tag "seven" .
		d:x ab:tag "eight" .
	`)
	fmt.Println(result[0:54])
	// Output:
	// <?xml version="1.0"?><data modified="2" milliseconds="
}
