package tests

import (
	"strings"
	"testing"

	"github.com/tmcphillips/blazegraph-util/sparql"
	"github.com/tmcphillips/blazegraph-util/util"
)

var sr = sparql.ResultSet{
	sparql.Head{[]string{"s", "o"}},
	sparql.Results{[]sparql.Binding{{
		"s": {"uri", "http://tmcphill.net/data#x"},
		"o": {"literal", "seven"},
	}, {
		"s": {"uri", "http://tmcphill.net/data#y"},
		"o": {"literal", "eight"},
	}}}}

func TestResult_Vars(t *testing.T) {
	util.StringEquals(t, strings.Join(sr.Vars(), ", "), "s, o")
}

func TestResult_Bindings(t *testing.T) {

	util.StringEquals(t, sr.Bindings()[0]["s"].Type, "uri")
	util.StringEquals(t, sr.Bindings()[0]["s"].Value, "http://tmcphill.net/data#x")
	util.StringEquals(t, sr.Bindings()[0]["o"].Type, "literal")
	util.StringEquals(t, sr.Bindings()[0]["o"].Value, "seven")

	util.StringEquals(t, sr.Bindings()[1]["s"].Type, "uri")
	util.StringEquals(t, sr.Bindings()[1]["s"].Value, "http://tmcphill.net/data#y")
	util.StringEquals(t, sr.Bindings()[1]["o"].Type, "literal")
	util.StringEquals(t, sr.Bindings()[1]["o"].Value, "eight")
}

func TestResult_BoundValues(t *testing.T) {
	util.StringEquals(t,
		strings.Join(sr.BoundValues(0), ", "),
		"http://tmcphill.net/data#x, seven")
	util.StringEquals(t,
		strings.Join(sr.BoundValues(1), ", "),
		"http://tmcphill.net/data#y, eight")
}
