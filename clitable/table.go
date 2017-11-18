package clitable

import (
	"fmt"
	"io"
	"os"
)

type colSize map[int]int
type row []string

// Table represents the table to be printed
type Table struct {
	cs   colSize
	rows []row
}

func (c colSize) setIfBigger(k, v int) {
	cv, e := c[k]
	if (e && cv < v) || !e {
		c[k] = v
	}
}

// New creates a new table
func New() *Table {
	t := new(Table)
	t.cs = make(colSize)
	return t
}

// AddRow adds a row to the table
func (t *Table) AddRow(cols ...string) *Table {
	var r row
	for i, col := range cols {
		t.cs.setIfBigger(i, len(col))
		r = append(r, col)
	}
	t.rows = append(t.rows, r)
	return t
}

// Print prints the table to standard output
func (t *Table) Print() {
	t.Fprint(os.Stdout)
}

// Errprint prints the table to error output
func (t *Table) Errprint() {
	t.Fprint(os.Stderr)
}

// Fprint prints the table to any io.Writer of your choice
func (t *Table) Fprint(w io.Writer) {
	t.printEmpty(w)
	for _, row := range t.rows {
		t.printRow(row, w)
		t.printEmpty(w)
	}
}

func (t *Table) printRow(row row, w io.Writer) {
	for k, col := range row {
		c := string(col)
		fmt.Fprint(w, "|", pad(c+"", t.cs[k]+1, ' '))
	}
	fmt.Fprintf(w, "|\n")
}

func (t *Table) printEmpty(w io.Writer) {
	for i := 0; i < len(t.cs); i++ {
		fmt.Fprint(w, "+"+pad("", t.cs[i]+1, '-'))
	}
	fmt.Fprintf(w, "+\n")
}

func pad(str string, dlen int, padchar rune) string {
	if len(str) < dlen {
		diff := dlen - len(str)
		app := str
		for i := 0; i < diff; i++ {
			app += string(padchar)
		}
		return app
	}
	return str
}
