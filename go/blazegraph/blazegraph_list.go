package blazegraph

import (
	"fmt"

	"github.com/cirss/geist/cli"
)

func List(cc *cli.CommandContext) (err error) {

	count := cc.Flags.String("count", "none", "Include count of triples in each dataset [none, estimate, exact]")

	var helped bool
	if helped, err = cc.ParseFlags(); helped || err != nil {
		return
	}

	return doList(cc, *count)
}

func doList(cc *cli.CommandContext, count string) (err error) {

	bc := cc.Resource("BlazegraphClient").(*BlazegraphClient)

	datasets, err := bc.ListDatasets()
	if err != nil {
		fmt.Fprintf(cc.ErrWriter, err.Error())
		return
	}

	for _, dataset := range datasets {
		if count == "none" {
			fmt.Fprintf(cc.OutWriter, "%s\n", dataset)
		} else {
			count, _ := bc.CountTriples(dataset, false)
			fmt.Fprintf(cc.OutWriter, "%-10s %d\n", dataset, count)
		}
	}
	return
}
