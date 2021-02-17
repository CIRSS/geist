package main

import (
	"flag"
	"fmt"

	"github.com/cirss/geist/blazegraph"
)

func handleDestroySubcommand(args []string, flags *flag.FlagSet) {
	flags.Usage = func() {}
	flags.SetOutput(errorMessageWriter)
	dataset := flags.String("dataset", "", "`name` of RDF dataset to destroy (required)")
	if helpRequested(args, flags) {
		return
	}
	if err := flags.Parse(args[1:]); err != nil {
		showCommandUsage(args, flags)
		return
	}
	if len(*dataset) == 0 {
		fmt.Fprintln(errorMessageWriter, "name of dataset must be given using the -dataset flag")
		showCommandUsage(args, flags)
		return
	}
	doDestroy(*dataset)
}

func doDestroy(name string) {
	bc := blazegraph.NewBlazegraphClient(*options.url)
	bc.DestroyDataSet(name)
	// if err != nil {
	// 	fmt.Fprintln(Main.ErrWriter, err.Error())
	// }
	// fmt.Fprintln(Main.OutWriter, string(response))
}

func helpRequested(args []string, flags *flag.FlagSet) bool {
	if len(args) > 1 && args[1] == "help" {
		showCommandDescription()
		showCommandUsage(args, flags)
		return true
	}
	return false
}

func showCommandDescription() {
	fmt.Fprintf(Main.OutWriter,
		"\nDeletes an RDF dataset and corresponding Blazegraph namespace, all RDF graphs\n"+
			"in the dataset, and all triples in each of those graphs.\n")
}

func showCommandUsage(args []string, flags *flag.FlagSet) {
	fmt.Fprintf(Main.OutWriter, "\nUsage: blazegraph %s <flags>\n\n", args[0])
	fmt.Fprint(Main.OutWriter, "Flags:\n")
	flags.PrintDefaults()
	fmt.Fprintln(Main.OutWriter)
}
