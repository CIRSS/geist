package util

import (
	"testing"
)

func TestTrimEachLine_EmptyString(t *testing.T) {
	s := ""
	StringEquals(t, TrimEachLine(s), "")
}

func TestTrimEachLine_Spaces(t *testing.T) {
	s := "      "
	StringEquals(t, TrimEachLine(s), "")
}

func TestTrimEachLine_SpacesAndFinalNewline(t *testing.T) {
	s := "      \n"
	StringEquals(t, TrimEachLine(s), "\n")
}

func TestTrimEachLine_ThreeNewlines(t *testing.T) {
	s := "\n\n\n"
	StringEquals(t, TrimEachLine(s), "\n\n\n")
}

func TestTrimEachLine_ThreeLinesOfSpacesWithFinalNewline(t *testing.T) {
	s := "    \n  \n   \n"
	StringEquals(t, TrimEachLine(s), "\n\n\n")
}

func TestTrimEachLine_ThreeLinesOfSpaces_NoFinalNewline(t *testing.T) {
	s := "    \n  \n   "
	StringEquals(t, TrimEachLine(s), "\n\n")
}

func TestTrimEachLine_Letters(t *testing.T) {
	s := "abcdefg"
	StringEquals(t, TrimEachLine(s), "abcdefg")
}

func TestTrimEachLine_LettersAndFinalNewline(t *testing.T) {
	s := "abcdefg\n"
	StringEquals(t, TrimEachLine(s), "abcdefg\n")
}

func TestTrimEachLine_ThreeLinesOfLettersWithFinalNewline(t *testing.T) {
	s := "abcdefg\nhijklmnop\nqrstuv\n"
	StringEquals(t, TrimEachLine(s), "abcdefg\nhijklmnop\nqrstuv\n")
}

func TestTrimEachLine_ThreeLinesOfLetters_NoFinalNewline(t *testing.T) {
	s := "abcdefg\nhijklmnop\nqrstuv"
	StringEquals(t, TrimEachLine(s), "abcdefg\nhijklmnop\nqrstuv")
}
