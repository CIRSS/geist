package main

import (
	"os"
	"strings"
)

func runWithArgs(commandLine string) {
	Main.ErrWriter = os.Stdout
	os.Args = strings.Fields(commandLine)
	Main.Run()
}

func Example_drop_then_dump() {
	runWithArgs("bgi drop")
	runWithArgs("bgi dump")
	// Output:
	//
}

func Example_drop_load_then_dump() {
	runWithArgs("bgi drop")
	runWithArgs("bgi -f testdata/in.nt load")
	runWithArgs("bgi dump")
	// Output:
	// <http://tmcphill.net/data#y> <http://tmcphill.net/tags#tag> "eight" .
	// <http://tmcphill.net/tags#tag> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/1999/02/22-rdf-syntax-ns#Property> .
	// <http://tmcphill.net/tags#tag> <http://www.w3.org/2000/01/rdf-schema#subPropertyOf> <http://tmcphill.net/tags#tag> .
	// <http://tmcphill.net/data#x> <http://tmcphill.net/tags#tag> "seven" .
}
