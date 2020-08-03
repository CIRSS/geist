package tests

import (
	"testing"

	"github.com/tmcphillips/blazegraph-util/sparqlrep"
)

func assertEqual(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Error("Actual string different from expected")
		t.Error("Expected : " + expected)
		t.Error("Actual   : " + actual)
		t.Fail()
	}
}

func TestEscapeNewlines_JustSpaces(t *testing.T) {
	assertEqual(t, sparqlrep.EscapeNewlines(""), "")
	assertEqual(t, sparqlrep.EscapeNewlines(" "), " ")
	assertEqual(t, sparqlrep.EscapeNewlines("   "), "   ")
}

func TestEscapeNewLines_SpacesAndNewlines(t *testing.T) {
	assertEqual(t, sparqlrep.EscapeNewlines(" \n"), " \\n")
	assertEqual(t, sparqlrep.EscapeNewlines("\n "), "\\n ")
	assertEqual(t, sparqlrep.EscapeNewlines(" \n "), " \\n ")
	assertEqual(t, sparqlrep.EscapeNewlines("\n \n"), "\\n \\n")
	assertEqual(t, sparqlrep.EscapeNewlines(" \n\n"), " \\n\\n")
	assertEqual(t, sparqlrep.EscapeNewlines("\n\n "), "\\n\\n ")
	assertEqual(t, sparqlrep.EscapeNewlines(" \n \n "), " \\n \\n ")
	assertEqual(t, sparqlrep.EscapeNewlines("\n\n\n"), "\\n\\n\\n")
	assertEqual(t, sparqlrep.EscapeNewlines(" \n \n \n "), " \\n \\n \\n ")
}

func TestEscapeNewlines_JustNewlines(t *testing.T) {
	assertEqual(t, sparqlrep.EscapeNewlines("\n"), "\\n")
	assertEqual(t, sparqlrep.EscapeNewlines("\n\n"), "\\n\\n")
	assertEqual(t, sparqlrep.EscapeNewlines("\n\n\n"), "\\n\\n\\n")
}

func TestEscapeNewlines_TextAndNewlines(t *testing.T) {
	assertEqual(t, sparqlrep.EscapeNewlines("foo\n"), "foo\\n")
	assertEqual(t, sparqlrep.EscapeNewlines("\nbar"), "\\nbar")
	assertEqual(t, sparqlrep.EscapeNewlines("foo\nbar"), "foo\\nbar")
	assertEqual(t, sparqlrep.EscapeNewlines("foo\nbar\nbaz"), "foo\\nbar\\nbaz")
	assertEqual(t, sparqlrep.EscapeNewlines("foo\n\nbar"), "foo\\n\\nbar")
	assertEqual(t, sparqlrep.EscapeNewlines("\n\nbar"), "\\n\\nbar")
}

func TestEscapeNewlined_OneRawStringArgumentToFunc(t *testing.T) {
	assertEqual(t, sparqlrep.EscapeNewlines("up `foo`"), "up `foo`")
	assertEqual(t, sparqlrep.EscapeNewlines("up `foo bar`"), "up `foo bar`")
	assertEqual(t, sparqlrep.EscapeNewlines("up `foo\nbar`"), "up `foo\\nbar`")
	assertEqual(t, sparqlrep.EscapeNewlines("up `foo\nbar`"), "up `foo\\nbar`")
}
