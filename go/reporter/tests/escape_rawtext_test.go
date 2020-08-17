package tests

import (
	"testing"

	"github.com/tmcphillips/blazegraph-util/reporter"
)

func escapeRawText(dp reporter.DelimiterPair, text string) string {
	result, _ := reporter.EscapeRawText(dp, text)
	return result
}

func TestEscapeRawText_AllOneRawString(t *testing.T) {
	assertEqual(t, escapeRawText(reporter.GraveDelimiters, "``"), "\"\"")
	assertEqual(t, escapeRawText(reporter.GraveDelimiters, "`foo`"), "\"foo\"")
	assertEqual(t, escapeRawText(reporter.GraveDelimiters, "`foo bar`"), "\"foo bar\"")
	assertEqual(t, escapeRawText(reporter.GraveDelimiters, "`foo\nbar`"), "\"foo\\nbar\"")
}
func TestEscapeRawText_OneRawStringArgumentToFunc(t *testing.T) {
	assertEqual(t, escapeRawText(reporter.GraveDelimiters, "up ``"), "up \"\"")
	assertEqual(t, escapeRawText(reporter.GraveDelimiters, "up `foo`"), "up \"foo\"")
	assertEqual(t, escapeRawText(reporter.GraveDelimiters, "up `foo bar`"), "up \"foo bar\"")
	assertEqual(t, escapeRawText(reporter.GraveDelimiters, "up `foo\nbar`"), "up \"foo\\nbar\"")
}

func TestEscapeRawText_TwoRawStrings(t *testing.T) {
	assertEqual(t, escapeRawText(reporter.GraveDelimiters, "`` ``"), "\"\" \"\"")
	assertEqual(t, escapeRawText(reporter.GraveDelimiters, "`foo` `bar`"), "\"foo\" \"bar\"")
	assertEqual(t, escapeRawText(reporter.GraveDelimiters, "`foo\n` `bar`"), "\"foo\\n\" \"bar\"")
	assertEqual(t, escapeRawText(reporter.GraveDelimiters, "`foo\n` baz `bar`"), "\"foo\\n\" baz \"bar\"")
}

func TestEscapeRawText_TwoRawStrings_WithNewlineBetween(t *testing.T) {
	assertEqual(t, escapeRawText(reporter.GraveDelimiters, "`` \n ``"), "\"\" \n \"\"")
}

func TestEscapeRawText_EightDigits(t *testing.T) {
	assertEqual(t, escapeRawText(reporter.GraveDelimiters, "01`34`67"), "01\"34\"67")
}

func TestEscapeRawText_TwoCharDelimiters_TenDigits(t *testing.T) {
	delimiters := reporter.DelimiterPair{Start: "<%", End: "%>"}
	assertEqual(t, escapeRawText(delimiters, "01<%45%>89"), "01\"45\"89")
}

func TestEscapeRawText_MultilineText_NoRawText(t *testing.T) {
	text := `
	foo
	  bar
	baz
	`
	assertEqual(t, escapeRawText(reporter.GraveDelimiters, text), text)
}

func TestEscapeRawText_MultilineText_SingleLineRawText(t *testing.T) {
	delimiters := reporter.DelimiterPair{Start: "<%", End: "%>"}
	text := `
foo
  <%bar%>
baz
`
	assertEqual(t, escapeRawText(delimiters, text), `
foo
  "bar"
baz
`)
}

func TestEscapeRawText_MultilineText_MultilineRawText(t *testing.T) {
	delimiters := reporter.DelimiterPair{Start: "<%", End: "%>"}
	text := `
foo
  <%bar
zing%>
baz
`
	assertEqual(t, escapeRawText(delimiters, text), `
foo
  "bar\nzing"
baz
`)
}
