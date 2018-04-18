package db

import (
	"errors"
	"reflect"
	"testing"
)

// where := WhereEq{"length", 4}
// err := db.SelectWhere("words", where, cb, "word")

func TestSelectCols(t *testing.T) {
	db, err := Open("../testdata/words.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if have, want := db.Select("words", nil, "word", "nosuch"), errors.New(`no such column: "nosuch"`); !reflect.DeepEqual(have, want) {
		t.Errorf("have %v, want %v", have, want)
	}
}

func TestSelectSimple(t *testing.T) {
	db, err := Open("../testdata/words.sqlite")
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

	if have, want := words[0], "hangdog"; have != want {
		t.Errorf("have:\n%#v\nwant:\n%#v", have, want)
	}
	if have, want := words[999], "ideologist"; have != want {
		t.Errorf("have:\n%#v\nwant:\n%#v", have, want)
	}
}

func TestSelectWithoutRowid(t *testing.T) {
	db, err := Open("../testdata/withoutrowid.sqlite")
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
	db, err := Open("../testdata/alter.sqlite")
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
	db, err := Open("../testdata/words.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var count int64 = 1
	cb := func(r Row) {
		var n [3]int64
		if err := r.Scan(nil, &n[0], &n[1], &n[2]); err != nil {
			t.Fatal(err)
		}
		if have, want := n, [3]int64{count, count, count}; have != want {
			t.Errorf("have %v, want %v", have, want)
		}
		count++
	}
	if err := db.Select("words", cb, "word", "rowid", "oid", "_rowid_"); err != nil {
		t.Fatal(err)
	}
}
