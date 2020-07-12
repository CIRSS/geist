package sparql

import (
	"strings"
	"testing"
)

var sr = SparqlResult{
	Head{[]string{"s, o"}},
	Results{[]Binding{{
		"s": {"uri", "http://tmcphill.net/data#x"},
		"o": {"literal", "seven"},
	}, {
		"s": {"uri", "http://tmcphill.net/data#y"},
		"o": {"literal", "eight"},
	}}}}

func TestSparqlResult_Vars(t *testing.T) {
	AssertStringEquals(t, strings.Join(sr.Vars(), ", "), "s, o")
}

func TestSparqlResult_Bindings(t *testing.T) {

	AssertStringEquals(t, sr.Bindings()[0]["s"].Type, "uri")
	AssertStringEquals(t, sr.Bindings()[0]["s"].Value, "http://tmcphill.net/data#x")
	AssertStringEquals(t, sr.Bindings()[0]["o"].Type, "literal")
	AssertStringEquals(t, sr.Bindings()[0]["o"].Value, "seven")

	AssertStringEquals(t, sr.Bindings()[1]["s"].Type, "uri")
	AssertStringEquals(t, sr.Bindings()[1]["s"].Value, "http://tmcphill.net/data#y")
	AssertStringEquals(t, sr.Bindings()[1]["o"].Type, "literal")
	AssertStringEquals(t, sr.Bindings()[1]["o"].Value, "eight")
}
