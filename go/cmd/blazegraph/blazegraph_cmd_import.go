package main

import (
	"fmt"

	"github.com/cirss/geist/cli"
)

func handleImportSubcommand(cc *cli.CommandContext) (err error) {
	cc.Flags.String("dataset", "kb", "`name` of RDF dataset to import triples into")
	file := cc.Flags.String("file", "-", "File containing triples to import")
	format := cc.Flags.String("format", "ttl", "Format of triples to import [jsonld, nt, ttl, or xml]")
	if cc.ShowHelpIfRequested() {
		return
	}
	if err = cc.Flags.Parse(cc.Args[1:]); err != nil {
		cc.ShowCommandUsage()
		return
	}
	return doImport(cc, *file, *format)
}

func doImport(cc *cli.CommandContext, file string, format string) (err error) {
	bc := BlazegraphClient(cc)
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
