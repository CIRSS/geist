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

func (bc *BlazegraphClient) DeleteAllTriples() (responseBody []byte) {
	request, _ := http.NewRequest("DELETE", bc.endpoint, nil)
	response, _ := bc.httpClient.Do(request)
	responseBody, _ = ioutil.ReadAll(response.Body)
	response.Body.Close()
	return responseBody
}

func (bc *BlazegraphClient) PostRequest(contentType string, acceptType string,
	requestBody string) (responseBody []byte) {
	request, _ := http.NewRequest("POST", bc.endpoint, strings.NewReader(requestBody))
	request.Header.Add("Content-Type", contentType)
	request.Header.Add("Accept", acceptType)
	response, _ := bc.httpClient.Do(request)
	responseBody, _ = ioutil.ReadAll(response.Body)
	response.Body.Close()
	return responseBody
}

func (bc *BlazegraphClient) PostNewData(data string) (responseBody []byte) {
	return bc.PostRequest("application/x-turtle", "text/plain", data)
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

type SparqlResult struct {
	Head struct {
		Vars []string
	}
	Results struct {
		Bindings []struct {
			S struct {
				Type  string
				Value string
			}
			O struct {
				Type  string
				Value string
			}
		}
	}
}

func (bc *BlazegraphClient) PostSparqlQuery(query string) (responseBody []byte) {
	return bc.PostRequest("application/sparql-query", "application/json", query)
}

func (bc *BlazegraphClient) SparqlQuery(query string) (SparqlResult, error) {
	responseBody := bc.PostSparqlQuery(query)
	var sr SparqlResult
	err := json.Unmarshal(responseBody, &sr)
	return sr, err
}
