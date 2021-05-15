package geist

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"
)

type SparqlClient struct {
	RestClient
	Parameters string
}

func NewSparqlClient(endpoint string) *SparqlClient {
	sc := new(SparqlClient)
	sc.HttpClient = &http.Client{}
	sc.Endpoint = endpoint
	return sc
}

func (sc *SparqlClient) PostSparqlRequest(contentType string, acceptType string,
	requestBody []byte) (responseBody []byte, err error) {
	return sc.PostRequest(sc.Endpoint+sc.Parameters, contentType, acceptType, requestBody)
}

func (sc *SparqlClient) PostData(format string, data []byte) (responseBody []byte, err error) {
	responseBody, err = sc.PostSparqlRequest(format, "text/plain", data)
	return
}

func (sc SparqlClient) Select(query string) (rs *ResultSet, err error) {
	responseBody, err := sc.PostSparqlRequest("application/sparql-query", "application/json", []byte(query))
	if err != nil {
		err = NewGeistError("Error posting SPARQL request", err, false)
		return
	}

	err = json.Unmarshal(responseBody, &rs)
	if err != nil {
		print()
		print(string(responseBody))
		print()
	}
	return
}

func (sc *SparqlClient) SelectCSV(query string) (csv string, err error) {
	responseBody, err := sc.PostSparqlRequest("application/sparql-query", "text/csv", []byte(query))
	csv = string(responseBody)
	return
}

func (sc *SparqlClient) SelectXML(query string) (csv string, err error) {
	responseBody, err := sc.PostSparqlRequest("application/sparql-query", "sparql-results+xml", []byte(query))
	csv = string(responseBody)
	return
}

func (sc *SparqlClient) SelectAll() (rs *ResultSet, err error) {
	rs, err = sc.Select(
		`SELECT ?s ?p ?o
		 WHERE
		 { ?s ?p ?o }`,
	)
	return
}

func (sc *SparqlClient) Construct(format string, query string) (triples []byte, err error) {
	triples, err = sc.PostSparqlRequest("application/sparql-query", format, []byte(query))
	return
}

func (sc *SparqlClient) ConstructAll(format string, sorted bool) (triples string, err error) {

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
		ntriplesSlice := strings.Split(strings.Trim(triples, "\n"), "\n")
		sort.Strings(ntriplesSlice)
		triples = strings.Join(ntriplesSlice, "\n")
	}

	triples = strings.Trim(triples, " \n")

	return
}
