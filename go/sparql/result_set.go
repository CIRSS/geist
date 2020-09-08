package sparql

import (
	"encoding/json"

	"github.com/cirss/geist"
)

type ResultSet struct {
	Head    Head    `json:"head"`
	Results Results `json:"results"`
}

type Head struct {
	Vars []string `json:"vars"`
}

type Results struct {
	Bindings []Binding `json:"bindings"`
}

type Binding map[string]struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (b Binding) DelimitedValue(name string) (delimitedValue string) {
	value := b[name].Value
	switch b[name].Type {
	case "uri":
		delimitedValue = "<" + value + ">"
	case "literal":
		delimitedValue = "\"" + value + "\""
	}
	return
}

func (sr *ResultSet) Variables() []string {
	return sr.Head.Vars
}

func (sr *ResultSet) Bindings() []Binding {
	return sr.Results.Bindings
}

func (rs *ResultSet) ColumnCount() int {
	return len(rs.Head.Vars)
}

func (rs *ResultSet) RowCount() int {
	return len(rs.Results.Bindings)
}

func (rs *ResultSet) JSONString() (jsonString string, err error) {
	jsonBytes, err := json.Marshal(*rs)
	jsonString = string(jsonBytes)
	return
}

func (sr *ResultSet) Row(rowIndex int) []string {
	rowValues := make([]string, sr.ColumnCount())
	binding := sr.Bindings()[rowIndex]
	for columnIndex, varName := range sr.Variables() {
		rowValues[columnIndex] = binding[varName].Value
	}
	return rowValues
}

func (rs *ResultSet) appendRows(rows [][]string) [][]string {
	for i, _ := range rs.Results.Bindings {
		rows = append(rows, rs.Row(i))
	}
	return rows
}

func (rs *ResultSet) Rows() [][]string {
	var rows [][]string
	return rs.appendRows(rows)
}

func (rs *ResultSet) FormattedTable(columnSeparator bool) string {
	table := [][]string{rs.Head.Vars}
	table = rs.appendRows(table)
	return geist.WriteStringTable(table, columnSeparator)
}

func (sr *ResultSet) Column(columnIndex int) []string {
	variable := sr.Variables()[columnIndex]
	columnValues := make([]string, sr.RowCount())
	for rowIndex, binding := range sr.Bindings() {
		columnValues[rowIndex] = binding[variable].Value
	}
	return columnValues
}
