package blazegraph

import (
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
		"prefix": func(prefix string, uri string) (err error) {
			return
		},
		"select": func(queryTemplate string, args ...interface{}) (rs sparql.ResultSet) {
			queryReportTemplate := reporter.NewReportTemplate(queryTemplate)

			var data interface{}
			if len(args) == 1 {
				data = args[0]
			}
			query, err := queryReportTemplate.Expand(data)
			if err != nil {
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
		"join": func(elems []string, sep string) (js string) {
			js = strings.Join(elems, sep)
			return
		},
	}

	rp.SetFuncs(funcs)
	report, err = rp.Expand(nil)
	if err != nil {
		return
	}

	// report = util.TrimByLine(report)
	return
}
