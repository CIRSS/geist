package main

import (
	"flag"
	"io"

	"github.com/cirss/geist/blazegraph"
)

type BGCommandContext struct {
	InReader    io.Reader
	OutWriter   io.Writer
	ErrWriter   io.Writer
	args        []string
	flags       *flag.FlagSet
	client      *blazegraph.BlazegraphClient
	instanceUrl *string
}

func (bc *BGCommandContext) BlazegraphClient() *blazegraph.BlazegraphClient {
	if bc.client == nil {
		bc.client = blazegraph.NewBlazegraphClient(*bc.instanceUrl)
	}
	return bc.client
}
