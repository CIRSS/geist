package blaze

import (
	"fmt"

	"github.com/cirss/go-cli/go/cli"
)

func List(cc *cli.CommandContext) (err error) {

	// declare command flags
	count := cc.Flags.String("count", "none", "Include count of triples in each dataset [none, estimate, exact]")

	// parse flags
	var helped bool
	if helped, err = cc.ParseFlags(); helped || err != nil {
		return
	}

	bc := cc.Resource("BlazegraphClient").(*BlazegraphClient)

	datasets, err := bc.ListDatasets()
	if err != nil {
		fmt.Fprintf(cc.ErrWriter, err.Error())
		return
	}

	for _, dataset := range datasets {
		if *count == "none" {
			fmt.Fprintf(cc.OutWriter, "%s\n", dataset)
		} else {
			count, _ := bc.CountTriples(dataset, false)
			fmt.Fprintf(cc.OutWriter, "%-10s %d\n", dataset, count)
		}
	}
	return
}
