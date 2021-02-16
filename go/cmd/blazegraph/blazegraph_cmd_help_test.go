package main

import (
	"strings"
	"testing"

	"github.com/cirss/geist/util"
)

func TestBlazegraphCmd_no_command(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph")
	util.LineContentsEqual(t, outputBuffer.String(), `
		Error: no blazegraph command given

		Usage: blazegraph <command> [<args>]

		Available commands:

		help    - Show help
		create  - Create a new dataset
		destroy - Destroy a dataset
		export  - Export contents of a dataset
		import  - Import data into a dataset
		report  - Expand a report using a dataset
		select  - Perform a select query on a dataset

		See 'blazegraph help <command>' for help with a specific command.
	`)
}

func TestBlazegraphCmd_help_command_with_no_argument(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph help")
	util.LineContentsEqual(t, outputBuffer.String(), `

		Usage: blazegraph <command> [<args>]

		Available commands:

		help    - Show help
		create  - Create a new dataset
		destroy - Destroy a dataset
		export  - Export contents of a dataset
		import  - Import data into a dataset
		report  - Expand a report using a dataset
		select  - Perform a select query on a dataset

		See 'blazegraph help <command>' for help with a specific command.
	`)
}

func TestBlazegraphCmd_unsupported_command(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph not-a-command")
	util.LineContentsEqual(t, outputBuffer.String(), `
		Error: 'not-a-command' is not a blazegraph command

		Usage: blazegraph <command> [<args>]

		Available commands:

		help    - Show help
		create  - Create a new dataset
		destroy - Destroy a dataset
		export  - Export contents of a dataset
		import  - Import data into a dataset
		report  - Expand a report using a dataset
		select  - Perform a select query on a dataset

		See 'blazegraph help <command>' for help with a specific command.
	`)
}
