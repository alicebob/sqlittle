package sqlit

import (
	"fmt"
	"reflect"
	"testing"
)

func TestVarint(t *testing.T) {
	for i, cas := range []struct {
		b []byte
		l int
		n int64
	}{
		{
			b: []byte("\x00"),
			l: 1,
			n: 0,
		},
		{
			b: []byte("\xFF\x00"),
			l: 2,
			n: 16256, // 0b00111111_10000000
		},
		{
			b: []byte("\xFF\x7F"),
			l: 2,
			n: 0x3FFF, // 0b00111111_11111111
		},
		{
			b: []byte("\xBF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"),
			l: 9,
			n: 0x7FFFFFFFFFFFFFFF,
		},
		{
			b: []byte("\xBF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFFignored"),
			l: 9,
			n: 0x7FFFFFFFFFFFFFFF,
		},
		{
			// int64 overflow
			b: []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"),
			l: 9,
			n: -1,
		},
	} {
		n, l := readVarint(cas.b)
		if have, want := l, cas.l; have != want {
			t.Errorf("case %d: have %d, want %d", i, have, want)
		}
		if have, want := n, cas.n; have != want {
			t.Errorf("case %d: have %d, want %d", i, have, want)
		}
	}
}

func TestTablesSingle(t *testing.T) {
	f, err := openFile("./test/single.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	master, err := f.pageMaster()
	if err != nil {
		t.Fatal(err)
	}

	rows, err := master.Rows(f)
	if err != nil {
		t.Fatal(err)
	}
	if have, want := rows, 1; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestTablesFour(t *testing.T) {
	db, err := openFile("./test/four.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	master, err := db.pageMaster()
	if err != nil {
		t.Fatal(err)
	}

	rowCount, err := master.Rows(db)
	if err != nil {
		t.Fatal(err)
	}
	if have, want := rowCount, 4; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}

	aap, err := db.Table("aap")
	if err != nil {
		t.Fatal(err)
	}
	if aap == nil {
		t.Fatal("no table found")
	}
	var rows []interface{}
	if _, err := aap.root.Iter(
		db,
		func(rowid int64, pl Payload) (bool, error) {
			c, err := addOverflow(db, pl)
			if err != nil {
				return false, err
			}
			e, err := parseRecord(c)
			if err != nil {
				return false, err
			}
			rows = append(rows, e)
			return false, nil
		}); err != nil {
		t.Fatal(err)
	}
	if have, want := rows, []interface{}{
		[]interface{}{"world"},
		[]interface{}{"universe"},
		[]interface{}{"town"},
	}; !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestTableLong(t *testing.T) {
	// starts table interior page
	db, err := openFile("./test/long.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	bottles, err := db.Table("bottles")
	if err != nil {
		t.Fatal(err)
	}
	if bottles == nil {
		t.Fatal("no table found")
	}

	rowCount, err := bottles.root.Rows(db)
	if err != nil {
		t.Fatal(err)
	}
	if have, want := rowCount, 1000; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}

	var rows []interface{}
	if _, err := bottles.root.Iter(
		db,
		func(rowid int64, pl Payload) (bool, error) {
			c, err := addOverflow(db, pl)
			if err != nil {
				return false, err
			}
			e, err := parseRecord(c)
			if err != nil {
				return false, err
			}
			if rowid == 42 {
				rows = append(rows, e)
				return true, nil
			}
			return false, nil
		}); err != nil {
		t.Fatal(err)
	}
	if have, want := rows, []interface{}{
		[]interface{}{"bottles of beer on the wall 42"},
	}; !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestTableOverflow(t *testing.T) {
	// record overflow
	testline := ""
	for i := 1; ; i++ {
		testline += fmt.Sprintf("%d", i)
		if i == 1000 {
			break
		}
		testline += "longline"
	}

	db, err := openFile("./test/overflow.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mytable, err := db.Table("mytable")
	if err != nil {
		t.Fatal(err)
	}
	if mytable == nil {
		t.Fatal("no table found")
	}

	rowCount, err := mytable.root.Rows(db)
	if err != nil {
		t.Fatal(err)
	}
	if have, want := rowCount, 1; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}

	var rows []interface{}
	if _, err := mytable.root.Iter(
		db,
		func(rowid int64, pl Payload) (bool, error) {
			c, err := addOverflow(db, pl)
			if err != nil {
				return false, err
			}
			e, err := parseRecord(c)
			if err != nil {
				return false, err
			}
			rows = append(rows, e)
			return false, nil
		}); err != nil {
		t.Fatal(err)
	}
	if have, want := rows, []interface{}{
		[]interface{}{testline},
	}; !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestTableValues(t *testing.T) {
	// different value types
	db, err := openFile("./test/values.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	things, err := db.Table("things")
	if err != nil {
		t.Fatal(err)
	}
	if things == nil {
		t.Fatal("no table found")
	}

	rowCount, err := things.root.Rows(db)
	if err != nil {
		t.Fatal(err)
	}
	if have, want := rowCount, 17; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}

	var rows []Row
	if _, err := things.root.Iter(
		db,
		func(rowid int64, pl Payload) (bool, error) {
			c, err := addOverflow(db, pl)
			if err != nil {
				return false, err
			}
			e, err := parseRecord(c)
			if err != nil {
				return false, err
			}
			rows = append(rows, e)
			return false, nil
		}); err != nil {
		t.Fatal(err)
	}
	if have, want := rows, []Row{
		{nil, int64(0), int64(0)},
		{"", int64(1), int64(0)},
		{"", int64(0), int64(0)},
		{"", int64(80), int64(0)},
		{"", -int64(80), int64(0)},
		{"", int64(1 << 14), int64(0)},
		{"", -int64(1 << 14), int64(0)},
		{"", int64(1 << 20), int64(0)},
		{"", -int64(1 << 20), int64(0)},
		{"", int64(1 << 30), int64(0)},
		{"", -int64(1 << 30), int64(0)},
		{"", int64(1 << 42), int64(0)},
		{"", -int64(1 << 42), int64(0)},
		{"", int64(1 << 53), int64(0)},
		{"", -int64(1 << 53), int64(0)},
		{"", int64(0), float64(3.14)},
		{"", -int64(0), -float64(3.14)},
	}; !reflect.DeepEqual(have, want) {
		t.Errorf("have:\n%#v\nwant:\n%#v", have, want)
	}
}

func TestIndexSingle(t *testing.T) {
	// scan a whole index (single page)
	db, err := openFile("./test/index.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	hello, err := db.Index("hello_index")
	if err != nil {
		t.Fatal(err)
	}
	if hello == nil {
		t.Fatal("no index found")
	}

	rowCount, err := hello.root.Rows(db)
	if err != nil {
		t.Fatal(err)
	}
	if have, want := rowCount, 3; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}

	var rows []Row
	if _, err := hello.root.Iter(
		db,
		func(pl Payload) (bool, error) {
			pf, err := addOverflow(db, pl)
			if err != nil {
				return false, err
			}
			_, row, err := parseIndexRow(pf)
			rows = append(rows, row)
			return false, err
		}); err != nil {
		t.Fatal(err)
	}
	if have, want := rows, []Row{
		{"town"},
		{"universe"},
		{"world"},
	}; !reflect.DeepEqual(have, want) {
		t.Errorf("have:\n%#v\nwant:\n%#v", have, want)
	}
}

func TestIndexWords(t *testing.T) {
	// scan a whole index, with internal page
	db, err := openFile("./test/words.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	index, err := db.Index("words_index_1")
	if err != nil {
		t.Fatal(err)
	}
	if index == nil {
		t.Fatal("no index found")
	}

	rowCount, err := index.root.Rows(db)
	if err != nil {
		t.Fatal(err)
	}
	if have, want := rowCount, 1000; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}

	var rows []Row
	if _, err := index.root.Iter(
		db,
		func(pl Payload) (bool, error) {
			pf, err := addOverflow(db, pl)
			if err != nil {
				return false, err
			}
			_, row, err := parseIndexRow(pf)
			rows = append(rows, row)
			return false, err
		}); err != nil {
		t.Fatal(err)
	}
	if have, want := len(rows), 1000; have != want {
		t.Errorf("have:\n%#v\nwant:\n%#v", have, want)
	}
}
