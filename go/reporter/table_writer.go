package reporter

import (
	"fmt"
	"strings"
	"text/tabwriter"
)

type TableWriter struct {
	buffer       *strings.Builder
	tabWriter    *tabwriter.Writer
	MaxRowLen    int
	RowSeparator string
}

func NewTableWriter(colSeparators bool) *TableWriter {
	tw := new(TableWriter)
	tw.buffer = new(strings.Builder)
	var flags uint
	if colSeparators {
		flags = tabwriter.Debug
		tw.RowSeparator = " \t "
	} else {
		tw.RowSeparator = "  \t "
	}
	tw.tabWriter = tabwriter.NewWriter(tw.buffer, 0, 0, 0, ' ', flags)
	return tw
}

func (tw *TableWriter) Write(row []byte) (n int, err error) {
	return tw.tabWriter.Write(row)
}

func (tw *TableWriter) String() string {
	tw.tabWriter.Flush()
	return tw.buffer.String()
}

func (tw *TableWriter) updateMaxRowLen(r string) {
	l := len(r)
	if tw.MaxRowLen < l {
		tw.MaxRowLen = l
	}
}

func (tw *TableWriter) WriteRow(rowArray []string) {
	rowString := strings.Join(rowArray, tw.RowSeparator)
	tw.updateMaxRowLen(rowString)
	fmt.Fprintln(tw, rowString)
}

func (tw *TableWriter) WriteStringArray(sa [][]string) {
	for _, row := range sa {
		tw.WriteRow(row)
	}
}

func WriteStringTable(sa [][]string, colSeparators bool) string {
	tw := NewTableWriter(colSeparators)
	tw.WriteStringArray(sa)
	s := tw.String()
	dashCount := tw.MaxRowLen - 1
	if colSeparators {
		dashCount += len(sa[0]) - 1
	}
	dashedLine := "\n" + strings.Repeat("-", dashCount)
	f := strings.Index(s, "\n")
	return s[:f] + dashedLine + s[f:]
}
