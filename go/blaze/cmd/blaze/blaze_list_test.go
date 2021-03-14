package main

import (
	"strings"
	"testing"

	"github.com/cirss/geist/go/util"
)

func TestBlazegraphCmd_list_empty_store(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blaze destroy --all --quiet")

	assertExitCode(t, "blaze list", 0)

	util.LineContentsEqual(t, outputBuffer.String(), ``)
}

func TestBlazegraphCmd_list_default_dataset(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blaze destroy --all --quiet")
	run("blaze create --quiet")

	t.Run("no count", func(t *testing.T) {
		outputBuffer.Reset()
		assertExitCode(t, "blaze list", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`kb
			`)
	})

	t.Run("exact count", func(t *testing.T) {
		outputBuffer.Reset()
		assertExitCode(t, "blaze list --count exact", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`kb         0
			`)
	})

	t.Run("estimate count", func(t *testing.T) {
		outputBuffer.Reset()
		assertExitCode(t, "blaze list --count estimate", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`kb         0
			`)
	})
}

func TestBlazegraphCmd_list_custom_dataset(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blaze destroy --all --quiet")
	run("blaze create --quiet --dataset foo")

	assertExitCode(t, "blaze list", 0)

	util.LineContentsEqual(t, outputBuffer.String(),
		`foo
		`)
}

func TestBlazegraphCmd_list_custom_datasets(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blaze destroy --all --quiet")
	run("blaze create --quiet --dataset foo")
	run("blaze create --quiet --dataset bar")
	run("blaze create --quiet --dataset baz")

	t.Run("no count", func(t *testing.T) {
		outputBuffer.Reset()
		assertExitCode(t, "blaze list", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`bar
			 baz
			 foo
			`)
	})

	t.Run("exact count", func(t *testing.T) {
		outputBuffer.Reset()
		assertExitCode(t, "blaze list --count exact", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`bar        0
			 baz        0
			 foo        0
			`)
	})

	t.Run("estimate count", func(t *testing.T) {
		outputBuffer.Reset()
		assertExitCode(t, "blaze list --count estimate", 0)
		util.LineContentsEqual(t, outputBuffer.String(),
			`bar        0
			 baz        0
			 foo        0
			`)
	})
}

var expectedListHelpOutput = string(
	`
	blaze list: Lists the names of the RDF datasets in the Blazegraph instance.

	usage: blaze list [<flags>]

	flags:
		-count string
				Include count of triples in each dataset [none, estimate, exact] (default "none")
		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
		-quiet
				Discard normal command output
		-silent
				Discard normal and error command output

	`)

func TestBlazegraphCmd_list_help(t *testing.T) {
	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer
	assertExitCode(t, "blaze list help", 0)
	util.LineContentsEqual(t, outputBuffer.String(), expectedListHelpOutput)
}

func TestBlazegraphCmd_help_list(t *testing.T) {
	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer
	assertExitCode(t, "blaze help list", 0)
	util.LineContentsEqual(t, outputBuffer.String(), expectedListHelpOutput)
}

func TestBlazegraphCmd_list_bad_flag(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	assertExitCode(t, "blaze list --not-a-flag", 1)

	util.LineContentsEqual(t, outputBuffer.String(),
		`blaze list: flag provided but not defined: -not-a-flag

		usage: blaze list [<flags>]

		flags:
			-count string
					Include count of triples in each dataset [none, estimate, exact] (default "none")
			-instance URL
					URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
			-quiet
					Discard normal command output
			-silent
					Discard normal and error command output

		`)
}
