package sqlittle

import (
	"errors"
	"reflect"
	"testing"
)

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

func TestSelectRowidNonRowir(t *testing.T) {
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
