package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cirss/geist/blazegraph"
)

func handleDestroySubcommand(flags *flag.FlagSet) {
	addCommonOptions(flags)
	dataset := flags.String("dataset", "", "Dataset to destroy")
	if err := flags.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		flags.Usage()
		return
	}
	if len(*dataset) == 0 {
		fmt.Fprintln(Main.ErrWriter, "Error: Name of dataset to destroy not provided.")
		flags.Usage()
		return
	}
	doDestroy(*dataset)
}

func doDestroy(name string) {
	bc := blazegraph.NewBlazegraphClient(*options.url)
	bc.DestroyDataSet(name)
	// if err != nil {
	// 	fmt.Fprintln(Main.ErrWriter, err.Error())
	// }
	// fmt.Fprintln(Main.OutWriter, string(response))
}
