package blazegraph

import (
	"fmt"

	"github.com/cirss/geist/cli"
)

func Import(cc *cli.CommandContext) (err error) {

	// declare command flags
	cc.Flags.String("dataset", "kb", "`name` of RDF dataset to import triples into")
	file := cc.Flags.String("file", "-", "File containing triples to import")
	format := cc.Flags.String("format", "ttl", "Format of triples to import [jsonld, nt, ttl, or xml]")

	// parse flags
	var helped bool
	if helped, err = cc.ParseFlags(); helped || err != nil {
		return
	}

	bc := cc.Resource("BlazegraphClient").(*BlazegraphClient)

	data, err := cc.ReadFileOrStdin(*file)
	if err != nil {
		fmt.Fprintf(cc.ErrWriter, err.Error())
		return
	}

	switch *format {

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
		fmt.Fprintf(cc.ErrWriter, err.Error())
	}
	return
}
