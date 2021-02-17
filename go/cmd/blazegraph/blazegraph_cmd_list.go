package main

import (
	"flag"
	"fmt"

	"github.com/cirss/geist/blazegraph"
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
	bc := blazegraph.NewBlazegraphClient(*options.url)
	list, re := bc.ListDatasets()
	if re != nil {
		fmt.Fprintf(Main.ErrWriter, re.Error())
		return
	}
	fmt.Fprint(Main.OutWriter, string(list))
}
