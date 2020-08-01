package tests

import (
	"testing"

	"github.com/tmcphillips/blazegraph-util/sparqlrep"
)

func TestEscapeRawText_AllOneRawString(t *testing.T) {
	assertEqual(t, sparqlrep.EscapeRawText(sparqlrep.GraveDelimiters, "``"), "\"\"")
	assertEqual(t, sparqlrep.EscapeRawText(sparqlrep.GraveDelimiters, "`foo`"), "\"foo\"")
	assertEqual(t, sparqlrep.EscapeRawText(sparqlrep.GraveDelimiters, "`foo bar`"), "\"foo bar\"")
	assertEqual(t, sparqlrep.EscapeRawText(sparqlrep.GraveDelimiters, "`foo\nbar`"), "\"foo\\nbar\"")
}
func TestEscapeRawText_OneRawStringArgumentToFunc(t *testing.T) {
	assertEqual(t, sparqlrep.EscapeRawText(sparqlrep.GraveDelimiters, "up ``"), "up \"\"")
	assertEqual(t, sparqlrep.EscapeRawText(sparqlrep.GraveDelimiters, "up `foo`"), "up \"foo\"")
	assertEqual(t, sparqlrep.EscapeRawText(sparqlrep.GraveDelimiters, "up `foo bar`"), "up \"foo bar\"")
	assertEqual(t, sparqlrep.EscapeRawText(sparqlrep.GraveDelimiters, "up `foo\nbar`"), "up \"foo\\nbar\"")
}

func TestEscapeRawText_TwoRawStrings(t *testing.T) {
	assertEqual(t, sparqlrep.EscapeRawText(sparqlrep.GraveDelimiters, "`` ``"), "\"\" \"\"")
	assertEqual(t, sparqlrep.EscapeRawText(sparqlrep.GraveDelimiters, "`foo` `bar`"), "\"foo\" \"bar\"")
	assertEqual(t, sparqlrep.EscapeRawText(sparqlrep.GraveDelimiters, "`foo\n` `bar`"), "\"foo\\n\" \"bar\"")
	assertEqual(t, sparqlrep.EscapeRawText(sparqlrep.GraveDelimiters, "`foo\n` baz `bar`"), "\"foo\\n\" baz \"bar\"")
}

func TestEscapeRawText_TwoRawStrings_WithNewlineBetween(t *testing.T) {
	assertEqual(t, sparqlrep.EscapeRawText(sparqlrep.GraveDelimiters, "`` \n ``"), "\"\" \n \"\"")
}

func TestEscapeRawText_EightDigits(t *testing.T) {
	assertEqual(t, sparqlrep.EscapeRawText(sparqlrep.GraveDelimiters, "01`34`67"), "01\"34\"67")
}

func TestEscapeRawText_TwoCharDelimiters_TenDigits(t *testing.T) {
	delimiters := sparqlrep.DelimiterPair{"<%", "%>"}
	assertEqual(t, sparqlrep.EscapeRawText(delimiters, "01<%45%>89"), "01\"45\"89")
}

func TestEscapeRawText_MultilineText_NoRawText(t *testing.T) {
	text := `
	foo
	  bar
	baz
	`
	assertEqual(t, sparqlrep.EscapeRawText(sparqlrep.GraveDelimiters, text), text)
}

func TestEscapeRawText_MultilineText_SingleLineRawText(t *testing.T) {
	delimiters := sparqlrep.DelimiterPair{"<%", "%>"}
	text := `
foo
  <%bar%>
baz
`
	assertEqual(t, sparqlrep.EscapeRawText(delimiters, text), `
foo
  "bar"
baz
`)
}

func TestEscapeRawText_MultilineText_MultilineRawText(t *testing.T) {
	delimiters := sparqlrep.DelimiterPair{"<%", "%>"}
	text := `
foo
  <%bar
zing%>
baz
`
	assertEqual(t, sparqlrep.EscapeRawText(delimiters, text), `
foo
  "bar\nzing"
baz
`)
}
