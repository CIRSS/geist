package bg

import (
	"io/ioutil"
	"net/http"
	"strings"
)

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

func (bc *BlazegraphClient) PostSparqlQuery(query string) string {
	request, _ := http.NewRequest("POST", bc.endpoint, strings.NewReader(query))
	request.Header.Add("Content-Type", "application/sparql-query")
	request.Header.Add("Accept", "application/json")
	response, _ := bc.httpClient.Do(request)
	body, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	return string(body)
}

func (bc *BlazegraphClient) GetAllTriples() string {
	result := bc.PostSparqlQuery(`
		SELECT ?s ?p ?o
		WHERE
		{}
	`)
	return result
}
