package blazegraph

import (
	"fmt"

	"github.com/cirss/geist/go/geist"
	"github.com/cirss/go-cli/go/cli"
)

func Query(cc *cli.CommandContext) (err error) {

	// declare command flags
	cc.Flags.String("dataset", "kb", "`name` of RDF dataset to query")
	dryrun := cc.Flags.Bool("dryrun", false, "Output query but do not execute it")
	file := cc.Flags.String("file", "-", "File containing the SPARQL query to execute")
	format := cc.Flags.String("format", "json", "Format of result set to produce [csv, json, table, or xml]")
	separators := cc.Flags.Bool("columnseparators", true, "Display column separators in table format")

	// parse flags
	var helped bool
	if helped, err = cc.ParseFlags(); helped || err != nil {
		return
	}

	bc := cc.Resource("BlazegraphClient").(*BlazegraphClient)

	queryText, err := cc.ReadFileOrStdin(*file)
	if err != nil {
		fmt.Fprintf(cc.ErrWriter, err.Error())
		return
	}

	queryTemplate := geist.NewTemplate("query", string(queryText), nil, bc)
	err = queryTemplate.Parse()
	if err != nil {
		fmt.Fprintf(cc.ErrWriter, "Error expanding query template:\n")
		fmt.Fprintf(cc.ErrWriter, "%s\n", err.Error())
		return
	}

	q, err := queryTemplate.Expand(nil)

	if err != nil {
		fmt.Fprintf(cc.ErrWriter, "Error expanding query template: ")
		fmt.Fprintf(cc.ErrWriter, "%s\n", err.Error())
		return
	}

	if *dryrun {
		fmt.Print(string(q))
		return
	}

	switch *format {

	case "csv":
		resultCSV, _ := bc.SelectCSV(string(q))
		if err != nil {
			break
		}
		fmt.Fprintf(cc.OutWriter, resultCSV)
		return

	case "json":
		rs, e := bc.Select(string(q))
		err = e
		if err != nil {
			break
		}
		resultJSON, _ := rs.JSONString()
		fmt.Fprintf(cc.OutWriter, resultJSON)
		return

	case "table":
		rs, e := bc.Select(string(q))
		err = e
		if err != nil {
			break
		}
		table := rs.FormattedTable(*separators)
		fmt.Fprintf(cc.OutWriter, table)
		return

	case "xml":
		resultXML, e := bc.SelectXML(string(q))
		err = e
		if err != nil {
			break
		}
		fmt.Fprintf(cc.OutWriter, resultXML)
		return
	}

	if err != nil {
		fmt.Fprintf(cc.ErrWriter, err.Error())
	}

	return

}
