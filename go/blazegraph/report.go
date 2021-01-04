package blazegraph

import (
	"errors"
	"strings"
	"text/template"

	"github.com/cirss/geist"
)

func prependPrefixes(rp *geist.Template, text string) string {
	sb := strings.Builder{}
	for prefix, uri := range rp.Properties.Prefixes {
		sb.WriteString("PREFIX " + prefix + ": " + "<" + uri + ">" + "\n")
	}
	sb.WriteString(text)
	return sb.String()
}

func (bc *Client) selectFunc(rp *geist.Template, queryText string, args []interface{}) (rs interface{}, err error) {

	var data interface{}
	if len(args) == 1 {
		data = args[0]
	}

	query, re := rp.ExpandSubreport("select", prependPrefixes(rp, queryText), data)
	if re != nil {
		return
	}
	return bc.Select(query)
}

func (bc *Client) runQueryFunc(rp *geist.Template, queryText string, args []interface{}) (rs interface{}, err error) {

	var data interface{}
	if len(args) == 1 {
		data = args[0]
	}
	reportTemplate := geist.NewTemplate("include", string(queryText), nil)
	reportTemplate.Properties = rp.Properties
	reportTemplate.Parse()
	rs, err = reportTemplate.Expand(data)
	print(rs)
	return
}

func (bc *Client) ExpandReport(rp *geist.Template) (report string, err error) {

	funcs := template.FuncMap{
		"subject": func(s string) string {
			if s[0] == '?' {
				return s
			}
			return "<" + s + ">"
		},
		"prefix": func(prefix string, uri string) (s string, err error) {
			rp.Properties.Prefixes[prefix] = uri
			return "", nil
		},
		"select": func(queryText string, args ...interface{}) (interface{}, error) {
			return bc.selectFunc(rp, queryText, args)
		},
		"query": func(name string, args ...string) (s string, err error) {
			if len(args) == 0 {
				err = errors.New("No body provided for query " + name)
				return
			}
			body := geist.GetParameterAppendedBody(args)
			rp.Properties.Queries[name] = body
			rp.AddFunction(name, func(args ...interface{}) (rs interface{}, err error) {
				queryText := rp.Properties.Queries[name]
				query, err := rp.ExpandSubreport(name, prependPrefixes(rp, queryText), args)
				if err != nil {
					return
				}
				rs, err = bc.Select(query)
				return
			})
			return "", nil
		},

		"rule": func(name string, args ...string) (s string, err error) {
			if len(args) == 0 {
				err = errors.New("No body provided for rule " + name)
				return
			}
			body := geist.GetParameterAppendedBody(args)
			rp.Properties.Rules[name] = body
			rp.AddFunction(name, func(args ...interface{}) (rs interface{}, err error) {
				ruleText := rp.Properties.Rules[name]
				rs, err = rp.ExpandSubreport(name, ruleText, args)
				return
			})
			return "", nil
		},
	}

	rp.AddFuncs(funcs)
	rp.Parse()
	report, err = rp.Expand(nil)

	return
}
