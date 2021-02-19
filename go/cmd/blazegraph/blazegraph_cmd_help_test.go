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

		Usage: blazegraph <command> [<flags>]

		Commands:

		create   - Create a new RDF dataset
		destroy  - Delete an RDF dataset
		export   - Export contents of a dataset
		help     - Show help
		import   - Import data into a dataset
		list     - List RDF datasets
		query    - Perform a SPARQL query on a dataset
		report   - Expand a report using a dataset
		status   - Check the status of the Blazegraph instance

		Common flags:

		-instance URL
			  URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")

		See 'blazegraph help <command>' for help with one of the above commands.
	`)
}

func TestBlazegraphCmd_help_command_with_no_argument(t *testing.T) {

	var outputBuffer strings.Builder
	Main.OutWriter = &outputBuffer
	Main.ErrWriter = &outputBuffer

	run("blazegraph help")
	util.LineContentsEqual(t, outputBuffer.String(), `

		Usage: blazegraph <command> [<flags>]

		Commands:

		create   - Create a new RDF dataset
		destroy  - Delete an RDF dataset
		export   - Export contents of a dataset
		help     - Show help
		import   - Import data into a dataset
		list     - List RDF datasets
		query    - Perform a SPARQL query on a dataset
		report   - Expand a report using a dataset
		status   - Check the status of the Blazegraph instance

		Common flags:

		-instance URL
			  URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")

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

		Usage: blazegraph <command> [<flags>]

		Commands:

		create   - Create a new RDF dataset
		destroy  - Delete an RDF dataset
		export   - Export contents of a dataset
		help     - Show help
		import   - Import data into a dataset
		list     - List RDF datasets
		query    - Perform a SPARQL query on a dataset
		report   - Expand a report using a dataset
		status   - Check the status of the Blazegraph instance

		Common flags:

		-instance URL
			  URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")

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

		Usage: blazegraph <command> [<flags>]

		Commands:

		create   - Create a new RDF dataset
		destroy  - Delete an RDF dataset
		export   - Export contents of a dataset
		help     - Show help
		import   - Import data into a dataset
		list     - List RDF datasets
		query    - Perform a SPARQL query on a dataset
		report   - Expand a report using a dataset
		status   - Check the status of the Blazegraph instance

		Common flags:

		-instance URL
			  URL of Blazegraph instance (default "http://127.0.0.1:9999/blazegraph")

		See 'blazegraph help <command>' for help with one of the above commands.

	`)
}
