package main

import (
	"fmt"
	"os"

	"github.com/tmcphillips/blazegraph-util/bg"
	"github.com/tmcphillips/main-wrapper/mw"
)

// MW wraps the main() function.  It enables tests to manipulate the
// input and output streams used by main(), and provides a new FlagSet
// for each execution so that main() can be called by multiple tests.
var MW mw.MainWrapper

func init() {
	MW = mw.NewMainWrapper("bgi", main)
}

// Exercises the template package
func main() {

	var err error

	flags := MW.InitFlagSet()
	err = flags.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		flags.Usage()
		return
	}

	command := os.Args[1]

	switch command {
	case "drop":
		bc := bg.NewBlazegraphClient()
		bc.DeleteAllTriples()
	}

}
