package main

import (
	"strings"
	"testing"

	"github.com/tmcphillips/blazegraph-util/util"
)

func TestBlazegraphCmd_report_static_content(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	t.Run("constant-template", func(t *testing.T) {
		outputBuffer.Reset()
		template := `A constant template.`
		Main.InReader = strings.NewReader(template)
		run("blazegraph report")
		util.LineContentsEqual(t, outputBuffer.String(), `
			A constant template.
		`)
	})

	t.Run("function-with-quoted-argument", func(t *testing.T) {
		outputBuffer.Reset()
		template := `
			{{up "A constant template"}}
		`
		Main.InReader = strings.NewReader(template)
		run("blazegraph report")
		util.LineContentsEqual(t, outputBuffer.String(), `
			A CONSTANT TEMPLATE
		`)
	})

	t.Run("function-with-delimited-one-line-argument", func(t *testing.T) {
		outputBuffer.Reset()
		template := `
			{{up <% A constant template %>}}
		`
		Main.InReader = strings.NewReader(template)
		run("blazegraph report")
		util.LineContentsEqual(t, outputBuffer.String(), `
			A CONSTANT TEMPLATE
		`)
	})

	t.Run("function-with-delimited-two-line-argument", func(t *testing.T) {
		outputBuffer.Reset()
		template := `
			{{up <% A constant 
				template %>}}
		`
		Main.InReader = strings.NewReader(template)
		run("blazegraph report")
		util.LineContentsEqual(t, outputBuffer.String(), `
			A CONSTANT
			TEMPLATE
		`)
	})
}

func TestBlazegraphCmd_report_two_triples(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph drop")

	Main.InReader = strings.NewReader(`
		<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
		<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
	`)
	run("blazegraph import --format ttl")

	t.Run("select-piped-to-table", func(t *testing.T) {
		outputBuffer.Reset()
		template := `
			Example select query with tabular output in report

			{{select <%  
					prefix ab: <http://tmcphill.net/tags#>
					SELECT ?s ?o
					WHERE
					{ ?s ab:tag ?o }
				%> | tabulate }}
		`
		Main.InReader = strings.NewReader(template)
		run("blazegraph report")
		util.LineContentsEqual(t, outputBuffer.String(), `
			Example select query with tabular output in report

			s                          | o
			----------------------------------
			http://tmcphill.net/data#x | seven
			http://tmcphill.net/data#y | eight
		`)
	})

	t.Run("select-to-variable-to-table", func(t *testing.T) {
		outputBuffer.Reset()
		template := `
			Example select query with tabular output in report

			{{ with $rs := (select <%
					prefix ab: <http://tmcphill.net/tags#>
					SELECT ?s ?o
					WHERE
					{ ?s ab:tag ?o }
				%>) }} {{ tabulate $rs }} {{end}}
		`
		Main.InReader = strings.NewReader(template)
		run("blazegraph report")
		util.LineContentsEqual(t, outputBuffer.String(), `
			Example select query with tabular output in report

			s                          | o
			----------------------------------
			http://tmcphill.net/data#x | seven
			http://tmcphill.net/data#y | eight
		`)
	})

}

// func TestBlazegraphCmd_report_address_book(t *testing.T) {

// 	var outputBuffer strings.Builder
// 	Main.OutWriter = &outputBuffer
// 	Main.ErrWriter = &outputBuffer

// 	run("blazegraph drop")

// 	Main.InReader = strings.NewReader(`
// 		<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
// 		<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
// 	`)
// 	run("blazegraph import --format ttl")

// 	t.Run("constant-template", func(t *testing.T) {
// 		outputBuffer.Reset()
// 		template := `A constant template.`
// 		Main.InReader = strings.NewReader(template)
// 		run("blazegraph report")
// 		util.LineContentsEqual(t, outputBuffer.String(), `
// 			A constant template.
// 		`)
// 	})
// }
