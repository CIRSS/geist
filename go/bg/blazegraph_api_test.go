package sparqlrep

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ExampleBlazegraph_EmptyRequest_HttpGet_StatusOK() {
	response, _ := http.Get(SparqlEndpoint)
	fmt.Println(response.Status)
	// Output:
	// 200 OK
}

func ExampleBlazegraph_EmptyRequest_HttpClientGet_StatusOK() {
	client := &http.Client{}
	response, _ := client.Get(SparqlEndpoint)
	fmt.Println(response.Status)
	// Output:
	// 200 OK
}

func ExampleBlazegraph_EmptyRequest_HttpClientDoGet_StatusOK() {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", SparqlEndpoint, nil)
	response, _ := client.Do(request)
	fmt.Println(response.Status)
	// Output:
	// 200 OK
}

func ExampleBlazegraph_EmptyRequest_HttpClientDoGet_ResponseContentType_RdfXml() {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", SparqlEndpoint, nil)
	response, _ := client.Do(request)
	fmt.Println(response.Header["Content-Type"])
	// Output:
	// [application/rdf+xml]
}

func ExampleBlazegraph_EmptyRequest_HttpClientDoGet_ResponseContentType_JSON() {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", SparqlEndpoint, nil)
	request.Header.Add("Accept", "application/json")
	response, _ := client.Do(request)
	fmt.Println(response.Header["Content-Type"])
	// Output:
	// [application/sparql-results+json]
}

func ExampleBlazegraph_EmptyRequest_Body_First100Bytes() {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", SparqlEndpoint, nil)
	request.Header.Add("Accept", "application/json")
	response, _ := client.Do(request)
	b, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	fmt.Println(string(b[0:100]))
	// Output:
	// {
	//   "head" : {
	//     "vars" : [ "subject", "predicate", "object", "context" ]
	//   },
	//   "results" : {
}

func ExampleBlazegraph_PostSparqlRequest_SelectAllTriples() {
	client := &http.Client{}
	request, _ := http.NewRequest("POST", SparqlEndpoint, strings.NewReader(`
		SELECT ?s ?p ?o
		WHERE
		{}
	`))
	request.Header.Add("Content-Type", "application/sparql-query")
	request.Header.Add("Accept", "application/json")
	response, _ := client.Do(request)
	b, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	fmt.Println(string(b))
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
