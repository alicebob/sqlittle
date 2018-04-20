package db

import (
	"reflect"
	"testing"

	"github.com/andreyvit/diff"
	"github.com/davecgh/go-spew/spew"
)

func TestIndexedSelectEq(t *testing.T) {
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
	if err := db.IndexedSelectEq("albums", "albums_name", Row{"Abbey Road"}, cb, "name"); err != nil {
		t.Fatal(err)
	}
	want := []string{"Abbey Road"}
	if have, want := words, want; !reflect.DeepEqual(have, want) {
		t.Errorf("diff:\n%s", diff.LineDiff(spew.Sdump(want), spew.Sdump(have)))
	}
}
