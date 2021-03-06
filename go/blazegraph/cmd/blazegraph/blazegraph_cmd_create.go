package main

import (
	"fmt"

	"github.com/cirss/geist/blazegraph"
	"github.com/cirss/geist/cli"
)

func handleCreateSubcommand(cc *cli.CommandContext) (err error) {

	dataset := cc.Flags.String("dataset", "kb", "`name` of RDF dataset to create")
	infer := cc.Flags.String("infer", "none", "Inference to perform on update [none, rdfs, owl]")

	var helped bool
	if helped, err = cc.ParseFlags(); helped || err != nil {
		return
	}

	if len(*dataset) == 0 {
		fmt.Fprintln(cc.ErrorMessageWriter, "name of dataset must be given using the -dataset flag")
		cc.ShowCommandUsage()
		return
	}

	p := blazegraph.NewDatasetProperties(*dataset)
	p.Inference = *infer
	_, err = BlazegraphClient(cc).CreateDataSet(p)
	return
}
