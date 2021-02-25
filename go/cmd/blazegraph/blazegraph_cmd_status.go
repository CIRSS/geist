package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/cirss/geist/cli"
)

const retryPeriod = 100

func handleStatusSubcommand(cc *cli.CommandContext) (err error) {

	timeout := cc.Flags.Int("timeout", 0, "Total number of `milliseconds` to wait for Blazegraph instance to respond")
	if cc.ShowHelpIfRequested() {
		return
	}
	if err = cc.Flags.Parse(cc.Args[1:]); err != nil {
		cc.ShowCommandUsage()
		return
	}

	bc := BlazegraphClient(cc)

	var status string

	maxRetries := *timeout / retryPeriod

	var retries int
	for retries = 0; retries < maxRetries; retries++ {
		status, err = bc.GetStatus()
		if err != nil {
			fmt.Fprintln(Main.ErrWriter, err.Error())
			switch err.(type) {
			case *url.Error:
				time.Sleep(retryPeriod * time.Millisecond)
				continue

			default:
				return
			}
			return
		}
	}

	if err != nil {
		if retries >= maxRetries {
			fmt.Fprintf(Main.ErrWriter, "Exceeded timeout connecting to Blazegraph instance\n", retries)
		}
		return
	}

	fmt.Fprintln(Main.OutWriter, status)
	return
}
