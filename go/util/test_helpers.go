package util

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
		t.Log("assertStringEquals:\n\nexpected:\n" + expected + "\nactual:\n" + actual + "\n")
		t.FailNow()
	}
}

func RemoveBlankLine(s string) string {
	var buffer strings.Builder
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		if len(line) > 0 {
			buffer.WriteString(line)
			buffer.WriteString("\n")
		}
	}
	return buffer.String()
}

func TrimEachLine(s string) string {
	var sb strings.Builder
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		trimmedLine := trim(line)
		sb.WriteString(trimmedLine)
		sb.WriteString("\n")
	}
	return sb.String()
}

func TrimByLine(s string) string {
	s = trim(s)
	lastLineBlank := false
	var buffer strings.Builder
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		trimmedLine := trim(line)
		lineLength := len(trimmedLine)
		if lineLength > 0 {
			buffer.WriteString(trimmedLine)
			buffer.WriteString("\n")
			lastLineBlank = false
			continue
		}
		if !lastLineBlank {
			buffer.WriteString("\n")
			lastLineBlank = true
		}
	}
	return buffer.String()
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
