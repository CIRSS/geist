package main

import (
	"strings"
	"testing"

	"github.com/cirss/geist/util"
)

func TestBlazegraphCmd_destroy_help(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	assertExitCode(t, "blazegraph destroy help", 0)
	util.LineContentsEqual(t, outputBuffer.String(), `

		Deletes an RDF dataset and corresponding Blazegraph namespace, all RDF graphs
		in the dataset, and all triples in each of those graphs.

		Usage: blazegraph destroy [<flags>]

		Flags:

		-all
				destroy ALL datasets in the Blazegraph instance

		-dataset name
				name of RDF dataset to destroy (default "kb")

		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")

	`)
}

func TestBlazegraphCmd_help_destroy(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	assertExitCode(t, "blazegraph help destroy", 0)
	util.LineContentsEqual(t, outputBuffer.String(), `

		Deletes an RDF dataset and corresponding Blazegraph namespace, all RDF graphs
		in the dataset, and all triples in each of those graphs.

		Usage: blazegraph destroy [<flags>]

		Flags:

		-all
				destroy ALL datasets in the Blazegraph instance

		-dataset name
				name of RDF dataset to destroy (default "kb")

		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")

	`)
}

func TestBlazegraphCmd_destroy_no_dataset_argument(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	assertExitCode(t, "blazegraph destroy -dataset", 1)
	util.LineContentsEqual(t, outputBuffer.String(), `

		flag needs an argument: -dataset

		Usage: blazegraph destroy [<flags>]

		Flags:

		-all
				destroy ALL datasets in the Blazegraph instance

		-dataset name
				name of RDF dataset to destroy (default "kb")

		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")

	`)
}

func TestBlazegraphCmd_destroy_bad_flag(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	assertExitCode(t, "blazegraph destroy --not-a-flag", 1)
	util.LineContentsEqual(t, outputBuffer.String(), `

		flag provided but not defined: -not-a-flag

		Usage: blazegraph destroy [<flags>]

		Flags:

		-all
				destroy ALL datasets in the Blazegraph instance

		-dataset name
				name of RDF dataset to destroy (default "kb")

		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")

	`)
}
