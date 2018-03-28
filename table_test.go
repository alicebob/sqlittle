package sqlit

import (
	"testing"
)

func TestVarint(t *testing.T) {
	for i, cas := range []struct {
		b []byte
		n int64
	}{
		{
			b: []byte("\x00"),
			n: 0,
		},
		{
			b: []byte("\xFF\x00"),
			n: 16256, // 0b00111111_10000000
		},
		{
			b: []byte("\xFF\x7F"),
			n: 0x3FFF, // 0b00111111_11111111
		},
		{
			b: []byte("\xBF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"),
			n: 0x7FFFFFFFFFFFFFFF,
		},
		{
			b: []byte("\xBF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFFignored"),
			n: 0x7FFFFFFFFFFFFFFF,
		},
		{
			// int64 overflow
			b: []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"),
			n: -1,
		},
	} {
		n, _ := readVarint(cas.b)
		if have, want := n, cas.n; have != want {
			t.Errorf("case %d: have %d, want %d", i, have, want)
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
	if have, want := master.Rows(), 1; have != want {
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

	if have, want := master.Rows(), 4; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}
