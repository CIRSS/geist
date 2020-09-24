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

type DatasetProperties struct {
	Name      string
	Inference bool
	Quads     bool
}

func NewClient(url string) *Client {
	bc := new(Client)
	bc.Url = url
	bc.SparqlEndpoint = bc.Url + "/sparql"
	bc.NamespaceEndpoint = bc.Url + "/namespace"
	bc.HttpClient = &http.Client{}
	return bc
}

func (sc *Client) CreateDataSet(properties DatasetProperties) (responseBody []byte, err error) {

	body := "com.bigdata.rdf.sail.namespace=" + properties.Name + "\n"

	if properties.Inference {
		body += "com.bigdata.rdf.sail.truthMaintenance=true\n"
		body += "com.bigdata.rdf.store.AbstractTripleStore.axiomsClass=com.bigdata.rdf.axioms.OwlAxioms\n"
	} else {
		body += "com.bigdata.rdf.sail.truthMaintenance=false\n"
		body += "com.bigdata.rdf.store.AbstractTripleStore.axiomsClass=com.bigdata.rdf.axioms.NoAxioms\n"
	}

	if properties.Quads {
		body += "com.bigdata.rdf.store.AbstractTripleStore.quads=true\n"
	} else {
		body += "com.bigdata.rdf.store.AbstractTripleStore.quads=false\n"
	}

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
