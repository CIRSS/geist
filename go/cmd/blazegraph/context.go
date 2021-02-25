package main

import (
	"github.com/cirss/geist/blazegraph"
	"github.com/cirss/geist/cli"
)

func BlazegraphClient(cc *cli.CommandContext) (bc *blazegraph.BlazegraphClient) {
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
