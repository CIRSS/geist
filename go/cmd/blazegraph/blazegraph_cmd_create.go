package main

import (
	"flag"
	"fmt"

	"github.com/cirss/geist/blazegraph"
)

func handleCreateSubcommand(args []string, flags *flag.FlagSet) {
	dataset := flags.String("dataset", "", "Dataset to create")
	infer := flags.String("infer", "none", "Inference to perform on update [none, rdfs, owl]")
	if err := flags.Parse(args[1:]); err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		flags.Usage()
		return
	}
	if len(*dataset) == 0 {
		fmt.Fprintln(Main.ErrWriter, "Error: Name of dataset to create not provided.")
		flags.Usage()
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
