package main

import (
	"flag"
	"fmt"
)

func handleListSubcommand(args []string, flags *flag.FlagSet) {
	count := flags.String("count", "none", "Include count of triples in each dataset [none, estimate, exact]")
	if helpRequested(args, flags) {
		return
	}
	if err := flags.Parse(args[1:]); err != nil {
		showCommandUsage(args, flags)
		return
	}
	doList(*count)
}

func doList(count string) {
	bc := context.blazegraphClient()
	datasets, re := bc.ListDatasets()
	if re != nil {
		fmt.Fprintf(Main.ErrWriter, re.Error())
		return
	}
	for _, dataset := range datasets {
		if count == "none" {
			fmt.Fprintf(Main.OutWriter, "%s\n", dataset)
		} else {
			count, _ := bc.CountTriples(dataset, false)
			fmt.Fprintf(Main.OutWriter, "%-10s %d\n", dataset, count)
		}
	}
}
