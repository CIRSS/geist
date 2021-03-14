package main

import (
	"os"

	"github.com/cirss/geist/go/blaze"
	"github.com/cirss/go-cli/go/cli"
)

var Main *cli.ProgramContext

func init() {
	Main = cli.NewProgramContext("blaze", main)
}

func main() {
	cc := blaze.NewBlazeCommandContext(Main)
	cc.InvokeCommand(os.Args)
}
