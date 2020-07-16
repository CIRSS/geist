package bg

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
		return nil, err
	}
	responseBody, _ = ioutil.ReadAll(response.Body)
	if err != nil {
		return responseBody, err
	}
	response.Body.Close()
	return
}

func (bc *BlazegraphClient) PostRequest(contentType string, acceptType string,
	requestBody []byte) (responseBody []byte) {
	request, _ := http.NewRequest("POST", bc.endpoint, bytes.NewReader(requestBody))
	request.Header.Add("Content-Type", contentType)
	request.Header.Add("Accept", acceptType)
	response, _ := bc.httpClient.Do(request)
	responseBody, _ = ioutil.ReadAll(response.Body)
	response.Body.Close()
	return
}

func (bc *BlazegraphClient) PostNewData(data []byte) (responseBody []byte) {
	responseBody = bc.PostRequest("application/x-turtle", "text/plain", data)
	return
}

func (bc *BlazegraphClient) PostNewStringData(data string) (responseBody []byte) {
	responseBody = bc.PostNewData([]byte(data))
	return
}

func (bc *BlazegraphClient) RequestAllTriples() (responseBody []byte) {
	responseBody = bc.PostSparqlQuery(
		`SELECT ?s ?p ?o
		 WHERE
		 { ?s ?p ?o }`,
	)
	return
}

func (bc *BlazegraphClient) RequestAllTriplesAsJSON() interface{} {
	responseBody := bc.RequestAllTriples()
	var resultJSON interface{}
	json.Unmarshal(responseBody, &resultJSON)
	return resultJSON
}

func (bc *BlazegraphClient) DumpAsNTriples() string {
	responseBody := bc.RequestAllTriples()
	var sr sparql.SparqlResult
	json.Unmarshal(responseBody, &sr)
	var dump strings.Builder
	for _, b := range sr.Bindings() {
		triple := fmt.Sprintf("%s %s %s .\n",
			b.DelimitedValue("s"), b.DelimitedValue("p"), b.DelimitedValue("o"))
		dump.WriteString(triple)
	}
	return dump.String()
}

func (bc *BlazegraphClient) PostSparqlQuery(query string) (responseBody []byte) {
	return bc.PostRequest("application/sparql-query", "application/json", []byte(query))
}

func (bc *BlazegraphClient) SparqlQuery(query string) (sr sparql.SparqlResult, err error) {
	responseBody := bc.PostSparqlQuery(query)
	err = json.Unmarshal(responseBody, &sr)
	return
}
