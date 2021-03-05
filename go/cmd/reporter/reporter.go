package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/cirss/geist/cli"
)

// Program wraps the main() function.  It enables tests to manipulate the
// input and output streams used by main(), and provides a new FlagSet
// for each execution so that main() can be called by multiple tests.
var Program *cli.ProgramContext

func init() {
	Program = cli.NewProgramContext("sparqlrep", main)
}

// Exercises the template package
func main() {

	var err error

	flags := Program.InitFlagSet()
	err = flags.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		flags.Usage()
		return
	}

	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}

}
