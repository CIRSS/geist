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

	run("blazegraph destroy help")
	util.LineContentsEqual(t, outputBuffer.String(), `

		Deletes an RDF dataset and corresponding Blazegraph namespace, all RDF graphs
		in the dataset, and all triples in each of those graphs.

		Usage: blazegraph destroy <flags>

		Flags:

		-dataset name
	    	name of RDF dataset to destroy (required)

		-url string
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")

	`)
}

func TestBlazegraphCmd_help_destroy(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph help destroy")
	util.LineContentsEqual(t, outputBuffer.String(), `

		Deletes an RDF dataset and corresponding Blazegraph namespace, all RDF graphs
		in the dataset, and all triples in each of those graphs.

		Usage: blazegraph destroy <flags>

		Flags:

		-dataset name
	    	name of RDF dataset to destroy (required)

		-url string
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")

	`)
}

func TestBlazegraphCmd_destroy_no_flags(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy")
	util.LineContentsEqual(t, outputBuffer.String(), `

		name of dataset must be given using the -dataset flag

		Usage: blazegraph destroy <flags>

		Flags:

		-dataset name
	    	name of RDF dataset to destroy (required)

		-url string
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")

	`)
}

func TestBlazegraphCmd_destroy_bad_flag(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --not-a-flag")
	util.LineContentsEqual(t, outputBuffer.String(), `

		flag provided but not defined: -not-a-flag

		Usage: blazegraph destroy <flags>

		Flags:

		-dataset name
	    	name of RDF dataset to destroy (required)

		-url string
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")

	`)
}