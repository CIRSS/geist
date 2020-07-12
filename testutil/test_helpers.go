package testutil

import (
	"encoding/json"
	"testing"

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

func AssertStringEquals(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Log("assertStringEquals:\n\nexpected: " + expected + "\nactual:   " + actual + "\n")
		t.Fail()
	}
}

func AssertJSONEquals(t *testing.T, actual interface{}, expected string) {
	actualCanonical, _ := CanonicalJSON(actual)
	expectedCanonical, _ := CanonicalJSONFromString(expected)
	if actualCanonical != expectedCanonical {
		t.Log("assertEquivalentJSON:\n" +
			"\nexpected: " + expectedCanonical +
			"\nactual:   " + actualCanonical + "\n")
		t.Fail()
	}
}
