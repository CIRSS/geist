package sparql

import (
	"testing"
)

// var sr SparqlResult

// func init() {

func TestSparqlResult(t *testing.T) {

	vars := []string{"s, o"}

	head := HeadT{vars}

	bindings := []BindingT{{
		"s": TypedValueT{Type: "uri", Value: "http://tmcphill.net/data#x"},
		"o": TypedValueT{Type: "literal", Value: "seven"},
	}, {
		"s": TypedValueT{Type: "uri", Value: "http://tmcphill.net/data#y"},
		"o": TypedValueT{Type: "literal", Value: "eight"},
	}}

	results := ResultsT{bindings}

	sr := SparqlResult{head, results}
	t.Log(sr)

}
