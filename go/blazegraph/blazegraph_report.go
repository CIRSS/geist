package blazegraph

import (
	"errors"
	"strings"
	"text/template"

	"github.com/tmcphillips/blazegraph-util/reporter"
	"github.com/tmcphillips/blazegraph-util/sparql"
)

func (bc *Client) ExpandReport(rp *reporter.ReportTemplate) (report string, err error) {

	funcs := template.FuncMap{
		"up": func(s string) string {
			return strings.ToUpper(s)
		},
		"prefix": func(prefix string, uri string) (s string, err error) {
			rp.Properties.Prefixes[prefix] = uri
			return "", nil
		},
		"macro": func(name string, body string) (s string, err error) {
			macroTemplate := reporter.NewReportTemplate(name, body, &reporter.MacroDelimiters)
			macroTemplate.SetFuncs(rp.Properties.Funcs)
			err = macroTemplate.Parse(false)
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
		"subquery": func(name string, body string) (s string, err error) {
			rp.Properties.Queries[name] = body
			return "", nil
		},
		"runquery": func(name string, args ...interface{}) (rs sparql.ResultSet, err error) {
			queryTemplateString := rp.Properties.Queries[name]
			sb := strings.Builder{}
			for prefix, uri := range rp.Properties.Prefixes {
				sb.WriteString("PREFIX " + prefix + ": " + "<" + uri + ">" + "\n")
			}
			sb.WriteString(queryTemplateString)

			queryReportTemplate := reporter.NewReportTemplate(name, sb.String(), nil)
			queryReportTemplate.Properties = rp.Properties
			queryReportTemplate.Parse(false)
			var data interface{}
			if len(args) == 1 {
				data = args[0]
			}
			query, re := queryReportTemplate.Expand(data)
			if re != nil {
				return
			}
			rs, _ = bc.Select(query)
			return
		},
		"select": func(queryTemplateString string, args ...interface{}) (rs sparql.ResultSet) {
			sb := strings.Builder{}
			for prefix, uri := range rp.Properties.Prefixes {
				sb.WriteString("PREFIX " + prefix + ": " + "<" + uri + ">" + "\n")
			}
			sb.WriteString(queryTemplateString)

			queryReportTemplate := reporter.NewReportTemplate("select", sb.String(), nil)
			queryReportTemplate.Properties = rp.Properties
			queryReportTemplate.Parse(false)
			var data interface{}
			if len(args) == 1 {
				data = args[0]
			}
			query, re := queryReportTemplate.Expand(data)
			if re != nil {
				return
			}
			rs, _ = bc.Select(query)
			return
		},
		// "query": func(name string, args ...interface{}) (rs sparql.ResultSet, err error) {

		// 	sb := strings.Builder{}
		// 	for prefix, uri := range rp.Properties.Prefixes {
		// 		sb.WriteString("PREFIX " + prefix + ": " + "<" + uri + ">" + "\n")
		// 	}
		// 	sb.WriteString(queryTemplateString)

		// 	macroTemplate := rp.Properties.Macros[name]
		// 	var data interface{}
		// 	if len(args) == 1 {
		// 		data = args[0]
		// 	}
		// 	query, err := macroTemplate.Expand(data)

		// 	print(query)
		// 	rs, _ = bc.Select(query)
		// 	return
		// },
		"tabulate": func(rs sparql.ResultSet) (table string) {
			table = rs.FormattedTable(true)
			return
		},
		"rows": func(rs sparql.ResultSet) (rows [][]string) {
			rows = rs.Rows()
			return
		},
		"column": func(columnIndex int, rs sparql.ResultSet) (column []string) {
			column = rs.Column(columnIndex)
			return
		},
		"vector": func(rs sparql.ResultSet) (vector []string, err error) {
			if rs.RowCount() == 1 {
				vector = rs.Row(0)
			} else if rs.ColumnCount() == 1 {
				vector = rs.Column(0)
			} else {
				err = errors.New("Result set is not a vector.")
			}
			return
		},
		"value": func(rs sparql.ResultSet) (value string, err error) {
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
	}

	rp.SetFuncs(funcs)
	rp.Parse(true)
	report, err = rp.Expand(nil)

	return
}
