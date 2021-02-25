package main

import (
	"fmt"

	"github.com/cirss/geist"
	"github.com/cirss/geist/cli"
)

func handleReportSubcommand(cc *cli.CommandContext) (err error) {
	cc.Flags.String("dataset", "", "`name` of RDF dataset to create report from")
	file := cc.Flags.String("file", "-", "File containing report template to expand")
	if cc.ShowHelpIfRequested() {
		return
	}
	if err = cc.Flags.Parse(cc.Args[1:]); err != nil {
		cc.ShowCommandUsage()
		return
	}
	return doReport(cc, *file)
}

func doReport(cc *cli.CommandContext, file string) (err error) {
	bc := BlazegraphClient(cc)
	reportTemplate, err := readFileOrStdin(file)
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
