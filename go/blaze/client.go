package blaze

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/cirss/geist/go/geist"
)

var DefaultUrl = "http://127.0.0.1:9999/blazegraph"
var DefaultDataset = "kb"

type InstanceStatus struct {
	InstanceUrl            string
	SparqlEndpoint         string
	BlazegraphBuildVersion string
	QueryStartCount        int
	RunningQueriesCount    int
	QueryDoneCount         int
	QueryErrorCount        int
}

type BlazegraphClient struct {
	geist.SparqlClient
	NamespaceEndpoint string
	Url               string
}

func NewBlazegraphClient(instanceUrl string) *BlazegraphClient {
	bc := new(BlazegraphClient)
	bc.Url = instanceUrl
	bc.NamespaceEndpoint = bc.Url + "/namespace"
	bc.HttpClient = &http.Client{}
	bc.SetDataset(DefaultDataset)
	return bc
}

func (bc *BlazegraphClient) SetDataset(dataset string) {
	bc.SparqlEndpoint = bc.SparqlEndpointForDataset(dataset)
}

func (bc *BlazegraphClient) SparqlEndpointForDataset(dataset string) string {
	return bc.NamespaceEndpoint + "/" + dataset + "/sparql"
}

func (bc *BlazegraphClient) CreateDataSet(dp *DatasetProperties) (response string, err error) {

	requestBody := dp.String()

	responseBody, err := bc.PostRequest(bc.NamespaceEndpoint,
		"text/plain", "text/plain", []byte(requestBody))
	response = string(responseBody)
	if err != nil {
		return
	}

	responseTokens := strings.Split(string(responseBody), ": ")
	switch responseTokens[0] {
	case "CREATED":
		break
	case "EXISTS":
		message := fmt.Sprintf("dataset %s already exists", responseTokens[1])
		err = geist.NewGeistError(message, err, false)
		break
	default:
		err = geist.NewGeistError(string(responseBody), err, false)
	}

	return
}

func (bc *BlazegraphClient) DestroyDataSet(name string) (response string, err error) {

	responseBody, err := bc.DeleteRequest(bc.NamespaceEndpoint + "/" + name)
	response = string(responseBody)
	if err != nil {
		return
	}

	if strings.Contains(response, "com.bigdata.rdf.sail.webapp.DatasetNotFoundException") {
		message := fmt.Sprintf("dataset %s does not exist", name)
		err = geist.NewGeistError(message, nil, false)
		return
	}

	expectedResponse := "DELETED: " + strings.Trim(name, " ")
	if response != expectedResponse {
		message := fmt.Sprintf("error destroying dataset %s:\n%s", name, response)
		err = geist.NewGeistError(message, nil, false)
		return
	}

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

func (bc *BlazegraphClient) ListDatasets() (datasets []string, err error) {
	responseBody, err := bc.GetRequest(bc.NamespaceEndpoint,
		"text/plain", "text/plain")
	if err != nil {
		return
	}
	// fmt.Println(string(responseBody))
	re := regexp.MustCompile(`Namespace> "([^"]+)"`)
	submatchall := re.FindAllStringSubmatch(string(responseBody), -1)
	for _, element := range submatchall {
		datasets = append(datasets, element[1])
	}
	sort.Strings(datasets)
	return
}

func ExtractStringUsingRegEx(s string, regex string) string {
	re := regexp.MustCompile(regex)
	submatch := re.FindStringSubmatch(s)
	return submatch[1]
}

func ExtractIntUsingRegEx(s string, regex string) (value int, err error) {
	re := regexp.MustCompile(regex)
	submatch := re.FindStringSubmatch(s)
	return strconv.Atoi(submatch[1])
}

func (bc *BlazegraphClient) GetStatus() (statusJSON string, err error) {

	responseBody, err := bc.GetRequest(bc.Url+"/status",
		"text/plain", "text/plain")

	if err != nil {
		err = geist.NewGeistError("Error posting SPARQL request", err, false)
		return
	}

	statusString := string(responseBody)
	status := InstanceStatus{}
	status.InstanceUrl = bc.Url
	status.SparqlEndpoint = bc.SparqlEndpoint
	status.BlazegraphBuildVersion = ExtractStringUsingRegEx(statusString, `span id="buildVersion">([0-9\.]+)</span`)
	status.QueryStartCount, _ = ExtractIntUsingRegEx(statusString, `queryStartCount=([0-9]+)`)
	status.RunningQueriesCount, _ = ExtractIntUsingRegEx(statusString, `runningQueriesCount=([0-9]+)`)
	status.QueryDoneCount, _ = ExtractIntUsingRegEx(statusString, `queryDoneCount=([0-9]+)`)
	status.QueryErrorCount, _ = ExtractIntUsingRegEx(statusString, `queryErrorCount=([0-9]+)`)
	statusJSONBytes, err := json.MarshalIndent(status, "", "    ")
	statusJSON = string(statusJSONBytes)
	return
}

func (bc *BlazegraphClient) CountTriples(dataset string, exact bool) (count int, err error) {
	responseBody, err := bc.GetRequest(bc.SparqlEndpointForDataset(dataset)+"?ESTCARD",
		"text/plain", "application/xml")
	if err != nil {
		return
	}
	count, err = ExtractIntUsingRegEx(string(responseBody), `rangeCount="([0-9]+)"`)
	return
}
