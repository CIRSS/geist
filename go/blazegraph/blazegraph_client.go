package blazegraph

import (
	"io/ioutil"
	"net/http"

	"github.com/cirss/geist/sparql"
)

var DefaultUrl = "http://127.0.0.1:9999/blazegraph"

type Client struct {
	sparql.Client
	Url               string
	NamespaceEndpoint string
}

func NewClient(url string) *Client {
	bc := new(Client)
	bc.Url = url
	bc.SparqlEndpoint = bc.Url + "/sparql"
	bc.NamespaceEndpoint = bc.Url + "/namespace"
	bc.HttpClient = &http.Client{}
	return bc
}

func (sc *Client) CreateDataSet(name string) (responseBody []byte, err error) {

	body :=
		`
		com.bigdata.rdf.sail.namespace=kb
		com.bigdata.rdf.sail.truthMaintenance=false
		com.bigdata.rdf.store.AbstractTripleStore.quads=true
		com.bigdata.rdf.store.AbstractTripleStore.statementIdentifiers=false
		com.bigdata.rdf.store.AbstractTripleStore.axiomsClass=com.bigdata.rdf.axioms.NoAxioms
		`
	responseBody, err = sc.PostRequest(sc.NamespaceEndpoint,
		"text/plain", "text/plain", []byte(body))
	return
}

func (sc *Client) DestroyDataSet(name string) (responseBody []byte, err error) {
	request, _ := http.NewRequest("DELETE", sc.NamespaceEndpoint+"/"+name, nil)
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
