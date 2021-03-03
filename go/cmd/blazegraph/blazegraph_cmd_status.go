package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/cirss/geist/cli"
)

const retryPeriod = 100

func handleStatusSubcommand(cc *cli.CommandContext) (err error) {

	timeout := cc.Flags.Int("timeout", 0, "Number of `milliseconds` to wait for Blazegraph instance to respond")

	if cc.ShowHelpIfRequested() {
		return
	}
	cc.Flags.Usage()

	if err = parseFlags(cc); err != nil {
		return
	}

	bc := BlazegraphClient(cc)

	var status string

	maxRetries := max(*timeout/retryPeriod, 1)
	var retries int
	for retries = 0; retries < maxRetries; retries++ {
		status, err = bc.GetStatus()
		if err != nil {
			fmt.Fprintln(cc.ErrWriter, err.Error())
			switch err.(type) {
			case *url.Error:
				time.Sleep(retryPeriod * time.Millisecond)
				continue

			default:
				return
			}
			return
		}
		break
	}

	if err != nil {
		if retries >= maxRetries {
			fmt.Fprintf(cc.ErrWriter, "Exceeded timeout connecting to Blazegraph instance\n")
		}
		return
	}

	fmt.Fprintln(cc.OutWriter, status)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
