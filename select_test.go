package sqlittle

import (
	"errors"
	"reflect"
	"testing"

	"github.com/andreyvit/diff"
	"github.com/davecgh/go-spew/spew"
)

func TestSelectCols(t *testing.T) {
	db, err := Open("testdata/words.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if have, want := db.Select("words", nil, "word", "nosuch"), errors.New(`no such column: "nosuch"`); !reflect.DeepEqual(have, want) {
		t.Errorf("have %v, want %v", have, want)
	}
}

func TestSelectSimple(t *testing.T) {
	db, err := Open("testdata/words.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var words []string
	cb := func(r Row) {
		var w string
		if err := r.Scan(nil, &w); err != nil {
			t.Fatal(err)
		}
		words = append(words, w)
	}
	if err := db.Select("words", cb, "length", "wORd"); err != nil {
		t.Fatal(err)
	}
	if have, want := len(words), 1000; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := words[0], "hangdog"; have != want {
		t.Errorf("have:\n%#v\nwant:\n%#v", have, want)
	}
	if have, want := words[999], "ideologist"; have != want {
		t.Errorf("have:\n%#v\nwant:\n%#v", have, want)
	}
}

func TestSelectWithoutRowid(t *testing.T) {
	db, err := Open("testdata/withoutrowid.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var words []string
	cb := func(r Row) {
		var w string
		if err := r.Scan(nil, &w); err != nil {
			t.Fatal(err)
		}
		words = append(words, w)
	}
	if err := db.Select("words", cb, "length", "word"); err != nil {
		t.Fatal(err)
	}
	if have, want := len(words), 1000; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	// without rowids are ordered
	if have, want := words[0], "Adams"; have != want {
		t.Errorf("have:\n%#v\nwant:\n%#v", have, want)
	}
	if have, want := words[999], "yeshivahs"; have != want {
		t.Errorf("have:\n%#v\nwant:\n%#v", have, want)
	}
}

func TestSelectAlter(t *testing.T) {
	// table has a column added with a DEFAULT. These values won't be present
	// in the Row values.
	db, err := Open("testdata/alter.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var count = 0
	cb := func(r Row) {
		var n int
		if err := r.Scan(&n); err != nil {
			t.Fatal(err)
		}
		if have, want := n, 42; have != want {
			t.Fatalf("have:\n%#v\nwant:\n%#v", have, want)
		}
		count++
	}
	if err := db.Select("words", cb, "something"); err != nil {
		t.Fatal(err)
	}
	if have, want := count, 1000; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}

func TestSelectColumnRowid(t *testing.T) {
	// special column named "rowid"
	db, err := Open("testdata/words.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var count int64 = 1
	cb := func(r Row) {
		var n [4]int64
		if err := r.Scan(nil, &n[0], &n[1], &n[2], &n[3]); err != nil {
			t.Fatal(err)
		}
		if have, want := n, [...]int64{count, count, count, count}; have != want {
			t.Errorf("have %v, want %v", have, want)
		}
		count++
	}
	if err := db.Select("words", cb, "word", "rowid", "oid", "_rowid_", "rOwId"); err != nil {
		t.Fatal(err)
	}
}

func TestSelectFunky(t *testing.T) {
	// funkykey has columns in a different order than the definition
	db, err := Open("testdata/funkykey.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var rows [][]string
	cb := func(r Row) {
		var w [4]string
		if err := r.Scan(&w[0], &w[1], &w[2], &w[3]); err != nil {
			t.Fatal(err)
		}
		rows = append(rows, w[:])
	}
	if err := db.Select("fuz", cb, "a", "b", "c", "d"); err != nil {
		t.Fatal(err)
	}
	// ordered by (c, a)
	want := [][]string{
		[]string{"algebraic", "begotten", "colder", "destinies"},
		[]string{"allegory", "beagle", "consequent", "duffers"},
		[]string{"angle", "billiards", "crotchety", "delta"},
	}
	if have, want := rows, want; !reflect.DeepEqual(have, want) {
		t.Errorf("diff:\n%s", diff.LineDiff(spew.Sdump(want), spew.Sdump(have)))
	}
}

func TestSelectRowid(t *testing.T) {
	db, err := Open("testdata/words.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	row, err := db.SelectRowid("words", 42, "word", "length")
	if err != nil {
		t.Fatal(err)
	}
	var word string
	var l int
	if err := row.Scan(&word, &l); err != nil {
		t.Fatal(err)
	}
	if have, want := word, "aniseed"; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
	if have, want := l, 7; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}

func TestSelectRowidNonRowid(t *testing.T) {
	db, err := Open("testdata/withoutrowid.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	_, err = db.SelectRowid("words", 42, "word", "length")
	if have, want := err, errors.New("can't use SelectRowid on a WITHOUT ROWID table"); !reflect.DeepEqual(have, want) {
		t.Errorf("have %v, want %v", have, want)
	}
}

func TestPKSelectNoPK(t *testing.T) {
	db, err := Open("testdata/single.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if have, want := db.PKSelect("hello", Row{"foo"}, nil, "col"), errors.New(`table has no primary key`); !reflect.DeepEqual(have, want) {
		t.Errorf("have %v, want %v", have, want)
	}
}

func TestPKSelect(t *testing.T) {
	db, err := Open("testdata/primarykey.sqlite")
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
	db, err := Open("testdata/funkykey.sqlite")
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
