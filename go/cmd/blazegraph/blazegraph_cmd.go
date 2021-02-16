package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/cirss/geist"
	"github.com/cirss/geist/blazegraph"
	"github.com/tmcphillips/main-wrapper/mw"
)

// Main wraps the main() function.  It enables tests to manipulate the
// input and output streams used by main(), and provides a new FlagSet
// for each execution so that main() can be called by multiple tests.
var Main mw.MainWrapper

var options struct {
	url *string
}

func init() {
	Main = mw.NewMainWrapper("geist", main)
}

func main() {

	var err error

	if len(os.Args) < 2 {
		fmt.Fprintln(Main.ErrWriter, "Usage: geist <command> [args...]")
		return
	}

	command := os.Args[1]
	flags := Main.InitFlagSet()

	switch command {

	case "create":
		addCommonOptions(flags)
		dataset := flags.String("dataset", "", "Dataset to create")
		infer := flags.String("infer", "none", "Inference to perform on update [none, rdfs, owl]")
		if err = flags.Parse(os.Args[2:]); err != nil {
			fmt.Fprintf(Main.ErrWriter, err.Error())
			flags.Usage()
			return
		}
		if len(*dataset) == 0 {
			fmt.Fprintln(Main.ErrWriter, "Error: Name of dataset to create not provided.")
			flags.Usage()
			return
		}
		doCreate(*dataset, *infer)

	case "destroy":
		addCommonOptions(flags)
		dataset := flags.String("dataset", "", "Dataset to destroy")
		if err = flags.Parse(os.Args[2:]); err != nil {
			fmt.Fprintf(Main.ErrWriter, err.Error())
			flags.Usage()
			return
		}
		if len(*dataset) == 0 {
			fmt.Fprintln(Main.ErrWriter, "Error: Name of dataset to destroy not provided.")
			flags.Usage()
			return
		}
		doDestroy(*dataset)

	case "export":
		addCommonOptions(flags)
		format := flags.String("format", "nt", "Format for doExported triples")
		sort := flags.Bool("sort", false, "Sort the exported triples if true")
		if err = flags.Parse(os.Args[2:]); err != nil {
			fmt.Fprintf(Main.ErrWriter, err.Error())
			flags.Usage()
			return
		}
		doExport(*format, *sort)

	case "import":
		addCommonOptions(flags)
		file := flags.String("file", "-", "File containing triples to import")
		format := flags.String("format", "ttl", "Format of triples to import")
		if err = flags.Parse(os.Args[2:]); err != nil {
			fmt.Fprintf(Main.ErrWriter, err.Error())
			flags.Usage()
			return
		}
		doImport(*file, *format)

	case "report":
		addCommonOptions(flags)
		file := flags.String("file", "-", "File containing report template to expand")
		if err = flags.Parse(os.Args[2:]); err != nil {
			fmt.Fprintf(Main.ErrWriter, err.Error())
			flags.Usage()
			return
		}
		doReport(*file)

	case "select":
		addCommonOptions(flags)
		dryrun := flags.Bool("dryrun", false, "Output query but do not execute it")
		file := flags.String("file", "-", "File containing select query to execute")
		format := flags.String("format", "json", "Format of result set to produce")
		separators := flags.Bool("columnseparators", true, "Display column separators in table format")
		if err = flags.Parse(os.Args[2:]); err != nil {
			fmt.Println(err)
			flags.Usage()
			return
		}
		doSelectQuery(*dryrun, *file, *format, *separators)

	default:
		fmt.Fprintf(Main.ErrWriter, "Unrecognized command: %s\n", command)
	}
}

func addCommonOptions(flags *flag.FlagSet) {
	options.url = flags.String("url", blazegraph.DefaultUrl, "URL of Blazegraph instance")
}

func doCreate(name string, infer string) {
	bc := blazegraph.NewBlazegraphClient(*options.url)
	p := blazegraph.NewDatasetProperties(name)
	p.Inference = infer
	bc.CreateDataSet(p)
	// if err != nil {
	// 	fmt.Fprintln(Main.ErrWriter, err.Error())
	// }
	// fmt.Fprintln(Main.OutWriter, string(response))
}

func doDestroy(name string) {
	bc := blazegraph.NewBlazegraphClient(*options.url)
	bc.DestroyDataSet(name)
	// if err != nil {
	// 	fmt.Fprintln(Main.ErrWriter, err.Error())
	// }
	// fmt.Fprintln(Main.OutWriter, string(response))
}

func doReport(file string) {
	bc := blazegraph.NewBlazegraphClient(*options.url)
	reportTemplate, err := readFileOrStdin(file)
	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		return
	}
	rt := geist.NewTemplate("main", string(reportTemplate), nil, bc)
	report, re := bc.ExpandReport(rt)
	if re != nil {
		fmt.Fprintf(Main.ErrWriter, re.Error())
		return
	}
	fmt.Fprint(Main.OutWriter, report)
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
	bc := blazegraph.NewBlazegraphClient(*options.url)
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

func doSelectQuery(dryrun bool, file string, format string, columnSeparators bool) {

	bc := blazegraph.NewBlazegraphClient(*options.url)
	queryText, err := readFileOrStdin(file)
	if err != nil {
		fmt.Fprintf(Main.ErrWriter, err.Error())
		return
	}

	queryTemplate := geist.NewTemplate("query", string(queryText), nil, bc)
	err = queryTemplate.Parse()
	if err != nil {
		fmt.Fprintf(Main.ErrWriter, "Error expanding query template:\n")
		fmt.Fprintf(Main.ErrWriter, "%s\n", err.Error())
		return
	}

	q, err := queryTemplate.Expand(nil)

	if err != nil {
		fmt.Fprintf(Main.ErrWriter, "Error expanding query template: ")
		fmt.Fprintf(Main.ErrWriter, "%s\n", err.Error())
		return
	}

	if dryrun {
		fmt.Print(string(q))
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
	bc := blazegraph.NewBlazegraphClient(*options.url)
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
