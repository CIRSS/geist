package cli

import (
	"flag"
	"fmt"
	"io"
)

type errorMessageWriterStruct struct {
	errorStream io.Writer
}

func (emw errorMessageWriterStruct) Write(p []byte) (n int, err error) {
	fmt.Fprintln(emw.errorStream)
	return emw.errorStream.Write(p)
}

type NullWriter struct {
	w io.Writer
}

func (nw NullWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

type CommandContext struct {
	programContext     *ProgramContext
	commands           *CommandCollection
	Descriptor         *CommandDescriptor
	Args               []string
	Flags              *flag.FlagSet
	Quiet              *bool
	InReader           io.Reader
	OutWriter          io.Writer
	ErrWriter          io.Writer
	ErrorMessageWriter errorMessageWriterStruct
	Properties         map[string]interface{}
}

func NewCommandContext(commands *CommandCollection, pc *ProgramContext) (cc *CommandContext) {

	cc = new(CommandContext)
	cc.programContext = pc
	cc.commands = commands

	cc.Flags = pc.InitFlagSet()
	cc.Flags.Usage = func() {}

	cc.ErrWriter = pc.ErrWriter
	cc.OutWriter = pc.OutWriter

	cc.Properties = make(map[string]interface{})
	cc.ErrorMessageWriter.errorStream = cc.ErrWriter

	cc.Quiet = cc.Flags.Bool("quiet", false, "Discard normal command output")

	return
}

func (cc *CommandContext) Lookup(commandName string) (*CommandDescriptor, bool) {
	return cc.commands.Lookup(commandName)
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
	fmt.Fprintf(cc.OutWriter, "\nUsage: %s %s [<flags>]\n\n", cc.programContext.Name, cc.Descriptor.Name)
	fmt.Fprint(cc.OutWriter, "Flags:\n\n")
	cc.Flags.PrintDefaults()
	fmt.Fprintln(cc.OutWriter)
}

func (cc *CommandContext) ShowProgramUsage() {
	fmt.Fprintf(cc.OutWriter, "Usage: %s <command> [<flags>]\n\n", cc.programContext.Name)
	fmt.Fprint(cc.OutWriter, "Commands:\n\n")
	for _, sc := range cc.commands.commandList {
		fmt.Fprintf(cc.OutWriter, "  %-7s  - %s\n", sc.Name, sc.Summary)
	}
	fmt.Fprint(cc.OutWriter, "\nCommon flags:\n\n")
	cc.Flags.PrintDefaults()
	fmt.Fprintf(cc.OutWriter, "\nSee '%s help <command>' for help with one of the above commands.\n\n", cc.programContext.Name)
	return
}

func (cc *CommandContext) ParseFlags2() (err error) {

	cc.Flags.SetOutput(cc.ErrorMessageWriter)
	if err = cc.Flags.Parse(cc.Args[1:]); err != nil {
		cc.Flags.SetOutput(cc.ErrWriter)
		cc.ShowCommandUsage()
		return
	}
	cc.Flags.SetOutput(cc.ErrWriter)

	if *cc.Quiet {
		cc.OutWriter = NullWriter{}
	}

	return
}

func (cc *CommandContext) ParseFlags() (helpShown bool, err error) {

	if cc.ShowHelpIfRequested() {
		helpShown = true
		return
	}

	err = cc.ParseFlags2()

	return
}
