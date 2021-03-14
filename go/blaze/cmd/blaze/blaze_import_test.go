package main

import (
	"strings"
	"testing"

	"github.com/cirss/geist/go/util"
)

func TestBlazegraphCmd_import_two_triples(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	t.Run("import_nt", func(t *testing.T) {

		run("blaze destroy --dataset kb --quiet")
		run("blaze create --quiet --dataset kb")

		Main.InReader = strings.NewReader(`
			<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
			<http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
		`)

		assertExitCode(t, "blaze import --format nt", 0)

		outputBuffer.Reset()
		run("blaze export --format nt --sort=true")
		util.LineContentsEqual(t, outputBuffer.String(),
			`<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
			 <http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
			`)
	})

	t.Run("import_ttl", func(t *testing.T) {

		run("blaze destroy --dataset kb --quiet")
		run("blaze create --quiet --dataset kb")

		Main.InReader = strings.NewReader(
			`@prefix data: <http://tmcphill.net/data#> .
			 @prefix tags: <http://tmcphill.net/tags#> .

			 data:y tags:tag "eight" .
			 data:x tags:tag "seven" .
			`)

		assertExitCode(t, "blaze import --format ttl", 0)

		outputBuffer.Reset()
		run("blaze export --format nt --sort=true")
		util.LineContentsEqual(t, outputBuffer.String(),
			`<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
			 <http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
			`)
	})

	t.Run("import_jsonld", func(t *testing.T) {

		run("blaze destroy --dataset kb --quiet")
		run("blaze create --quiet --dataset kb")

		Main.InReader = strings.NewReader(
			`
			[
				{
					"@id": "http://tmcphill.net/data#x",
					"http://tmcphill.net/tags#tag": "seven"
				},
				{
					"@id": "http://tmcphill.net/data#y",
					"http://tmcphill.net/tags#tag": "eight"
				}
			]
			`)

		assertExitCode(t, "blaze import --format jsonld", 0)

		outputBuffer.Reset()
		run("blaze export --format nt --sort=true")
		util.LineContentsEqual(t, outputBuffer.String(),
			`<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven"^^<http://www.w3.org/2001/XMLSchema#string> .
			 <http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight"^^<http://www.w3.org/2001/XMLSchema#string> .
			`)
	})

	t.Run("import_ttl", func(t *testing.T) {

		run("blaze destroy --dataset kb --quiet")
		run("blaze create --quiet --dataset kb")

		Main.InReader = strings.NewReader(
			`@prefix data: <http://tmcphill.net/data#> .
			 @prefix tags: <http://tmcphill.net/tags#> .

			 data:y tags:tag "eight" .
			 data:x tags:tag "seven" .
			`)

		assertExitCode(t, "blaze import --format ttl", 0)

		outputBuffer.Reset()
		run("blaze export --format nt --sort=true")
		util.LineContentsEqual(t, outputBuffer.String(),
			`<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
			 <http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
			`)
	})

	t.Run("import_xml", func(t *testing.T) {

		run("blaze destroy --dataset kb --quiet")
		run("blaze create --quiet --dataset kb")

		Main.InReader = strings.NewReader(
			`<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#">

			 <rdf:Description rdf:about="http://tmcphill.net/data#y">
				<tag xmlns="http://tmcphill.net/tags#">eight</tag>
		 	 </rdf:Description>

 			 <rdf:Description rdf:about="http://tmcphill.net/data#x">
				<tag xmlns="http://tmcphill.net/tags#">seven</tag>
  			 </rdf:Description>

			 </rdf:RDF>
			`)

		assertExitCode(t, "blaze import --format xml", 0)

		outputBuffer.Reset()
		run("blaze export --format nt --sort=true")
		util.LineContentsEqual(t, outputBuffer.String(),
			`<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
			 <http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
			`)
	})
}

func TestBlazegraphCmd_import_specific_dataset(t *testing.T) {

	triples_ttl :=
		`<http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
		 <http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
		`
	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	// t.Run("default", func(t *testing.T) {
	// 	outputBuffer.Reset()
	// 	run("blaze destroy --silent")
	// 	run("blaze create --quiet")
	// 	Main.InReader = strings.NewReader(triples_ttl)
	// 	assertExitCode(t, "blaze import", 0)
	// 	run("blaze export --sort=true")
	// 	util.LineContentsEqual(t, outputBuffer.String(), triples_ttl)
	// })

	t.Run("single_custom", func(t *testing.T) {
		outputBuffer.Reset()
		run("blaze destroy --dataset foo --silent")
		run("blaze create --dataset foo --quiet")
		Main.InReader = strings.NewReader(triples_ttl)
		assertExitCode(t, "blaze import --dataset foo", 0)
		run("blaze export --dataset foo --sort=true")
		util.LineContentsEqual(t, outputBuffer.String(), triples_ttl)
	})

}

var expectedImportHelpOutput = string(
	`
	blaze import: Imports triples in the specified format into an RDF dataset.

	usage: blaze import [<flags>]

	flags:
		-dataset name
				name of RDF dataset to import triples into (default "kb")
		-file string
				File containing triples to import (default "-")
		-format string
				Format of triples to import [jsonld, nt, ttl, or xml] (default "ttl")
		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
		-quiet
				Discard normal command output
		-silent
				Discard normal and error command output

	`)

func TestBlazegraphCmd_import_help(t *testing.T) {
	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer
	assertExitCode(t, "blaze import help", 0)
	util.LineContentsEqual(t, outputBuffer.String(), expectedImportHelpOutput)
}

func TestBlazegraphCmd_help_import(t *testing.T) {
	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer
	assertExitCode(t, "blaze help import", 0)
	util.LineContentsEqual(t, outputBuffer.String(), expectedImportHelpOutput)
}

func TestBlazegraphCmd_import_bad_flag(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	assertExitCode(t, "blaze import --not-a-flag", 1)

	util.LineContentsEqual(t, outputBuffer.String(),
		`blaze import: flag provided but not defined: -not-a-flag

		usage: blaze import [<flags>]

		flags:
			-dataset name
					name of RDF dataset to import triples into (default "kb")
			-file string
					File containing triples to import (default "-")
			-format string
					Format of triples to import [jsonld, nt, ttl, or xml] (default "ttl")
			-instance URL
					URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
			-quiet
					Discard normal command output
			-silent
					Discard normal and error command output

	`)
}
