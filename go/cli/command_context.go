package cli

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
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
	commands           *CommandSet
	Descriptor         *CommandDescriptor
	Args               []string
	Flags              *flag.FlagSet
	Quiet              *bool
	InReader           io.Reader
	OutWriter          io.Writer
	ErrWriter          io.Writer
	ErrorMessageWriter errorMessageWriterStruct
	Properties         map[string]interface{}
	Providers          map[string]func(cc *CommandContext) interface{}
}

func NewCommandContext(pc *ProgramContext, commands *CommandSet) (cc *CommandContext) {

	cc = new(CommandContext)
	cc.programContext = pc
	cc.commands = commands

	cc.Flags = pc.InitFlagSet()
	cc.Flags.Usage = func() {}

	cc.InReader = pc.InReader
	cc.ErrWriter = pc.ErrWriter
	cc.OutWriter = pc.OutWriter

	cc.Properties = make(map[string]interface{})
	cc.Providers = make(map[string]func(cc *CommandContext) interface{})
	cc.ErrorMessageWriter.errorStream = cc.ErrWriter

	cc.Quiet = cc.Flags.Bool("quiet", false, "Discard normal command output")

	return
}

func (cc *CommandContext) AddProvider(resource string, f func(cc *CommandContext) interface{}) {
	cc.Providers[resource] = f
}

func (cc *CommandContext) Resource(resourceName string) (resource interface{}) {

	provider, exists := cc.Providers[resourceName]
	if !exists {
		panic("No resource provider for " + resourceName)
	}

	resource = provider(cc)

	return
}

func (cc *CommandContext) Lookup(commandName string) (*CommandDescriptor, bool) {
	return cc.commands.Lookup(commandName)
}

func (cc *CommandContext) ShowHelpIfRequested() bool {
	if len(cc.Args) > 1 && cc.Args[1] == "help" {
		cc.ShowCommandDescription()
		cc.ShowCommandUsage(cc.OutWriter)
		return true
	}
	return false
}

func (cc *CommandContext) ShowCommandDescription() {
	fmt.Fprintf(cc.OutWriter, "\n%s %s: %s\n",
		cc.programContext.Name, cc.Descriptor.Name, cc.Descriptor.Description)
}

func (cc *CommandContext) ShowCommandUsage(w io.Writer) {
	fmt.Fprintf(w, "\nusage: %s %s [<flags>]\n\n", cc.programContext.Name, cc.Descriptor.Name)
	fmt.Fprint(w, "flags:\n")
	cc.Flags.SetOutput(w)
	cc.Flags.PrintDefaults()
	fmt.Fprintln(w)
}

func (cc *CommandContext) ShowProgramUsage() {
	fmt.Fprintf(cc.OutWriter, "usage: %s <command> [<flags>]\n\n", cc.programContext.Name)
}

func (cc *CommandContext) ShowProgramCommands() {
	fmt.Fprint(cc.OutWriter, "commands:\n")
	for _, sc := range cc.commands.commandList {
		fmt.Fprintf(cc.OutWriter, "  %-7s  - %s\n", sc.Name, sc.Summary)
	}
	fmt.Fprint(cc.OutWriter, "\nflags:\n")
	cc.Flags.PrintDefaults()
	fmt.Fprintf(cc.OutWriter,
		"\nSee '%s help <command>' for help with one of the above commands.\n\n",
		cc.programContext.Name)
	return
}

func (cc *CommandContext) ParseCommandFlags() (err error) {

	var errBuffer = new(strings.Builder)

	cc.Flags.SetOutput(errBuffer)
	if err = cc.Flags.Parse(cc.Args[1:]); err != nil {
		err = NewCLIError(err)
		fmt.Fprintf(cc.ErrWriter, "%s %s: %s",
			cc.programContext.Name, cc.Descriptor.Name, errBuffer.String())
		cc.ShowCommandUsage(cc.ErrWriter)

		return
	}

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

	err = cc.ParseCommandFlags()
	if err != nil {
		return
	}

	if len(cc.Flags.Args()) > 0 {
		fmt.Fprintf(cc.ErrWriter, "%s %s: unused argument: %s\n",
			cc.programContext.Name, cc.Descriptor.Name, cc.Flags.Args()[0])
		cc.ShowCommandUsage(cc.ErrWriter)
		err = NewCLIError(nil)
		return
	}

	return
}

func (cc *CommandContext) InvokeCommand(args []string) {

	if len(args) < 2 {
		fmt.Fprintf(cc.programContext.ErrWriter, "%s: no command given\n\n",
			cc.programContext.Name)
		cc.ShowProgramUsage()
		cc.ShowProgramCommands()
		cc.programContext.ExitIfNonzero(1)
		return
	}

	commandName := args[1]
	descriptor, exists := cc.commands.Lookup(commandName)
	cc.Descriptor = descriptor
	if !exists {
		fmt.Fprintf(cc.programContext.ErrWriter, "%s: unrecognized command: %s\n\n",
			cc.programContext.Name, commandName)
		cc.ShowProgramUsage()
		cc.ShowProgramCommands()
		cc.programContext.ExitIfNonzero(1)
		return
	}

	cc.Args = args[1:]
	err := cc.Descriptor.Handler(cc)

	if err != nil {
		switch err.(type) {
		case CLIError:
			break
		default:
			fmt.Fprintf(cc.ErrWriter, "%s %s: %s\n",
				cc.programContext.Name, cc.Descriptor.Name, err.Error())
		}
		cc.programContext.ExitIfNonzero(1)
		return
	}
}

func Help(cc *CommandContext) (err error) {
	if len(cc.Args) < 2 {
		fmt.Fprintln(cc.OutWriter)
		cc.ShowProgramUsage()
		cc.ShowProgramCommands()
		return
	}
	commandName := cc.Args[1]
	if commandName == "help" {
		return
	}
	if c, exists := cc.Lookup(commandName); exists {
		cc.Descriptor = c
		cc.Args = []string{commandName, "help"}
		c.Handler(cc)
	} else {
		fmt.Fprintf(cc.ErrWriter, "%s help: unrecognized %s command: %s\n\n",
			cc.programContext.Name, cc.programContext.Name, commandName)
		cc.ShowProgramUsage()
		cc.ShowProgramCommands()
		err = NewCLIError(nil)
	}
	return
}

func (cc *CommandContext) ReadFileOrStdin(filePath string) (bytes []byte, err error) {
	var r io.Reader
	if filePath == "-" {
		r = cc.InReader
	} else {
		r, _ = os.Open(filePath)
	}
	return ioutil.ReadAll(r)
}
