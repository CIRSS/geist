package main

import (
	"os"
	"strings"
	"testing"

	"github.com/cirss/geist/util"
)

func run(commandLine string) int {
	Main.ExitCode = 0
	os.Args = strings.Fields(commandLine)
	exitCode := Main.Run()
	return exitCode
}

func assertExitCode(t *testing.T, commandLine string, expected int) {
	actual := run(commandLine)
	util.IntEquals(t, actual, expected)
}
