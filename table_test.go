package sqlit

import (
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
	f, err := openFile("./test/single.sql")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	master, err := f.pageMaster()
	if err != nil {
		t.Fatal(err)
	}

	rows, err := master.Rows()
	if err != nil {
		t.Fatal(err)
	}
	if have, want := rows, 1; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestTablesFour(t *testing.T) {
	f, err := openFile("./test/four.sql")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	master, err := f.pageMaster()
	if err != nil {
		t.Fatal(err)
	}

	rowCount, err := master.Rows()
	if err != nil {
		t.Fatal(err)
	}
	if have, want := rowCount, 4; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}

	aap, err := f.Table("aap")
	if err != nil {
		t.Fatal(err)
	}
	if aap == nil {
		t.Fatal("no table found")
	}
	var rows []interface{}
	if err := aap.root.Iter(func(rowid int64, c []byte) (bool, error) {
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
