package main

import (
	"strings"
	"testing"

	"github.com/cirss/geist/go/util"
)

func TestBlazegraphCmd_create_default_dataset(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --all --quiet")

	assertExitCode(t, "blazegraph create", 0)
	util.LineContentsEqual(t, outputBuffer.String(),
		`Successfully created dataset kb
		`)

	outputBuffer.Reset()
	assertExitCode(t, "blazegraph list", 0)
	util.LineContentsEqual(t, outputBuffer.String(),
		`kb
		`)
}

func TestBlazegraphCmd_create_dataset_quiet(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --all --quiet")

	assertExitCode(t, "blazegraph create --quiet", 0)
	util.LineContentsEqual(t, outputBuffer.String(),
		``)

	outputBuffer.Reset()
	assertExitCode(t, "blazegraph list", 0)
	util.LineContentsEqual(t, outputBuffer.String(),
		`kb
		`)
}

func TestBlazegraphCmd_create_default_already_exists(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --all --quiet")
	run("blazegraph create --quiet")

	assertExitCode(t, "blazegraph create", 1)
	util.LineContentsEqual(t, outputBuffer.String(),
		`blazegraph create: create dataset failed: dataset kb already exists
		`)
}

func TestBlazegraphCmd_create_custom_dataset(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --all --quiet")

	assertExitCode(t, "blazegraph create --dataset foo", 0)
	util.LineContentsEqual(t, outputBuffer.String(),
		`Successfully created dataset foo
		`)

	outputBuffer.Reset()
	assertExitCode(t, "blazegraph list", 0)
	util.LineContentsEqual(t, outputBuffer.String(),
		`foo
		`)
}

func TestBlazegraphCmd_create_custom_already_exists(t *testing.T) {
	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer
	run("blazegraph destroy --all --quiet")
	assertExitCode(t, "blazegraph create --dataset foo --quiet", 0)
}

func TestBlazegraphCmd_create_missing_dataset_name(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --all --quiet")

	assertExitCode(t, "blazegraph create --dataset", 1)
	util.LineContentsEqual(t, outputBuffer.String(),
		`blazegraph create: flag needs an argument: -dataset

		usage: blazegraph create [<flags>]

		flags:
          -dataset name
            	name of RDF dataset to create (default "kb")
          -infer string
            	Inference to perform on update [none, rdfs, owl] (default "none")
          -instance URL
            	URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
          -quiet
            	Discard normal command output
		-silent
			Discard normal and error command output

		`)
}

func TestBlazegraphCmd_create_dataset_name_without_flag(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --all --quiet")

	assertExitCode(t, "blazegraph create foo", 1)
	util.LineContentsEqual(t, outputBuffer.String(),
		`blazegraph create: unused argument: foo

		usage: blazegraph create [<flags>]

		flags:
		  -dataset name
				name of RDF dataset to create (default "kb")
		  -infer string
				Inference to perform on update [none, rdfs, owl] (default "none")
		  -instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
		  -quiet
				Discard normal command output
		  -silent
				Discard normal and error command output

		`)
}

var expectedCreateHelpOutput = string(
	`
	blazegraph create: Creates a new RDF dataset and corresponding Blazegraph namespace.

	usage: blazegraph create [<flags>]

	flags:
		-dataset name
				name of RDF dataset to create (default "kb")
		-infer string
				Inference to perform on update [none, rdfs, owl] (default "none")
		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
		-quiet
				Discard normal command output
		-silent
				Discard normal and error command output

	`)

func TestBlazegraphCmd_create_help(t *testing.T) {
	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer
	assertExitCode(t, "blazegraph create help", 0)
	util.LineContentsEqual(t, outputBuffer.String(), expectedCreateHelpOutput)
}

func TestBlazegraphCmd_help_create(t *testing.T) {
	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer
	assertExitCode(t, "blazegraph help create", 0)
	util.LineContentsEqual(t, outputBuffer.String(), expectedCreateHelpOutput)
}

func TestBlazegraphCmd_create_bad_flag(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	assertExitCode(t, "blazegraph create --not-a-flag", 1)

	util.LineContentsEqual(t, outputBuffer.String(),
		`blazegraph create: flag provided but not defined: -not-a-flag

		usage: blazegraph create [<flags>]

		flags:
			-dataset name
					name of RDF dataset to create (default "kb")
			-infer string
					Inference to perform on update [none, rdfs, owl] (default "none")
			-instance URL
					URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
			-quiet
				Discard normal command output
			-silent
				Discard normal and error command output

		`)
}
