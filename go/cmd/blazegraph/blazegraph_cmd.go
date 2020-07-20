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

	if len(os.Args) < 2 {
		fmt.Fprintln(Main.ErrWriter, "Usage: blazegraph <command> [args...]")
		return
	}

	command := os.Args[1]
	flags := Main.InitFlagSet()

	switch command {

	case "drop":
		if len(os.Args) > 2 {
			fmt.Fprintln(Main.ErrWriter,
				"Error: The 'blazegraph drop' command takes no arguments and no flags.")
			return
		}
		drop()

	case "dump":
		format := flags.String("format", "n-triples", "Format for dumped triples")
		if err = flags.Parse(os.Args[2:]); err != nil {
			fmt.Println(err)
			flags.Usage()
			return
		}
		dump(*format)

	case "load":
		file := flags.String("file", "-", "File containing triples to load")
		format := flags.String("format", "n-triples", "Format of triples to load")
		if err = flags.Parse(os.Args[2:]); err != nil {
			fmt.Println(err)
			flags.Usage()
			return
		}
		load(*file, *format)

	default:
		fmt.Printf("Unrecognized command: %s\n", command)
	}

}

func drop() {
	bc := blazegraph.NewClient()
	bc.DeleteAll()
}

func load(file string, format string) {
	bc := blazegraph.NewClient()
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch format {

	case "turtle":
		_, err = bc.PostData("application/x-turtle", data)

	case "json-ld":
		_, err = bc.PostData("application/ld+json", data)
	}

	if err != nil {
		fmt.Println(err)
		return
	}
}

func dump(format string) {
	bc := blazegraph.NewClient()
	var err error
	var triples string

	switch format {
	case "n-triples":
		triples, err = bc.ConstructAll("text/plain")
	case "json-ld":
		triples, err = bc.ConstructAll("application/ld+json")
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(Main.OutWriter, triples)
}
