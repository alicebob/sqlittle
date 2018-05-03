// +build ci

package ci

import (
	"reflect"
	"testing"

	sdb "github.com/alicebob/sqlittle/db"
	"github.com/alicebob/sqlittle/sql"
)

// table definitions are read and changes are picked up
func TestDefs(t *testing.T) {
	file, close := tmpfile(t)
	defer close()

	if _, err := sqlite(file, "CREATE TABLE foo (n int)"); err != nil {
		t.Fatal(err)
	}
	if _, err := sqlite(file, "CREATE INDEX foo_index ON foo (n DESC)"); err != nil {
		t.Fatal(err)
	}

	db, err := sdb.OpenFile(file)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	{
		db.RLock()
		table, err := db.Table("foo")
		if err != nil {
			t.Fatal(err)
		}
		ct, err := table.Def()
		if err != nil {
			t.Fatal(err)
		}
		if have, want := ct, (&sql.CreateTableStmt{
			Table: "foo",
			Columns: []sql.ColumnDef{
				{Name: "n", Type: "int", Null: true},
			},
		}); !reflect.DeepEqual(have, want) {
			t.Fatalf("have %#v, want %#v", have, want)
		}

		index, err := db.Index("foo_index")
		if err != nil {
			t.Fatal(err)
		}
		ci, err := index.Def()
		if err != nil {
			t.Fatal(err)
		}
		if have, want := ci, (&sql.CreateIndexStmt{
			Index: "foo_index",
			Table: "foo",
			IndexedColumns: []sql.IndexedColumn{
				{Column: "n", SortOrder: sql.Desc},
			},
		}); !reflect.DeepEqual(have, want) {
			t.Fatalf("have %#v, want %#v", have, want)
		}
		db.RUnlock()
	}

	if _, err := sqlite(file, "ALTER TABLE foo ADD COLUMN o varchar"); err != nil {
		t.Fatal(err)
	}

	{
		db.RLock()
		table, err := db.Table("foo")
		if err != nil {
			t.Fatal(err)
		}
		ct, err := table.Def()
		if err != nil {
			t.Fatal(err)
		}
		if have, want := ct, (&sql.CreateTableStmt{
			Table: "foo",
			Columns: []sql.ColumnDef{
				{Name: "n", Type: "int", Null: true},
				{Name: "o", Type: "varchar", Null: true},
			},
		}); !reflect.DeepEqual(have, want) {
			t.Fatalf("have %#v, want %#v", have, want)
		}
		db.RUnlock()
	}
}
