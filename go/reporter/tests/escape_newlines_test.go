package tests

import (
	"testing"

	"github.com/cirss/geist/reporter"
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
	assertEqual(t, reporter.EscapeNewlines(""), "")
	assertEqual(t, reporter.EscapeNewlines(" "), " ")
	assertEqual(t, reporter.EscapeNewlines("   "), "   ")
}

func TestEscapeNewLines_SpacesAndNewlines(t *testing.T) {
	assertEqual(t, reporter.EscapeNewlines(" \n"), " \\n")
	assertEqual(t, reporter.EscapeNewlines("\n "), "\\n ")
	assertEqual(t, reporter.EscapeNewlines(" \n "), " \\n ")
	assertEqual(t, reporter.EscapeNewlines("\n \n"), "\\n \\n")
	assertEqual(t, reporter.EscapeNewlines(" \n\n"), " \\n\\n")
	assertEqual(t, reporter.EscapeNewlines("\n\n "), "\\n\\n ")
	assertEqual(t, reporter.EscapeNewlines(" \n \n "), " \\n \\n ")
	assertEqual(t, reporter.EscapeNewlines("\n\n\n"), "\\n\\n\\n")
	assertEqual(t, reporter.EscapeNewlines(" \n \n \n "), " \\n \\n \\n ")
}

func TestEscapeNewlines_JustNewlines(t *testing.T) {
	assertEqual(t, reporter.EscapeNewlines("\n"), "\\n")
	assertEqual(t, reporter.EscapeNewlines("\n\n"), "\\n\\n")
	assertEqual(t, reporter.EscapeNewlines("\n\n\n"), "\\n\\n\\n")
}

func TestEscapeNewlines_TextAndNewlines(t *testing.T) {
	assertEqual(t, reporter.EscapeNewlines("foo\n"), "foo\\n")
	assertEqual(t, reporter.EscapeNewlines("\nbar"), "\\nbar")
	assertEqual(t, reporter.EscapeNewlines("foo\nbar"), "foo\\nbar")
	assertEqual(t, reporter.EscapeNewlines("foo\nbar\nbaz"), "foo\\nbar\\nbaz")
	assertEqual(t, reporter.EscapeNewlines("foo\n\nbar"), "foo\\n\\nbar")
	assertEqual(t, reporter.EscapeNewlines("\n\nbar"), "\\n\\nbar")
}

func TestEscapeNewlined_OneRawStringArgumentToFunc(t *testing.T) {
	assertEqual(t, reporter.EscapeNewlines("up `foo`"), "up `foo`")
	assertEqual(t, reporter.EscapeNewlines("up `foo bar`"), "up `foo bar`")
	assertEqual(t, reporter.EscapeNewlines("up `foo\nbar`"), "up `foo\\nbar`")
	assertEqual(t, reporter.EscapeNewlines("up `foo\nbar`"), "up `foo\\nbar`")
}
