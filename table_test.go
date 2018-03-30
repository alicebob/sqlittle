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

func TestRecord(t *testing.T) {
	for i, cas := range []struct {
		e    string
		want []interface{}
	}{
		{
			e: "\x06\x17\x17\x17\x01Wtablehellohello\x02CREATE TABLE hello (who varchar(255))",
			want: []interface{}{
				"table",
				"hello",
				"hello",
				int64(2),
				"CREATE TABLE hello (who varchar(255))",
			},
		},
		{
			// type 8: int 0
			e:    "\x02\b",
			want: []interface{}{int64(0)},
		},
		{
			// type 9: int 1
			e:    "\x02\t",
			want: []interface{}{int64(1)},
		},
		{
			// type 1: 8 bit
			e:    "\x02\x01P",
			want: []interface{}{int64(80)},
		},
		{
			// type 1: 8 bit
			e:    "\x02\x01\xb0",
			want: []interface{}{-int64(80)},
		},
		{
			// type 2: 16 bit
			e:    "\x02\x02@\x00",
			want: []interface{}{int64(1 << 14)},
		},
		{
			// type 2: 16 bit
			e:    "\x02\x02\xc0\x00",
			want: []interface{}{-int64(1 << 14)},
		},
		{
			// type 3: 24 bit
			e:    "\x02\x03\x7f\x00\x00",
			want: []interface{}{int64(0x7f0000)},
		},
		{
			// type 3: 24 bit
			e:    "\x02\x03\xff\xff\xff",
			want: []interface{}{-int64(1)},
		},
		{
			// type 4: 32 bit
			e:    "\x02\x04\x7f\x00\x00\x00",
			want: []interface{}{int64(0x7f000000)},
		},
		{
			// type 4: 32 bit
			e:    "\x02\x04\xff\xff\xff\xff",
			want: []interface{}{-int64(1)},
		},
		{
			// type 5: 48 bit
			e:    "\x02\x05\x7f\x00\x00\x00\x00\x00",
			want: []interface{}{int64(0x7f0000000000)},
		},
		{
			// type 5: 48 bit
			e:    "\x02\x05\xff\xff\xff\xff\xff\xff",
			want: []interface{}{-int64(1)},
		},
		{
			// type 6: 64 bit
			e:    "\x02\x06\x7f\x00\x00\x00\x00\x00\x00\x00",
			want: []interface{}{int64(0x7f00000000000000)},
		},
		{
			// type 6: 64 bit
			e:    "\x02\x06\xff\xff\xff\xff\xff\xff\xff\xff",
			want: []interface{}{-int64(1)},
		},
		{
			// type 7: float
			e:    "\x02\x07\x00\x00\x00\x00\x00\x00\x00\x00",
			want: []interface{}{0.0},
		},
		{
			// type 7: float
			e:    "\x02\x07\x40\x09\x21\xfb\x54\x44\x2d\x18",
			want: []interface{}{3.141592653589793},
		},
	} {
		e, err := parseRecord([]byte(cas.e))
		if err != nil {
			t.Fatalf("case %d: %s", i, err)
		}
		if have, want := e, cas.want; !reflect.DeepEqual(have, want) {
			t.Errorf("case %d: have\n-[[%#v]], want\n-[[%#v]]", i, have, want)
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
