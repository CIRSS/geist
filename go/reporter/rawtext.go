package reporter

import (
	"errors"
	"regexp"
	"strings"
)

const (
	doubleQuote       = `"`
	escDoubleQuote    = `\"`
	newline           = "\n"
	escapedNewline    = `\n`
	escapedLineEnding = `[\t ]*\n`
)

// DelimiterPair defines the start and end delimiter for a report text region.
type DelimiterPair struct {
	Start string
	End   string
}

// EscapeDoubleQuotes substitutes an escaped double-quote character sequence
// (\") for each actual doublequote character in the argument and returns the
// updated string.
func EscapeDoubleQuotes(text string) string {
	return strings.ReplaceAll(text, doubleQuote, escDoubleQuote)
}

// EscapeNewlines substitutes an escaped newline character sequence (\\n) for
// for each actual newline character (\n) in the argument and returns the
// updated string.
func EscapeNewlines(text string) string {
	return strings.ReplaceAll(text, newline, escapedNewline)
}

func RemoveNewlines(text string) string {
	return strings.ReplaceAll(text, newline, "")
}

func RemoveEscapedLineEndings(text string) string {
	re := regexp.MustCompile(`[ \t]*\\[\\]?[ \t]*\n[ \t]*`)
	return re.ReplaceAllString(text, "")
}

func RestoreNewlines(text string) string {
	return strings.ReplaceAll(text, escapedNewline, newline)
}

// EscapeRawText finds substrings delimited by the given DelimiterPair
// and within each escapes newlines and replaces the starting and
// end delimiters with double quotes.
func EscapeRawText(dp DelimiterPair, text string) (escapedText string, err error) {

	for {
		rawTextStart := strings.Index(text, dp.Start)
		if rawTextStart == -1 {
			break
		}

		rawTextEnd := strings.Index(text[rawTextStart+1:], dp.End)
		if rawTextEnd == -1 {
			err = errors.New("Unmatched raw string delimiter")
			break
		}

		rawTextEnd += rawTextStart + len(dp.End) + 1
		rawText := text[rawTextStart:rawTextEnd]
		rawText = EscapeNewlines(rawText)
		rawText = EscapeDoubleQuotes(rawText)
		rawText = rawText[len(dp.Start) : len(rawText)-len(dp.End)]
		quotedRawText := doubleQuote + rawText + doubleQuote

		text = text[0:rawTextStart] + quotedRawText + text[rawTextEnd:]
	}

	return text, err
}
