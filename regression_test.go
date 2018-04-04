package sqlittle

import (
	"io/ioutil"
	"testing"
)

func TestIssue1(t *testing.T) {
	db, err := OpenFile("./test/issue_1.sqlite")
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
	f, err := ioutil.ReadFile("./test/issue_3.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	if have, want := Fuzz(f), 0; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIssue4(t *testing.T) {
	f, err := ioutil.ReadFile("./test/issue_4.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	if have, want := Fuzz(f), 0; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}
