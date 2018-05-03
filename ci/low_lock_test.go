// +build ci

package ci

import (
	"fmt"
	"reflect"
	"testing"

	sdb "github.com/alicebob/sqlittle/db"
)

// locking picks up changes in the file
func TestLockReload(t *testing.T) {
	file, close := tmpfile(t)
	defer close()

	if _, err := sqlite(file, "CREATE TABLE number (n)"); err != nil {
		t.Fatal(err)
	}

	little, err := sdb.OpenFile(file)
	if err != nil {
		t.Fatal(err)
	}
	defer little.Close()

	loadNumbers := func() []string {
		if err := little.RLock(); err != nil {
			t.Fatal(err)
		}
		defer func() {
			if err := little.RUnlock(); err != nil {
				t.Fatal(err)
			}
		}()

		table, err := little.Table("number")
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
}

// A readlock should make sqlite's write fail
func TestLockWrite(t *testing.T) {
	file, close := tmpfile(t)
	defer close()

	// Locking and unlocking files in the same process is a bit of a mess.  We
	// call out to the sqlite3 binary to make sure sqlittle's lock mumbling and
	// sqlite's lock mumbling don't interfere.
	if _, err := sqlite(file, "CREATE TABLE number (n)"); err != nil {
		t.Fatal(err)
	}

	little, err := sdb.OpenFile(file)
	if err != nil {
		t.Fatal(err)
	}
	defer little.Close()

	if err := little.RLock(); err != nil {
		t.Fatal(err)
	}

	// insert via sqlite should fail with SQLITE_BUSY (5)
	if _, err := sqlite(file, `INSERT INTO number VALUES ("one")`); err == nil {
		t.Fatal("expected an error")
	} else if have, want := err.Error(), "exit status 5: Error: database is locked\n"; have != want {
		t.Fatalf("have %#v, want %#v", have, want)
	}

	if err := little.RUnlock(); err != nil {
		t.Fatal(err)
	}

	// sqlite insert should be fine now
	if _, err := sqlite(file, `INSERT INTO number VALUES ("two")`); err != nil {
		t.Fatal(err)
	}

	// sqlittle should pick up those rows
}
