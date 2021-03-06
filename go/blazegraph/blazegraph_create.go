package blazegraph

import (
	"fmt"

	"github.com/cirss/geist/cli"
)

func Create(cc *cli.CommandContext) (err error) {

	// declare command flags
	dataset := cc.Flags.String("dataset", "kb", "`name` of RDF dataset to create")
	infer := cc.Flags.String("infer", "none", "Inference to perform on update [none, rdfs, owl]")

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

	p := NewDatasetProperties(*dataset)
	p.Inference = *infer

	bc := cc.Resource("BlazegraphClient").(*BlazegraphClient)

	_, err = bc.CreateDataSet(p)
	return
}
