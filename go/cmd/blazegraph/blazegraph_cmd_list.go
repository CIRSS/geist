package main

import (
	"flag"
	"fmt"
)

func handleListSubcommand(args []string, flags *flag.FlagSet) (err error) {
	count := flags.String("count", "none", "Include count of triples in each dataset [none, estimate, exact]")
	if helpRequested(args, flags) {
		return
	}
	if err = flags.Parse(args[1:]); err != nil {
		showCommandUsage(args, flags)
		return
	}
	return doList(*count)
}

func doList(count string) (err error) {
	bc := context.blazegraphClient()
	datasets, err := bc.ListDatasets()
	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
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
	return
}
