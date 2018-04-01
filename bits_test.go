package sqlittle

import (
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
		// Error cases
		{
			b: []byte("\xFF"),
			l: -1,
			n: 0,
		},
		{
			b: []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"),
			l: -1,
			n: 0,
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
