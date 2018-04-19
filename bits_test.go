package sqlittle

import (
	"testing"
)

func TestVarint(t *testing.T) {
	test := func(eb []byte, el int, en int64) {
		t.Helper()
		n, l := readVarint(eb)
		if have, want := l, el; have != want {
			t.Errorf("have %d, want %d", have, want)
		}
		if have, want := n, en; have != want {
			t.Errorf("have %d, want %d", have, want)
		}
	}
	test([]byte("\x00"), 1, 0)
	test([]byte("\xFF\x00"), 2, 16256)  // 0b00111111_10000000
	test([]byte("\xFF\x7F"), 2, 0x3FFF) // 0b00111111_11111111
	test([]byte("\xBF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), 9, 0x7FFFFFFFFFFFFFFF)
	test([]byte("\xBF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFFignored"), 9, 0x7FFFFFFFFFFFFFFF)
	// int64 overflow
	test([]byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), 9, -1)
	// Error cases
	test([]byte("\xFF"), -1, 0)
	test([]byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), -1, 0)
}
