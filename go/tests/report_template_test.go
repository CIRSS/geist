package tests

import (
	"strings"
	"testing"
	"text/template"

	"github.com/cirss/geist"
	"github.com/cirss/geist/util"
)

func TestReportTemplate_AnonymousStructInstance(t *testing.T) {

	sweaters := struct {
		Material string
		Count    uint
	}{
		Material: "cotton",
		Count:    42,
	}

	rt := geist.NewTemplate(
		"main",
		`
		{{.Count}} items
		are made of
		{{.Material}}
		`, nil)

	rt.Parse()
	actual, _ := rt.Expand(sweaters)
	util.LineContentsEqual(t, actual,
		`
		42 items
		are made of
		cotton
	`)
}

func TestReportTemplate_AnonymousStructInstance_MissingBrace(t *testing.T) {
	rt := geist.NewTemplate("main", "{{.Foo}", nil)
	err := rt.Parse()
	util.LineContentsEqual(t, err.Error(),
		`template: main:1: unexpected "}" in operand`)
}

func TestReportTemplate_AnonymousStructInstance_MissingField(t *testing.T) {
	rt := geist.NewTemplate("main", "{{.Foo}}", nil)
	rt.Parse()
	_, err := rt.Expand(struct{ Bar string }{Bar: "baz"})
	util.LineContentsEqual(t, err.Error(),
		`template: main:1:2: executing "main" at <.Foo>: can't evaluate field Foo in type struct { Bar string }`)
}

func TestReportTemplate_AnonymousStructInstance_NilData(t *testing.T) {
	rt := geist.NewTemplate("main", "{{.Foo}}", nil)
	rt.Parse()
	actual, _ := rt.Expand(nil)
	util.LineContentsEqual(t, actual, `
		<no value>
	`)
}

func TestReportTemplate_MultilineVariableValue(t *testing.T) {

	rt := geist.NewTemplate(
		"main",
		`
		{{with $result := <%
			foo
			bar
		%>}}{{$result}}{{end}}
		`, &geist.JSPDelimiters)
	rt.Parse()
	actual, _ := rt.Expand(nil)
	util.LineContentsEqual(t, actual, `
		foo
		bar
	`)
}

func TestReportTemplate_MultilineVariableValue_MissingEnd(t *testing.T) {
	rt := geist.NewTemplate(
		"main",
		`
		{{with $result := <%
			foo
			bar
		%>}}{{$result}}
		`, &geist.JSPDelimiters)
	err := rt.Parse()
	util.LineContentsEqual(t, err.Error(),
		`template: main:2: unexpected EOF`)
}

func TestReportTemplate_MultilineVariableValue_WrongVariableName(t *testing.T) {
	rt := geist.NewTemplate(
		"main",
		`
		{{with $result := <%
			foo
			bar
		%>}}{{$wrongVariableName}}{{end}}
		`, &geist.JSPDelimiters)
	err := rt.Parse()
	util.LineContentsEqual(t, err.Error(),
		`template: main:1: undefined variable "$wrongVariableName"`)
}

func TestReportTemplate_UnmatchedRawStringDelimiter(t *testing.T) {
	rt := geist.NewTemplate(
		"main",
		`
		{{with $result := '''
			foo
			bar
		}}{{$result}}{{end}}
	`, nil)
	err := rt.Parse()
	util.LineContentsEqual(t, err.Error(),
		`Unmatched raw string delimiter`)
}

func TestReportTemplate_MultilineFunctionArgument(t *testing.T) {

	funcs := template.FuncMap{
		"up": func(s string) string {
			return strings.ToUpper(s)
		},
	}

	rt := geist.NewTemplate(
		"main",
		`{{with $result := up '''
				foo
				bar
		''' }}{{$result}}{{end}}
	`, nil)
	rt.AddFuncs(funcs)
	rt.Parse()
	actual, _ := rt.Expand(nil)
	util.LineContentsEqual(t, actual, `
		FOO
		BAR
	`)
}

func TestReportTemplate_RangeOverStringSlice(t *testing.T) {

	colors := []string{"red", "blue", "yellow"}

	rt := geist.NewTemplate(
		"main",
		`
		{{range .}} 			\
			the color is {{.}}
		{{end}}					\
		`, nil)
	rt.Parse()
	actual, _ := rt.Expand(colors)
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

	rt := geist.NewTemplate(
		"main",
		`
		Name   | City      | Phone
		-------|-----------|--------------
		{{range .}}{{index . 0}}    | {{index . 1}} | {{index . 2}}
		{{end}}
	`, nil)
	rt.Parse()
	actual, _ := rt.Expand(contacts)
	util.LineContentsEqual(t, actual,
		`
        Name   | City      | Phone
        -------|-----------|--------------
        Tim    | Oakland   | 530-219-4754
        Bob    | Concord   | 510-320-9943
        Joe    | San Diego | 213-101-9313
		`)
}

func TestReportTemplate_TableOfValues_IndexOutOfRange(t *testing.T) {

	contacts := [][]string{
		{"Tim", "Oakland  ", "530-219-4754"},
		{"Bob", "Concord  ", "510-320-9943"},
		{"Joe", "San Diego", "213-101-9313"},
	}

	rt := geist.NewTemplate(
		"main",
		`
		Name   | City      | Phone
		-------|-----------|--------------
		{{range .}}{{index . 0}}    | {{index . 1}} | {{index . 3}}
		{{end}}
	`, nil)
	rt.Parse()
	_, err := rt.Expand(contacts)
	util.LineContentsEqual(t, err.Error(),
		`template: main:3:48: executing "main" at <index . 3>: error calling index: reflect: slice index out of range`)
}
