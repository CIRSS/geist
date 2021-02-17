package blazegraph

import (
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/cirss/geist"
)

var DefaultUrl = "http://127.0.0.1:9999/blazegraph"

type BlazegraphClient struct {
	geist.SparqlClient
	Url               string
	NamespaceEndpoint string
}

func NewBlazegraphClient(url string) *BlazegraphClient {
	bc := new(BlazegraphClient)
	bc.Url = url
	bc.SparqlEndpoint = bc.Url + "/sparql"
	bc.NamespaceEndpoint = bc.Url + "/namespace"
	bc.HttpClient = &http.Client{}
	return bc
}

func (sc *BlazegraphClient) CreateDataSet(dp *DatasetProperties) (responseBody []byte, err error) {

	body := dp.String()

	responseBody, err = sc.PostRequest(sc.NamespaceEndpoint,
		"text/plain", "text/plain", []byte(body))
	return
}

func (sc *BlazegraphClient) DestroyDataSet(name string) (responseBody []byte, err error) {
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

func (bc *BlazegraphClient) selectFunc(rp *geist.Template, queryText string, args []interface{}) (rs interface{}, err error) {

	var data interface{}
	if len(args) == 1 {
		data = args[0]
	}

	query, re := rp.ExpandSubreport("select", geist.PrependPrefixes(rp, queryText), data)
	if re != nil {
		return
	}
	return bc.Select(query)
}

func (bc *BlazegraphClient) ExpandReport(rp *geist.Template) (report string, err error) {

	funcs := template.FuncMap{
		"prefix": func(prefix string, uri string) (s string, err error) {
			rp.Properties.Prefixes[prefix] = uri
			return "", nil
		},
		"select": func(queryText string, args ...interface{}) (interface{}, error) {
			return bc.selectFunc(rp, queryText, args)
		},
	}

	rp.AddFuncs(funcs)
	rp.Parse()
	report, err = rp.Expand(nil)

	return
}

func (sc *BlazegraphClient) ListDatasets() (responseBody []byte, err error) {
	responseBody, err = sc.GetRequest(sc.NamespaceEndpoint,
		"text/plain", "text/plain")
	return
}
