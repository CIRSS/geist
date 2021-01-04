package main

import (
	"strings"
	"testing"

	"github.com/cirss/geist/util"
)

func TestBlazegraphCmd_report_query_uses_rule(t *testing.T) {

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
