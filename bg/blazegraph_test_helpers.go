package bg

import (
	"encoding/json"

	"github.com/ucarion/jcs"
)

var SparqlEndpoint = "http://127.0.0.1:9999/blazegraph/sparql"

func CanonicalJSONFromString(jsonString string) (string, error) {
	var originalJSON interface{}
	json.Unmarshal([]byte(jsonString), &originalJSON)
	return CanonicalJSON(originalJSON)
}

func CanonicalJSON(originalJSON interface{}) (string, error) {
	canonicalJSONString, err := jcs.Format(originalJSON)
	return canonicalJSONString, err
}
