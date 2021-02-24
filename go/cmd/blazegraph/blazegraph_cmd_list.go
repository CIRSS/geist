package main

import (
	"fmt"
)

func handleListSubcommand(cc *BGCommandContext) (err error) {
	count := cc.flags.String("count", "none", "Include count of triples in each dataset [none, estimate, exact]")
	if helpRequested(cc) {
		return
	}
	if err = cc.flags.Parse(cc.args[1:]); err != nil {
		showCommandUsage(cc)
		return
	}
	return doList(cc, *count)
}

func doList(cc *BGCommandContext, count string) (err error) {
	bc := cc.BlazegraphClient()
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
