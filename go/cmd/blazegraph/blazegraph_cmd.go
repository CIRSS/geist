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
		bc := blazegraph.NewClient()
		bc.DeleteAllTriples()

	case "dump":
		bc := blazegraph.NewClient()
		dump, err := bc.DumpAsNTriples()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(dump)

	case "load":
		df := *dataFile
		bc := blazegraph.NewClient()
		data, err := ioutil.ReadFile(df)
		if err != nil {
			fmt.Println(err)
			return
		}
		bc.PostTurtleBytes(data)

	case "load-jsonld":
		df := *dataFile
		bc := blazegraph.NewClient()
		data, err := ioutil.ReadFile(df)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = bc.PostJSONLDBytes(data)
		if err != nil {
			fmt.Println(err)
			return
		}

	default:
		fmt.Printf("Unrecognized command: %s\n", command)
	}

}
