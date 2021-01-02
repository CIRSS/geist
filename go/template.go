package geist

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	textTemplate "text/template"

	"github.com/cirss/geist/util"
)

type Properties struct {
	Delimiters DelimiterPair
	Funcs      textTemplate.FuncMap
	Prefixes   map[string]string
	Macros     map[string]*Template
	Queries    map[string]string
	Rules      map[string]string
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

// Template is customizable to create reports from different kinds of templates.
type Template struct {
	Name         string
	Text         string
	TextTemplate *textTemplate.Template
	Properties   Properties
}

// NewTemplate returns a geist.Template with the given customizations.
func NewTemplate(name string, text string, delimiters *DelimiterPair) *Template {
	rt := new(Template)
	rt.Name = name
	text = util.Trim(text)
	rt.Text = util.TrimEachLine(text)
	rt.Properties.Funcs = textTemplate.FuncMap{}
	if delimiters != nil {
		rt.Properties.Delimiters = *delimiters
	} else {
		rt.Properties.Delimiters = DefaultDelimiters
	}
	rt.Properties.Prefixes = map[string]string{}
	rt.Properties.Macros = map[string]*Template{}
	rt.Properties.Queries = map[string]string{}
	rt.Properties.Rules = map[string]string{}
	rt.addStandardFunctions()
	return rt
}

func (rt *Template) CompileFunctions(text string) (remainder string) {

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
	compileTemplate := NewTemplate("compile", compileText, &TripleSingleQuoteDelimiters)
	compileTemplate.Properties = rt.Properties
	compileTemplate.Parse()
	var buffer strings.Builder
	compileTemplate.TextTemplate.Execute(&buffer, nil)
	//	rt.Properties = compileTemplate.Properties

	if len(text) >= compileBlockEnd+6 {
		remainder = text[compileBlockEnd+6:]
	} else {
		remainder = ""
	}
	return
}

func (rp *Template) Parse() (err error) {
	text := rp.CompileFunctions(rp.Text)
	text, err = EscapeRawText(rp.Properties.Delimiters, text)
	if err != nil {
		return
	}
	text = RemoveEscapedLineEndings(text)
	rp.TextTemplate = textTemplate.New(rp.Name)
	rp.TextTemplate.Funcs(rp.Properties.Funcs)
	_, err = rp.TextTemplate.Parse(text)
	return
}

func (rp *Template) AddFunction(name string, function interface{}) {
	rp.Properties.Funcs[name] = function
}

func (rp *Template) AddFuncs(newFunctions textTemplate.FuncMap) {
	for name, function := range newFunctions {
		rp.AddFunction(name, function)
	}
}

func (rp *Template) Expand(data interface{}) (result string, err error) {
	var buffer strings.Builder
	err = rp.TextTemplate.Execute(&buffer, data)
	if err != nil {
		return
	}
	result = buffer.String()
	result = RestoreNewlines(result)
	result = util.Trim(result)
	return
}

func (rp *Template) ExpandSubreport(name string, text string, data interface{}) (report string, err error) {
	reportTemplate := NewTemplate("include", string(text), nil)
	reportTemplate.Properties = rp.Properties
	err = reportTemplate.Parse()
	if err != nil {
		return
	}
	report, err = reportTemplate.Expand(data)
	return
}

func (rp *Template) expandMacro(name string, args []interface{}) (result interface{}, err error) {
	macroTemplate := rp.Properties.Macros[name]
	result, err = macroTemplate.Expand(args)
	return
}

func GetParameterAppendedBody(args []string) (body string) {
	argsCount := len(args)
	body = args[argsCount-1]
	var parameters []string
	parameterCount := argsCount - 1
	if parameterCount > 0 {
		parameters = args[0:parameterCount]
		buffer := strings.Builder{}
		buffer.WriteString("{{ with $args := . }}\n")
		for index, parameter := range parameters {
			buffer.WriteString(fmt.Sprintf("{{ with $%s := index $args %d }}\n", parameter, index))
		}
		buffer.WriteString(body)
		buffer.WriteString(strings.Repeat("\n{{end}}", parameterCount+1))
		body = buffer.String()
	}
	return
}

func (rp *Template) addStandardFunctions() {

	funcs := textTemplate.FuncMap{
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
		"macro": func(name string, args ...string) (s string, err error) {
			if len(args) == 0 {
				err = errors.New("No body provided for macro " + name)
				return
			}
			body := GetParameterAppendedBody(args)
			macroTemplate := NewTemplate(name, body, &MacroDelimiters)
			macroTemplate.AddFuncs(rp.Properties.Funcs)
			err = macroTemplate.Parse()
			if err != nil {
				return
			}
			rp.Properties.Macros[name] = macroTemplate
			rp.AddFunction(name, func(args ...interface{}) (result interface{}, err error) {
				return rp.expandMacro(name, args)
			})

			return "", nil
		},
		"tabulate": func(rs DataTable) (table string) {
			table = rs.FormattedTable(true)
			return
		},
		"rows": func(rs DataTable) (rows [][]string) {
			rows = rs.Rows()
			return
		},
		"column": func(columnIndex int, rs *DataTable) (column []string) {
			column = (*rs).Column(columnIndex)
			return
		},
		"vector": func(rs DataTable) (vector []string, err error) {
			if rs.RowCount() == 1 {
				vector = rs.Row(0)
			} else if rs.ColumnCount() == 1 {
				vector = rs.Column(0)
			} else {
				err = errors.New("Result set is not a vector.")
			}
			return
		},
		"value": func(rs DataTable) (value string, err error) {
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
		"sp": func(args ...int) (s string) {
			count := 1
			if len(args) > 0 {
				count = args[0]
			}
			return strings.Repeat(" ", count)
		},
	}

	rp.AddFuncs(funcs)

}
