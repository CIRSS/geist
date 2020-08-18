package reporter

import (
	"strings"
	"text/template"

	"github.com/tmcphillips/blazegraph-util/util"
)

type Properties struct {
	Delimiters DelimiterPair
	Funcs      template.FuncMap
	Prefixes   map[string]string
	Macros     map[string]string
}

var GraveDelimiters DelimiterPair
var JSPDelimiters DelimiterPair
var TripleSingleQuoteDelimiters DelimiterPair
var DefaultDelimiters DelimiterPair

func init() {
	GraveDelimiters = DelimiterPair{"`", "`"}
	JSPDelimiters = DelimiterPair{Start: "<%", End: "%>"}
	TripleSingleQuoteDelimiters = DelimiterPair{Start: "'''", End: "'''"}
	DefaultDelimiters = TripleSingleQuoteDelimiters
}

// ReportTemplate is customizable to create reports from different kinds of templates.
type ReportTemplate struct {
	Name         string
	Text         string
	TextTemplate *template.Template
	Properties   Properties
}

// NewReportTemplate returns a ReportTemplate with the given customizations.
func NewReportTemplate(name string, text string, delimiters *DelimiterPair) *ReportTemplate {
	// print(text)
	rt := new(ReportTemplate)
	rt.Name = name
	rt.Text = text
	if delimiters != nil {
		rt.Properties.Delimiters = *delimiters
	} else {
		rt.Properties.Delimiters = DefaultDelimiters
	}
	rt.Properties.Prefixes = map[string]string{}
	rt.Properties.Macros = map[string]string{}
	return rt
}

func (rp *ReportTemplate) Parse(removeNewlines bool) (re *ReportError) {
	var err error

	text := util.TrimEachLine(rp.Text)
	text, err = EscapeRawText(rp.Properties.Delimiters, text)
	if err != nil {
		re = NewReportError(err, &text)
		return
	}

	if removeNewlines {
		text = RemoveNewlines(text)
	}

	rp.TextTemplate = template.New(rp.Name)
	if rp.Properties.Funcs != nil {
		rp.TextTemplate.Funcs(rp.Properties.Funcs)
	}

	_, err = rp.TextTemplate.Parse(text)
	if err != nil {
		re = NewReportError(err, &text)
	}

	return
}

func (rp *ReportTemplate) SetFuncs(funcs template.FuncMap) {
	rp.Properties.Funcs = funcs
}

type ReportError struct {
	TextTemplateError error
	Template          *string
}

func NewReportError(e error, t *string) *ReportError {
	return &ReportError{TextTemplateError: e, Template: t}
}

func (re *ReportError) Error() string {
	return re.TextTemplateError.Error()
}

func (rp *ReportTemplate) Expand(data interface{}) (result string, re *ReportError) {
	var buffer strings.Builder
	err := rp.TextTemplate.Execute(&buffer, data)
	if err != nil {
		re = NewReportError(err, &rp.Text)
		return
	}
	result = buffer.String()
	result = InsertNewlines(result)
	return
}
