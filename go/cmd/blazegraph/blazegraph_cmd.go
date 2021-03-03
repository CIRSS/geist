package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/cirss/geist/blazegraph"
	"github.com/cirss/geist/cli"
)

type NullWriter struct {
	w io.Writer
}

func (nw NullWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

type ErrorMessageWriter struct {
	errorStream io.Writer
}

func (emw ErrorMessageWriter) Write(p []byte) (n int, err error) {
	fmt.Fprintln(emw.errorStream)
	return emw.errorStream.Write(p)
}

// Main wraps the main() function.  It enables tests to manipulate the
// input and output streams used by main(), and provides a new FlagSet
// for each execution so that main() can be called by multiple tests.
var Main cli.MainWrapper

func init() {
	Main = cli.NewMainWrapper("blazegraph", main)
}

var quiet *bool

var errorMessageWriter ErrorMessageWriter
var commandCollection *cli.CommandCollection

func init() {

	commandCollection = cli.NewCommandCollection([]cli.CommandDescriptor{
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
}

func main() {

	errorMessageWriter.errorStream = Main.ErrWriter

	cc := cli.NewCommandContext(commandCollection)

	cc.ErrWriter = Main.ErrWriter
	cc.OutWriter = Main.OutWriter
	cc.Flags = Main.InitFlagSet()
	cc.Flags.Usage = func() {}

	cc.Flags.String("instance", blazegraph.DefaultUrl, "`URL` of Blazegraph instance")
	quiet = cc.Flags.Bool("quiet", false, "Discard normal command output")

	if len(os.Args) < 2 {
		fmt.Fprint(Main.ErrWriter, "\nno blazegraph command given\n\n")
		cc.ShowProgramUsage()
		Main.ExitIfNonzero(1)
		return
	}

	commandName := os.Args[1]
	descriptor, exists := commandCollection.Lookup(commandName)
	cc.Descriptor = descriptor
	if !exists {
		fmt.Fprintf(Main.ErrWriter, "\nnot a blazegraph command: %s\n\n", commandName)
		cc.ShowProgramUsage()
		Main.ExitIfNonzero(1)
		return
	}

	cc.Args = os.Args[1:]
	err := cc.Descriptor.Handler(cc)
	if err != nil {
		Main.ExitIfNonzero(1)
		return
	}
}

func parseFlags(cc *cli.CommandContext) (err error) {

	cc.Flags.SetOutput(errorMessageWriter)
	if err = cc.Flags.Parse(cc.Args[1:]); err != nil {
		cc.Flags.SetOutput(cc.ErrWriter)
		cc.ShowCommandUsage()
		return
	}
	cc.Flags.SetOutput(cc.ErrWriter)

	if *quiet {
		cc.OutWriter = NullWriter{}
	}

	return
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
