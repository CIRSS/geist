package blazegraph

import (
	"strings"
	"text/template"

	"github.com/tmcphillips/blazegraph-util/reporter"
)

func prependPrefixes(rp *reporter.ReportTemplate, text string) string {
	sb := strings.Builder{}
	for prefix, uri := range rp.Properties.Prefixes {
		sb.WriteString("PREFIX " + prefix + ": " + "<" + uri + ">" + "\n")
	}
	sb.WriteString(text)
	return sb.String()
}

func (bc *Client) selectFunc(rp *reporter.ReportTemplate, queryText string, args []interface{}) (rs interface{}, err error) {

	var data interface{}
	if len(args) == 1 {
		data = args[0]
	}

	query, re := rp.ExpandSubreport("select", prependPrefixes(rp, queryText), false, data)
	if re != nil {
		return
	}
	return bc.Select(query)
}

func (bc *Client) runQueryFunc(rp *reporter.ReportTemplate, queryText string, args []interface{}) (rs interface{}, err error) {

	var data interface{}
	if len(args) == 1 {
		data = args[0]
	}
	reportTemplate := reporter.NewReportTemplate("include", string(queryText), nil)
	reportTemplate.Properties = rp.Properties
	reportTemplate.Parse(false)
	rs, err = reportTemplate.Expand(data)
	print(rs)
	return
}

func (bc *Client) ExpandReport(rp *reporter.ReportTemplate) (report string, err error) {

	funcs := template.FuncMap{
		"prefix": func(prefix string, uri string) (s string, err error) {
			rp.Properties.Prefixes[prefix] = uri
			return "", nil
		},
		"runquery": func(name string, args ...interface{}) (rs interface{}, err error) {
			queryText := rp.Properties.Queries[name]
			var data interface{}
			if len(args) == 1 {
				data = args[0]
			}
			query, err := rp.ExpandSubreport(name, prependPrefixes(rp, queryText), false, data)
			if err != nil {
				return
			}
			rs, err = bc.Select(query)
			return
		},
		"select": func(queryText string, args ...interface{}) (interface{}, error) {
			return bc.selectFunc(rp, queryText, args)
		},
		"query": func(name string, body string) (s string, err error) {
			rp.Properties.Queries[name] = body
			rp.AddFunction(name, func(args ...interface{}) (rs interface{}, err error) {
				queryText := rp.Properties.Queries[name]
				var data interface{}
				if len(args) == 1 {
					data = args[0]
				}
				query, err := rp.ExpandSubreport(name, prependPrefixes(rp, queryText), false, data)
				if err != nil {
					return
				}
				rs, err = bc.Select(query)
				return
			})
			return "", nil
		},
	}

	rp.AddFuncs(funcs)
	rp.Parse(true)
	report, err = rp.Expand(nil)

	return
}
