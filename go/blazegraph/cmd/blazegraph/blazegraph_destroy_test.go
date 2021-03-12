package main

import (
	"strings"
	"testing"

	"github.com/cirss/geist/go/util"
)

func TestBlazegraphCmd_destroy_default_dataset(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --all --quiet")
	run("blazegraph create --quiet")

	assertExitCode(t, "blazegraph destroy", 0)
	util.LineContentsEqual(t, outputBuffer.String(),
		`Successfully destroyed dataset kb
		`)
}

func TestBlazegraphCmd_destroy_default_dataset_quiet(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --all --quiet")
	run("blazegraph create --quiet")

	assertExitCode(t, "blazegraph destroy --quiet", 0)
	util.LineContentsEqual(t, outputBuffer.String(),
		``)
}

func TestBlazegraphCmd_destroy_nonexistent_default_dataset(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --all --quiet")

	assertExitCode(t, "blazegraph destroy", 1)
	util.LineContentsEqual(t, outputBuffer.String(),
		`blazegraph destroy: destroy dataset failed: dataset kb does not exist
		`)
}

func TestBlazegraphCmd_destroy_nonexistent_default_dataset_quiet(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --all --quiet")

	assertExitCode(t, "blazegraph destroy --quiet", 1)
	util.LineContentsEqual(t, outputBuffer.String(),
		`blazegraph destroy: destroy dataset failed: dataset kb does not exist
		`)
}

func TestBlazegraphCmd_destroy_custom_dataset(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --all --quiet")
	run("blazegraph create --dataset foo --quiet")

	assertExitCode(t, "blazegraph destroy --dataset foo", 0)
	util.LineContentsEqual(t, outputBuffer.String(),
		`Successfully destroyed dataset foo
		`)
}

func TestBlazegraphCmd_destroy_nonexistent_custom_dataset(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --all --quiet")

	assertExitCode(t, "blazegraph destroy --dataset foo", 1)
	util.LineContentsEqual(t, outputBuffer.String(),
		`blazegraph destroy: destroy dataset failed: dataset foo does not exist
		`)
}

func TestBlazegraphCmd_destroy_all_of_several_datasets(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --all --quiet")
	run("blazegraph create --dataset foo --quiet")
	run("blazegraph create --dataset bar --quiet")
	run("blazegraph create --dataset baz --quiet")

	assertExitCode(t, "blazegraph destroy --all", 0)
	util.LineContentsEqual(t, outputBuffer.String(),
		`Successfully destroyed dataset bar
		 Successfully destroyed dataset baz
		 Successfully destroyed dataset foo
		`)
}

func TestBlazegraphCmd_destroy_one_of_several_datasets(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --all --quiet")
	run("blazegraph create --dataset foo --quiet")
	run("blazegraph create --dataset bar --quiet")
	run("blazegraph create --dataset baz --quiet")

	assertExitCode(t, "blazegraph destroy --dataset bar", 0)
	util.LineContentsEqual(t, outputBuffer.String(),
		`Successfully destroyed dataset bar
		`)

	outputBuffer.Reset()
	run("blazegraph list")
	util.LineContentsEqual(t, outputBuffer.String(),
		`baz
		 foo
		`)
}

func TestBlazegraphCmd_destroy_missing_dataset_name(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	assertExitCode(t, "blazegraph destroy --dataset", 1)
	util.LineContentsEqual(t, outputBuffer.String(),
		`blazegraph destroy: flag needs an argument: -dataset

        usage: blazegraph destroy [<flags>]

        flags:
          -all
            	destroy ALL datasets in the Blazegraph instance
          -dataset name
            	name of RDF dataset to destroy (default "kb")
          -instance URL
            	URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
          -quiet
            	Discard normal command output

		`)
}

var expectedDestroyHelpOutput = string(
	`
	blazegraph destroy: Deletes an RDF dataset and corresponding Blazegraph namespace, all RDF graphs
        in the dataset, and all triples in each of those graphs.

        usage: blazegraph destroy [<flags>]

        flags:
          -all
            	destroy ALL datasets in the Blazegraph instance
          -dataset name
            	name of RDF dataset to destroy (default "kb")
          -instance URL
            	URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
          -quiet
            	Discard normal command output

		`)

func TestBlazegraphCmd_destroy_help(t *testing.T) {
	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer
	assertExitCode(t, "blazegraph destroy help", 0)
	util.LineContentsEqual(t, outputBuffer.String(), expectedDestroyHelpOutput)
}

func TestBlazegraphCmd_help_destroy(t *testing.T) {
	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer
	assertExitCode(t, "blazegraph help destroy", 0)
	util.LineContentsEqual(t, outputBuffer.String(), expectedDestroyHelpOutput)
}

func TestBlazegraphCmd_destroy_no_dataset_argument(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	assertExitCode(t, "blazegraph destroy -dataset", 1)
	util.LineContentsEqual(t, outputBuffer.String(),
		`blazegraph destroy: flag needs an argument: -dataset

		usage: blazegraph destroy [<flags>]

		flags:
			-all
					destroy ALL datasets in the Blazegraph instance
			-dataset name
					name of RDF dataset to destroy (default "kb")
			-instance URL
					URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
			-quiet
					Discard normal command output

	`)
}

func TestBlazegraphCmd_destroy_bad_flag(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	assertExitCode(t, "blazegraph destroy --not-a-flag", 1)
	util.LineContentsEqual(t, outputBuffer.String(),
		`blazegraph destroy: flag provided but not defined: -not-a-flag

		usage: blazegraph destroy [<flags>]

		flags:
			-all
					destroy ALL datasets in the Blazegraph instance
			-dataset name
					name of RDF dataset to destroy (default "kb")
			-instance URL
					URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
			-quiet
					Discard normal command output

	`)
}
