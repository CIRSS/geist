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

var subcommands map[string]func(flags *flag.FlagSet)

func init() {
	subcommands = map[string]func(flags *flag.FlagSet){
		"create":  handleCreateSubcommand,
		"destroy": handleDestroySubcommand,
		"export":  handleExportSubcommand,
		"import":  handleImportSubcommand,
		"report":  handleReportSubcommand,
		"select":  handleSelectSubcommand,
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(Main.ErrWriter, "Usage: geist <command> [args...]")
		return
	}

	flags := Main.InitFlagSet()

	command := os.Args[1]
	if commandHandler, exists := subcommands[command]; exists {
		commandHandler(flags)
	} else {
		fmt.Fprintf(Main.ErrWriter, "Unrecognized command: %s\n", command)
	}
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
