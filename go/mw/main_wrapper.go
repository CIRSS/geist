package mw

import (
	"flag"
	"io"
	"os"
)

// MainWrapper enables tests to manipulate the input and
// output streams used by main(), and provides a new FlagSet
// for each execution.
type MainWrapper struct {
	cmd       string
	mainFunc  func()
	Flags     *flag.FlagSet
	InReader  io.Reader
	OutWriter io.Writer
	ErrWriter io.Writer
	ExitCode  int
	TestMode  bool
}

// NewMainWrapper creates and initilizes an instance
// of MainWrapper for calling the provided main function
// and using the standard I/O streams by default.
func NewMainWrapper(cmd string, main func()) MainWrapper {
	mw := MainWrapper{}
	mw.cmd = cmd
	mw.mainFunc = main
	mw.InReader = os.Stdin
	mw.OutWriter = os.Stdout
	mw.ErrWriter = os.Stderr
	return mw
}

func (mw *MainWrapper) InitFlagSet() *flag.FlagSet {
	mw.Flags = flag.NewFlagSet(mw.cmd, 0)
	mw.Flags.SetOutput(mw.ErrWriter)
	return mw.Flags
}

// Run invokes the wrapped main() function after
// instantiating a new FlagSet.
func (mw *MainWrapper) Run() int {
	mw.TestMode = true
	mw.mainFunc()
	return mw.ExitCode
}

func (mw *MainWrapper) ExitIfNonzero(code int) {
	mw.ExitCode = code
	if !mw.TestMode {
		os.Exit(code)
	}
}
