package main

import (
	"fmt"

	"github.com/cirss/geist/cli"
)

func handleStatusSubcommand(cc *cli.CommandContext) (err error) {
	if cc.ShowHelpIfRequested() {
		return
	}
	if err = cc.Flags.Parse(cc.Args[1:]); err != nil {
		cc.ShowCommandUsage()
		return
	}
	return doStatus(cc)
}

func doStatus(cc *cli.CommandContext) (err error) {
	bc := BlazegraphClient(cc)
	status, err := bc.GetStatus()
	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		return
	}
	fmt.Fprintln(Main.OutWriter, status)
	return
}
