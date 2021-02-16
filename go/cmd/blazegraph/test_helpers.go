package main

import (
	"os"
	"strings"
)

func run(commandLine string) {
	os.Args = strings.Fields(commandLine)
	Main.Run()
}
