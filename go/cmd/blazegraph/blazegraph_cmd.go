package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/tmcphillips/blazegraph-util/sparql"

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

	case "export":
		format := flags.String("format", "nt", "Format for exported triples")
		if err = flags.Parse(os.Args[2:]); err != nil {
			fmt.Fprintf(Main.ErrWriter, err.Error())
			flags.Usage()
			return
		}
		export(*format)

	case "import":
		file := flags.String("file", "-", "File containing triples to import")
		format := flags.String("format", "ttl", "Format of triples to import")
		if err = flags.Parse(os.Args[2:]); err != nil {
			fmt.Fprintf(Main.ErrWriter, err.Error())
			flags.Usage()
			return
		}
		inport(*file, *format)

	case "query":
		file := flags.String("file", "-", "File containing query to execute")
		format := flags.String("format", "json", "Format of result set to produce")
		if err = flags.Parse(os.Args[2:]); err != nil {
			fmt.Println(err)
			flags.Usage()
			return
		}
		query(*file, *format)

	default:
		fmt.Fprintf(Main.ErrWriter, "Unrecognized command: %s\n", command)
	}

}

func drop() {
	bc := blazegraph.NewClient()
	bc.DeleteAll()
}

func readFileOrStdin(filePath string) (bytes []byte, err error) {
	var r io.Reader
	if filePath == "-" {
		r = Main.InReader
	} else {
		r, _ = os.Open(filePath)
	}
	return ioutil.ReadAll(r)
}

func inport(file string, format string) {
	bc := blazegraph.NewClient()
	data, err := readFileOrStdin(file)
	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		return
	}

	switch format {

	case "jsonld":
		_, err = bc.PostData("application/ld+json", data)

	case "ttl":
		_, err = bc.PostData("application/x-turtle", data)
	}

	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		return
	}
}

func query(file string, format string) {
	bc := blazegraph.NewClient()
	q, err := readFileOrStdin(file)
	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		return
	}

	switch format {

	case "csv":
		resultCSV, _ := bc.SelectCSV(string(q))
		if err != nil {
			break
		}
		fmt.Fprintf(Main.OutWriter, resultCSV)
		return

	case "json":
		var rs sparql.ResultSet
		rs, err = bc.Select(string(q))
		if err != nil {
			break
		}
		resultJSON, _ := rs.JSONString()
		fmt.Fprintf(Main.OutWriter, resultJSON)
		return

	case "xml":
		resultXML, err := bc.SelectXML(string(q))
		if err != nil {
			break
		}
		fmt.Fprintf(Main.OutWriter, resultXML)
		return
	}

	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		return
	}

}

func export(format string) {
	bc := blazegraph.NewClient()
	var err error
	var triples string

	switch format {
	case "jsonld":
		triples, err = bc.ConstructAll("application/ld+json")
	case "nt":
		triples, err = bc.ConstructAll("text/plain")
	case "ttl":
		triples, err = bc.ConstructAll("application/x-turtle")
	case "xml":
		triples, err = bc.ConstructAll("application/rdf+xml")
	}

	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		return
	}

	fmt.Fprintf(Main.OutWriter, triples)
}
