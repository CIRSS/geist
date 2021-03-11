package tests

import (
	"testing"

	"github.com/cirss/geist/go/geist"
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
	assertEqual(t, geist.EscapeNewlines(""), "")
	assertEqual(t, geist.EscapeNewlines(" "), " ")
	assertEqual(t, geist.EscapeNewlines("   "), "   ")
}

func TestEscapeNewLines_SpacesAndNewlines(t *testing.T) {
	assertEqual(t, geist.EscapeNewlines(" \n"), " \\n")
	assertEqual(t, geist.EscapeNewlines("\n "), "\\n ")
	assertEqual(t, geist.EscapeNewlines(" \n "), " \\n ")
	assertEqual(t, geist.EscapeNewlines("\n \n"), "\\n \\n")
	assertEqual(t, geist.EscapeNewlines(" \n\n"), " \\n\\n")
	assertEqual(t, geist.EscapeNewlines("\n\n "), "\\n\\n ")
	assertEqual(t, geist.EscapeNewlines(" \n \n "), " \\n \\n ")
	assertEqual(t, geist.EscapeNewlines("\n\n\n"), "\\n\\n\\n")
	assertEqual(t, geist.EscapeNewlines(" \n \n \n "), " \\n \\n \\n ")
}

func TestEscapeNewlines_JustNewlines(t *testing.T) {
	assertEqual(t, geist.EscapeNewlines("\n"), "\\n")
	assertEqual(t, geist.EscapeNewlines("\n\n"), "\\n\\n")
	assertEqual(t, geist.EscapeNewlines("\n\n\n"), "\\n\\n\\n")
}

func TestEscapeNewlines_TextAndNewlines(t *testing.T) {
	assertEqual(t, geist.EscapeNewlines("foo\n"), "foo\\n")
	assertEqual(t, geist.EscapeNewlines("\nbar"), "\\nbar")
	assertEqual(t, geist.EscapeNewlines("foo\nbar"), "foo\\nbar")
	assertEqual(t, geist.EscapeNewlines("foo\nbar\nbaz"), "foo\\nbar\\nbaz")
	assertEqual(t, geist.EscapeNewlines("foo\n\nbar"), "foo\\n\\nbar")
	assertEqual(t, geist.EscapeNewlines("\n\nbar"), "\\n\\nbar")
}

func TestEscapeNewlined_OneRawStringArgumentToFunc(t *testing.T) {
	assertEqual(t, geist.EscapeNewlines("up `foo`"), "up `foo`")
	assertEqual(t, geist.EscapeNewlines("up `foo bar`"), "up `foo bar`")
	assertEqual(t, geist.EscapeNewlines("up `foo\nbar`"), "up `foo\\nbar`")
	assertEqual(t, geist.EscapeNewlines("up `foo\nbar`"), "up `foo\\nbar`")
}
