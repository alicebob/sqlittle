package db

import (
	"fmt"
	"testing"
)

func TestSelect(t *testing.T) {
	db, err := Open("../test/words.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	cb := func(r Row) { fmt.Printf("row: %+v\n", r) }
	if err := db.Select("words", cb, "length"); err != nil {
		t.Fatal(err)
	}

	// where := WhereEq{"length", 4}
	// err := db.SelectWhere("words", where, cb, "word")
}
