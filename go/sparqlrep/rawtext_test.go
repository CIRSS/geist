package sparqlrep

import "testing"

func assertEqual(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Error("Actual string different from expected")
		t.Error("Expected : " + expected)
		t.Error("Actual   : " + actual)
		t.Fail()
	}
}

func TestEscapeNewlines_JustSpaces(t *testing.T) {
	assertEqual(t, EscapeNewlines(""), "")
	assertEqual(t, EscapeNewlines(" "), " ")
	assertEqual(t, EscapeNewlines("   "), "   ")
}

func TestEscapeNewlines_JustNewlines(t *testing.T) {
	assertEqual(t, EscapeNewlines("\n"), "\\n")
	assertEqual(t, EscapeNewlines("\n\n"), "\\n\\n")
	assertEqual(t, EscapeNewlines("\n\n\n"), "\\n\\n\\n")
}

func TestEscapeNewlines_SpacesAndNewlines(t *testing.T) {
	assertEqual(t, EscapeNewlines(" \n"), " \\n")
	assertEqual(t, EscapeNewlines("\n "), "\\n ")
	assertEqual(t, EscapeNewlines(" \n "), " \\n ")
	assertEqual(t, EscapeNewlines("\n \n"), "\\n \\n")
	assertEqual(t, EscapeNewlines(" \n\n"), " \\n\\n")
	assertEqual(t, EscapeNewlines("\n\n "), "\\n\\n ")
	assertEqual(t, EscapeNewlines(" \n \n "), " \\n \\n ")
	assertEqual(t, EscapeNewlines("\n\n\n"), "\\n\\n\\n")
	assertEqual(t, EscapeNewlines(" \n \n \n "), " \\n \\n \\n ")
}

func TestEscapeNewlines_TextAndNewlines(t *testing.T) {
	assertEqual(t, EscapeNewlines("foo\n"), "foo\\n")
	assertEqual(t, EscapeNewlines("\nbar"), "\\nbar")
	assertEqual(t, EscapeNewlines("foo\nbar"), "foo\\nbar")
	assertEqual(t, EscapeNewlines("foo\nbar\nbaz"), "foo\\nbar\\nbaz")
	assertEqual(t, EscapeNewlines("foo\n\nbar"), "foo\\n\\nbar")
	assertEqual(t, EscapeNewlines("\n\nbar"), "\\n\\nbar")
}

func TestEscapeNewlined_OneRawStringArgumentToFunc(t *testing.T) {
	assertEqual(t, EscapeNewlines("up `foo`"), "up `foo`")
	assertEqual(t, EscapeNewlines("up `foo bar`"), "up `foo bar`")
	assertEqual(t, EscapeNewlines("up `foo\nbar`"), "up `foo\\nbar`")
	assertEqual(t, EscapeNewlines("up `foo\nbar`"), "up `foo\\nbar`")
}

func TestEscapeRawText_AllOneRawString(t *testing.T) {
	assertEqual(t, EscapeRawText(GraveDelimiters, "``"), "\"\"")
	assertEqual(t, EscapeRawText(GraveDelimiters, "`foo`"), "\"foo\"")
	assertEqual(t, EscapeRawText(GraveDelimiters, "`foo bar`"), "\"foo bar\"")
	assertEqual(t, EscapeRawText(GraveDelimiters, "`foo\nbar`"), "\"foo\\nbar\"")
}

func TestEscapeRawText_OneRawStringArgumentToFunc(t *testing.T) {
	assertEqual(t, EscapeRawText(GraveDelimiters, "up ``"), "up \"\"")
	assertEqual(t, EscapeRawText(GraveDelimiters, "up `foo`"), "up \"foo\"")
	assertEqual(t, EscapeRawText(GraveDelimiters, "up `foo bar`"), "up \"foo bar\"")
	assertEqual(t, EscapeRawText(GraveDelimiters, "up `foo\nbar`"), "up \"foo\\nbar\"")
}

func TestEscapeRawText_TwoRawStrings(t *testing.T) {
	assertEqual(t, EscapeRawText(GraveDelimiters, "`` ``"), "\"\" \"\"")
	assertEqual(t, EscapeRawText(GraveDelimiters, "`foo` `bar`"), "\"foo\" \"bar\"")
	assertEqual(t, EscapeRawText(GraveDelimiters, "`foo\n` `bar`"), "\"foo\\n\" \"bar\"")
	assertEqual(t, EscapeRawText(GraveDelimiters, "`foo\n` baz `bar`"), "\"foo\\n\" baz \"bar\"")
}

func TestEscapeRawText_TwoRawStrings_WithNewlineBetween(t *testing.T) {
	assertEqual(t, EscapeRawText(GraveDelimiters, "`` \n ``"), "\"\" \n \"\"")
}

func TestEscapeRawText_EightDigits(t *testing.T) {
	assertEqual(t, EscapeRawText(GraveDelimiters, "01`34`67"), "01\"34\"67")
}

func TestEscapeRawText_TwoCharDelimiters_TenDigits(t *testing.T) {
	delimiters := DelimiterPair{"<%", "%>"}
	assertEqual(t, EscapeRawText(delimiters, "01<%45%>89"), "01\"45\"89")
}

func TestEscapeRawText_MultilineText_NoRawText(t *testing.T) {
	text := `
	foo
	  bar
	baz
	`
	assertEqual(t, EscapeRawText(GraveDelimiters, text), text)
}

func TestEscapeRawText_MultilineText_SingleLineRawText(t *testing.T) {
	delimiters := DelimiterPair{"<%", "%>"}
	text := `
foo
  <%bar%>
baz
`
	assertEqual(t, EscapeRawText(delimiters, text), `
foo
  "bar"
baz
`)
}

func TestEscapeRawText_MultilineText_MultilineRawText(t *testing.T) {
	delimiters := DelimiterPair{"<%", "%>"}
	text := `
foo
  <%bar
zing%>
baz
`
	assertEqual(t, EscapeRawText(delimiters, text), `
foo
  "bar\nzing"
baz
`)
}
