package main

import (
	"flag"
	"fmt"

	"github.com/cirss/geist"
)

func handleReportSubcommand(args []string, flags *flag.FlagSet) {
	flags.String("dataset", "", "`name` of RDF dataset to create report from")
	file := flags.String("file", "-", "File containing report template to expand")
	if helpRequested(args, flags) {
		return
	}
	if err := flags.Parse(args[1:]); err != nil {
		showCommandUsage(args, flags)
		return
	}
	doReport(*file)
}

func doReport(file string) {
	bc := context.blazegraphClient()
	reportTemplate, err := readFileOrStdin(file)
	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		return
	}
	rt := geist.NewTemplate("main", string(reportTemplate), nil, bc)
	report, re := bc.ExpandReport(rt)
	if re != nil {
		fmt.Fprintf(Main.ErrWriter, re.Error())
		return
	}
	fmt.Fprint(Main.OutWriter, report)
}
