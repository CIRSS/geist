package sparqltests

import (
	"strings"
	"testing"

	"github.com/tmcphillips/blazegraph-util/assert"
	"github.com/tmcphillips/blazegraph-util/sparql"
)

var sr = sparql.Result{
	sparql.Head{[]string{"s", "o"}},
	sparql.Results{[]sparql.Binding{{
		"s": {"uri", "http://tmcphill.net/data#x"},
		"o": {"literal", "seven"},
	}, {
		"s": {"uri", "http://tmcphill.net/data#y"},
		"o": {"literal", "eight"},
	}}}}

func TestResult_Vars(t *testing.T) {
	assert.StringEquals(t, strings.Join(sr.Vars(), ", "), "s, o")
}

func TestResult_Bindings(t *testing.T) {

	assert.StringEquals(t, sr.Bindings()[0]["s"].Type, "uri")
	assert.StringEquals(t, sr.Bindings()[0]["s"].Value, "http://tmcphill.net/data#x")
	assert.StringEquals(t, sr.Bindings()[0]["o"].Type, "literal")
	assert.StringEquals(t, sr.Bindings()[0]["o"].Value, "seven")

	assert.StringEquals(t, sr.Bindings()[1]["s"].Type, "uri")
	assert.StringEquals(t, sr.Bindings()[1]["s"].Value, "http://tmcphill.net/data#y")
	assert.StringEquals(t, sr.Bindings()[1]["o"].Type, "literal")
	assert.StringEquals(t, sr.Bindings()[1]["o"].Value, "eight")
}

func TestResult_BoundValues(t *testing.T) {
	assert.StringEquals(t,
		strings.Join(sr.BoundValues(0), ", "),
		"http://tmcphill.net/data#x, seven")
	assert.StringEquals(t,
		strings.Join(sr.BoundValues(1), ", "),
		"http://tmcphill.net/data#y, eight")
}
