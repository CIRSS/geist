package main

import (
	"flag"
	"fmt"
)

func handleListSubcommand(args []string, flags *flag.FlagSet) {
	if helpRequested(args, flags) {
		return
	}
	if err := flags.Parse(args[1:]); err != nil {
		showCommandUsage(args, flags)
		return
	}
	doList()
}

func doList() {
	bc := context.blazegraphClient()
	datasets, re := bc.ListDatasets()
	if re != nil {
		fmt.Fprintf(Main.ErrWriter, re.Error())
		return
	}
	for _, dataset := range datasets {
		fmt.Fprintln(Main.OutWriter, dataset)
	}
}
