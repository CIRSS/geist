package main

import (
	"errors"
	"fmt"
)

func handleHelpSubcommand(cc *BGCommandContext) (err error) {
	if len(cc.args) < 2 {
		fmt.Fprintln(Main.OutWriter)
		showProgramUsage(cc.flags)
		return
	}
	command := cc.args[1]
	if command == "help" {
		return
	}
	if c, exists := commandmap[command]; exists {
		cc.args = []string{command, "help"}
		c.handler(cc)
	} else {
		fmt.Fprintf(Main.ErrWriter, "\nnot a blazegraph command: %s\n\n", command)
		showProgramUsage(cc.flags)
		err = errors.New("Not a blazegraph command")
	}
	return
}
