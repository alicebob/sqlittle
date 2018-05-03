// +build ci

package ci

import (
	"fmt"
	"reflect"
	"testing"

	sdb "github.com/alicebob/sqlittle/db"
)

// change the pagesize
func TestPagesize(t *testing.T) {
	file, close := tmpfile(t)
	defer close()

	if _, err := sqlite(file, "CREATE TABLE number (n)"); err != nil {
		t.Fatal(err)
	}
	if _, err := sqlite(file, "PRAGMA page_size=512"); err != nil {
		t.Fatal(err)
	}

	db, err := sdb.OpenFile(file)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	loadNumbers := func() []string {
		if err := db.RLock(); err != nil {
			t.Fatal(err)
		}
		defer func() {
			if err := db.RUnlock(); err != nil {
				t.Fatal(err)
			}
		}()

		table, err := db.Table("number")
		if err != nil {
			t.Fatal(err)
		}
		var found []string
		if err := table.Scan(func(_ int64, r sdb.Record) bool {
			found = append(found, r[0].(string))
			return false
		}); err != nil {
			t.Fatal(err)
		}
		return found
	}

	// nothing in the table yet
	if have, want := loadNumbers(), []string(nil); !reflect.DeepEqual(have, want) {
		t.Fatalf("have %#v, want %#v", have, want)
	}

	// insert via sqlite
	words := []string{"one", "two", "three"}
	for _, n := range words {
		if _, err := sqlite(file, fmt.Sprintf(`INSERT INTO number VALUES (%q)`, n)); err != nil {
			t.Fatal(err)
		}
	}

	// sqlittle should pick up those rows
	if have, want := loadNumbers(), words; !reflect.DeepEqual(have, want) {
		t.Fatalf("have %#v, want %#v", have, want)
	}

	// change table size
	if _, err := sqlite(file, "PRAGMA page_size=16384"); err != nil {
		t.Fatal(err)
	}

	// new pagesize should be picked up
	if have, want := loadNumbers(), words; !reflect.DeepEqual(have, want) {
		t.Fatalf("have %#v, want %#v", have, want)
	}
}
