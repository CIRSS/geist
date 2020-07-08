package sparqlrep

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ExampleBlazegraph_PostData_EmptyBody() {
	client := &http.Client{}
	request, _ := http.NewRequest("POST", SparqlEndpoint, strings.NewReader(""))
	request.Header.Add("Content-Type", "application/x-turtle")
	response, _ := client.Do(request)
	b, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	fmt.Println(string(b)[0:54])
	// Output:
	// <?xml version="1.0"?><data modified="0" milliseconds="
}

func ExampleBlazegraph_PostData_NamespaceDeclarationsOnly() {
	client := &http.Client{}
	request, _ := http.NewRequest("POST", SparqlEndpoint, strings.NewReader(`
		@prefix ab:    <http://learningsparql.com/ns/addressbook#> .
		@prefix d:     <http://learningsparql.com/ns/data#> .
	`))
	request.Header.Add("Content-Type", "application/x-turtle")
	response, _ := client.Do(request)
	b, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	fmt.Println(string(b)[0:54])
	// Output:
	// <?xml version="1.0"?><data modified="0" milliseconds="
}

func ExampleBlazegraph_PostData_TwoTriples() {
	client := &http.Client{}
	request, _ := http.NewRequest("POST", SparqlEndpoint, strings.NewReader(`
		@prefix ab:    <http://learningsparql.com/ns/addressbook#> .
		@prefix d:     <http://learningsparql.com/ns/data#> .

		d:y ab:tag "seven" .
		d:x ab:tag "eight" .
	`))
	request.Header.Add("Content-Type", "application/x-turtle")
	response, _ := client.Do(request)
	b, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	fmt.Println(string(b)[0:54])
	// Output:
	// <?xml version="1.0"?><data modified="2" milliseconds="
}
