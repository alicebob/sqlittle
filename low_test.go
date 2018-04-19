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

func TestLowWithoutRowid(t *testing.T) {
	db, err := OpenFile("./testdata/withoutrowid.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	table, err := db.Table("words")
	if err != nil {
		t.Fatal(err)
	}
	rows := 0
	table.WithoutRowidScan(func(r Record) bool {
		rows++
		return false
	})
	if have, want := rows, 1000; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}

	row, err := table.WithoutRowidPK(Record{"crankiest"})
	if err != nil {
		t.Fatal(err)
	}
	if have, want := row, (Record{"crankiest", int64(9)}); !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestLowWithoutRowid2(t *testing.T) {
	db, err := OpenFile("./testdata/funkykey.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	table, err := db.Table("fuz")
	if err != nil {
		t.Fatal(err)
	}

	row, err := table.WithoutRowidPK(Record{"consequent", "allegory"})
	if err != nil {
		t.Fatal(err)
	}
	// note that this is not the column order
	if have, want := row, (Record{"consequent", "allegory", "beagle", "delta"}); !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, want %#v", have, want)
	}
}
