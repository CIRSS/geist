package main

import (
	"flag"
	"fmt"

	"github.com/cirss/geist/blazegraph"
)

func handleDestroySubcommand(args []string, flags *flag.FlagSet) (err error) {
	dataset := flags.String("dataset", "kb", "`name` of RDF dataset to destroy")
	all := flags.Bool("all", false, "destroy ALL datasets in the Blazegraph instance")
	if helpRequested(args, flags) {
		return
	}
	if err = flags.Parse(args[1:]); err != nil {
		showCommandUsage(args, flags)
		return
	}
	if len(*dataset) == 0 {
		fmt.Fprintln(errorMessageWriter, "name of dataset must be given using the -dataset flag")
		showCommandUsage(args, flags)
		return
	}
	if *all {
		return doDestroyAll()
	} else {
		return doDestroy(*dataset)
	}
	return
}

func doDestroyAll() (err error) {
	bc := blazegraph.NewBlazegraphClient(*context.instanceUrl)
	datasets, err := bc.ListDatasets()
	if err != nil {
		return
	}
	for _, dataset := range datasets {
		_, err = bc.DestroyDataSet(dataset)
		if err != nil {
			return
		}
	}
	return
}

func doDestroy(name string) (err error) {
	bc := context.blazegraphClient()
	_, err = bc.DestroyDataSet(name)
	return
}
