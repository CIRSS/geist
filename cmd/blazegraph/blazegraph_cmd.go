package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tmcphillips/blazegraph-util/blazegraph"
	"github.com/tmcphillips/main-wrapper/mw"
)

// MW wraps the main() function.  It enables tests to manipulate the
// input and output streams used by main(), and provides a new FlagSet
// for each execution so that main() can be called by multiple tests.
var Main mw.MainWrapper

func init() {
	Main = mw.NewMainWrapper("bgi", main)
}

// Exercises the template package
func main() {

	var err error

	flags := Main.InitFlagSet()
	dataFile := flags.String("f", "-", "file to read or write")
	err = flags.Parse(os.Args[1:])
	if err != nil || len(os.Args) < 2 {
		if err != nil {
			fmt.Println(err)
		}
		flags.Usage()
		return
	}

	command := flags.Args()[0]

	switch command {

	case "drop":
		bc := blazegraph.NewBlazegraphClient()
		bc.DeleteAllTriples()

	case "dump":
		bc := blazegraph.NewBlazegraphClient()
		dump := bc.DumpAsNTriples()
		fmt.Println(dump)

	case "load":
		df := *dataFile
		bc := blazegraph.NewBlazegraphClient()
		data, err := ioutil.ReadFile(df)
		if err != nil {
			fmt.Println(err)
			return
		}
		bc.PostNewData(data)

	default:
		fmt.Printf("Unrecognized command: %s\n", command)
	}

}
