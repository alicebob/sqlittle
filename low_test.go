package sqlittle

import (
	"reflect"
	"testing"

	"github.com/alicebob/sqlittle/sql"
)

func TestLowEmpty(t *testing.T) {
	// table without any row
	db, err := OpenFile("./testdata/empty.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	table, err := db.Table("foo")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := table.Rowid(42); err != nil {
		t.Fatal(err)
	}

	count := 0
	if err := table.Scan(func(int64, Record) bool {
		count++
		return false
	}); err != nil {
		t.Fatal(err)
	}
	if have, want := count, 0; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestLowDefs(t *testing.T) {
	db, err := OpenFile("./testdata/words.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	table, err := db.Table("words")
	if err != nil {
		t.Fatal(err)
	}
	def, err := table.Def()
	if err != nil {
		t.Fatal(err)
	}
	if have, want := def, (&sql.CreateTableStmt{
		Table: "words",
		Columns: []sql.ColumnDef{
			{Name: "word", Type: "varchar", Null: true},
			{Name: "length", Type: "int", Null: true},
		},
	}); !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, want %#v", have, want)
	}

	index, err := db.Index("words_index_2")
	if err != nil {
		t.Fatal(err)
	}
	idef, err := index.Def()
	if err != nil {
		t.Fatal(err)
	}
	if have, want := idef, (&sql.CreateIndexStmt{
		Index: "words_index_2",
		Table: "words",
		IndexedColumns: []sql.IndexedColumn{
			{Column: "length"},
			{Column: "word"},
		},
	}); !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestLowScanEq(t *testing.T) {
	db, err := OpenFile("./testdata/withoutrowid.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	table, err := db.WithoutRowidTable("words")
	if err != nil {
		t.Fatal(err)
	}

	var found Record
	if err := table.ScanEq(
		Record{"crankiest"},
		func(r Record) bool {
			found = r
			return false
		}); err != nil {
		t.Fatal(err)
	}
	if have, want := found, (Record{"crankiest", int64(9)}); !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestLowWithoutRowid2(t *testing.T) {
	db, err := OpenFile("./testdata/funkykey.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	table, err := db.WithoutRowidTable("fuz")
	if err != nil {
		t.Fatal(err)
	}

	var found Record
	if err := table.ScanEq(
		Record{"consequent", "allegory"},
		func(r Record) bool {
			found = r
			return false
		},
	); err != nil {
		t.Fatal(err)
	}
	// note that this is not the column order
	if have, want := found, (Record{"consequent", "allegory", "beagle", "duffers"}); !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, want %#v", have, want)
	}
}
