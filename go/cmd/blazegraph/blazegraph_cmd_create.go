package main

import (
	"fmt"

	"github.com/cirss/geist/blazegraph"
	"github.com/cirss/geist/cli"
)

func handleCreateSubcommand(cc *cli.CommandContext) (err error) {
	dataset := cc.Flags.String("dataset", "kb", "`name` of RDF dataset to create")
	infer := cc.Flags.String("infer", "none", "Inference to perform on update [none, rdfs, owl]")
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
	return doCreate(cc, *dataset, *infer)
}

func doCreate(cc *cli.CommandContext, name string, infer string) (err error) {
	p := blazegraph.NewDatasetProperties(name)
	p.Inference = infer
	_, err = BlazegraphClient(cc).CreateDataSet(p)
	return
}
