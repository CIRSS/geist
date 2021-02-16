package main

import (
	"flag"
	"fmt"

	"github.com/cirss/geist/blazegraph"
)

func handleExportSubcommand(args []string, flags *flag.FlagSet) {
	format := flags.String("format", "nt", "Format for doExported triples")
	sort := flags.Bool("sort", false, "Sort the exported triples if true")
	if err := flags.Parse(args[1:]); err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		flags.Usage()
		return
	}
	doExport(*format, *sort)
}

func doExport(format string, sorted bool) {
	bc := blazegraph.NewBlazegraphClient(*options.url)
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
