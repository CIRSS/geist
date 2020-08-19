package tables

type DataTable interface {
	ColumnCount() int
	RowCount() int
	Row(rowIndex int) []string
	Column(columnIndex int) []string
	Rows() [][]string
	FormattedTable(columnSeparator bool) string
}
