package main

import (
	"fmt"

	"github.com/cirss/geist"
)

func handleReportSubcommand(cc *BGCommandContext) (err error) {
	cc.flags.String("dataset", "", "`name` of RDF dataset to create report from")
	file := cc.flags.String("file", "-", "File containing report template to expand")
	if helpRequested(cc) {
		return
	}
	if err = cc.flags.Parse(cc.args[1:]); err != nil {
		showCommandUsage(cc)
		return
	}
	return doReport(cc, *file)
}

func doReport(cc *BGCommandContext, file string) (err error) {
	bc := cc.BlazegraphClient()
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
