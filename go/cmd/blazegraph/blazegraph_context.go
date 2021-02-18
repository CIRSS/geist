package main

import "github.com/cirss/geist/blazegraph"

type BlazegraphContext struct {
	client      *blazegraph.BlazegraphClient
	instanceUrl *string
}

func (bc *BlazegraphContext) blazegraphClient() *blazegraph.BlazegraphClient {
	if bc.client == nil {
		bc.client = blazegraph.NewBlazegraphClient(*context.instanceUrl)
	}
	return bc.client
}
