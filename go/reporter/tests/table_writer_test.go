package tests

import (
	"fmt"
	"testing"

	"github.com/tmcphillips/blazegraph-util/reporter"
	"github.com/tmcphillips/blazegraph-util/util"
)

func TestTableWriter_NoData(t *testing.T) {
	tw := reporter.NewTableWriter(true)
	util.StringEquals(t, tw.String(), "")
}

func TestTableWriter_OneTabDelimitedRow(t *testing.T) {
	tw := reporter.NewTableWriter(true)
	fmt.Fprintln(tw, "a\tb\tc\t")
	util.StringEquals(t, tw.String(), "a|b|c|\n")
}

func TestTableWriter_WriteStringArray(t *testing.T) {

	contacts := [][]string{
		{"Timothy", "Oakland", "530-219-4754"},
		{"Bob", "Concord", "510-320-9943"},
		{"Joseph", "San Diego", "01-213-101-9313"},
	}

	tw := reporter.NewTableWriter(true)
	tw.WriteStringArray(contacts)
	util.LineContentsEqual(t, tw.String(),
		`
		Timothy | Oakland   | 530-219-4754
		Bob     | Concord   | 510-320-9943
		Joseph  | San Diego | 01-213-101-9313
	`)
}

func TestTableWriter_WriteStringTable(t *testing.T) {

	contacts := [][]string{
		{"Name", "City", "Phone"},
		{"Timothy", "Oakland", "530-219-4754"},
		{"Bob", "Concord", "510-320-9943"},
		{"Joseph", "San Diego", "01-213-101-9313"},
	}

	actual := reporter.WriteStringTable(contacts, false)

	util.LineContentsEqual(t, actual,
		`
		Name      City        Phone
		-------------------------------------
		Timothy   Oakland     530-219-4754
		Bob       Concord     510-320-9943
		Joseph    San Diego   01-213-101-9313
	`)
}

func TestTableWriter_WriteStringTableWithSeparators(t *testing.T) {

	contacts := [][]string{
		{"Name", "City", "Phone"},
		{"Timothy", "Oakland", "530-219-4754"},
		{"Bob", "Concord", "510-320-9943"},
		{"Joseph", "San Diego", "01-213-101-9313"},
	}

	actual := reporter.WriteStringTable(contacts, true)

	util.LineContentsEqual(t, actual,
		`
		Name    | City      | Phone
		-------------------------------------
		Timothy | Oakland   | 530-219-4754
		Bob     | Concord   | 510-320-9943
		Joseph  | San Diego | 01-213-101-9313
	`)
}
