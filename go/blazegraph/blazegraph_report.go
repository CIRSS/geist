package blazegraph

import (
	"errors"
	"strings"
	"text/template"

	"github.com/tmcphillips/blazegraph-util/reporter"
	"github.com/tmcphillips/blazegraph-util/sparql"
)

func (bc *Client) ExpandReport(rp *reporter.ReportTemplate) (report string, re *reporter.ReportError) {

	funcs := template.FuncMap{
		"up": func(s string) string {
			return strings.ToUpper(s)
		},
		"prefix": func(prefix string, uri string) (s string, err error) {
			rp.Properties.Prefixes[prefix] = uri
			return "", nil
		},
		"select": func(queryTemplateString string, args ...interface{}) (rs sparql.ResultSet) {
			sb := strings.Builder{}
			for prefix, uri := range rp.Properties.Prefixes {
				sb.WriteString("PREFIX " + prefix + ": " + "<" + uri + ">" + "\n")
			}
			sb.WriteString(queryTemplateString)

			queryReportTemplate := reporter.NewReportTemplate(sb.String())
			queryReportTemplate.Properties = rp.Properties
			var data interface{}
			if len(args) == 1 {
				data = args[0]
			}
			query, re := queryReportTemplate.Expand(data, false)
			if re != nil {
				return
			}
			rs, _ = bc.Select(query)
			return
		},
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
	report, re = rp.Expand(nil, true)
	if re != nil {
		return
	}

	return
}
