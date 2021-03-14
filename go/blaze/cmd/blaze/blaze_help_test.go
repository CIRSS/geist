package main

import (
	"strings"
	"testing"

	"github.com/cirss/geist/go/util"
)

func TestBlazegraphCmd_no_command(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	assertExitCode(t, "blazegraph", 1)

	util.LineContentsEqual(t, outputBuffer.String(),
		`blaze: no command given

		usage: blaze <command> [<flags>]

		commands:
		create   - Create a new RDF dataset
		destroy  - Delete an RDF dataset
		export   - Export contents of a dataset
		help     - Show help
		import   - Import data into a dataset
		list     - List RDF datasets
		query    - Perform a SPARQL query on a dataset
		report   - Expand a report using a dataset
		status   - Check the status of the Blazegraph instance

		flags:
		-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
		-quiet
				Discard normal command output
		-silent
				Discard normal and error command output

		See 'blaze help <command>' for help with one of the above commands.

		`)
}

func TestBlazegraphCmd_help_command_with_no_argument(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	assertExitCode(t, "blaze help", 0)

	util.LineContentsEqual(t, outputBuffer.String(),
		`
		usage: blaze <command> [<flags>]

		commands:
			create   - Create a new RDF dataset
			destroy  - Delete an RDF dataset
			export   - Export contents of a dataset
			help     - Show help
			import   - Import data into a dataset
			list     - List RDF datasets
			query    - Perform a SPARQL query on a dataset
			report   - Expand a report using a dataset
			status   - Check the status of the Blazegraph instance

		flags:
			-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
			-quiet
				Discard normal command output
			-silent
				Discard normal and error command output

		See 'blaze help <command>' for help with one of the above commands.

		`)
}

func TestBlazegraphCmd_unsupported_command(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	assertExitCode(t, "blaze not-a-command", 1)

	util.LineContentsEqual(t, outputBuffer.String(),
		`blaze: unrecognized command: not-a-command

		usage: blaze <command> [<flags>]

		commands:
			create   - Create a new RDF dataset
			destroy  - Delete an RDF dataset
			export   - Export contents of a dataset
			help     - Show help
			import   - Import data into a dataset
			list     - List RDF datasets
			query    - Perform a SPARQL query on a dataset
			report   - Expand a report using a dataset
			status   - Check the status of the Blazegraph instance

		flags:
			-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
			-quiet
				Discard normal command output
			-silent
				Discard normal and error command output

		See 'blaze help <command>' for help with one of the above commands.

		`)
}

func TestBlazegraphCmd_help_unsupported_command(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	assertExitCode(t, "blaze help not-a-command", 1)

	util.LineContentsEqual(t, outputBuffer.String(),
		`blaze help: unrecognized blaze command: not-a-command

		usage: blaze <command> [<flags>]

		commands:
			create   - Create a new RDF dataset
			destroy  - Delete an RDF dataset
			export   - Export contents of a dataset
			help     - Show help
			import   - Import data into a dataset
			list     - List RDF datasets
			query    - Perform a SPARQL query on a dataset
			report   - Expand a report using a dataset
			status   - Check the status of the Blazegraph instance

		flags:
			-instance URL
				URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")
			-quiet
				Discard normal command output
			-silent
				Discard normal and error command output

		See 'blaze help <command>' for help with one of the above commands.

		`)
}
