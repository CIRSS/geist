package cli

import (
	"flag"
	"fmt"
	"io"
)

type CommandContext struct {
	commands   *CommandCollection
	Descriptor *CommandDescriptor
	Args       []string
	Flags      *flag.FlagSet
	InReader   io.Reader
	OutWriter  io.Writer
	ErrWriter  io.Writer
	Properties map[string]interface{}
}

func NewCommandContext(commands *CommandCollection) (cc *CommandContext) {
	cc = new(CommandContext)
	cc.commands = commands
	cc.Properties = make(map[string]interface{})
	return
}

func (cc *CommandContext) ShowHelpIfRequested() bool {
	if len(cc.Args) > 1 && cc.Args[1] == "help" {
		cc.ShowCommandDescription()
		cc.ShowCommandUsage()
		return true
	}
	return false
}

func (cc *CommandContext) ShowCommandDescription() {
	fmt.Fprintf(cc.OutWriter, "\n%s\n", cc.Descriptor.Description)
}

func (cc *CommandContext) ShowCommandUsage() {
	fmt.Fprintf(cc.OutWriter, "\nUsage: blazegraph %s [<flags>]\n\n", cc.Descriptor.Name)
	fmt.Fprint(cc.OutWriter, "Flags:\n\n")
	cc.Flags.PrintDefaults()
	fmt.Fprintln(cc.OutWriter)
}

func (cc *CommandContext) ShowProgramUsage() {
	fmt.Fprint(cc.OutWriter, "Usage: blazegraph <command> [<flags>]\n\n")
	fmt.Fprint(cc.OutWriter, "Commands:\n\n")
	for _, sc := range cc.commands.commandList {
		fmt.Fprintf(cc.OutWriter, "  %-7s  - %s\n", sc.Name, sc.Summary)
	}
	fmt.Fprint(cc.OutWriter, "\nCommon flags:\n\n")
	cc.Flags.PrintDefaults()
	fmt.Fprint(cc.OutWriter, "\nSee 'blazegraph help <command>' for help with one of the above commands.\n\n")
	return
}
