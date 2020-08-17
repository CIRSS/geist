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

	rt := reporter.NewReportTemplate(
		`
		{{.Count}} items \n
		are made of	     \n
		{{.Material}}    \n
		`)

	actual, _ := rt.Expand(sweaters, true)
	util.LineContentsEqual(t, actual,
		`
		42 items
		are made of
		cotton
	`)
}

func TestReportTemplate_AnonymousStructInstance_MissingBrace(t *testing.T) {
	rt := reporter.NewReportTemplate("{{.Foo}")
	_, err := rt.Expand(nil, true)
	util.LineContentsEqual(t, err.Error(),
		`template: main:1: unexpected "}" in operand`)
}

func TestReportTemplate_AnonymousStructInstance_MissingField(t *testing.T) {
	rt := reporter.NewReportTemplate("{{.Foo}}")
	_, err := rt.Expand(struct{ Bar string }{Bar: "baz"}, true)
	util.LineContentsEqual(t, err.Error(),
		`template: main:1:2: executing "main" at <.Foo>: can't evaluate field Foo in type struct { Bar string }`)
}

func TestReportTemplate_AnonymousStructInstance_NilData(t *testing.T) {
	rt := reporter.NewReportTemplate("{{.Foo}}")
	actual, _ := rt.Expand(nil, true)
	util.LineContentsEqual(t, actual, `
		<no value>
	`)
}

func TestReportTemplate_MultilineVariableValue(t *testing.T) {

	rt := reporter.NewReportTemplate(
		`
		{{with $result := <%
			foo
			bar
		%>}}{{$result}}{{end}}
	`)
	rt.SetDelimiters(reporter.JSPDelimiters)

	actual, _ := rt.Expand(nil, true)
	util.LineContentsEqual(t, actual, `
		foo
		bar
	`)
}

func TestReportTemplate_UnmatchedRawStringDelimiter(t *testing.T) {
	rt := reporter.NewReportTemplate(
		`
		{{with $result := '''
			foo
			bar
		}}{{$result}}{{end}}
	`)
	_, err := rt.Expand(nil, true)

	util.LineContentsEqual(t, err.Error(),
		`ReportTemplate: Unmatched raw string delimiter`)
}

func TestReportTemplate_MultilineFunctionArgument(t *testing.T) {

	funcs := template.FuncMap{
		"up": func(s string) string {
			return strings.ToUpper(s)
		},
	}

	rt := reporter.NewReportTemplate(
		`{{with $result := up '''
				foo
				bar
		''' }}{{$result}}{{end}}
	`)
	rt.SetFuncs(funcs)

	actual, _ := rt.Expand(nil, true)
	util.LineContentsEqual(t, actual, `
		FOO
		BAR
	`)
}

func TestReportTemplate_RangeOverStringSlice(t *testing.T) {

	colors := []string{"red", "blue", "yellow"}

	rt := reporter.NewReportTemplate(
		`
		{{range .}} 
			the color is {{.}} \n
		{{end}}
		`)

	actual, _ := rt.Expand(colors, true)
	util.LineContentsEqual(t, actual,
		`
		the color is red
		the color is blue
		the color is yellow
		`)
}

func TestReportTemplate_TableOfValues(t *testing.T) {

	contacts := [][]string{
		{"Tim", "Oakland  ", "530-219-4754"},
		{"Bob", "Concord  ", "510-320-9943"},
		{"Joe", "San Diego", "213-101-9313"},
	}

	rt := reporter.NewReportTemplate(
		`
		Name   | City      | Phone									\n
		-------|-----------|--------------							\n
		{{range .}}{{index . 0}}    | {{index . 1}} | {{index . 2}}	\n
		{{end}}
	`)

	actual, _ := rt.Expand(contacts, true)
	util.LineContentsEqual(t, actual,
		`
        Name   | City      | Phone
        -------|-----------|--------------
        Tim    | Oakland   | 530-219-4754
        Bob    | Concord   | 510-320-9943
        Joe    | San Diego | 213-101-9313
		`)
}
