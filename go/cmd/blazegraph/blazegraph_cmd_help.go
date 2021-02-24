package main

import (
	"errors"
	"fmt"
)

func handleHelpSubcommand(cc *Context) (err error) {
	if len(cc.args) < 2 {
		fmt.Fprintln(Main.OutWriter)
		showProgramUsage(cc.flags)
		return
	}
	commandName := cc.args[1]
	if commandName == "help" {
		return
	}
	if c, exists := commandCollection.commandMap[commandName]; exists {
		cc.args = []string{commandName, "help"}
		c.handler(cc)
	} else {
		fmt.Fprintf(Main.ErrWriter, "\nnot a blazegraph command: %s\n\n", commandName)
		showProgramUsage(cc.flags)
		err = errors.New("Not a blazegraph command")
	}
	return
}
