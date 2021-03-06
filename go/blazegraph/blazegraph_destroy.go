package blazegraph

import (
	"fmt"

	"github.com/cirss/geist/cli"
)

func Destroy(cc *cli.CommandContext) (err error) {

	// declare command flags
	dataset := cc.Flags.String("dataset", "kb", "`name` of RDF dataset to destroy")
	all := cc.Flags.Bool("all", false, "destroy ALL datasets in the Blazegraph instance")

	// parse flags
	var helped bool
	if helped, err = cc.ParseFlags(); helped || err != nil {
		return
	}

	if len(*dataset) == 0 {
		fmt.Fprintln(cc.ErrorMessageWriter, "name of dataset must be given using the -dataset flag")
		cc.ShowCommandUsage()
		return
	}

	bc := cc.Resource("BlazegraphClient").(*BlazegraphClient)

	if *all {
		var datasets []string
		datasets, err = bc.ListDatasets()
		if err != nil {
			return
		}
		for _, dataset := range datasets {
			_, err = bc.DestroyDataSet(dataset)
			if err != nil {
				return
			}
		}
	} else {
		_, err = bc.DestroyDataSet(*dataset)
	}
	return
}
