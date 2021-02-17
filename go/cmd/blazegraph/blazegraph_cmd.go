package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/cirss/geist/blazegraph"
	"github.com/tmcphillips/main-wrapper/mw"
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

var options struct {
	url *string
}

func init() {
	Main = mw.NewMainWrapper("blazegraph", main)
}

type command struct {
	name        string
	handler     func(args []string, flags *flag.FlagSet)
	summary     string
	description string
}

var commands []*command
var commandmap map[string]*command
var errorMessageWriter ErrorMessageWriter

func init() {

	commands = []*command{
		{"help", handleHelpSubcommand, "Show help", ""},
		{"list", handleListSubcommand, "List RDF datasets",
			"Lists the names of the RDF datasets in the Blazegraph instance."},
		{"create", handleCreateSubcommand, "Create a new RDF dataset",
			"Creates an RDF dataset and corresponding Blazegraph namespace."},
		{"destroy", handleDestroySubcommand, "Delete an RDF dataset",
			"Deletes an RDF dataset and corresponding Blazegraph namespace, all RDF graphs\n" +
				"in the dataset, and all triples in each of those graphs."},
		{"export", handleExportSubcommand, "Export contents of a dataset",
			"Exports all triples in an RDF dataset in the requested format."},
		{"import", handleImportSubcommand, "Import data into a dataset",
			"Imports triples in the specified format into an RDF dataset."},
		{"report", handleReportSubcommand, "Expand a report using a dataset",
			"Expands the provided report template using the identified RDF dataset."},
		{"select", handleSelectSubcommand, "Perform a select query on a dataset",
			"Performs a select query on the identified RDF dataset."},
	}

	commandmap = make(map[string]*command)

	for _, command := range commands {
		commandmap[command.name] = command
	}
}

func main() {

	errorMessageWriter.errorStream = Main.ErrWriter

	if len(os.Args) < 2 {
		fmt.Fprint(Main.ErrWriter, "\nno blazegraph command given\n\n")
		showProgramUsage()
		return
	}

	flags := Main.InitFlagSet()
	flags.Usage = func() {}
	flags.SetOutput(errorMessageWriter)
	options.url = flags.String("url", blazegraph.DefaultUrl, "URL of Blazegraph instance")
	command := os.Args[1]
	arguments := os.Args[1:]
	if c, exists := commandmap[command]; exists {
		c.handler(arguments, flags)
	} else {
		fmt.Fprintf(Main.ErrWriter, "\nnot a blazegraph command: %s\n\n", command)
		showProgramUsage()
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
