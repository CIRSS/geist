package main

import (
	"flag"
	"fmt"

	"github.com/cirss/geist/blazegraph"
)

func handleDestroySubcommand(args []string, flags *flag.FlagSet) {
	dataset := flags.String("dataset", "kb", "`name` of RDF dataset to destroy")
	all := flags.Bool("all", false, "destroy ALL datasets in the Blazegraph instance")
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
	if *all {
		doDestroyAll()
	} else {
		doDestroy(*dataset)
	}
}

func doDestroyAll() {
	bc := blazegraph.NewBlazegraphClient(*options.url)
	datasets, err := bc.ListDatasets()
	if err == nil {
		for _, dataset := range datasets {
			bc.DestroyDataSet(dataset)
		}
	}
}

func doDestroy(name string) {
	bc := blazegraph.NewBlazegraphClient(*options.url)
	bc.DestroyDataSet(name)
}
