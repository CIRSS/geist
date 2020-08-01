package sparqlrep

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func uppercase(s string) string {
	return strings.ToUpper(s)
}

func ExampleReportParser_Expand() {

	dp := DelimiterPair{
		Start: "<%",
		End:   "%>",
	}

	t := `{{with $result := <%
food
bar
%>}}{{$result}}{{end}}`

	rt := NewReportTemplate(dp, nil, t)
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

	dp := DelimiterPair{
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

	rt := NewReportTemplate(dp, funcs, t)
	err := rt.Expand(os.Stdout, nil)

	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// FOO
	// BAR
}
