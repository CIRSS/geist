package main

import (
	"flag"
	"fmt"
)

func handleExportSubcommand(args []string, flags *flag.FlagSet) {
	flags.String("dataset", "kb", "`name` of RDF dataset to export")
	format := flags.String("format", "nt", "Format for exported triples [jsonld, nt, ttl, or xml]")
	sort := flags.Bool("sort", false, "Sort the exported triples if true")
	if helpRequested(args, flags) {
		return
	}
	if err := flags.Parse(args[1:]); err != nil {
		showCommandUsage(args, flags)
		return
	}
	doExport(*format, *sort)
}

func doExport(format string, sorted bool) {
	bc := context.blazegraphClient()
	var err error
	var triples string

	switch format {
	case "jsonld":
		triples, err = bc.ConstructAll("application/ld+json", sorted)
	case "nt":
		triples, err = bc.ConstructAll("text/plain", sorted)
		if sorted {
		}
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
}
