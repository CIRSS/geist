package main

import (
	"os"

	"github.com/cirss/geist/go/blazegraph"
	"github.com/cirss/go-cli/go/cli"
)

var Main *cli.ProgramContext

func init() {
	Main = cli.NewProgramContext("blazegraph", main)
}

func main() {
	cc := blazegraph.NewBlazeCommandContext(Main)
	cc.InvokeCommand(os.Args)
}
