package sqlittle

import (
	"reflect"
	"testing"
)

func TestDBWal(t *testing.T) {
	file := "./test/wal_crashed.sqlite"
	db, err := OpenFile(file)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if err := db.RLock(); err != nil {
		t.Fatal(err)
	}
	defer db.RUnlock()
	m, err := db.master()
	if err != nil {
		t.Fatal(err)
	}
	if have, want := m, []sqliteMaster{
		{
			typ:      "table",
			name:     "words",
			tblName:  "words",
			rootPage: 2,
			sql:      "CREATE TABLE words (word varchar)",
		},
	}; !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, want %#v", have, want)
	}
	// if have, want := db.Tables, []string{"words"}; !reflect.DeepEqual(have, want) {
	// t.Errorf("have %#v, want %#v", have, want)
	// }
	// if have, want := db.Indexes, []string(nil); !reflect.DeepEqual(have, want) {
	// t.Errorf("have %#v, want %#v", have, want)
	// }
}
