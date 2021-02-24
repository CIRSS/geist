package main

import (
	"fmt"
)

func handleExportSubcommand(cc *Context) (err error) {
	cc.flags.String("dataset", "kb", "`name` of RDF dataset to export")
	format := cc.flags.String("format", "nt", "Format for exported triples [jsonld, nt, ttl, or xml]")
	sort := cc.flags.Bool("sort", false, "Sort the exported triples if true")
	if helpRequested(cc) {
		return
	}
	if err = cc.flags.Parse(cc.args[1:]); err != nil {
		showCommandUsage(cc)
		return
	}
	return doExport(cc, *format, *sort)
}

func doExport(cc *Context, format string, sorted bool) (err error) {
	bc := cc.BlazegraphClient()
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
