package main

import (
	"fmt"
)

func handleImportSubcommand(cc *BGCommandContext) (err error) {
	cc.flags.String("dataset", "kb", "`name` of RDF dataset to import triples into")
	file := cc.flags.String("file", "-", "File containing triples to import")
	format := cc.flags.String("format", "ttl", "Format of triples to import [jsonld, nt, ttl, or xml]")
	if helpRequested(cc) {
		return
	}
	if err = cc.flags.Parse(cc.args[1:]); err != nil {
		showCommandUsage(cc)
		return
	}
	return doImport(cc, *file, *format)
}

func doImport(cc *BGCommandContext, file string, format string) (err error) {
	bc := cc.BlazegraphClient()
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
