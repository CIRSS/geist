package main

import (
	"fmt"

	"github.com/cirss/geist/cli"
)

func handleExportSubcommand(cc *cli.CommandContext) (err error) {
	cc.Flags.String("dataset", "kb", "`name` of RDF dataset to export")
	format := cc.Flags.String("format", "nt", "Format for exported triples [jsonld, nt, ttl, or xml]")
	sort := cc.Flags.Bool("sort", false, "Sort the exported triples if true")
	if cc.ShowHelpIfRequested() {
		return
	}
	if err = cc.Flags.Parse(cc.Args[1:]); err != nil {
		cc.ShowCommandUsage()
		return
	}
	return doExport(cc, *format, *sort)
}

func doExport(cc *cli.CommandContext, format string, sorted bool) (err error) {
	bc := BlazegraphClient(cc)
	var triples string

	switch format {
	case "jsonld":
		triples, err = bc.ConstructAll("application/ld+json", sorted)
	case "nt":
		triples, err = bc.ConstructAll("text/plain", sorted)
	case "ttl":
		triples, err = bc.ConstructAll("application/x-turtle", sorted)
	case "xml":
		triples, err = bc.ConstructAll("application/rdf+xml", sorted)
	}

	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		return
	}

	fmt.Fprintf(Main.OutWriter, triples)
	return
}
