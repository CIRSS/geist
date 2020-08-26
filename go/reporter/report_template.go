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

func (rt *ReportTemplate) CompileFunctions(text string) (remainder string) {

	remainder = text

	compileBlockStart := strings.Index(text, "{{{")
	if compileBlockStart == -1 {
		return
	}

	compileBlockEnd := strings.Index(text[compileBlockStart+1:], "}}}")
	if compileBlockEnd == -1 {
		return
	}

	compileText := text[compileBlockStart+3 : compileBlockEnd+3]
	compileTemplate := NewReportTemplate("compile", compileText, &TripleSingleQuoteDelimiters)
	compileTemplate.Properties = rt.Properties
	compileTemplate.Parse()
	var buffer strings.Builder
	compileTemplate.TextTemplate.Execute(&buffer, nil)
	//	rt.Properties = compileTemplate.Properties

	remainder = text[compileBlockEnd+6:]
	return
}

func (rp *ReportTemplate) Parse() (err error) {

	text := util.TrimEachLine(rp.Text)

	text = rp.CompileFunctions(text)

	text, err = EscapeRawText(rp.Properties.Delimiters, text)
	if err != nil {
		return
	}

	text = RemoveEscapedLineEndings(text)

	rp.TextTemplate = template.New(rp.Name)
	rp.TextTemplate.Funcs(rp.Properties.Funcs)

	_, err = rp.TextTemplate.Parse(text)

	return
}

func (rp *ReportTemplate) AddFunction(name string, function interface{}) {
	rp.Properties.Funcs[name] = function
}

func (rp *ReportTemplate) AddFuncs(newFunctions template.FuncMap) {
	for name, function := range newFunctions {
		rp.AddFunction(name, function)
	}
}

func (rp *ReportTemplate) Expand(data interface{}) (result string, err error) {
	var buffer strings.Builder
	err = rp.TextTemplate.Execute(&buffer, data)
	if err != nil {
		return
	}
	result = buffer.String()
	result = RestoreNewlines(result)
	return
}

func (rp *ReportTemplate) ExpandSubreport(name string, text string, data interface{}) (report string, err error) {
	reportTemplate := NewReportTemplate("include", string(text), nil)
	reportTemplate.Properties = rp.Properties
	err = reportTemplate.Parse()
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
			s, err = rp.ExpandSubreport(fileName, string(text), nil)
			return
		},
		"up": func(s string) string {
			return strings.ToUpper(s)
		},
		"macro": func(name string, body string) (s string, err error) {
			macroTemplate := NewReportTemplate(name, body, &MacroDelimiters)
			macroTemplate.AddFuncs(rp.Properties.Funcs)
			err = macroTemplate.Parse()
			if err != nil {
				return
			}
			rp.Properties.Macros[name] = macroTemplate
			rp.AddFunction(name, func(args ...interface{}) (result interface{}, err error) {
				macroTemplate := rp.Properties.Macros[name]
				var data interface{}
				if len(args) == 1 {
					data = args[0]
				}
				result, err = macroTemplate.Expand(data)
				return
			})

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
