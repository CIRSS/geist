package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/cirss/geist/blazegraph"
	"github.com/cirss/geist/mw"
)

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
var Main mw.MainWrapper

func init() {
	Main = mw.NewMainWrapper("blazegraph", main)
}

var errorMessageWriter ErrorMessageWriter
var commandCollection *CommandCollection

func init() {

	commandCollection = NewCommandCollection([]CommandDescriptor{
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

	c := new(Context)

	c.ErrWriter = Main.ErrWriter
	c.OutWriter = Main.OutWriter
	c.flags = Main.InitFlagSet()
	c.flags.Usage = func() {}
	c.flags.SetOutput(errorMessageWriter)

	c.instanceUrl = c.flags.String("instance", blazegraph.DefaultUrl, "`URL` of Blazegraph instance")

	if len(os.Args) < 2 {
		fmt.Fprint(Main.ErrWriter, "\nno blazegraph command given\n\n")
		showProgramUsage(c.flags)
		Main.ExitIfNonzero(1)
		return
	}

	commandName := os.Args[1]
	commandDescriptor, exists := commandCollection.commandMap[commandName]
	if !exists {
		fmt.Fprintf(Main.ErrWriter, "\nnot a blazegraph command: %s\n\n", commandName)
		showProgramUsage(c.flags)
		Main.ExitIfNonzero(1)
		return
	}

	c.args = os.Args[1:]
	err := commandDescriptor.handler(c)
	if err != nil {
		Main.ExitIfNonzero(1)
		return
	}
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

func showProgramUsage(flags *flag.FlagSet) {
	fmt.Fprint(Main.OutWriter, "Usage: blazegraph <command> [<flags>]\n\n")
	fmt.Fprint(Main.OutWriter, "Commands:\n\n")
	for _, sc := range commandCollection.commandList {
		fmt.Fprintf(Main.OutWriter, "  %-7s  - %s\n", sc.name, sc.summary)
	}
	fmt.Fprint(Main.OutWriter, "\nCommon flags:\n")
	flags.PrintDefaults()
	fmt.Fprint(Main.OutWriter, "\nSee 'blazegraph help <command>' for help with one of the above commands.\n\n")
	return
}

func showCommandDescription(c *CommandDescriptor) {
	fmt.Fprintf(Main.OutWriter, "\n%s\n", c.description)
}

func showCommandUsage(cc *Context) {
	fmt.Fprintf(Main.OutWriter, "\nUsage: blazegraph %s [<flags>]\n\n", cc.args[0])
	fmt.Fprint(Main.OutWriter, "Flags:\n")
	cc.flags.PrintDefaults()
	fmt.Fprintln(Main.OutWriter)
}

func helpRequested(cc *Context) bool {
	if len(cc.args) > 1 && cc.args[1] == "help" {
		showCommandDescription(commandCollection.commandMap[cc.args[0]])
		showCommandUsage(cc)
		return true
	}
	return false
}
