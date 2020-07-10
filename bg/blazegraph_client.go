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

func (bc *BlazegraphClient) DeleteAllTriples() string {
	request, _ := http.NewRequest("DELETE", bc.endpoint, nil)
	response, _ := bc.httpClient.Do(request)
	responseBody, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	return string(responseBody)
}

func (bc *BlazegraphClient) PostRequest(contentType string, acceptType string,
	requestBody string) string {
	request, _ := http.NewRequest("POST", bc.endpoint, strings.NewReader(requestBody))
	request.Header.Add("Content-Type", contentType)
	request.Header.Add("Accept", acceptType)
	response, _ := bc.httpClient.Do(request)
	responseBody, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	return string(responseBody)
}

func (bc *BlazegraphClient) PostSparqlQuery(query string) (interface{}, error) {
	resultString := bc.PostRequest("application/sparql-query", "application/json", query)
	var resultJSON interface{}
	err := json.Unmarshal([]byte(resultString), &resultJSON)
	return resultJSON, err
}

func (bc *BlazegraphClient) PostNewData(data string) string {
	return bc.PostRequest("application/x-turtle", "text/plain", data)
}

func (bc *BlazegraphClient) SelectAllTriples() interface{} {
	resultJSON, _ := bc.PostSparqlQuery(
		`SELECT ?s ?p ?o
		 WHERE
		 { ?s ?p ?o }`,
	)
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

func (bc *BlazegraphClient) SparqlQuery(query string) (SparqlResult, error) {
	resultString := bc.PostRequest("application/sparql-query", "application/json", query)
	var result SparqlResult
	err := json.Unmarshal([]byte(resultString), &result)
	return result, err
}

// func SparqlResultVars(sr SparqlResult) string {
// 	var s string
// 	fmt.Sprintf("%s %s", sr.Head.Vars[0], sr.Head.Vars[1])
// 	return
