package blazegraph

import (
	"strings"
	"text/template"

	"github.com/tmcphillips/blazegraph-util/reporter"
)

func (bc *Client) ExpandReport(rp *reporter.ReportTemplate) (report string, err error) {

	funcs := template.FuncMap{
		"prefix": func(prefix string, uri string) (s string, err error) {
			rp.Properties.Prefixes[prefix] = uri
			return "", nil
		},
		"runquery": func(name string, args ...interface{}) (rs interface{}, err error) {
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
		"select": func(queryTemplateString string, args ...interface{}) (rs interface{}) {
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
	}

	rp.AddFuncs(funcs)
	rp.Parse(true)
	report, err = rp.Expand(nil)

	return
}
