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

// Main wraps the main() function.  It enables tests to manipulate the
// input and output streams used by main(), and provides a new FlagSet
// for each execution so that main() can be called by multiple tests.
var Main mw.MainWrapper

var options struct {
	url *string
}

func init() {
	Main = mw.NewMainWrapper("geist", main)
}

type command struct {
	name        string
	handler     func(args []string, flags *flag.FlagSet)
	description string
}

var commands []*command
var commandmap map[string]*command

func init() {

	commands = []*command{
		{"help", handleHelpSubcommand, "Show help"},
		{"create", handleCreateSubcommand, "Create a new dataset"},
		{"destroy", handleDestroySubcommand, "Destroy a dataset"},
		{"export", handleExportSubcommand, "Export contents of a dataset"},
		{"import", handleImportSubcommand, "Import data into a dataset"},
		{"report", handleReportSubcommand, "Expand a report using a dataset"},
		{"select", handleSelectSubcommand, "Perform a select query on a dataset"},
	}

	commandmap = make(map[string]*command)

	for _, command := range commands {
		commandmap[command.name] = command
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintf(Main.ErrWriter, "Error: no blazegraph command given\n")
		showUsage()
		return
	}

	flags := Main.InitFlagSet()
	options.url = flags.String("url", blazegraph.DefaultUrl, "URL of Blazegraph instance")
	command := os.Args[1]
	arguments := os.Args[1:]
	if c, exists := commandmap[command]; exists {
		c.handler(arguments, flags)
	} else {
		fmt.Fprintf(Main.ErrWriter, "Error: '%s' is not a blazegraph command\n", command)
		showUsage()
	}
}

func handleHelpSubcommand(args []string, flags *flag.FlagSet) {
	showUsage()
}

func showUsage() {
	fmt.Fprint(Main.OutWriter, "\nUsage: blazegraph <command> [<args>]\n\n")
	fmt.Fprint(Main.OutWriter, "Available commands:\n\n")
	for _, sc := range commands {
		fmt.Fprintf(Main.OutWriter, "%-7s - %s\n", sc.name, sc.description)
	}
	fmt.Fprint(Main.OutWriter, "\nSee 'blazegraph help <command>' for help with a specific command.\n\n")
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
