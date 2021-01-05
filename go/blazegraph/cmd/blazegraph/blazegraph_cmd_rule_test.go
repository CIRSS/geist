package main

import (
	"strings"
	"testing"

	"github.com/cirss/geist/util"
)

func TestBlazegraphCmd_rule_in_select(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --dataset kb")
	run("blazegraph create --dataset kb")

	Main.InReader = strings.NewReader(`
		<:x> <:foo> <:y> .
		<:y> <:bar> <:z> .
		<:z> <:baz> "baz" .
	`)
	run("blazegraph import --format ttl")

	outputBuffer.Reset()
	template := `

		{{{
			{{ rule "foo_bar_baz" "s" "o" '''
				{{_subject $s}} <:foo> ?y .
				?y <:bar> ?z .
				?z <:baz> {{_object $o}} .
			''' }}
		}}}

		{{ select '''
			SELECT DISTINCT ?o
			WHERE
			{ {{ foo_bar_baz "?s" "?o" }} }
			ORDER BY ?o
		''' | value }}
	`
	Main.InReader = strings.NewReader(template)
	run("blazegraph report")
	util.LineContentsEqual(t, outputBuffer.String(), `
		baz
	`)
}

func TestBlazegraphCmd_rule_in_query(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --dataset kb")
	run("blazegraph create --dataset kb")

	Main.InReader = strings.NewReader(`
		<:x> <:foo> <:y> .
		<:y> <:bar> <:z> .
		<:z> <:baz> "baz" .
	`)
	run("blazegraph import --format ttl")

	outputBuffer.Reset()
	template := `

		{{{
			{{ rule "foo_bar_baz" "s" "o" '''
				{{_subject $s}} <:foo> ?y .
				?y <:bar> ?z .
				?z <:baz> {{_object $o}} .
			''' }}

			{{query "select_foo_bar_baz" '''
				SELECT DISTINCT ?o
				WHERE
				{ {{ foo_bar_baz "?s" "?o" }} }
				ORDER BY ?o
			''' }}

		}}}

		{{ select_foo_bar_baz ":x" | value }}
`
	Main.InReader = strings.NewReader(template)
	run("blazegraph report")
	util.LineContentsEqual(t, outputBuffer.String(), `
		baz
	`)
}

func TestBlazegraphCmd_rule_in_rule(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --dataset kb")
	run("blazegraph create --dataset kb")

	Main.InReader = strings.NewReader(`
		<:x> <:foo> <:y> .
		<:y> <:bar> <:z> .
		<:z> <:baz> "baz" .
	`)
	run("blazegraph import --format ttl")

	outputBuffer.Reset()
	template := `

		{{{
			{{ rule "foo_bar_baz_rule_1" "s" "o" '''
				{{_subject $s}} <:foo> ?y .
				?y <:bar> ?z .
				?z <:baz> {{_object $o}} .
			''' }}

			{{ rule "foo_bar_baz_rule_2" "s" "o" '''
				{{ foo_bar_baz_rule_1 $s $o }}
			'''}}

			{{query "foo_bar_baz_query" '''
				SELECT DISTINCT ?s ?o
				WHERE
				{ {{ foo_bar_baz_rule_2 "?s" "?o" }} }
				ORDER BY ?o
			''' }}

		}}}

		{{ foo_bar_baz_query ":x" | tabulate }}
`
	Main.InReader = strings.NewReader(template)
	run("blazegraph report")
	util.LineContentsEqual(t, outputBuffer.String(), `
		s                                   | o
		=========================================
		http://127.0.0.1:9999/blazegraph/:x | baz
	`)
}

func TestBlazegraphCmd_rule_in_query_called_by_macro(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --dataset kb")
	run("blazegraph create --dataset kb")

	Main.InReader = strings.NewReader(`
		<:y> <:tag> "eight" .
		<:x> <:tag> "seven" .
	`)
	run("blazegraph import --format ttl")

	outputBuffer.Reset()
	template := `

		{{{
			{{ rule "hasTag" "s" "o" '''
				{{_subject $s}} <:tag> {{_object $o}}
			''' }}

			{{query "select_subjects" '''
				SELECT DISTINCT ?s
				WHERE
				{ {{ hasTag "?s" "?o" }} }
				ORDER BY ?s
			''' }}

			{{query "select_tags_for_subject" "Subject" '''
				SELECT ?tag
				WHERE { {{ hasTag $Subject "?tag" }} }
				ORDER BY ?tag
			''' }}

			{{macro "tabulate_tags_for_subject" "Subject" '''
				{{ select_tags_for_subject $Subject | tabulate }}
			''' }}
		}}}

		{{ range $Subject := select_subjects | vector }}
			{{ tabulate_tags_for_subject $Subject }}
		{{ end }}

`
	Main.InReader = strings.NewReader(template)
	run("blazegraph report")
	util.LineContentsEqual(t, outputBuffer.String(), `
		tag
		====
		seven

		tag
		====
		eight
	`)
}
