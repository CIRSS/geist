package blaze

import (
	"fmt"

	"github.com/cirss/geist/go/geist"
	"github.com/cirss/go-cli/pkg/cli"
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
		fmt.Fprintln(cc.ErrWriter, "name of dataset must be given using the -dataset flag")
		cc.ShowCommandUsage(cc.ErrWriter)
		return
	}

	p := NewDatasetProperties(*dataset)
	p.Inference = *infer

	bc := cc.Resource("BlazegraphClient").(*BlazegraphClient)

	_, err = bc.CreateDataSet(p)
	if err != nil {
		err = geist.NewGeistError("create dataset failed", err, true)
		return
	}

	fmt.Fprintf(cc.OutWriter, "Successfully created dataset %s\n", *dataset)

	return
}
