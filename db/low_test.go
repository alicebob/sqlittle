package db

import (
	"reflect"
	"testing"

	"github.com/alicebob/sqlittle/sql"
	"github.com/andreyvit/diff"
	"github.com/davecgh/go-spew/spew"
)

func TestLowEmpty(t *testing.T) {
	// table without any row
	db, err := OpenFile("./../testdata/empty.sqlite")
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
	db, err := OpenFile("./../testdata/words.sqlite")
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

func TestLowScanEqNonrowid(t *testing.T) {
	db, err := OpenFile("./../testdata/withoutrowid.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	table, err := db.NonRowidTable("words")
	if err != nil {
		t.Fatal(err)
	}

	var found Record
	if err := table.ScanEq(
		Key{KeyCol{V: "crankiest"}},
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
	db, err := OpenFile("./../testdata/funkykey.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	table, err := db.NonRowidTable("fuz")
	if err != nil {
		t.Fatal(err)
	}

	var found Record
	if err := table.ScanEq(
		Key{KeyCol{V: "consequent"}, KeyCol{V: "allegory"}},
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

func TestLowScanEq2(t *testing.T) {
	db, err := OpenFile("./../testdata/prefix.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	index, err := db.Index("words_prefix")
	if err != nil {
		t.Fatal(err)
	}

	var found []Record
	if err := index.ScanEq(
		Key{KeyCol{V: "who"}},
		func(r Record) bool {
			found = append(found, r)
			return false
		}); err != nil {
		t.Fatal(err)
	}
	want := []Record{
		Record{"who", int64(285)}, // whoop
		Record{"who", int64(381)}, // whoa
		Record{"who", int64(920)}, // whorls
	}
	if have, want := found, want; !reflect.DeepEqual(have, want) {
		t.Errorf("diff:\n%s", diff.LineDiff(spew.Sdump(want), spew.Sdump(have)))
	}
}

func TestLowScanEqDesc(t *testing.T) {
	db, err := OpenFile("./../testdata/prefix.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	index, err := db.Index("words_prefix_desc")
	if err != nil {
		t.Fatal(err)
	}

	var found []Record
	if err := index.ScanEq(
		Key{KeyCol{V: "who", Desc: true}},
		func(r Record) bool {
			found = append(found, r)
			return false
		}); err != nil {
		t.Fatal(err)
	}
	// only the prefix is DESC; rowids are still ascending
	want := []Record{
		Record{"who", int64(285)}, // whoop
		Record{"who", int64(381)}, // whoa
		Record{"who", int64(920)}, // whorls
	}
	if have, want := found, want; !reflect.DeepEqual(have, want) {
		t.Errorf("diff:\n%s", diff.LineDiff(spew.Sdump(want), spew.Sdump(have)))
	}
}

func TestLowScanRange(t *testing.T) {
	db, err := OpenFile("./../testdata/words.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	index, err := db.Index("words_index_1")
	if err != nil {
		t.Fatal(err)
	}

	var found []Record
	if err := index.ScanRange(
		Key{KeyCol{V: "tradition"}},
		Key{KeyCol{V: "twofer"}},
		func(r Record) bool {
			found = append(found, r)
			return false
		}); err != nil {
		t.Fatal(err)
	}
	want := []Record{
		Record{"tradition", int64(50)},
		Record{"training", int64(258)},
		Record{"tranquiler", int64(384)},
		Record{"transfigures", int64(927)},
		Record{"tremulously", int64(389)},
		Record{"tribalism", int64(105)},
		Record{"trigger's", int64(962)},
		Record{"triteness", int64(738)},
		Record{"trustee's", int64(819)},
		Record{"trustworthiness", int64(731)},
		Record{"twiddled", int64(211)},
		Record{"twirls", int64(547)},
		// doesn't match "twofer"
	}
	if have, want := found, want; !reflect.DeepEqual(have, want) {
		t.Errorf("diff:\n%s", diff.LineDiff(spew.Sdump(want), spew.Sdump(have)))
	}
}

func TestLowScanRangeDesc(t *testing.T) {
	db, err := OpenFile("./../testdata/prefix.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	index, err := db.Index("words_prefix_desc")
	if err != nil {
		t.Fatal(err)
	}

	var found []Record
	if err := index.ScanRange(
		Key{KeyCol{V: "blu", Desc: true}},
		Key{KeyCol{V: "bla", Desc: true}},
		func(r Record) bool {
			found = append(found, r)
			return false
		}); err != nil {
		t.Fatal(err)
	}
	want := []Record{
		Record{"blu", int64(524)}, // blurb's
		Record{"blo", int64(183)},
		Record{"blo", int64(563)},
		Record{"bli", int64(608)},
		Record{"bli", int64(821)},
		Record{"ble", int64(69)}, // bleeps
	}
	if have, want := found, want; !reflect.DeepEqual(have, want) {
		t.Errorf("diff:\n%s", diff.LineDiff(spew.Sdump(want), spew.Sdump(have)))
	}
}
