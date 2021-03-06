package blazegraph

import (
	"fmt"

	"github.com/cirss/geist/cli"
)

func Export(cc *cli.CommandContext) (err error) {

	cc.Flags.String("dataset", "kb", "`name` of RDF dataset to export")
	format := cc.Flags.String("format", "nt", "Format for exported triples [jsonld, nt, ttl, or xml]")
	sort := cc.Flags.Bool("sort", false, "Sort the exported triples if true")

	var helped bool
	if helped, err = cc.ParseFlags(); helped || err != nil {
		return
	}

	bc := cc.Resource("BlazegraphClient").(*BlazegraphClient)

	var triples string

	switch *format {
	case "jsonld":
		triples, err = bc.ConstructAll("application/ld+json", *sort)
	case "nt":
		triples, err = bc.ConstructAll("text/plain", *sort)
	case "ttl":
		triples, err = bc.ConstructAll("application/x-turtle", *sort)
	case "xml":
		triples, err = bc.ConstructAll("application/rdf+xml", *sort)
	}

	if err != nil {
		fmt.Fprintf(cc.ErrWriter, err.Error())
		return
	}

	fmt.Fprintf(cc.OutWriter, "%s", triples)
	if len(triples) > 0 {
		fmt.Fprintln(cc.OutWriter)

	}
	return
}
