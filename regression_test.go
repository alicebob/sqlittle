package sqlittle

import (
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
