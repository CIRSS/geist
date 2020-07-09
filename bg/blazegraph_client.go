package bg

import (
	"encoding/json"
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

func (bc *BlazegraphClient) deleteAllTriples() string {
	request, _ := http.NewRequest("DELETE", bc.endpoint, nil)
	response, _ := bc.httpClient.Do(request)
	responseBody, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	return string(responseBody)
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

func (bc *BlazegraphClient) PostSparqlQuery(query string) interface{} {
	resultString := bc.postRequest("application/sparql-query", "application/json", query)
	var resultJSON interface{}
	json.Unmarshal([]byte(resultString), &resultJSON)
	return resultJSON
}

func (bc *BlazegraphClient) PostNewData(data string) string {
	return bc.postRequest("application/x-turtle", "text/plain", data)
}

func (bc *BlazegraphClient) SelectAllTriples() interface{} {
	resultJSON := bc.PostSparqlQuery(
		`SELECT ?s ?p ?o
		 WHERE
		 { ?s ?p ?o }`,
	)
	return resultJSON
}
