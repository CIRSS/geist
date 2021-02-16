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

type subcommand struct {
	handler     func(args []string, flags *flag.FlagSet)
	description string
}

var subcommands map[string]subcommand

func init() {
	subcommands = map[string]subcommand{
		"help":    {handleHelpSubcommand, "Show help for a subcommand"},
		"create":  {handleCreateSubcommand, "Create a new dataset"},
		"destroy": {handleDestroySubcommand, "Destroy a dataset"},
		"export":  {handleExportSubcommand, "Export contents of a dataset"},
		"import":  {handleImportSubcommand, "Import data into a dataset"},
		"report":  {handleReportSubcommand, "Expand a report using a dataset"},
		"select":  {handleSelectSubcommand, "Perform a select query on a dataset"},
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(Main.ErrWriter, "Usage: geist <command> [args...]")
		return
	}

	flags := Main.InitFlagSet()

	command := os.Args[1]
	if subcommand, exists := subcommands[command]; exists {
		subcommand.handler(os.Args[1:], flags)
	} else {
		fmt.Fprintf(Main.ErrWriter, "Unrecognized command: %s\n", command)
	}
}

func handleHelpSubcommand(args []string, flags *flag.FlagSet) {

}

func addCommonOptions(flags *flag.FlagSet) {
	options.url = flags.String("url", blazegraph.DefaultUrl, "URL of Blazegraph instance")
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
