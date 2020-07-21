package blazegraph

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tmcphillips/blazegraph-util/sparql"
)

var SparqlEndpoint = "http://127.0.0.1:9999/blazegraph/sparql"

type Client struct {
	httpClient *http.Client
	endpoint   string
}

func NewClient() *Client {
	bc := new(Client)
	bc.httpClient = &http.Client{}
	bc.endpoint = "http://127.0.0.1:9999/blazegraph/sparql"
	return bc
}

func (bc *Client) DeleteAll() (responseBody []byte, err error) {
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

func (bc *Client) PostRequest(contentType string, acceptType string,
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

func (bc *Client) PostData(format string, data []byte) (responseBody []byte, err error) {
	responseBody, err = bc.PostRequest(format, "text/plain", data)
	return
}

func (bc *Client) Select(query string) (rs sparql.ResultSet, err error) {
	responseBody, err := bc.PostRequest("application/sparql-query", "application/json", []byte(query))
	err = json.Unmarshal(responseBody, &rs)
	return
}

func (bc *Client) SelectAll() (rs sparql.ResultSet, err error) {
	rs, err = bc.Select(
		`SELECT ?s ?p ?o
		 WHERE
		 { ?s ?p ?o }`,
	)
	return
}

func (bc *Client) Construct(format string, query string) (triples []byte, err error) {
	triples, err = bc.PostRequest("application/sparql-query", format, []byte(query))
	return
}

func (bc *Client) ConstructAll(format string) (triples string, err error) {
	responseBody, err := bc.Construct(format, `
		CONSTRUCT
		{ ?s ?p ?o }
		WHERE
		{ ?s ?p ?o }`,
	)
	if err == nil {
		triples = string(responseBody)
	}
	return
}

// func (bc *Client) DumpAsNTriples() (triples string, err error) {
// 	responseBody, err := bc.RequestAllTriples()
// 	if err != nil {
// 		return
// 	}

// 	var sr sparql.Result
// 	err = json.Unmarshal(responseBody, &sr)
// 	if err != nil {
// 		return
// 	}

// 	var dump strings.Builder
// 	for _, b := range sr.Bindings() {
// 		triple := fmt.Sprintf("%s %s %s .\n",
// 			b.DelimitedValue("s"), b.DelimitedValue("p"), b.DelimitedValue("o"))
// 		dump.WriteString(triple)
// 	}
// 	triples = dump.String()

// 	return
// }
