package sparql

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

type Client struct {
	HttpClient *http.Client
	Endpoint   string
}

func NewClient(endpoint string) *Client {
	sc := new(Client)
	sc.HttpClient = &http.Client{}
	sc.Endpoint = endpoint
	return sc
}

func (sc *Client) PostRequest(url string, contentType string, acceptType string,
	requestBody []byte) (responseBody []byte, err error) {

	// create the http requeest using the provided body
	request, err := http.NewRequest("POST", url, bytes.NewReader(requestBody))
	if err != nil {
		return
	}
	request.Header.Add("Content-Type", contentType)
	request.Header.Add("Accept", acceptType)

	// perform the request and obtain the response
	response, err := sc.HttpClient.Do(request)
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

func (sc *Client) PostSparqlRequest(contentType string, acceptType string,
	requestBody []byte) (responseBody []byte, err error) {
	return sc.PostRequest(sc.Endpoint, contentType, acceptType, requestBody)
}

func (sc *Client) DeleteAll() (responseBody []byte, err error) {
	request, _ := http.NewRequest("DELETE", sc.Endpoint, nil)
	response, err := sc.HttpClient.Do(request)
	if err != nil {
		return
	}
	if responseBody, err = ioutil.ReadAll(response.Body); err != nil {
		return
	}
	response.Body.Close()
	return
}

func (sc *Client) PostData(format string, data []byte) (responseBody []byte, err error) {
	responseBody, err = sc.PostSparqlRequest(format, "text/plain", data)
	return
}

func (sc *Client) Select(query string) (rs *ResultSet, err error) {
	responseBody, err := sc.PostSparqlRequest("application/sparql-query", "application/json", []byte(query))
	err = json.Unmarshal(responseBody, &rs)
	if err != nil {
		print()
		print(string(responseBody))
		print()
	}
	return
}

func (sc *Client) SelectCSV(query string) (csv string, err error) {
	responseBody, err := sc.PostSparqlRequest("application/sparql-query", "text/csv", []byte(query))
	csv = string(responseBody)
	return
}

func (sc *Client) SelectXML(query string) (csv string, err error) {
	responseBody, err := sc.PostSparqlRequest("application/sparql-query", "sparql-results+xml", []byte(query))
	csv = string(responseBody)
	return
}

func (sc *Client) SelectAll() (rs *ResultSet, err error) {
	rs, err = sc.Select(
		`SELECT ?s ?p ?o
		 WHERE
		 { ?s ?p ?o }`,
	)
	return
}

func (sc *Client) Construct(format string, query string) (triples []byte, err error) {
	triples, err = sc.PostSparqlRequest("application/sparql-query", format, []byte(query))
	return
}

func (sc *Client) ConstructAll(format string, sorted bool) (triples string, err error) {

	responseBody, err := sc.Construct(format, `
		CONSTRUCT
		{ ?s ?p ?o }
		WHERE
		{ ?s ?p ?o }`,
	)
	if err != nil {
		return
	}

	triples = string(responseBody)

	if sorted && format == "text/plain" {
		ntriplesSlice := strings.Split(triples, "\n")
		sort.Strings(ntriplesSlice)
		triples = strings.Join(ntriplesSlice, "\n")
	}

	return
}
