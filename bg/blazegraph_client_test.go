package bg

import (
	"fmt"
)

func ExampleBlazegraphClient_EmptyStore_GetAllTriplesAsJSON() {
	bc := NewBlazegraphClient()
	result := bc.GetAllTriplesAsJSON()
	fmt.Println(result)
	// Output:
	// {
	//   "head" : {
	//     "vars" : [ "s", "p", "o" ]
	//   },
	//   "results" : {
	//     "bindings" : [ { } ]
	//   }
	// }
}

func ExampleBlazegraph_Client_EmptyStore_PostTwoTriples() {
	bc := NewBlazegraphClient()
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
