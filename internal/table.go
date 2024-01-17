package internal

import (
	"github.com/hoehwa/gopkg/bubbletea/table"
)

// RenderTable render a table with the given data, columns and rows.
func RenderTable(cols table.Columns, rows table.Rows) {
	table.Init(cols, rows)
}
