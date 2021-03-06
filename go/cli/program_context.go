package cli

import (
	"flag"
	"io"
	"os"
)

// ProgramContext enables tests to manipulate the input and
// output streams used by main(), and provides a new FlagSet
// for each execution.
type ProgramContext struct {
	Name      string
	mainFunc  func()
	Flags     *flag.FlagSet
	InReader  io.Reader
	OutWriter io.Writer
	ErrWriter io.Writer
	ExitCode  int
	TestMode  bool
}

// NewProgramContext creates and initilizes an instance
// of ProgramContext for calling the provided main function
// and using the standard I/O streams by default.
func NewProgramContext(name string, main func()) *ProgramContext {
	pc := new(ProgramContext)
	pc.Name = name
	pc.mainFunc = main
	pc.InReader = os.Stdin
	pc.OutWriter = os.Stdout
	pc.ErrWriter = os.Stderr
	return pc
}

func (pc *ProgramContext) InitFlagSet() *flag.FlagSet {
	pc.Flags = flag.NewFlagSet(pc.Name, 0)
	pc.Flags.SetOutput(pc.ErrWriter)
	return pc.Flags
}

// Run invokes the wrapped main() function after
// instantiating a new FlagSet.
func (pc *ProgramContext) Run() int {
	pc.TestMode = true
	pc.mainFunc()
	return pc.ExitCode
}

func (pc *ProgramContext) ExitIfNonzero(code int) {
	pc.ExitCode = code
	if !pc.TestMode {
		os.Exit(code)
	}
}

func (pc *ProgramContext) NewCommandContext(commands *CommandSet) (cc *CommandContext) {
	return NewCommandContext(pc, commands)
}
