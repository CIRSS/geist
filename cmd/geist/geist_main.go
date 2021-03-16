package main

import (
	"os"

	"github.com/cirss/blaze/pkg/blaze"
	"github.com/cirss/go-cli/pkg/cli"
)

var Main *cli.ProgramContext

func init() {
	Main = cli.NewProgramContext("geist", main)
}

func main() {
	cc := blaze.NewBlazeCommandContext(Main)
	cc.InvokeCommand(os.Args)
}
