package main

import (
	"fmt"

	"github.com/cirss/geist"
	"github.com/cirss/geist/cli"
)

func handleQuerySubcommand(cc *cli.CommandContext) (err error) {
	cc.Flags.String("dataset", "kb", "`name` of RDF dataset to query")
	dryrun := cc.Flags.Bool("dryrun", false, "Output query but do not execute it")
	file := cc.Flags.String("file", "-", "File containing the SPARQL query to execute")
	format := cc.Flags.String("format", "json", "Format of result set to produce [csv, json, table, or xml]")
	separators := cc.Flags.Bool("columnseparators", true, "Display column separators in table format")
	if cc.ShowHelpIfRequested() {
		return
	}
	if err = cc.Flags.Parse(cc.Args[1:]); err != nil {
		cc.ShowCommandUsage()
		return
	}
	return doSelectQuery(cc, *dryrun, *file, *format, *separators)
}

func doSelectQuery(cc *cli.CommandContext, dryrun bool, file string, format string, columnSeparators bool) (err error) {
	bc := BlazegraphClient(cc)
	queryText, err := readFileOrStdin(file)
	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		return
	}

	queryTemplate := geist.NewTemplate("query", string(queryText), nil, bc)
	err = queryTemplate.Parse()
	if err != nil {
		fmt.Fprintf(Main.ErrWriter, "Error expanding query template:\n")
		fmt.Fprintf(Main.ErrWriter, "%s\n", err.Error())
		return
	}

	q, err := queryTemplate.Expand(nil)

	if err != nil {
		fmt.Fprintf(Main.ErrWriter, "Error expanding query template: ")
		fmt.Fprintf(Main.ErrWriter, "%s\n", err.Error())
		return
	}

	if dryrun {
		fmt.Print(string(q))
		return
	}

	switch format {

	case "csv":
		resultCSV, _ := bc.SelectCSV(string(q))
		if err != nil {
			break
		}
		fmt.Fprintf(Main.OutWriter, resultCSV)
		return

	case "json":
		rs, e := bc.Select(string(q))
		err = e
		if err != nil {
			break
		}
		resultJSON, _ := rs.JSONString()
		fmt.Fprintf(Main.OutWriter, resultJSON)
		return

	case "table":
		rs, e := bc.Select(string(q))
		err = e
		if err != nil {
			break
		}
		table := rs.FormattedTable(columnSeparators)
		fmt.Fprintf(Main.OutWriter, table)
		return

	case "xml":
		resultXML, e := bc.SelectXML(string(q))
		err = e
		if err != nil {
			break
		}
		fmt.Fprintf(Main.OutWriter, resultXML)
		return
	}

	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
	}

	return

}