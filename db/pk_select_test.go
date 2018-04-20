package db

import (
	"errors"
	"reflect"
	"testing"

	"github.com/andreyvit/diff"
	"github.com/davecgh/go-spew/spew"
)

func TestPKSelectNoPK(t *testing.T) {
	db, err := Open("../testdata/single.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if have, want := db.PKSelect("hello", Row{"foo"}, nil, "col"), errors.New(`table has no primary key`); !reflect.DeepEqual(have, want) {
		t.Errorf("have %v, want %v", have, want)
	}
}

func TestPKSelect(t *testing.T) {
	db, err := Open("../testdata/primarykey.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var rows []Row
	cb := func(r Row) {
		rows = append(rows, r)
	}
	if err := db.PKSelect("words", Row{"twofer"}, cb, "word", "rowid"); err != nil {
		t.Fatal(err)
	}
	want := []Row{Row{"twofer", int64(832)}}
	if have, want := rows, want; !reflect.DeepEqual(have, want) {
		t.Errorf("diff:\n%s", diff.LineDiff(spew.Sdump(want), spew.Sdump(have)))
	}
}

func TestPKSelectRowid(t *testing.T) {
	// PK is a rowid alias
	db, err := Open("../testdata/music.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var words []string
	cb := func(r Row) {
		w, _ := r.ScanString()
		words = append(words, w)
	}
	if err := db.PKSelect("albums", Row{int64(2)}, cb, "name"); err != nil {
		t.Fatal(err)
	}
	want := []string{"Abbey Road"}
	if have, want := words, want; !reflect.DeepEqual(have, want) {
		t.Errorf("diff:\n%s", diff.LineDiff(spew.Sdump(want), spew.Sdump(have)))
	}
}

func TestPKSelectNonRowid(t *testing.T) {
	// PKSelect() on a non-rowid table
	db, err := Open("../testdata/funkykey.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var words []string
	cb := func(r Row) {
		w, _ := r.ScanString()
		words = append(words, w)
	}
	if err := db.PKSelect("fuz", Row{"colder"}, cb, "d"); err != nil {
		t.Fatal(err)
	}
	want := []string{"destinies"}
	if have, want := words, want; !reflect.DeepEqual(have, want) {
		t.Errorf("diff:\n%s", diff.LineDiff(spew.Sdump(want), spew.Sdump(have)))
	}
}
