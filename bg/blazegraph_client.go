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

func (bc *BlazegraphClient) postRequest(contentType string, acceptType string,
	requestBody string) string {
	request, _ := http.NewRequest("POST", bc.endpoint, strings.NewReader(requestBody))
	request.Header.Add("Content-Type", contentType)
	request.Header.Add("Accept", acceptType)
	response, _ := bc.httpClient.Do(request)
	responseBody, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	return string(responseBody)
}

func (bc *BlazegraphClient) PostSparqlQuery(resultFormat string, query string) string {
	return bc.postRequest("application/sparql-query", resultFormat, query)
}

func (bc *BlazegraphClient) PostNewData(data string) string {
	return bc.postRequest("application/x-turtle", "text/plain", data)
}

func (bc *BlazegraphClient) GetAllTriplesAsJSON() string {
	result := bc.PostSparqlQuery(
		"application/json",
		`SELECT ?s ?p ?o
		 WHERE
		 {}`,
	)
	return result
}
