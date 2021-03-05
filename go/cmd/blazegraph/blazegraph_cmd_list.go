package main

import (
	"fmt"

	"github.com/cirss/geist/cli"
)

func handleListSubcommand(cc *cli.CommandContext) (err error) {

	count := cc.Flags.String("count", "none", "Include count of triples in each dataset [none, estimate, exact]")

	var helped bool
	if helped, err = cc.ParseFlags(); helped || err != nil {
		return
	}

	return doList(cc, *count)
}

func doList(cc *cli.CommandContext, count string) (err error) {
	bc := BlazegraphClient(cc)
	datasets, err := bc.ListDatasets()
	if err != nil {
		fmt.Fprintf(Program.ErrWriter, err.Error())
		return
	}

	for _, dataset := range datasets {
		if count == "none" {
			fmt.Fprintf(Program.OutWriter, "%s\n", dataset)
		} else {
			count, _ := bc.CountTriples(dataset, false)
			fmt.Fprintf(Program.OutWriter, "%-10s %d\n", dataset, count)
		}
	}
	return
}
