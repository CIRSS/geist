package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/cirss/geist/blazegraph"
	"github.com/cirss/geist/util"
)

func TestBlazegraphCmd_status_success(t *testing.T) {

	var outBuffer strings.Builder
	var errBuffer strings.Builder
	Main.OutWriter = &outBuffer
	Main.ErrWriter = &errBuffer

	assertExitCode(t, "blazegraph status", 0)

	util.LineContentsEqual(t, errBuffer.String(), "")

	var status blazegraph.InstanceStatus
	fmt.Println(outBuffer.String())
	err := json.Unmarshal([]byte(outBuffer.String()), &status)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	fmt.Printf("%v\n", status)
	util.StringEquals(t, status.InstanceUrl, "http://127.0.0.1:9999/blazegraph")
	util.StringEquals(t, status.SparqlEndpoint, "http://127.0.0.1:9999/blazegraph/sparql")
	util.StringEquals(t, status.BlazegraphBuildVersion, "2.1.5")
}

func TestBlazegraphCmd_status_failure(t *testing.T) {

	var outBuffer strings.Builder
	var errBuffer strings.Builder
	Main.OutWriter = &outBuffer
	Main.ErrWriter = &errBuffer

	assertExitCode(t, "blazegraph status --instance http://not-a-blazegraph-instance", 1)

	util.LineContentsEqual(t, outBuffer.String(), "")
	util.LineContentsEqual(t, errBuffer.String(),
		`blazegraph status: Exceeded timeout connecting to Blazegraph instance
		`)
}

func TestBlazegraphCmd_status_quiet_success(t *testing.T) {

	var outBuffer strings.Builder
	var errBuffer strings.Builder
	Main.OutWriter = &outBuffer
	Main.ErrWriter = &errBuffer

	assertExitCode(t, "blazegraph status --quiet", 0)

	util.LineContentsEqual(t, outBuffer.String(), "")
	util.LineContentsEqual(t, errBuffer.String(), "")
}

func TestBlazegraphCmd_status_quiet_failure(t *testing.T) {

	var outBuffer strings.Builder
	var errBuffer strings.Builder
	Main.OutWriter = &outBuffer
	Main.ErrWriter = &errBuffer

	assertExitCode(t, "blazegraph status --quiet --instance http://not-a-blazegraph-instance", 1)

	util.LineContentsEqual(t, outBuffer.String(), "")
	util.LineContentsEqual(t, errBuffer.String(),
		`blazegraph status: Exceeded timeout connecting to Blazegraph instance
		`)
}

var expectedStatusHelpOutput = string(
	`
	blazegraph status: Requests the status of the Blazegraph instance, optionally waiting until
	the instance is fully running. Returns status in JSON format.

	usage: blazegraph status [<flags>]

	flags:
		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
		-quiet
				Discard normal command output
		-timeout milliseconds
				Number of milliseconds to wait for Blazegraph instance to respond

	`)

func TestBlazegraphCmd_status_help(t *testing.T) {
	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer
	assertExitCode(t, "blazegraph status help", 0)
	util.LineContentsEqual(t, outputBuffer.String(), expectedStatusHelpOutput)
}

func TestBlazegraphCmd_help_status(t *testing.T) {
	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer
	assertExitCode(t, "blazegraph help status", 0)
	util.LineContentsEqual(t, outputBuffer.String(), expectedStatusHelpOutput)
}

func TestBlazegraphCmd_status_bad_flag(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	assertExitCode(t, "blazegraph status --not-a-flag", 1)

	util.LineContentsEqual(t, outputBuffer.String(),
		`blazegraph status: flag provided but not defined: -not-a-flag

		usage: blazegraph status [<flags>]

		flags:
		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
		-quiet
				Discard normal command output
		-timeout milliseconds
				Number of milliseconds to wait for Blazegraph instance to respond

		`)
}
