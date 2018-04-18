package sqlittle

import (
	"io/ioutil"
	"testing"
)

func TestIssue1(t *testing.T) {
	// go-fuzz: invalid table definition
	db, err := OpenFile("./testdata/issue_1.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	_, err = db.Table("a")
	if have, want := err, ErrInvalidDef; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIssue3(t *testing.T) {
	// go-fuzz: file has a `file change counter` with value 0
	f, err := ioutil.ReadFile("./testdata/issue_3.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	if have, want := Fuzz(f), 0; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIssue4(t *testing.T) {
	// go-fuzz: there is a cell which is longer than the payload, which points
	// to page 0
	f, err := ioutil.ReadFile("./testdata/issue_4.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	if have, want := Fuzz(f), 0; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIssue5(t *testing.T) {
	// go-fuzz: internal table btree contains a pointer to itself
	f, err := ioutil.ReadFile("./testdata/issue_5.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	if have, want := fuzz(f), ErrRecursion; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIssue7(t *testing.T) {
	// go-fuzz: negative payload length
	f, err := ioutil.ReadFile("./testdata/issue_7.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	if have, want := fuzz(f), ErrCorrupted; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}
