package db

import (
	"testing"
)

// where := WhereEq{"length", 4}
// err := db.SelectWhere("words", where, cb, "word")

func TestSelect(t *testing.T) {
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
