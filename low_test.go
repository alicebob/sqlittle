package sqlittle

import (
	"testing"
)

func TestLowEmpty(t *testing.T) {
	// table without any row
	db, err := OpenFile("./test/empty.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	table, err := db.Table("foo")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := table.Rowid(42); err != nil {
		t.Fatal(err)
	}

	count := 0
	if err := table.Scan(func(int64, Record) bool {
		count++
		return false
	}); err != nil {
		t.Fatal(err)
	}
	if have, want := count, 0; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}
