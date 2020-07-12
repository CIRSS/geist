package sparql

import (
	"strings"
	"testing"

	tu "github.com/tmcphillips/blazegraph-util/testutil"
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
	tu.AssertStringEquals(t, strings.Join(sr.Vars(), ", "), "s, o")
}

func TestSparqlResult_Bindings(t *testing.T) {

	tu.AssertStringEquals(t, sr.Bindings()[0]["s"].Type, "uri")
	tu.AssertStringEquals(t, sr.Bindings()[0]["s"].Value, "http://tmcphill.net/data#x")
	tu.AssertStringEquals(t, sr.Bindings()[0]["o"].Type, "literal")
	tu.AssertStringEquals(t, sr.Bindings()[0]["o"].Value, "seven")

	tu.AssertStringEquals(t, sr.Bindings()[1]["s"].Type, "uri")
	tu.AssertStringEquals(t, sr.Bindings()[1]["s"].Value, "http://tmcphill.net/data#y")
	tu.AssertStringEquals(t, sr.Bindings()[1]["o"].Type, "literal")
	tu.AssertStringEquals(t, sr.Bindings()[1]["o"].Value, "eight")
}
