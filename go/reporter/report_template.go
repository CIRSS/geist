package reporter

import (
	"errors"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/tmcphillips/blazegraph-util/tables"
	"github.com/tmcphillips/blazegraph-util/util"
)

type Properties struct {
	Delimiters DelimiterPair
	Funcs      template.FuncMap
	Prefixes   map[string]string
	Macros     map[string]*ReportTemplate
	Queries    map[string]string
}

var GraveDelimiters DelimiterPair
var JSPDelimiters DelimiterPair
var TripleSingleQuoteDelimiters DelimiterPair
var MacroDelimiters DelimiterPair
var DefaultDelimiters DelimiterPair

func init() {
	GraveDelimiters = DelimiterPair{"`", "`"}
	JSPDelimiters = DelimiterPair{Start: "<%", End: "%>"}
	TripleSingleQuoteDelimiters = DelimiterPair{Start: "'''", End: "'''"}
	MacroDelimiters = DelimiterPair{Start: "<?", End: "?>"}
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
	rt := new(ReportTemplate)
	rt.Name = name
	rt.Text = text
	rt.Properties.Funcs = template.FuncMap{}
	if delimiters != nil {
		rt.Properties.Delimiters = *delimiters
	} else {
		rt.Properties.Delimiters = DefaultDelimiters
	}
	rt.Properties.Prefixes = map[string]string{}
	rt.Properties.Macros = map[string]*ReportTemplate{}
	rt.Properties.Queries = map[string]string{}
	rt.addStandardFunctions()
	return rt
}

func (rp *ReportTemplate) Parse(removeNewlines bool) (err error) {

	text := util.TrimEachLine(rp.Text)
	text, err = EscapeRawText(rp.Properties.Delimiters, text)
	if err != nil {
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

	return
}

func (rp *ReportTemplate) AddFuncs(newFunctions template.FuncMap) {
	for key, value := range newFunctions {
		rp.Properties.Funcs[key] = value
	}
}

func (rp *ReportTemplate) Expand(data interface{}) (result string, err error) {
	var buffer strings.Builder
	err = rp.TextTemplate.Execute(&buffer, data)
	if err != nil {
		return
	}
	result = buffer.String()
	result = InsertNewlines(result)
	return
}

func (rp *ReportTemplate) ExpandSubreport(name string, text string, removeNewlines bool, data interface{}) (report string, err error) {
	reportTemplate := NewReportTemplate("include", string(text), nil)
	reportTemplate.Properties = rp.Properties
	reportTemplate.Parse(removeNewlines)
	if err != nil {
		return
	}
	report, err = reportTemplate.Expand(data)
	return
}

func (rp *ReportTemplate) addStandardFunctions() {

	funcs := template.FuncMap{
		"include": func(fileName string) (s string, err error) {
			text, err := ioutil.ReadFile(fileName)
			if err != nil {
				return
			}
			s, err = rp.ExpandSubreport(fileName, string(text), true, nil)
			return
		},
		"up": func(s string) string {
			return strings.ToUpper(s)
		},
		"macro": func(name string, body string) (s string, err error) {
			macroTemplate := NewReportTemplate(name, body, &MacroDelimiters)
			macroTemplate.AddFuncs(rp.Properties.Funcs)
			err = macroTemplate.Parse(true)
			if err != nil {
				return
			}
			rp.Properties.Macros[name] = macroTemplate
			return "", nil
		},
		"expand": func(name string, args ...interface{}) (result interface{}, err error) {
			macroTemplate := rp.Properties.Macros[name]
			var data interface{}
			if len(args) == 1 {
				data = args[0]
			}
			result, err = macroTemplate.Expand(data)
			return
		},
		"query": func(name string, body string) (s string, err error) {
			rp.Properties.Queries[name] = body
			return "", nil
		},
		"tabulate": func(rs tables.DataTable) (table string) {
			table = rs.FormattedTable(true)
			return
		},
		"rows": func(rs tables.DataTable) (rows [][]string) {
			rows = rs.Rows()
			return
		},
		"column": func(columnIndex int, rs *tables.DataTable) (column []string) {
			column = (*rs).Column(columnIndex)
			return
		},
		"vector": func(rs tables.DataTable) (vector []string, err error) {
			if rs.RowCount() == 1 {
				vector = rs.Row(0)
			} else if rs.ColumnCount() == 1 {
				vector = rs.Column(0)
			} else {
				err = errors.New("Result set is not a vector.")
			}
			return
		},
		"value": func(rs tables.DataTable) (value string, err error) {
			if rs.RowCount() != 1 || rs.ColumnCount() != 1 {
				err = errors.New("Result set is not a single value.")
			}
			value = rs.Column(0)[0]
			return
		},
		"join": func(elems []string, sep string) (js string) {
			js = strings.Join(elems, sep)
			return
		},
		"nl": func() (s string) {
			return escapedNewline
		},
	}

	rp.AddFuncs(funcs)

}
