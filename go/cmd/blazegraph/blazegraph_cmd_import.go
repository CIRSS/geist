package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cirss/geist/blazegraph"
)

func handleImportSubcommand(flags *flag.FlagSet) {
	addCommonOptions(flags)
	file := flags.String("file", "-", "File containing triples to import")
	format := flags.String("format", "ttl", "Format of triples to import")
	if err := flags.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		flags.Usage()
		return
	}
	doImport(*file, *format)
}

func doImport(file string, format string) {
	bc := blazegraph.NewBlazegraphClient(*options.url)
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
		return
	}
}
