package main

import (
	"strings"
	"testing"

	"github.com/cirss/geist/util"
)

func TestBlazegraphCmd_create_help(t *testing.T) {

	var outputBuffer strings.Builder
	Program.OutWriter = &outputBuffer
	Program.ErrWriter = &outputBuffer

	assertExitCode(t, "blazegraph create help", 0)

	util.LineContentsEqual(t, outputBuffer.String(),
		`
		Creates an RDF dataset and corresponding Blazegraph namespace.

		Usage: blazegraph create [<flags>]

		Flags:

		-dataset name
				name of RDF dataset to create (default "kb")
		-infer string
				Inference to perform on update [none, rdfs, owl] (default "none")
		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
		-quiet
            	Discard normal command output

		`)
}

func TestBlazegraphCmd_help_create(t *testing.T) {

	var outputBuffer strings.Builder
	Program.OutWriter = &outputBuffer
	Program.ErrWriter = &outputBuffer

	assertExitCode(t, "blazegraph help create", 0)

	util.LineContentsEqual(t, outputBuffer.String(),
		`
		Creates an RDF dataset and corresponding Blazegraph namespace.

		Usage: blazegraph create [<flags>]

		Flags:

		-dataset name
				name of RDF dataset to create (default "kb")
		-infer string
				Inference to perform on update [none, rdfs, owl] (default "none")
		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
		-quiet
            	Discard normal command output

	`)
}

func TestBlazegraphCmd_create_bad_flag(t *testing.T) {

	var outputBuffer strings.Builder
	Program.OutWriter = &outputBuffer
	Program.ErrWriter = &outputBuffer

	assertExitCode(t, "blazegraph create --not-a-flag", 1)

	util.LineContentsEqual(t, outputBuffer.String(),
		`
		flag provided but not defined: -not-a-flag

		Usage: blazegraph create [<flags>]

		Flags:

		-dataset name
				name of RDF dataset to create (default "kb")
		-infer string
				Inference to perform on update [none, rdfs, owl] (default "none")
		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
		-quiet
			Discard normal command output

		`)
}
