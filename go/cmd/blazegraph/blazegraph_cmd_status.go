package main

import (
	"fmt"
)

func handleStatusSubcommand(cc *Context) (err error) {
	if helpRequested(cc) {
		return
	}
	if err = cc.flags.Parse(cc.args[1:]); err != nil {
		showCommandUsage(cc)
		return
	}
	return doStatus(cc)
}

func doStatus(cc *Context) (err error) {
	bc := cc.BlazegraphClient()
	status, err := bc.GetStatus()
	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		return
	}
	fmt.Fprintln(Main.OutWriter, status)
	return
}
