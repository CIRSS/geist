package blazegraph

import (
	"io/ioutil"
	"net/http"

	"github.com/cirss/geist/sparql"
)

var DefaultBlazegraphUrl = "http://127.0.0.1:9999/blazegraph"
var DefaultSparqlEndpoint = DefaultBlazegraphUrl + "/sparql"
var DefaultNamespaceEndpoint = DefaultBlazegraphUrl + "/namespace"

type Client struct {
	sparql.Client
}

func NewClient() *Client {
	bc := new(Client)
	bc.HttpClient = &http.Client{}
	bc.Endpoint = DefaultSparqlEndpoint
	return bc
}

func (sc *Client) CreateDataSet(name string) (responseBody []byte, err error) {

	body :=
		`com.bigdata.rdf.sail.namespace=kb
com.bigdata.rdf.sail.truthMaintenance=false
com.bigdata.rdf.store.AbstractTripleStore.quads=true
com.bigdata.rdf.store.AbstractTripleStore.statementIdentifiers=false
com.bigdata.rdf.store.AbstractTripleStore.axiomsClass=com.bigdata.rdf.axioms.NoAxioms
`
	responseBody, err = sc.PostRequest(DefaultNamespaceEndpoint,
		"text/plain", "text/plain", []byte(body))
	return
}

func (sc *Client) DestroyDataSet(name string) (responseBody []byte, err error) {
	request, _ := http.NewRequest("DELETE", DefaultNamespaceEndpoint+"/"+name, nil)
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
