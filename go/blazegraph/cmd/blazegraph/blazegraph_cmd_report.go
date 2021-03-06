package main

import (
	"fmt"

	"github.com/cirss/geist"
	"github.com/cirss/geist/cli"
)

func handleReportSubcommand(cc *cli.CommandContext) (err error) {
	cc.Flags.String("dataset", "", "`name` of RDF dataset to create report from")
	file := cc.Flags.String("file", "-", "File containing report template to expand")

	var helped bool
	if helped, err = cc.ParseFlags(); helped || err != nil {
		return
	}

	bc := BlazegraphClient(cc)

	reportTemplate, err := readFileOrStdin(*file)
	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		return
	}

	rt := geist.NewTemplate("main", string(reportTemplate), nil, bc)

	report, err := bc.ExpandReport(rt)
	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		return
	}

	fmt.Fprint(Main.OutWriter, report)
	return
}
