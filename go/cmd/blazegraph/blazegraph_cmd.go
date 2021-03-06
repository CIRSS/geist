package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/cirss/geist/blazegraph"
	"github.com/cirss/geist/cli"
)

var Main *cli.ProgramContext

func init() {
	Main = cli.NewProgramContext("blazegraph", main)
}

func main() {

	commands := cli.NewCommandSet([]cli.CommandDescriptor{
		{"create", handleCreateSubcommand, "Create a new RDF dataset",
			"Creates an RDF dataset and corresponding Blazegraph namespace."},
		{"destroy", handleDestroySubcommand, "Delete an RDF dataset",
			"Deletes an RDF dataset and corresponding Blazegraph namespace, all RDF graphs\n" +
				"in the dataset, and all triples in each of those graphs."},
		{"export", handleExportSubcommand, "Export contents of a dataset",
			"Exports all triples in an RDF dataset in the requested format."},
		{"help", cli.HandleHelpSubcommand, "Show help", ""},
		{"import", handleImportSubcommand, "Import data into a dataset",
			"Imports triples in the specified format into an RDF dataset."},
		{"list", handleListSubcommand, "List RDF datasets",
			"Lists the names of the RDF datasets in the Blazegraph instance."},
		{"query", handleQuerySubcommand, "Perform a SPARQL query on a dataset",
			"Performs a SPARQL query on the identified RDF dataset."},
		{"report", handleReportSubcommand, "Expand a report using a dataset",
			"Expands the provided report template using the identified RDF dataset."},
		{"status", handleStatusSubcommand, "Check the status of the Blazegraph instance",
			"Requests the status of the Blazegraph instance, optionally waiting until\n" +
				"the instance is fully running. Returns status in JSON format."},
	})

	cc := Main.NewCommandContext(commands)

	cc.Flags.String("instance", blazegraph.DefaultUrl, "`URL` of Blazegraph instance")

	cc.InvokeCommand(os.Args)
}

func readFileOrStdin(filePath string) (bytes []byte, err error) {
	var r io.Reader
	if filePath == "-" {
		r = Main.InReader
	} else {
		r, _ = os.Open(filePath)
	}
	return ioutil.ReadAll(r)
}
