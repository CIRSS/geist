package main

import (
	"flag"
	"fmt"
)

func handleStatusSubcommand(args []string, flags *flag.FlagSet) {
	if helpRequested(args, flags) {
		return
	}
	if err := flags.Parse(args[1:]); err != nil {
		showCommandUsage(args, flags)
		return
	}
	doStatus()
}

func doStatus() {
	bc := context.blazegraphClient()
	status, err := bc.GetStatus()
	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		return
	}
	fmt.Fprintln(Main.OutWriter, status)
}
