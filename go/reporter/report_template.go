package reporter

import (
	"strings"
	"text/template"
)

var GraveDelimiters DelimiterPair
var JSPDelimiters DelimiterPair
var TripleSingleQuoteDelimiters DelimiterPair

func init() {
	GraveDelimiters = DelimiterPair{"`", "`"}
	JSPDelimiters = DelimiterPair{Start: "<%", End: "%>"}
	TripleSingleQuoteDelimiters = DelimiterPair{Start: "'''", End: "'''"}
}

// ReportTemplate is customizable to create reports from different kinds of templates.
type ReportTemplate struct {
	template *template.Template
}

// NewReportTemplate returns a ReportTemplate with the given customizations.
func NewReportTemplate(delimiters DelimiterPair, funcs template.FuncMap, text string) *ReportTemplate {
	text = EscapeRawText(delimiters, text)
	text = RemoveNewlines(text)
	template := template.New("main")
	if funcs != nil {
		template = template.Funcs(funcs)
	}
	template, _ = template.Parse(text)
	rt := new(ReportTemplate)
	rt.template = template
	return rt
}

func (rp *ReportTemplate) Expand(data interface{}) (result string, err error) {
	var buffer strings.Builder
	err = rp.template.Execute(&buffer, data)
	if err != nil {
		return
	}
	result = buffer.String()
	result = InsertNewlines(result)
	return
}
