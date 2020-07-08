package bg

import (
	"fmt"
)

func ExampleBlazegraphClient_GetAllTriples() {
	bc := NewBlazegraphClient()
	result := bc.GetAllTriples()
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
