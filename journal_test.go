package sqlittle

import (
	"testing"
)

func TestJournal(t *testing.T) {
	for file, expect := range map[string]bool{
		"./test/journal_truncate.sqlite-journal": false,
		"./test/journal_persist.sqlite-journal":  false,
		"./test/journal_hot.sqlite-journal":      true,
		"./test/nosuch":                          false,
	} {
		valid, err := validJournal(file)
		if err != nil {
			t.Fatal(err)
		}
		if have, want := valid, expect; have != want {
			t.Errorf("%q: have %v, want %v", file, have, want)
		}
	}
}

func TestOpenHot(t *testing.T) {
	for file, expect := range map[string]error{
		"./test/journal_truncate.sqlite": nil,
		"./test/journal_persist.sqlite":  nil,
		"./test/journal_hot.sqlite":      ErrHotJournal,
	} {
		db, err := OpenFile(file)
		if have, want := err, expect; have != want {
			t.Errorf("have %#v, want %#v", have, want)
			continue
		}
		db.Close()
	}
}
