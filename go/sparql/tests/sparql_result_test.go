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

func TestResultSet_Vars(t *testing.T) {
	util.StringEquals(t, strings.Join(sr.Variables(), ", "), "s, o")
}

func TestResultSet_Bindings(t *testing.T) {

	util.StringEquals(t, sr.Bindings()[0]["s"].Type, "uri")
	util.StringEquals(t, sr.Bindings()[0]["s"].Value, "http://tmcphill.net/data#x")
	util.StringEquals(t, sr.Bindings()[0]["o"].Type, "literal")
	util.StringEquals(t, sr.Bindings()[0]["o"].Value, "seven")

	util.StringEquals(t, sr.Bindings()[1]["s"].Type, "uri")
	util.StringEquals(t, sr.Bindings()[1]["s"].Value, "http://tmcphill.net/data#y")
	util.StringEquals(t, sr.Bindings()[1]["o"].Type, "literal")
	util.StringEquals(t, sr.Bindings()[1]["o"].Value, "eight")
}
func TestResultSet_Row(t *testing.T) {

	util.StringEquals(t,
		strings.Join(sr.Row(0), ", "),
		"http://tmcphill.net/data#x, seven")

	util.StringEquals(t,
		strings.Join(sr.Row(1), ", "),
		"http://tmcphill.net/data#y, eight")
}

func TestResultSet_Rows(t *testing.T) {

	rows := sr.Rows()

	util.StringEquals(t,
		strings.Join(rows[0], ", "),
		"http://tmcphill.net/data#x, seven")

	util.StringEquals(t,
		strings.Join(rows[1], ", "),
		"http://tmcphill.net/data#y, eight")
}

func TestResultSet_Column(t *testing.T) {

	util.StringEquals(t,
		strings.Join(sr.Column(0), ", "),
		"http://tmcphill.net/data#x, http://tmcphill.net/data#y")

	util.StringEquals(t,
		strings.Join(sr.Column(1), ", "),
		"seven, eight")
}
