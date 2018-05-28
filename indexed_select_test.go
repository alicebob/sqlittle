package sqlittle

import (
	"reflect"
	"testing"

	"github.com/andreyvit/diff"
	"github.com/davecgh/go-spew/spew"
)

func init() {
	spew.Config.DisablePointerAddresses = true
	spew.Config.DisableCapacities = true
	spew.Config.SortKeys = true
}

func TestIndexedSelectRowid(t *testing.T) {
	db, err := Open("testdata/words.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var words []string
	cb := func(r Row) {
		var w string
		if err := r.Scan(&w); err != nil {
			t.Fatal(err)
		}
		words = append(words, w)
	}
	if err := db.IndexedSelect("words", "words_index_1", cb, "wORd"); err != nil {
		t.Fatal(err)
	}
	if have, want := len(words), 1000; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := words[0], "Adams"; have != want {
		t.Errorf("have:\n%#v\nwant:\n%#v", have, want)
	}
	if have, want := words[999], "yeshivahs"; have != want {
		t.Errorf("have:\n%#v\nwant:\n%#v", have, want)
	}
}

func TestIndexedSelectWithoutRowid(t *testing.T) {
	db, err := Open("testdata/music.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var names []string
	cb := func(r Row) {
		var w string
		if err := r.Scan(&w); err != nil {
			t.Fatal(err)
		}
		names = append(names, w)
	}
	if err := db.IndexedSelect("tracks", "tracks_length", cb, "NAME"); err != nil {
		t.Fatal(err)
	}
	// SELECT name FROM tracks ORDER BY length
	want := []string{
		"Norwegian Wood",
		"Drive My Car",
		"Something",
		"You Wont See Me",
		"Maxwells Silver Hammer",
		"Come Together",
	}
	if have, want := names, want; !reflect.DeepEqual(have, want) {
		t.Errorf("diff:\n%s", diff.LineDiff(spew.Sdump(want), spew.Sdump(have)))
	}
}

func TestIndexedSelectEq(t *testing.T) {
	db, err := Open("testdata/music.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var words []string
	cb := func(r Row) {
		w, _ := r.ScanString()
		words = append(words, w)
	}
	if err := db.IndexedSelectEq("albums", "albums_name", Key{"Abbey Road"}, cb, "name"); err != nil {
		t.Fatal(err)
	}
	want := []string{"Abbey Road"}
	if have, want := words, want; !reflect.DeepEqual(have, want) {
		t.Errorf("diff:\n%s", diff.LineDiff(spew.Sdump(want), spew.Sdump(have)))
	}
}

func TestIndexedSelectEqNonRowid(t *testing.T) {
	db, err := Open("testdata/music.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var words []string
	cb := func(r Row) {
		w, _ := r.ScanString()
		words = append(words, w)
	}
	if err := db.IndexedSelectEq("tracks", "tracks_length", Key{int64(121)}, cb, "name"); err != nil {
		t.Fatal(err)
	}
	want := []string{"Norwegian Wood"}
	if have, want := words, want; !reflect.DeepEqual(have, want) {
		t.Errorf("diff:\n%s", diff.LineDiff(spew.Sdump(want), spew.Sdump(have)))
	}
}

func TestIndexedSelectDesc(t *testing.T) {
	// DESC column should be automatically detected
	db, err := Open("testdata/prefix.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var words []string
	cb := func(r Row) {
		w, _ := r.ScanString()
		words = append(words, w)
	}
	if err := db.IndexedSelectEq("words", "words_prefix_desc", Key{"thi"}, cb,
		"word"); err != nil {
		t.Fatal(err)
	}
	want := []string{
		"thinking",
		"thirteen",
		"thickest",
		"third's",
	}
	if have, want := words, want; !reflect.DeepEqual(have, want) {
		t.Errorf("diff:\n%s", diff.LineDiff(spew.Sdump(want), spew.Sdump(have)))
	}
}

func TestIndexedSelectExprWhere(t *testing.T) {
	// index with a WHERE expression
	db, err := Open("testdata/expr.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var words []string
	cb := func(r Row) {
		w, _ := r.ScanString()
		words = append(words, w)
	}
	if err := db.IndexedSelect(
		"expr",
		"expr_where",
		cb,
		"name",
	); err != nil {
		t.Fatal(err)
	}
	want := []string{
		"longestnameever",
		"qqq",
	}
	if have, want := words, want; !reflect.DeepEqual(have, want) {
		t.Errorf("diff:\n%s", diff.LineDiff(spew.Sdump(want), spew.Sdump(have)))
	}
}

func TestIndexedSelectExprCol(t *testing.T) {
	// index with an expression column
	db, err := Open("testdata/expr.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var words []string
	cb := func(r Row) {
		w, _ := r.ScanString()
		words = append(words, w)
	}
	if err := db.IndexedSelect(
		"expr",
		"expr_name",
		cb,
		"name",
	); err != nil {
		t.Fatal(err)
	}
	want := []string{
		"aap",
		"foo",
		"longestnameever", // substr() expression, but we get the value from the row
		"qqq",
	}
	if have, want := words, want; !reflect.DeepEqual(have, want) {
		t.Errorf("diff:\n%s", diff.LineDiff(spew.Sdump(want), spew.Sdump(have)))
	}
}
