package bg

import (
	"encoding/json"
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
	requestBody string) (responseBody []byte) {
	request, _ := http.NewRequest("POST", bc.endpoint, strings.NewReader(requestBody))
	request.Header.Add("Content-Type", contentType)
	request.Header.Add("Accept", acceptType)
	response, _ := bc.httpClient.Do(request)
	responseBody, _ = ioutil.ReadAll(response.Body)
	response.Body.Close()
	return
}

func (bc *BlazegraphClient) PostNewData(data string) (responseBody []byte) {
	responseBody = bc.PostRequest("application/x-turtle", "text/plain", data)
	return
}

func (bc *BlazegraphClient) SelectAllTriples() interface{} {
	responseBody := bc.PostSparqlQuery(
		`SELECT ?s ?p ?o
		 WHERE
		 { ?s ?p ?o }`,
	)
	var resultJSON interface{}
	json.Unmarshal(responseBody, &resultJSON)
	return resultJSON
}

func (bc *BlazegraphClient) PostSparqlQuery(query string) (responseBody []byte) {
	return bc.PostRequest("application/sparql-query", "application/json", query)
}

func (bc *BlazegraphClient) SparqlQuery(query string) (sr sparql.SparqlResult, err error) {
	responseBody := bc.PostSparqlQuery(query)
	err = json.Unmarshal(responseBody, &sr)
	return
}
