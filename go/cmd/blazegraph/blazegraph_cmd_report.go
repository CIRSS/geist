package main

import (
	"flag"
	"fmt"

	"github.com/cirss/geist"
)

func handleReportSubcommand(args []string, flags *flag.FlagSet) (err error) {
	flags.String("dataset", "", "`name` of RDF dataset to create report from")
	file := flags.String("file", "-", "File containing report template to expand")
	if helpRequested(args, flags) {
		return
	}
	if err = flags.Parse(args[1:]); err != nil {
		showCommandUsage(args, flags)
		return
	}
	return doReport(*file)
}

func doReport(file string) (err error) {
	bc := context.blazegraphClient()
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
