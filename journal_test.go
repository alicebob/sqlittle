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
