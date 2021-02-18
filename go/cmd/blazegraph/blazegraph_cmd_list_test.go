package main

import (
	"strings"
	"testing"

	"github.com/cirss/geist/util"
)

func TestBlazegraphCmd_list_empty_store(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --all")
	run("blazegraph list")

	util.LineContentsEqual(t, outputBuffer.String(), ``)
}

func TestBlazegraphCmd_list_default_dataset(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --all")
	run("blazegraph create")
	run("blazegraph list")

	util.LineContentsEqual(t, outputBuffer.String(), "kb")
}

func TestBlazegraphCmd_list_custom_dataset(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --all")
	run("blazegraph create --dataset foo")
	run("blazegraph list")

	util.LineContentsEqual(t, outputBuffer.String(), "foo")
}

func TestBlazegraphCmd_list_custom_datasets(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph destroy --all")
	run("blazegraph create --dataset foo")
	run("blazegraph create --dataset bar")
	run("blazegraph create --dataset baz")
	run("blazegraph list")

	util.LineContentsEqual(t, outputBuffer.String(), `
		bar
		baz
		foo
	`)
}
func TestBlazegraphCmd_list_help(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph list help")
	util.LineContentsEqual(t, outputBuffer.String(), `

		Lists the names of the RDF datasets in the Blazegraph instance.

		Usage: blazegraph list [<flags>]

		Flags:

		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")

	`)
}

func TestBlazegraphCmd_help_list(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph help list")
	util.LineContentsEqual(t, outputBuffer.String(), `

	Lists the names of the RDF datasets in the Blazegraph instance.

	Usage: blazegraph list [<flags>]

	Flags:

	-instance URL
			URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")

	`)
}

func TestBlazegraphCmd_list_bad_flag(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph list --not-a-flag")
	util.LineContentsEqual(t, outputBuffer.String(), `

		flag provided but not defined: -not-a-flag

		Usage: blazegraph list [<flags>]

		Flags:

		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")

	`)
}
