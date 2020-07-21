package main

import (
	"os"
	"strings"
)

func runWithArgs(commandLine string) {
	os.Args = strings.Fields(commandLine)
	Main.Run()
}
