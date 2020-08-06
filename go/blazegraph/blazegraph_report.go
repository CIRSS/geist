package blazegraph

import (
	"strings"
	"text/template"

	"github.com/tmcphillips/blazegraph-util/reporter"
	"github.com/tmcphillips/blazegraph-util/sparql"
	"github.com/tmcphillips/blazegraph-util/util"
)

func (bc *Client) ExpandReport(reportTemplate string) (report string, err error) {

	funcs := template.FuncMap{
		"up": func(s string) string {
			return strings.ToUpper(s)
		},
		"select": func(query string) (rs sparql.ResultSet) {
			rs, _ = bc.Select(query)
			return
		},
		"tabulate": func(rs sparql.ResultSet) (table string) {
			table = rs.Table(true)
			return
		},
		"rows": func(rs sparql.ResultSet) (values [][]string) {
			variablesAndValues := rs.ValueTable()
			values = variablesAndValues[1:]
			return
		},
		"join": func(elems []string, sep string) (js string) {
			js = strings.Join(elems, sep)
			return
		},
	}

	rt := reporter.NewReportTemplate(reporter.TripleSingleQuoteDelimiters,
		funcs, reportTemplate)
	report, err = rt.Expand(nil)
	if err != nil {
		return
	}

	report = util.TrimByLine(report)
	return
}
