package blaze

import "github.com/cirss/go-cli/pkg/cli"

func NewBlazeCommandContext(pc *cli.ProgramContext) (cc *cli.CommandContext) {

	commands := cli.NewCommandSet([]cli.CommandDescriptor{
		{"create", Create, "Create a new RDF dataset",
			"Creates a new RDF dataset and corresponding Blazegraph namespace."},
		{"destroy", Destroy, "Delete an RDF dataset",
			"Deletes an RDF dataset and corresponding Blazegraph namespace, all RDF graphs\n" +
				"in the dataset, and all triples in each of those graphs."},
		{"export", Export, "Export contents of a dataset",
			"Exports all triples in an RDF dataset in the requested format."},
		{"help", cli.Help, "Show help", ""},
		{"import", Import, "Import data into a dataset",
			"Imports triples in the specified format into an RDF dataset."},
		{"list", List, "List RDF datasets",
			"Lists the names of the RDF datasets in the Blazegraph instance."},
		{"query", Query, "Perform a SPARQL query on a dataset",
			"Performs a SPARQL query on the identified RDF dataset."},
		{"report", Report, "Expand a report using a dataset",
			"Expands the provided report template using the identified RDF dataset."},
		{"status", Status, "Check the status of the Blazegraph instance",
			"Requests the status of the Blazegraph instance, optionally waiting until\n" +
				"the instance is fully running. Returns status in JSON format."},
	})

	cc = pc.NewCommandContext(commands)
	cc.AddProvider("BlazegraphClient", getBlazegraphClient)
	cc.Flags.String("instance", DefaultUrl, "`URL` of Blazegraph instance")

	return
}

func getBlazegraphClient(cc *cli.CommandContext) (bc interface{}) {
	bcc, exists := cc.Properties["blazegraph_client"]
	if exists {
		bc = bcc.(*BlazegraphClient)
	} else {
		instanceFlag := cc.Flags.Lookup("instance").Value.String()
		bc = NewBlazegraphClient(instanceFlag)
		cc.Properties["blazegraph_client"] = bc
	}
	return
}
