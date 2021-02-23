package main

import (
	"flag"
	"fmt"
)

func handleImportSubcommand(args []string, flags *flag.FlagSet) (err error) {
	flags.String("dataset", "kb", "`name` of RDF dataset to import triples into")
	file := flags.String("file", "-", "File containing triples to import")
	format := flags.String("format", "ttl", "Format of triples to import [jsonld, nt, ttl, or xml]")
	if helpRequested(args, flags) {
		return
	}
	if err = flags.Parse(args[1:]); err != nil {
		showCommandUsage(args, flags)
		return
	}
	return doImport(*file, *format)
}

func doImport(file string, format string) (err error) {
	bc := context.blazegraphClient()
	data, err := readFileOrStdin(file)
	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		return
	}

	switch format {

	case "jsonld":
		_, err = bc.PostData("application/ld+json", data)

	case "nt":
		_, err = bc.PostData("text/plain", data)

	case "ttl":
		_, err = bc.PostData("application/x-turtle", data)

	case "xml":
		_, err = bc.PostData("application/rdf+xml", data)
	}

	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
	}
	return
}
