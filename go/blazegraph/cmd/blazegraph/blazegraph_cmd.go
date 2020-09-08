package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/cirss/geist/blazegraph"
	"github.com/cirss/geist/geist"
	"github.com/tmcphillips/main-wrapper/mw"
)

// Main wraps the main() function.  It enables tests to manipulate the
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
		doDrop()

	case "export":
		format := flags.String("format", "nt", "Format for doExported triples")
		sort := flags.Bool("sort", false, "Sort the exported triples if true")
		if err = flags.Parse(os.Args[2:]); err != nil {
			fmt.Fprintf(Main.ErrWriter, err.Error())
			flags.Usage()
			return
		}
		doExport(*format, *sort)

	case "import":
		file := flags.String("file", "-", "File containing triples to import")
		format := flags.String("format", "ttl", "Format of triples to import")
		if err = flags.Parse(os.Args[2:]); err != nil {
			fmt.Fprintf(Main.ErrWriter, err.Error())
			flags.Usage()
			return
		}
		doImport(*file, *format)

	case "report":
		file := flags.String("file", "-", "File containing report template to expand")
		if err = flags.Parse(os.Args[2:]); err != nil {
			fmt.Fprintf(Main.ErrWriter, err.Error())
			flags.Usage()
			return
		}
		doReport(*file)

	case "select":
		file := flags.String("file", "-", "File containing select query to execute")
		format := flags.String("format", "json", "Format of result set to produce")
		separators := flags.Bool("columnseparators", true, "Display column separators in table format")
		if err = flags.Parse(os.Args[2:]); err != nil {
			fmt.Println(err)
			flags.Usage()
			return
		}
		doSelectQuery(*file, *format, *separators)

	default:
		fmt.Fprintf(Main.ErrWriter, "Unrecognized command: %s\n", command)
	}
}

func doReport(file string) {
	bc := blazegraph.NewClient()
	reportTemplate, err := readFileOrStdin(file)
	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		return
	}
	rt := geist.NewTemplate("main", string(reportTemplate), nil)
	report, re := bc.ExpandReport(rt)
	if re != nil {
		fmt.Fprintf(Main.ErrWriter, re.Error())
		return
	}
	fmt.Fprint(Main.OutWriter, report)
}

func doDrop() {
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

func doImport(file string, format string) {
	bc := blazegraph.NewClient()
	data, err := readFileOrStdin(file)
	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		return
	}

	switch format {

	case "jsonld":
		_, err = bc.PostData("application/ld+json", data)

	case "nt":
		_, err = bc.PostData("text/plain", data)

	case "ttl":
		_, err = bc.PostData("application/x-turtle", data)

	case "xml":
		_, err = bc.PostData("application/rdf+xml", data)
	}

	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		return
	}
}

func doSelectQuery(file string, format string, columnSeparators bool) {
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
		rs, err := bc.Select(string(q))
		if err != nil {
			break
		}
		resultJSON, _ := rs.JSONString()
		fmt.Fprintf(Main.OutWriter, resultJSON)
		return

	case "table":
		rs, err := bc.Select(string(q))
		if err != nil {
			break
		}
		table := rs.FormattedTable(columnSeparators)
		fmt.Fprintf(Main.OutWriter, table)
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

func doExport(format string, sorted bool) {
	bc := blazegraph.NewClient()
	var err error
	var triples string

	switch format {
	case "jsonld":
		triples, err = bc.ConstructAll("application/ld+json", sorted)
	case "nt":
		triples, err = bc.ConstructAll("text/plain", sorted)
		if sorted {
		}
	case "ttl":
		triples, err = bc.ConstructAll("application/x-turtle", sorted)
	case "xml":
		triples, err = bc.ConstructAll("application/rdf+xml", sorted)
	}

	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		return
	}

	fmt.Fprintf(Main.OutWriter, triples)
}
