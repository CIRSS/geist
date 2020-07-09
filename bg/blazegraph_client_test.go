package bg

import (
	"fmt"
	"testing"
)

func TestBlazegraphClient_GetAllTriplesAsJSON(t *testing.T) {
	bc := NewBlazegraphClient()
	bc.deleteAllTriples()
	assertJSONEquals(t, bc.GetAllTriplesAsJSON(),
		`{
			"head" : {
				"vars" : [ "s", "p", "o" ]
			},
			"results" : {
				"bindings" : [ ]
			}
		}`)
}

func ExampleBlazegraph_Client_EmptyStore_OneTriple() {
	bc := NewBlazegraphClient()
	bc.deleteAllTriples()
	result := bc.PostNewData(`
	@prefix ab:    <http://learningsparql.com/ns/addressbook#> .
	@prefix d:     <http://learningsparql.com/ns/data#> .

	d:y ab:tag "seven" .
	`)
	fmt.Println(result[0:54])
	// Output:
	// <?xml version="1.0"?><data modified="1" milliseconds="
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
