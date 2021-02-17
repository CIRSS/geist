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
		no blazegraph command given

		Usage: blazegraph <command> [<args>]

		Available commands:

		help     - Show help
		create   - Create a new RDF dataset
		destroy  - Delete an RDF dataset
		export   - Export contents of a dataset
		import   - Import data into a dataset
		report   - Expand a report using a dataset
		select   - Perform a select query on a dataset

		See 'blazegraph help <command>' for help with one of the above commands.
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

		help     - Show help
		create   - Create a new RDF dataset
		destroy  - Delete an RDF dataset
		export   - Export contents of a dataset
		import   - Import data into a dataset
		report   - Expand a report using a dataset
		select   - Perform a select query on a dataset

		See 'blazegraph help <command>' for help with one of the above commands.
	`)
}

func TestBlazegraphCmd_unsupported_command(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph not-a-command")
	util.LineContentsEqual(t, outputBuffer.String(), `
		not a blazegraph command: not-a-command

		Usage: blazegraph <command> [<args>]

		Available commands:

		help     - Show help
		create   - Create a new RDF dataset
		destroy  - Delete an RDF dataset
		export   - Export contents of a dataset
		import   - Import data into a dataset
		report   - Expand a report using a dataset
		select   - Perform a select query on a dataset

		See 'blazegraph help <command>' for help with one of the above commands.
	`)
}

func TestBlazegraphCmd_help_unsupported_command(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph help not-a-command")
	util.LineContentsEqual(t, outputBuffer.String(), `
		not a blazegraph command: not-a-command

		Usage: blazegraph <command> [<args>]

		Available commands:

		help     - Show help
		create   - Create a new RDF dataset
		destroy  - Delete an RDF dataset
		export   - Export contents of a dataset
		import   - Import data into a dataset
		report   - Expand a report using a dataset
		select   - Perform a select query on a dataset

		See 'blazegraph help <command>' for help with one of the above commands.
	`)
}
