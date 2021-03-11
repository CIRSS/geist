package main

import (
	"strings"
	"testing"

	"github.com/cirss/geist/go/util"
)

func TestBlazegraphCmd_report_static_content(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	t.Run("constant-template", func(t *testing.T) {
		outputBuffer.Reset()
		template := `A constant template.`
		Main.InReader = strings.NewReader(template)
		assertExitCode(t, "blazegraph report", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`A constant template.`)
	})

	t.Run("constant-template-containing-unquoted-percent-symbol", func(t *testing.T) {
		outputBuffer.Reset()
		template := `A constant template with % symbol`
		Main.InReader = strings.NewReader(template)
		assertExitCode(t, "blazegraph report", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`A constant template with % symbol`)
	})

	t.Run("constant-template-containing-doublequotes", func(t *testing.T) {
		outputBuffer.Reset()
		template := `
			"A constant template"
		`
		Main.InReader = strings.NewReader(template)
		assertExitCode(t, "blazegraph report", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`
			 "A constant template"
			`)
	})

	t.Run("constant-template-containing-quoted-percent-symbol", func(t *testing.T) {
		outputBuffer.Reset()
		template := `
			"A constant template with % symbol"
		`
		Main.InReader = strings.NewReader(template)
		assertExitCode(t, "blazegraph report", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`
			 "A constant template with % symbol"
			`)
	})

	t.Run("function-with-quoted-argument", func(t *testing.T) {
		outputBuffer.Reset()
		template := `
			{{up "A constant template"}}
		`
		Main.InReader = strings.NewReader(template)
		assertExitCode(t, "blazegraph report", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`
			 A CONSTANT TEMPLATE
			 `)
	})

	t.Run("function-with-delimited-one-line-argument", func(t *testing.T) {
		outputBuffer.Reset()
		template := `
			{{up '''A constant template'''}}
		`
		Main.InReader = strings.NewReader(template)
		assertExitCode(t, "blazegraph report", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`
			 A CONSTANT TEMPLATE
			`)
	})

	t.Run("function-with-delimited-one-line-argument-containing-single-quotes", func(t *testing.T) {
		outputBuffer.Reset()
		template := `
			{{up '''A 'constant' template'''}}
		`
		Main.InReader = strings.NewReader(template)
		assertExitCode(t, "blazegraph report", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`
			 A 'CONSTANT' TEMPLATE
			`)
	})

	t.Run("function-with-delimited-two-line-argument", func(t *testing.T) {
		outputBuffer.Reset()
		template := `
			{{up '''A constant
				template'''}}
		`
		Main.InReader = strings.NewReader(template)
		assertExitCode(t, "blazegraph report", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`
			 A CONSTANT
			 TEMPLATE
			`)
	})

	t.Run("function-with-delimited-two-line-argument-containing-double-quotes", func(t *testing.T) {
		outputBuffer.Reset()
		template := `
			{{up '''A "constant"
				template'''}}
		`
		Main.InReader = strings.NewReader(template)
		assertExitCode(t, "blazegraph report", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`
			 A "CONSTANT"
			 TEMPLATE
			`)
	})
}

func TestBlazegraphCmd_report_two_triples(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --dataset kb")
	run("blazegraph create --quiet --dataset kb")

	Main.InReader = strings.NewReader(`
		<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
		<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
	`)
	run("blazegraph import --format ttl")

	t.Run("select-piped-to-tabulate", func(t *testing.T) {
		outputBuffer.Reset()
		template := `
			Example select query with tabular output in report

			{{select '''
					prefix ab: <http://tmcphill.net/tags#>
					SELECT ?s ?o
					WHERE
					{ ?s ab:tag ?o }
					ORDER BY ?s
				''' | tabulate}}
		`
		Main.InReader = strings.NewReader(template)
		assertExitCode(t, "blazegraph report", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`
			 Example select query with tabular output in report

			 s                          | o
			 ==================================
			 http://tmcphill.net/data#x | seven
			 http://tmcphill.net/data#y | eight

			`)
	})

	t.Run("select-to-variable-to-tabulate", func(t *testing.T) {
		outputBuffer.Reset()
		template := `
				Example select query with tabular output in report

				{{with $tags := (select '''
						prefix ab: <http://tmcphill.net/tags#>
						SELECT ?s ?o
						WHERE
						{ ?s ab:tag ?o }
						ORDER BY ?s
					''')}}{{ tabulate $tags}}{{end}}
			`
		Main.InReader = strings.NewReader(template)
		assertExitCode(t, "blazegraph report", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`
			 Example select query with tabular output in report

			 s                          | o
			 ==================================
			 http://tmcphill.net/data#x | seven
			 http://tmcphill.net/data#y | eight

			`)
	})

	t.Run("select-to-dot-to-tabulate", func(t *testing.T) {
		outputBuffer.Reset()
		template := `
				Example select query with tabular output in report

				{{with (select '''
						prefix ab: <http://tmcphill.net/tags#>
						SELECT ?s ?o
						WHERE
						{ ?s ab:tag ?o }
						ORDER BY ?s
					''')}} {{tabulate .}} {{end}}
			`
		Main.InReader = strings.NewReader(template)
		assertExitCode(t, "blazegraph report", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`
			 Example select query with tabular output in report

			 s                          | o
			 ==================================
			 http://tmcphill.net/data#x | seven
			 http://tmcphill.net/data#y | eight

			`)
	})

	t.Run("select-to-variable-to-range", func(t *testing.T) {
		outputBuffer.Reset()
		template := `
				Example select query with tabular output in report

				{{ with (select '''
						prefix ab: <http://tmcphill.net/tags#>
						SELECT ?s ?o
						WHERE
						{ ?s ab:tag ?o }
						ORDER BY ?s
					''') }}											\
																	\
					Variables:
					{{join (.Head.Vars) ", "}}

					Values:
					{{range (rows .)}}{{ join . ", " | println}}{{end}}

				{{end}}
			`
		Main.InReader = strings.NewReader(template)
		assertExitCode(t, "blazegraph report", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`
			 Example select query with tabular output in report

			 Variables:
			 s, o

			 Values:
			 http://tmcphill.net/data#x, seven
			 http://tmcphill.net/data#y, eight



			`)
	})

}

func TestBlazegraphCmd_report_multiple_queries(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --dataset kb")
	run("blazegraph create --quiet --dataset kb")

	Main.InReader = strings.NewReader(`
		<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
		<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
	`)
	run("blazegraph import --format ttl")

	outputBuffer.Reset()
	template :=
		`{{prefix "ab" "http://tmcphill.net/tags#"}}	\
														\
		 {{with $subjects := (select '''

				SELECT ?s
				WHERE
				{ ?s ab:tag ?o }
				ORDER BY ?s

			''') | vector }}							\
														\
			{{range $subject := $subjects }} 			\
				{{with $objects := (select '''

						SELECT ?o
						WHERE
						{ <{{.}}> ab:tag ?o }
						ORDER BY ?o

					''' $subject)}} 					\
					{{tabulate $objects}}
				{{end}}									\
			{{end}}										\
		{{end }}`

	Main.InReader = strings.NewReader(template)

	assertExitCode(t, "blazegraph report", 0)

	util.LineContentsEqual(t, outputBuffer.String(),
		`o
		 ====
		 seven

		 o
		 ====
		 eight

		 `)
}

func TestBlazegraphCmd_report_macros(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --dataset kb")
	run("blazegraph create --quiet --dataset kb")

	Main.InReader = strings.NewReader(`
		<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
		<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
	`)
	run("blazegraph import --format ttl")

	outputBuffer.Reset()
	template :=
		`{{{											\
			{{ macro "M1" "Subject" '''{{select <?		\
				SELECT ?o								\
				WHERE { <{{.}}> ab:tag ?o }				\
				ORDER BY ?o								\
			?> $Subject | tabulate }}''' }}				\
		}}}												\
														\
		{{prefix "ab" "http://tmcphill.net/tags#"}} 	\
														\
		{{with $subjects := (select '''					\
				SELECT ?s								\
				WHERE									\
				{ ?s ab:tag ?o }						\
				ORDER BY ?s								\
			''') | vector }}							\
			{{range $subject := $subjects }}			\
				{{ M1 $subject }}

			{{end}}										\
		{{end}}											\
	`
	Main.InReader = strings.NewReader(template)

	assertExitCode(t, "blazegraph report", 0)

	util.LineContentsEqual(t, outputBuffer.String(),
		`o
		 ====
		 seven

		 o
		 ====
		 eight

		`)
}

func TestBlazegraphCmd_report_subqueries(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --dataset kb")
	run("blazegraph create --quiet --dataset kb")

	Main.InReader = strings.NewReader(`
		<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
		<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
	`)
	run("blazegraph import --format ttl")

	outputBuffer.Reset()
	template := `
		{{{
			{{ query "Q1" '''									\
				SELECT ?s										\
				WHERE											\
				{ ?s ab:tag ?o }								\
				ORDER BY ?s										\
			''' }}												\
																\
			{{ query "Q2" "Subject" '''	             			\
				SELECT ?o 										\
				WHERE { <{{$Subject}}> ab:tag ?o } 				\
				ORDER BY ?o 									\
			''' }}												\
		}}}														\
																\
		{{ prefix "ab" "http://tmcphill.net/tags#" }}			\
																\
		{{ range (Q1 | vector) }}								\
			{{ Q2 . | tabulate }}
		{{ end }}												\
	`
	Main.InReader = strings.NewReader(template)

	assertExitCode(t, "blazegraph report", 0)

	util.LineContentsEqual(t, outputBuffer.String(),
		`o
		 ====
		 seven

		 o
		 ====
		 eight

		`)
}

func TestBlazegraphCmd_report_address_book(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --dataset kb")
	run("blazegraph create --quiet --dataset kb")

	run("blazegraph import --format jsonld --file testdata/address-book.jsonld")

	t.Run("constant-template", func(t *testing.T) {
		outputBuffer.Reset()
		template := `
			Craig's email addresses
			=======================
			{{ range (select '''
				PREFIX ab: <http://learningsparql.com/ns/addressbook#>
				SELECT ?email
				WHERE
				{
					?person ab:firstname "Craig"^^<http://www.w3.org/2001/XMLSchema#string> .
					?person ab:email     ?email .
				}
			''' | vector) }}																	\
				{{ . }}
			{{end}}																				\
		`
		Main.InReader = strings.NewReader(template)

		assertExitCode(t, "blazegraph report", 0)

		util.LineContentsEqual(t, outputBuffer.String(),
			`
			 Craig's email addresses
			 =======================
			 c.ellis@usairwaysgroup.com
			 craigellis@yahoo.com
			`)
	})
}

func TestBlazegraphCmd_report_address_book_imports(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --dataset kb")
	run("blazegraph create --quiet --dataset kb")
	run("blazegraph import --format jsonld --file testdata/address-book.jsonld")

	t.Run("constant-template", func(t *testing.T) {
		outputBuffer.Reset()
		template :=
			`{{{
				{{ include "testdata/address-rules.gst" }}					\
			}}}
																			\
			{{ prefix "ab" "http://learningsparql.com/ns/addressbook#" }}	\
																			\
			Craig's email addresses
			=======================
																			\
			{{ range $Name := GetEmailForFirstName "Craig" | vector }}	\
				{{ $Name }}
			{{end}}															\
		`
		Main.InReader = strings.NewReader(template)

		assertExitCode(t, "blazegraph report", 0)

		util.LineContentsEqual(t, outputBuffer.String(),
			`
			Craig's email addresses
			=======================
			c.ellis@usairwaysgroup.com
			craigellis@yahoo.com
		`)
	})
}

func TestBlazegraphCmd_report_subquery_functions(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --dataset kb")
	run("blazegraph create --quiet --dataset kb")

	Main.InReader = strings.NewReader(`
		<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
		<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
	`)
	run("blazegraph import --format ttl")

	outputBuffer.Reset()
	template := `

		{{{
			{{ query "Q1" '''
				SELECT ?s
				WHERE
				{ ?s ab:tag ?o }
				ORDER BY ?s
			''' }}

			{{ query "Q2" "Subject" '''
				SELECT ?o
				WHERE { <{{$Subject}}> ab:tag ?o }
				ORDER BY ?o
			''' }}
		}}}
															\
		{{ prefix "ab" "http://tmcphill.net/tags#" }}		\
															\
		{{ range (Q1 | vector) }}							\
			{{ Q2 . | tabulate }}
		{{ end }}											\
	`
	Main.InReader = strings.NewReader(template)

	assertExitCode(t, "blazegraph report", 0)

	util.LineContentsEqual(t, outputBuffer.String(),
		`
		 o
		 ====
		 seven

		 o
		 ====
		 eight

		`)
}

func TestBlazegraphCmd_report_macro_functions(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --dataset kb")
	run("blazegraph create --quiet --dataset kb")

	Main.InReader = strings.NewReader(`
		<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
		<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
	`)
	run("blazegraph import --format ttl")

	outputBuffer.Reset()
	template := `

		{{{
			{{prefix "ab" "http://tmcphill.net/tags#"}}

			{{macro "M1" "Subject" '''{{select <?
				SELECT ?o
				WHERE { <{{.}}> ab:tag ?o }
				ORDER BY ?o
			?> $Subject | tabulate }}''' }}
		}}}													\\
															\\
		{{with $subjects := (select '''

				SELECT ?s
				WHERE
				{ ?s ab:tag ?o }
				ORDER BY ?s

			''') | vector }}								\\
															\\
			{{range $subject := $subjects }}				\\
				{{ M1 $subject }}

			{{end}}											\\
		{{end}}												\\
`
	Main.InReader = strings.NewReader(template)

	assertExitCode(t, "blazegraph report", 0)

	util.LineContentsEqual(t, outputBuffer.String(),
		`o
		 ====
		 seven

		 o
		 ====
		 eight

		`)
}

func TestBlazegraphCmd_report_macro_calls_query(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --dataset kb")
	run("blazegraph create --quiet --dataset kb")

	Main.InReader = strings.NewReader(`
		<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
		<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
	`)
	run("blazegraph import --format ttl")

	outputBuffer.Reset()
	template := `
		{{{
			{{prefix "ab" "http://tmcphill.net/tags#"}}

			{{query "select_subjects" '''
				SELECT DISTINCT ?s
				WHERE
				{ ?s ab:tag ?o }
				ORDER BY ?s
			''' }}

			{{query "select_tags_for_subject" "Subject" '''
				SELECT ?tag
				WHERE { <{{$Subject}}> ab:tag ?tag }
				ORDER BY ?tag
			''' }}

			{{macro "tabulate_tags_for_subject" "Subject" '''
				{{ select_tags_for_subject $Subject | tabulate }}
			''' }}
		}}}
																	\
		{{range $Subject := select_subjects | vector }}				\
			{{ tabulate_tags_for_subject $Subject }}

		{{end}}														\
`
	Main.InReader = strings.NewReader(template)

	assertExitCode(t, "blazegraph report", 0)

	util.LineContentsEqual(t, outputBuffer.String(),
		`tag
		 ====
		 seven

		 tag
		 ====
		 eight

		`)
}

var expectedReportHelpOutput = string(
	`
	blazegraph report: Expands the provided report template using the identified RDF dataset.

	usage: blazegraph report [<flags>]

	flags:
		-dataset name
				name of RDF dataset to create report from
		-file string
				File containing report template to expand (default "-")
		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
		-quiet
				Discard normal command output

	`)

func TestBlazegraphCmd_report_help(t *testing.T) {
	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer
	assertExitCode(t, "blazegraph report help", 0)
	util.LineContentsEqual(t, outputBuffer.String(), expectedReportHelpOutput)
}

func TestBlazegraphCmd_help_report(t *testing.T) {
	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer
	assertExitCode(t, "blazegraph help report", 0)
	util.LineContentsEqual(t, outputBuffer.String(), expectedReportHelpOutput)
}

func TestBlazegraphCmd_report_bad_flag(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	assertExitCode(t, "blazegraph report --not-a-flag", 1)

	util.LineContentsEqual(t, outputBuffer.String(),
		`blazegraph report: flag provided but not defined: -not-a-flag

		usage: blazegraph report [<flags>]

		flags:
			-dataset name
					name of RDF dataset to create report from
			-file string
					File containing report template to expand (default "-")
			-instance URL
					URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
			-quiet
					Discard normal command output

	`)
}
