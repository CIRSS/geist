package main

import (
	"fmt"
)

func handleDestroySubcommand(cc *Context) (err error) {
	dataset := cc.flags.String("dataset", "kb", "`name` of RDF dataset to destroy")
	all := cc.flags.Bool("all", false, "destroy ALL datasets in the Blazegraph instance")
	if helpRequested(cc) {
		return
	}
	if err = cc.flags.Parse(cc.args[1:]); err != nil {
		showCommandUsage(cc)
		return
	}
	if len(*dataset) == 0 {
		fmt.Fprintln(errorMessageWriter, "name of dataset must be given using the -dataset flag")
		showCommandUsage(cc)
		return
	}
	if *all {
		return doDestroyAll(cc)
	} else {
		return doDestroy(cc, *dataset)
	}
	return
}

func doDestroyAll(cc *Context) (err error) {
	bc := cc.BlazegraphClient()
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

func doDestroy(cc *Context, name string) (err error) {
	_, err = cc.BlazegraphClient().DestroyDataSet(name)
	return
}
