package blazegraph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tmcphillips/blazegraph-util/sparql"
)

var SparqlEndpoint = "http://127.0.0.1:9999/blazegraph/sparql"

type BlazegraphClient struct {
	httpClient *http.Client
	endpoint   string
}

func NewBlazegraphClient() *BlazegraphClient {
	bc := new(BlazegraphClient)
	bc.httpClient = &http.Client{}
	bc.endpoint = "http://127.0.0.1:9999/blazegraph/sparql"
	return bc
}

func (bc *BlazegraphClient) DeleteAllTriples() (responseBody []byte, err error) {
	request, _ := http.NewRequest("DELETE", bc.endpoint, nil)
	response, err := bc.httpClient.Do(request)
	if err != nil {
		return
	}
	if responseBody, err = ioutil.ReadAll(response.Body); err != nil {
		return
	}
	response.Body.Close()
	return
}

func (bc *BlazegraphClient) PostRequest(contentType string, acceptType string,
	requestBody []byte) (responseBody []byte, err error) {

	// create the http requeest using the provided body
	request, err := http.NewRequest("POST", bc.endpoint, bytes.NewReader(requestBody))
	if err != nil {
		return
	}
	request.Header.Add("Content-Type", contentType)
	request.Header.Add("Accept", acceptType)

	// perform the request and obtain the response
	response, err := bc.httpClient.Do(request)
	if err != nil {
		return
	}

	// read the response
	responseBody, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	response.Body.Close()
	return
}

func (bc *BlazegraphClient) PostJSONLDBytes(data []byte) (responseBody []byte, err error) {
	responseBody, err = bc.PostRequest("application/ld+json", "text/plain", data)
	return
}

func (bc *BlazegraphClient) PostTurtleBytes(data []byte) (responseBody []byte, err error) {
	responseBody, err = bc.PostRequest("application/x-turtle", "text/plain", data)
	return
}

func (bc *BlazegraphClient) PostTurtleString(data string) (responseBody []byte, err error) {
	responseBody, err = bc.PostTurtleBytes([]byte(data))
	return
}

func (bc *BlazegraphClient) RequestAllTriples() (responseBody []byte, err error) {
	responseBody, err = bc.PostSparqlQuery(
		`SELECT ?s ?p ?o
		 WHERE
		 { ?s ?p ?o }`,
	)
	return
}

func (bc *BlazegraphClient) RequestAllTriplesAsJSON() (resultJSON interface{}, err error) {
	responseBody, err := bc.RequestAllTriples()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseBody, &resultJSON)
	return
}

func (bc *BlazegraphClient) DumpAsNTriples() (triples string, err error) {
	responseBody, err := bc.RequestAllTriples()
	if err != nil {
		return
	}

	var sr sparql.SparqlResult
	err = json.Unmarshal(responseBody, &sr)
	if err != nil {
		return
	}

	var dump strings.Builder
	for _, b := range sr.Bindings() {
		triple := fmt.Sprintf("%s %s %s .\n",
			b.DelimitedValue("s"), b.DelimitedValue("p"), b.DelimitedValue("o"))
		dump.WriteString(triple)
	}
	triples = dump.String()

	return
}

func (bc *BlazegraphClient) PostSparqlQuery(query string) (responseBody []byte, err error) {
	return bc.PostRequest("application/sparql-query", "application/json", []byte(query))
}

func (bc *BlazegraphClient) SparqlQuery(query string) (sr sparql.SparqlResult, err error) {
	responseBody, err := bc.PostSparqlQuery(query)
	err = json.Unmarshal(responseBody, &sr)
	return
}
