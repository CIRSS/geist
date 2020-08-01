package tests

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/tmcphillips/blazegraph-util/sparqlrep"
)

func uppercase(s string) string {
	return strings.ToUpper(s)
}

func ExampleReportParser_Expand() {

	dp := sparqlrep.DelimiterPair{
		Start: "<%",
		End:   "%>",
	}

	t := `{{with $result := <%
food
bar
%>}}{{$result}}{{end}}`

	rt := sparqlrep.NewReportTemplate(dp, nil, t)
	err := rt.Expand(os.Stdout, nil)

	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// food
	// bar
}

func up(s string) string {
	return strings.ToUpper(s)
}

func ExampleReportParser_Expand_MultilineFunctionArgument() {

	dp := sparqlrep.DelimiterPair{
		Start: "<%",
		End:   "%>",
	}

	funcs := template.FuncMap{
		"up": up,
	}

	t := `{{with $result := up <%
foo
bar
%>}}{{$result}}{{end}}`

	rt := sparqlrep.NewReportTemplate(dp, funcs, t)
	err := rt.Expand(os.Stdout, nil)

	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// FOO
	// BAR
}
