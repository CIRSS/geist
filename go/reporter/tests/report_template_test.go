package tests

import (
	"strings"
	"testing"
	"text/template"

	"github.com/tmcphillips/blazegraph-util/reporter"
	"github.com/tmcphillips/blazegraph-util/util"
)

func TestReportTemplate_AnonymouStructInstance(t *testing.T) {

	sweaters := struct {
		Material string
		Count    uint
	}{
		Material: "cotton",
		Count:    42,
	}

	rt := reporter.NewReportTemplate(reporter.JSPDelimiters, nil,
		`
		{{.Count}} items 
		are made of 
		{{.Material}}
		`)

	actual, _ := rt.Expand(sweaters)
	util.LineContentsEqual(t, actual,
		`
		42 items
		are made of
		cotton
	`)
}

func TestReportTemplate_MultilineVariableValue(t *testing.T) {

	rt := reporter.NewReportTemplate(reporter.JSPDelimiters, nil,
		`
		{{with $result := <%
			foo
			bar
		%>}}{{$result}}{{end}}
	`)

	actual, _ := rt.Expand(nil)
	util.LineContentsEqual(t, actual, `
		foo
		bar
	`)
}

func TestReportTemplate_MultilineFunctionArgument(t *testing.T) {

	funcs := template.FuncMap{
		"up": func(s string) string {
			return strings.ToUpper(s)
		},
	}

	rt := reporter.NewReportTemplate(reporter.JSPDelimiters, funcs,
		`{{with $result := up <%
				foo
				bar
		%>}}{{$result}}{{end}}
	`)

	actual, _ := rt.Expand(nil)
	util.LineContentsEqual(t, actual, `
		FOO
		BAR
	`)
}

func TestReportTemplate_RangeOverStringSlice(t *testing.T) {

	colors := []string{"red", "blue", "yellow"}

	rt := reporter.NewReportTemplate(reporter.JSPDelimiters, nil,
		`
		{{range .}} the color is {{.}}
		{{end}}
		`)

	actual, _ := rt.Expand(colors)
	util.LineContentsEqual(t, actual,
		`
		the color is red
		the color is blue
		the color is yellow
		`)
}
