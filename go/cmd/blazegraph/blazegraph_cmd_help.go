package main

import (
	"errors"
	"flag"
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

func showProgramUsage(flags *flag.FlagSet) {
	fmt.Fprint(Main.OutWriter, "Usage: blazegraph <command> [<flags>]\n\n")
	fmt.Fprint(Main.OutWriter, "Commands:\n\n")
	for _, sc := range commands {
		fmt.Fprintf(Main.OutWriter, "  %-7s  - %s\n", sc.name, sc.summary)
	}
	fmt.Fprint(Main.OutWriter, "\nCommon flags:\n")
	flags.PrintDefaults()
	fmt.Fprint(Main.OutWriter, "\nSee 'blazegraph help <command>' for help with one of the above commands.\n\n")
	return
}

func helpRequested(cc *BGCommandContext) bool {
	if len(cc.args) > 1 && cc.args[1] == "help" {
		showCommandDescription(commandmap[cc.args[0]])
		showCommandUsage(cc)
		return true
	}
	return false
}

func showCommandDescription(c *command) {
	fmt.Fprintf(Main.OutWriter, "\n%s\n", c.description)
}

func showCommandUsage(cc *BGCommandContext) {
	fmt.Fprintf(Main.OutWriter, "\nUsage: blazegraph %s [<flags>]\n\n", cc.args[0])
	fmt.Fprint(Main.OutWriter, "Flags:\n")
	cc.flags.PrintDefaults()
	fmt.Fprintln(Main.OutWriter)
}
