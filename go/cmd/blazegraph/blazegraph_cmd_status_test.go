package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/cirss/geist/blazegraph"
	"github.com/cirss/geist/util"
)

func TestBlazegraphCmd_status(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	assertExitCode(t, "blazegraph status", 0)

	var status blazegraph.Status
	fmt.Println(outputBuffer.String())
	err := json.Unmarshal([]byte(outputBuffer.String()), &status)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	fmt.Printf("%v\n", status)
	util.StringEquals(t, status.InstanceUrl, "http://127.0.0.1:9999/blazegraph")
	util.StringEquals(t, status.SparqlEndpoint, "http://127.0.0.1:9999/blazegraph/sparql")
	util.StringEquals(t, status.BlazegraphBuildVersion, "2.1.5")
}

func TestBlazegraphCmd_status_help(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	assertExitCode(t, "blazegraph status help", 0)

	util.LineContentsEqual(t, outputBuffer.String(),
		`
		Requests the status of the Blazegraph instance, optionally waiting until
		the instance is fully running. Returns status in JSON format.

		Usage: blazegraph status [<flags>]

		Flags:

		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
		-quiet
				Discard normal command output
		-timeout milliseconds
				Number of milliseconds to wait for Blazegraph instance to respond

		`)
}

func TestBlazegraphCmd_help_status(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	assertExitCode(t, "blazegraph help status", 0)

	util.LineContentsEqual(t, outputBuffer.String(),
		`
		Requests the status of the Blazegraph instance, optionally waiting until
		the instance is fully running. Returns status in JSON format.

		Usage: blazegraph status [<flags>]

		Flags:

		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
		-quiet
				Discard normal command output
		-timeout milliseconds
				Number of milliseconds to wait for Blazegraph instance to respond

		`)
}

func TestBlazegraphCmd_status_bad_flag(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	assertExitCode(t, "blazegraph status --not-a-flag", 1)

	util.LineContentsEqual(t, outputBuffer.String(),
		`
		flag provided but not defined: -not-a-flag

		Usage: blazegraph status [<flags>]

		Flags:

		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
		-quiet
				Discard normal command output
		-timeout milliseconds
				Number of milliseconds to wait for Blazegraph instance to respond

		`)
}
