package sqlittle

import (
	"testing"
)

func TestJournal(t *testing.T) {
	for file, expect := range map[string]bool{
		"./testdata/journal_truncate.sqlite-journal": false,
		"./testdata/journal_persist.sqlite-journal":  false,
		"./testdata/journal_hot.sqlite-journal":      true,
		"./testdata/nosuch":                          false,
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
		"./testdata/journal_truncate.sqlite": nil,
		"./testdata/journal_persist.sqlite":  nil,
		"./testdata/journal_hot.sqlite":      ErrHotJournal,
	} {
		db, err := OpenFile(file)
		if have, want := err, expect; have != want {
			t.Errorf("have %#v, want %#v", have, want)
			continue
		}
		db.Close()
	}
}
