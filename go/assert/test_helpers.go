package assert

import (
	"encoding/json"
	"strings"
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

func trim(s string) string {
	return strings.Trim(s, " \n\r\t")
}

func LineContentsEqual(t *testing.T, actual string, expect string) {
	actualContent := strings.Split(trim(actual), "\n")
	expectContent := strings.Split(trim(expect), "\n")
	lineCount := len(actualContent)
	if lineCount != len(expectContent) {
		StringEquals(t, actual, expect)
	}
	for i := 0; i < lineCount; i++ {
		StringEquals(t, trim(actualContent[i]), trim(expectContent[i]))
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
