package main

import (
	"fmt"

	"github.com/cirss/geist/cli"
)

func handleDestroySubcommand(cc *cli.CommandContext) (err error) {
	dataset := cc.Flags.String("dataset", "kb", "`name` of RDF dataset to destroy")
	all := cc.Flags.Bool("all", false, "destroy ALL datasets in the Blazegraph instance")
	if cc.ShowHelpIfRequested() {
		return
	}
	if err = cc.Flags.Parse(cc.Args[1:]); err != nil {
		cc.ShowCommandUsage()
		return
	}
	if len(*dataset) == 0 {
		fmt.Fprintln(errorMessageWriter, "name of dataset must be given using the -dataset flag")
		cc.ShowCommandUsage()
		return
	}
	if *all {
		return doDestroyAll(cc)
	} else {
		return doDestroy(cc, *dataset)
	}
	return
}

func doDestroyAll(cc *cli.CommandContext) (err error) {
	bc := BlazegraphClient(cc)
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

func doDestroy(cc *cli.CommandContext, name string) (err error) {
	_, err = BlazegraphClient(cc).DestroyDataSet(name)
	return
}
