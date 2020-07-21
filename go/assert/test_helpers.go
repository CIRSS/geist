package assert

import (
	"encoding/json"
	"testing"

	"github.com/ucarion/jcs"
)

var SparqlEndpoint = "http://127.0.0.1:9999/blazegraph/sparql"

func canonicalJSON(jsonString string) (string, error) {
	var originalJSON interface{}
	json.Unmarshal([]byte(jsonString), &originalJSON)
	canonicalJSONString, err := jcs.Format(originalJSON)
	return canonicalJSONString, err
}

func StringEquals(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Log("assertStringEquals:\n\nexpected: " + expected + "\nactual:   " + actual + "\n")
		t.Fail()
	}
}

func JSONEquals(t *testing.T, actual string, expected string) {
	actualCanonical, _ := canonicalJSON(actual)
	expectedCanonical, _ := canonicalJSON(expected)
	if actualCanonical != expectedCanonical {
		t.Log("assertEquivalentJSON:\n" +
			"\nexpected: " + expectedCanonical +
			"\nactual:   " + actualCanonical + "\n")
		t.Fail()
	}
}
