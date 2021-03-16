package tests

import (
	"fmt"
	"strings"
	"testing"
	"text/template"

	"github.com/cirss/go-cli/pkg/util"
)

var outputBuffer strings.Builder

func TestTextTemplate_NamedStructInstance(t *testing.T) {
	outputBuffer.Reset()
	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}
	tmpl, _ := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	tmpl.Execute(&outputBuffer, sweaters)
	util.StringEquals(t, outputBuffer.String(), "17 items are made of wool")
}

func TestTemplate_AnonymouStructInstance(t *testing.T) {
	outputBuffer.Reset()
	sweaters := struct {
		Material string
		Count    uint
	}{
		Material: "cotton",
		Count:    42,
	}
	tmpl, _ := template.New("test").Parse("{{.Count}} items are made of {{.Material}}\n")
	tmpl.Execute(&outputBuffer, sweaters)
	util.StringEquals(t, outputBuffer.String(), "42 items are made of cotton\n")
}

func TestTemplate_Struct_MissingField_NoError_TruncatedTemplateExpansion(t *testing.T) {
	outputBuffer.Reset()
	sweaters := struct {
		Material string
		Count    uint
	}{
		Material: "cotton",
		Count:    42,
	}
	tmpl, _ := template.New("test").Parse("{{.Count}} items are made of {{.Materials}} SOME EXTRA STUFF")
	tmpl.Execute(&outputBuffer, sweaters)
	util.StringEquals(t, outputBuffer.String(), "42 items are made of ")
}

func TestTemplate_Struct_PrintPipelines(t *testing.T) {
	outputBuffer.Reset()
	sweaters := struct {
		Material string
		Count    uint
	}{
		Material: "cotton",
		Count:    42,
	}
	tmpl, _ := template.New("test").Parse(`{{print .Count}} items are made of {{print .Material}}`)
	tmpl.Execute(&outputBuffer, sweaters)
	util.StringEquals(t, outputBuffer.String(), "42 items are made of cotton")
}

func TestTemplate_Struct_PrintfPipelines(t *testing.T) {
	outputBuffer.Reset()
	sweaters := struct {
		Material string
		Count    uint
	}{
		Material: "cotton",
		Count:    42,
	}
	tmpl, _ := template.New("test").Parse(`{{printf "%d" .Count}} items are made of {{printf "%s" .Material}}`)
	tmpl.Execute(&outputBuffer, sweaters)
	util.StringEquals(t, outputBuffer.String(), "42 items are made of cotton")
}

func TestTemplate_Struct_PrintfPipelines_ArgumentsPiped(t *testing.T) {
	outputBuffer.Reset()
	sweaters := struct {
		Material string
		Count    uint
	}{
		Material: "cotton",
		Count:    42,
	}
	tmpl, _ := template.New("test").Parse(`{{.Count | printf "%d"}} items are made of {{.Material | printf "%s"}}`)
	tmpl.Execute(&outputBuffer, sweaters)
	util.StringEquals(t, outputBuffer.String(), "42 items are made of cotton")
}

func TestTemplate_Struct_PrintfPipeline(t *testing.T) {
	outputBuffer.Reset()
	sweaters := struct {
		Material string
		Count    uint
	}{
		Material: "cotton",
		Count:    42,
	}
	tmpl, _ := template.New("test").Parse(`{{printf "%d items are made of %s" .Count .Material}}`)
	tmpl.Execute(&outputBuffer, sweaters)
	util.StringEquals(t, outputBuffer.String(), "42 items are made of cotton")
}

func TestTemplate_Struct_PrintfPipeline_OneArgumentPiped(t *testing.T) {
	outputBuffer.Reset()
	sweaters := struct {
		Material string
		Count    uint
	}{
		Material: "cotton",
		Count:    42,
	}
	tmpl, _ := template.New("test").Parse(`{{.Material | printf "%d items are made of %s" .Count}}`)
	tmpl.Execute(&outputBuffer, sweaters)
	util.StringEquals(t, outputBuffer.String(), "42 items are made of cotton")
}

func TestTemplate_Struct_PrintfPipeline_OneArgumentViaDot(t *testing.T) {
	outputBuffer.Reset()
	sweaters := struct {
		Material string
		Count    uint
	}{
		Material: "cotton",
		Count:    42,
	}
	tmpl, _ := template.New("test").Parse(`{{with .Count}}{{printf "%d" .}}{{end}} items are made of {{print .Material}}`)
	tmpl.Execute(&outputBuffer, sweaters)
	util.StringEquals(t, outputBuffer.String(), "42 items are made of cotton")
}

func TestTemplate_Struct_PrintfPipeline_OneArgumentViaVariable(t *testing.T) {
	outputBuffer.Reset()
	sweaters := struct {
		Material string
		Count    uint
	}{
		Material: "cotton",
		Count:    42,
	}
	tmpl, _ := template.
		New("test").
		Parse(`{{with $c := .Count}}{{printf "%d" $c}}{{end}} items are made of {{print .Material}}`)
	tmpl.Execute(&outputBuffer, sweaters)
	util.StringEquals(t, outputBuffer.String(), "42 items are made of cotton")
}

func increment(i int) int {
	return i + 1
}

func TestTemplate_Execute_PrintfPipeline_OneArgumentViaVariableSetByFunction(t *testing.T) {
	outputBuffer.Reset()

	sweaters := struct {
		Material string
		Count    int
	}{
		Material: "cotton",
		Count:    42,
	}

	funcMap := template.FuncMap{
		"inc": increment,
	}

	tmpl, _ := template.
		New("test").
		Funcs(funcMap).
		Parse(`{{with $c := inc .Count}}{{printf "%d" $c}}{{end}} items are made of {{print .Material}}`)
	tmpl.Execute(&outputBuffer, sweaters)
	util.StringEquals(t, outputBuffer.String(), "43 items are made of cotton")
}

func TestTemplate_Execute_WrongFunctionArgumentType_Error(t *testing.T) {
	outputBuffer.Reset()

	sweaters := struct {
		Material string
		Count    int64
	}{
		Material: "cotton",
		Count:    42,
	}

	funcMap := template.FuncMap{
		"inc": increment,
	}

	tmpl, _ := template.
		New("test").
		Funcs(funcMap).
		Parse(`{{with $c := inc .Count}}{{printf "%d" $c}}{{end}} items are made of {{print .Material}}`)
	err := tmpl.Execute(&outputBuffer, sweaters)
	util.StringEquals(t, err.Error(),
		`template: test:1:17: executing "test" at <.Count>: wrong type for value; expected int; got int64`)
}

func TestTemplate_Execute_MultilinePipeline(t *testing.T) {
	outputBuffer.Reset()

	sweaters := struct {
		Material string
		Count    int
	}{
		Material: "cotton",
		Count:    42,
	}

	funcMap := template.FuncMap{
		"inc": increment,
	}

	tmpl, _ := template.
		New("test").
		Funcs(funcMap).
		Parse(
			`{{with $c := inc .Count}}{{$d := $c}}{{$e := $d}}
			 {{printf "%d" $d}}{{end}} items are made of {{print .Material}}
		`)
	err := tmpl.Execute(&outputBuffer, sweaters)
	if err != nil {
		fmt.Println(err)
	}
	util.LineContentsEqual(t, outputBuffer.String(),
		`
		43 items are made of cotton
		`)
}

func upper(s string) string {
	return strings.ToUpper(s)
}

func TestTemplate_Execute_MultilineFunctionArgument(t *testing.T) {
	outputBuffer.Reset()

	funcMap := template.FuncMap{
		"up": upper,
	}

	multilineTemplate := `{{"
food
bar
" | up}}`
	multilineTemplate = strings.ReplaceAll(multilineTemplate, "\n", "\\n")
	//	fmt.Println(multilineTemplate)
	tmpl, err := template.
		New("test").
		Funcs(funcMap).
		Parse(multilineTemplate)
	err = tmpl.Execute(&outputBuffer, "string")
	if err != nil {
		fmt.Println(err)
	}
	util.LineContentsEqual(t, outputBuffer.String(),
		`
		FOOD
		BAR
		`)
}

func TestTemplate_Execute_MultilineFunctionArgument_Two(t *testing.T) {
	outputBuffer.Reset()

	funcMap := template.FuncMap{
		"up": upper,
	}

	multilineTemplate := `{{with $result := SPARQL<<
food
bar
>> }}{{$result}}{{end}}`
	multilineTemplate = strings.ReplaceAll(multilineTemplate, "SPARQL<<", "up \"")
	multilineTemplate = strings.ReplaceAll(multilineTemplate, ">>", "\"")
	multilineTemplate = strings.ReplaceAll(multilineTemplate, "\n", "\\n")
	// fmt.Println(multilineTemplate)
	tmpl, err := template.
		New("test").
		Funcs(funcMap).
		Parse(multilineTemplate)
	err = tmpl.Execute(&outputBuffer, "string")
	if err != nil {
		fmt.Println(err)
	}
	util.LineContentsEqual(t, outputBuffer.String(),
		`
		FOOD
		BAR
		`)
}
