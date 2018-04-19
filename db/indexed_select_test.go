package db

import (
	"testing"
)

func TestIndexedSelectRowid(t *testing.T) {
	db, err := Open("../testdata/words.sqlite")
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
