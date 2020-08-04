package tests

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func ExampleTemplate_NamedStructInstance() {
	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}
	tmpl, _ := template.New("test").Parse("{{.Count}} items are made of {{.Material}}\n")
	tmpl.Execute(os.Stdout, sweaters)
	// Output:
	// 17 items are made of wool
}

func ExampleTemplate_AnonymouStructInstance() {
	sweaters := struct {
		Material string
		Count    uint
	}{
		Material: "cotton",
		Count:    42,
	}
	tmpl, _ := template.New("test").Parse("{{.Count}} items are made of {{.Material}}\n")
	tmpl.Execute(os.Stdout, sweaters)
	// Output:
	// 42 items are made of cotton
}

func ExampleTemplate_Struct_MissingField_NoError() {
	sweaters := struct {
		Material string
		Count    uint
	}{
		Material: "cotton",
		Count:    42,
	}
	tmpl, _ := template.New("test").Parse("{{.Count}} items are made of {{.Materials}}\n")
	tmpl.Execute(os.Stdout, sweaters)
	// Output:
	// 42 items are made of
}

func ExampleTemplate_Struct_PrintPipelines() {
	sweaters := struct {
		Material string
		Count    uint
	}{
		Material: "cotton",
		Count:    42,
	}
	tmpl, _ := template.New("test").Parse(`{{print .Count}} items are made of {{print .Material}}`)
	tmpl.Execute(os.Stdout, sweaters)
	// Output:
	// 42 items are made of cotton
}

func ExampleTemplate_Struct_PrintfPipelines() {
	sweaters := struct {
		Material string
		Count    uint
	}{
		Material: "cotton",
		Count:    42,
	}
	tmpl, _ := template.New("test").Parse(`{{printf "%d" .Count}} items are made of {{printf "%s" .Material}}`)
	tmpl.Execute(os.Stdout, sweaters)
	// Output:
	// 42 items are made of cotton
}

func ExampleTemplate_Struct_PrintfPipelines_ArgumentsPiped() {
	sweaters := struct {
		Material string
		Count    uint
	}{
		Material: "cotton",
		Count:    42,
	}
	tmpl, _ := template.New("test").Parse(`{{.Count | printf "%d"}} items are made of {{.Material | printf "%s"}}`)
	tmpl.Execute(os.Stdout, sweaters)
	// Output:
	// 42 items are made of cotton
}

func ExampleTemplate_Struct_PrintfPipeline() {
	sweaters := struct {
		Material string
		Count    uint
	}{
		Material: "cotton",
		Count:    42,
	}
	tmpl, _ := template.New("test").Parse(`{{printf "%d items are made of %s" .Count .Material}}`)
	tmpl.Execute(os.Stdout, sweaters)
	// Output:
	// 42 items are made of cotton
}

func ExampleTemplate_Struct_PrintfPipeline_OneArgumentPiped() {
	sweaters := struct {
		Material string
		Count    uint
	}{
		Material: "cotton",
		Count:    42,
	}
	tmpl, _ := template.New("test").Parse(`{{.Material | printf "%d items are made of %s" .Count}}`)
	tmpl.Execute(os.Stdout, sweaters)
	// Output:
	// 42 items are made of cotton
}

func ExampleTemplate_Struct_PrintfPipeline_OneArgumentViaDot() {
	sweaters := struct {
		Material string
		Count    uint
	}{
		Material: "cotton",
		Count:    42,
	}
	tmpl, _ := template.New("test").Parse(`{{with .Count}}{{printf "%d" .}}{{end}} items are made of {{print .Material}}`)
	tmpl.Execute(os.Stdout, sweaters)
	// Output:
	// 42 items are made of cotton
}

func ExampleTemplate_Struct_PrintfPipeline_OneArgumentViaVariable() {
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
	tmpl.Execute(os.Stdout, sweaters)
	// Output:
	// 42 items are made of cotton
}

func increment(i int) int {
	return i + 1
}

func ExampleTemplate_Execute_PrintfPipeline_OneArgumentViaVariableSetByFunction() {

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
	tmpl.Execute(os.Stdout, sweaters)
	// Output:
	// 43 items are made of cotton
}

func ExampleTemplate_Execute_WrongFunctionArgumentType_Error() {

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
	err := tmpl.Execute(os.Stdout, sweaters)
	fmt.Println(err)
	// Output:
	// template: test:1:17: executing "test" at <.Count>: wrong type for value; expected int; got int64
}

func ExampleTemplate_Execute_MultilinePipeline() {

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
		Parse(`
{{with $c := inc .Count}}{{$d := $c}}{{$e := $d}}
{{printf "%d" $d}}{{end}} items are made of {{print .Material}}
`)
	err := tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 43 items are made of cotton
}

func upper(s string) string {
	return strings.ToUpper(s)
}

func ExampleTemplate_Execute_MultilineFunctionArgument() {

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
	err = tmpl.Execute(os.Stdout, "string")
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// FOOD
	// BAR
}

func ExampleTemplate_Execute_MultilineFunctionArgument_Two() {

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
	err = tmpl.Execute(os.Stdout, "string")
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// FOOD
	// BAR
}
