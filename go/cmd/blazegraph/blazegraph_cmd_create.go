package main

import (
	"flag"
	"fmt"

	"github.com/cirss/geist/blazegraph"
)

func handleCreateSubcommand(args []string, flags *flag.FlagSet) {
	flags.Usage = func() {}
	flags.SetOutput(errorMessageWriter)
	dataset := flags.String("dataset", "", "Dataset to create")
	infer := flags.String("infer", "none", "Inference to perform on update [none, rdfs, owl]")
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
	doCreate(*dataset, *infer)
}

func doCreate(name string, infer string) {
	bc := blazegraph.NewBlazegraphClient(*options.url)
	p := blazegraph.NewDatasetProperties(name)
	p.Inference = infer
	bc.CreateDataSet(p)
	// if err != nil {
	// 	fmt.Fprintln(Main.ErrWriter, err.Error())
	// }
	// fmt.Fprintln(Main.OutWriter, string(response))
}
