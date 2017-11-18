package clitable

import (
	"testing"
)

func TestTable(t *testing.T) {
	table := New()
	if len(table.cs) != 0 {
		t.Error("initialization error")
	}

	table.AddRow("One", "Two", "Three")
	if table.cs[0] != 3 || table.cs[1] != 3 || table.cs[2] != 5 {
		t.Error("Counting error")
	}

	table.AddRow("One", "Two", "MoreThanThree")
	if table.cs[0] != 3 || table.cs[1] != 3 || table.cs[2] != 13 {
		t.Error("re-Counting error")
	}

	table.AddRow("a", "a", "a")
	if table.cs[0] != 3 || table.cs[1] != 3 || table.cs[2] != 13 {
		t.Error("re-Counting error")
	}
}
