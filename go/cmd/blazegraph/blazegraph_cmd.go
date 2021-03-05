package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/cirss/geist/blazegraph"
	"github.com/cirss/geist/cli"
)

var Program *cli.ProgramContext

func init() {
	Program = cli.NewProgramContext("blazegraph", main)
}

func main() {

	commandCollection := cli.NewCommandCollection([]cli.CommandDescriptor{
		{"create", handleCreateSubcommand, "Create a new RDF dataset",
			"Creates an RDF dataset and corresponding Blazegraph namespace."},
		{"destroy", handleDestroySubcommand, "Delete an RDF dataset",
			"Deletes an RDF dataset and corresponding Blazegraph namespace, all RDF graphs\n" +
				"in the dataset, and all triples in each of those graphs."},
		{"export", handleExportSubcommand, "Export contents of a dataset",
			"Exports all triples in an RDF dataset in the requested format."},
		{"help", handleHelpSubcommand, "Show help", ""},
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

	cc := cli.NewCommandContext(commandCollection, Program)

	cc.Flags.String("instance", blazegraph.DefaultUrl, "`URL` of Blazegraph instance")

	if len(os.Args) < 2 {
		fmt.Fprint(Program.ErrWriter, "\nno blazegraph command given\n\n")
		cc.ShowProgramUsage()
		Program.ExitIfNonzero(1)
		return
	}

	commandName := os.Args[1]
	descriptor, exists := commandCollection.Lookup(commandName)
	cc.Descriptor = descriptor
	if !exists {
		fmt.Fprintf(Program.ErrWriter, "\nnot a blazegraph command: %s\n\n", commandName)
		cc.ShowProgramUsage()
		Program.ExitIfNonzero(1)
		return
	}

	cc.Args = os.Args[1:]
	err := cc.Descriptor.Handler(cc)
	if err != nil {
		Program.ExitIfNonzero(1)
		return
	}
}

func readFileOrStdin(filePath string) (bytes []byte, err error) {
	var r io.Reader
	if filePath == "-" {
		r = Program.InReader
	} else {
		r, _ = os.Open(filePath)
	}
	return ioutil.ReadAll(r)
}
