package main

import (
	"flag"
	"fmt"

	"github.com/cirss/geist/blazegraph"
)

func handleCreateSubcommand(args []string, flags *flag.FlagSet) (err error) {
	dataset := flags.String("dataset", "kb", "`name` of RDF dataset to create")
	infer := flags.String("infer", "none", "Inference to perform on update [none, rdfs, owl]")
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
	return doCreate(*dataset, *infer)
}

func doCreate(name string, infer string) (err error) {
	bc := context.blazegraphClient()
	p := blazegraph.NewDatasetProperties(name)
	p.Inference = infer
	_, err = bc.CreateDataSet(p)
	return
}
