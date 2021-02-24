package main

import (
	"fmt"

	"github.com/cirss/geist/blazegraph"
)

func handleCreateSubcommand(cc *BGCommandContext) (err error) {
	dataset := cc.flags.String("dataset", "kb", "`name` of RDF dataset to create")
	infer := cc.flags.String("infer", "none", "Inference to perform on update [none, rdfs, owl]")
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
	return doCreate(cc, *dataset, *infer)
}

func doCreate(cc *BGCommandContext, name string, infer string) (err error) {
	p := blazegraph.NewDatasetProperties(name)
	p.Inference = infer
	_, err = cc.BlazegraphClient().CreateDataSet(p)
	return
}
