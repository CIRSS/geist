package main

import (
	"strings"
	"testing"

	"github.com/cirss/geist/util"
)

func TestBlazegraphCmd_create_help(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph create help")
	util.LineContentsEqual(t, outputBuffer.String(), `

		Creates an RDF dataset and corresponding Blazegraph namespace.

		Usage: blazegraph create <flags>

		Flags:

		-dataset string
				Dataset to create

		-infer string
				Inference to perform on update [none, rdfs, owl] (default "none")

		-url string
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")

	`)
}

func TestBlazegraphCmd_help_create(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph help create")
	util.LineContentsEqual(t, outputBuffer.String(), `

		Creates an RDF dataset and corresponding Blazegraph namespace.

		Usage: blazegraph create <flags>

		Flags:

		-dataset string
				Dataset to create

		-infer string
				Inference to perform on update [none, rdfs, owl] (default "none")

		-url string
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")

	`)
}

func TestBlazegraphCmd_create_no_flags(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph create")
	util.LineContentsEqual(t, outputBuffer.String(), `

		name of dataset must be given using the -dataset flag

		Usage: blazegraph create <flags>

		Flags:

		-dataset string
				Dataset to create

		-infer string
				Inference to perform on update [none, rdfs, owl] (default "none")

		-url string
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")

	`)
}

func TestBlazegraphCmd_create_bad_flag(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph create --not-a-flag")
	util.LineContentsEqual(t, outputBuffer.String(), `

		flag provided but not defined: -not-a-flag

		Usage: blazegraph create <flags>

		Flags:

		-dataset string
				Dataset to create

		-infer string
				Inference to perform on update [none, rdfs, owl] (default "none")

		-url string
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")

	`)
}
