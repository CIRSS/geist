package main

import (
	"os"

	"github.com/cirss/geist/go/blazegraph"
	"github.com/cirss/go-cli/go/cli"
)

var Main *cli.ProgramContext

func init() {
	Main = cli.NewProgramContext("blazegraph", main)
}

func main() {

	commands := cli.NewCommandSet([]cli.CommandDescriptor{
		{"create", blazegraph.Create, "Create a new RDF dataset",
			"Creates a new RDF dataset and corresponding Blazegraph namespace."},
		{"destroy", blazegraph.Destroy, "Delete an RDF dataset",
			"Deletes an RDF dataset and corresponding Blazegraph namespace, all RDF graphs\n" +
				"in the dataset, and all triples in each of those graphs."},
		{"export", blazegraph.Export, "Export contents of a dataset",
			"Exports all triples in an RDF dataset in the requested format."},
		{"help", cli.Help, "Show help", ""},
		{"import", blazegraph.Import, "Import data into a dataset",
			"Imports triples in the specified format into an RDF dataset."},
		{"list", blazegraph.List, "List RDF datasets",
			"Lists the names of the RDF datasets in the Blazegraph instance."},
		{"query", blazegraph.Query, "Perform a SPARQL query on a dataset",
			"Performs a SPARQL query on the identified RDF dataset."},
		{"report", blazegraph.Report, "Expand a report using a dataset",
			"Expands the provided report template using the identified RDF dataset."},
		{"status", blazegraph.Status, "Check the status of the Blazegraph instance",
			"Requests the status of the Blazegraph instance, optionally waiting until\n" +
				"the instance is fully running. Returns status in JSON format."},
	})

	cc := Main.NewCommandContext(commands)
	cc.AddProvider("BlazegraphClient", getBlazegraphClient)

	cc.Flags.String("instance", blazegraph.DefaultUrl, "`URL` of Blazegraph instance")

	cc.InvokeCommand(os.Args)
}

func getBlazegraphClient(cc *cli.CommandContext) (bc interface{}) {
	bcc, exists := cc.Properties["blazegraph_client"]
	if exists {
		bc = bcc.(*blazegraph.BlazegraphClient)
	} else {
		instanceFlag := cc.Flags.Lookup("instance").Value.String()
		bc = blazegraph.NewBlazegraphClient(instanceFlag)
		cc.Properties["blazegraph_client"] = bc
	}
	return
}
