package sparql

import (
	"strings"
	"testing"
)

// var sr SparqlResult

// func init() {

func TestSparqlResult(t *testing.T) {

	sr := SparqlResult{Head: HeadT{Vars: []string{"s, o"}}, Results: ResultsT{[]BindingT{{
		"s": {Type: "uri", Value: "http://tmcphill.net/data#x"},
		"o": {Type: "literal", Value: "seven"},
	}, {
		"s": {Type: "uri", Value: "http://tmcphill.net/data#y"},
		"o": {Type: "literal", Value: "eight"},
	}}}}

	AssertStringEquals(t, strings.Join(sr.Vars(), ", "), "s, o")

	AssertStringEquals(t, sr.Bindings()[0]["s"].Type, "uri")
	AssertStringEquals(t, sr.Bindings()[0]["s"].Value, "http://tmcphill.net/data#x")
	AssertStringEquals(t, sr.Bindings()[0]["o"].Type, "literal")
	AssertStringEquals(t, sr.Bindings()[0]["o"].Value, "seven")

	AssertStringEquals(t, sr.Bindings()[1]["s"].Type, "uri")
	AssertStringEquals(t, sr.Bindings()[1]["s"].Value, "http://tmcphill.net/data#y")
	AssertStringEquals(t, sr.Bindings()[1]["o"].Type, "literal")
	AssertStringEquals(t, sr.Bindings()[1]["o"].Value, "eight")
}
