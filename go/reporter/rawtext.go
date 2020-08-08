package reporter

import "strings"

const (
	doubleQuote    = `"`
	escDoubleQuote = `\"`
	newline        = "\n"
	escapedNewline = `\n`
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

// EscapeRawText finds substrings delimited by the given DelimiterPair
// and within each escapes newlines and replaces the starting and
// end delimiters with double quotes.
func EscapeRawText(dp DelimiterPair, text string) string {

	for {
		var rawTextStart, rawTextEnd int
		if rawTextStart = strings.Index(text, dp.Start); rawTextStart == -1 {
			break
		}
		if rawTextEnd = strings.Index(text[rawTextStart+1:], dp.End); rawTextEnd == -1 {
			break
		}
		rawTextEnd += rawTextStart + len(dp.End) + 1
		rawText := text[rawTextStart:rawTextEnd]
		rawText = EscapeNewlines(rawText)
		rawText = EscapeDoubleQuotes(rawText)
		rawText = rawText[len(dp.Start) : len(rawText)-len(dp.End)]

		text = text[0:rawTextStart] + doubleQuote + rawText + doubleQuote + text[rawTextEnd:]
	}

	return text
}
