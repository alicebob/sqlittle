// +build ci

package ci

import (
	"fmt"
	"testing"
	"time"

	sdb "github.com/alicebob/sqlittle/db"
)

func TestJournal(t *testing.T) {
	file, close := tmpfile(t)
	defer close()

	proc := openSqlite(t, file)
	defer proc.Close()

	proc.write(t, "CREATE TABLE number (n);\n")
	proc.write(t, "PRAGMA cache_size=5;\n")
	proc.write(t, "BEGIN;\n")
	for i := 0; i < 1000; i++ {
		proc.write(t, fmt.Sprintf("INSERT INTO number VALUES (\"number %d\");\n", i))
	}
	time.Sleep(1 * time.Second)

	// we should be able open the file in sqlittle
	db, err := sdb.OpenFile(file)
	if err != nil {
		t.Fatal(err)
	}
	db.Close()

	// -9! This should keep the -journal file around
	proc.Kill()

	time.Sleep(1 * time.Second)

	// journal file, but no more lock...
	_, err = sdb.OpenFile(file)
	if have, want := err, sdb.ErrHotJournal; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}
}
