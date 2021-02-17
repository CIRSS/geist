package main

import (
	"flag"
	"fmt"
)

func handleHelpSubcommand(args []string, flags *flag.FlagSet) {
	if len(args) < 2 {
		fmt.Fprintln(Main.OutWriter)
		showProgramUsage()
		return
	}
	command := args[1]
	if command == "help" {
		return
	}
	if c, exists := commandmap[command]; exists {
		c.handler([]string{command, "help"}, flags)
	} else {
		fmt.Fprintf(Main.ErrWriter, "\nnot a blazegraph command: %s\n\n", command)
		showProgramUsage()
	}
}

func showProgramUsage() {
	fmt.Fprint(Main.OutWriter, "Usage: blazegraph <command> [<args>]\n\n")
	fmt.Fprint(Main.OutWriter, "Available commands:\n\n")
	for _, sc := range commands {
		fmt.Fprintf(Main.OutWriter, "  %-7s  - %s\n", sc.name, sc.summary)
	}
	fmt.Fprint(Main.OutWriter, "\nSee 'blazegraph help <command>' for help with one of the above commands.\n\n")
	return
}

func helpRequested(args []string, flags *flag.FlagSet) bool {
	if len(args) > 1 && args[1] == "help" {
		showCommandDescription(commandmap[args[0]])
		showCommandUsage(args, flags)
		return true
	}
	return false
}

func showCommandDescription(c *command) {
	fmt.Fprintf(Main.OutWriter, "\n%s\n", c.description)
}

func showCommandUsage(args []string, flags *flag.FlagSet) {
	fmt.Fprintf(Main.OutWriter, "\nUsage: blazegraph %s <flags>\n\n", args[0])
	fmt.Fprint(Main.OutWriter, "Flags:\n")
	flags.PrintDefaults()
	fmt.Fprintln(Main.OutWriter)
}
