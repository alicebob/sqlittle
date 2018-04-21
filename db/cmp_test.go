package db

import (
	"testing"
)

func TestCmpString(t *testing.T) {
	test := func(a, b string, want int) {
		t.Helper()

		if have, want := CmpString(a)(b), want; have != want {
			t.Errorf("have %d, want %d", have, want)
		}
	}
	test("aap", "aaap", -1)
	test("aap", "aap", 0)
	test("aap", "noot", 1)
}

func TestCmpInt64(t *testing.T) {
	test := func(a int64, b interface{}, want int) {
		t.Helper()

		if have, want := CmpInt64(a)(b), want; have != want {
			t.Errorf("have %d, want %d", have, want)
		}
	}

	test(1, int64(-1), -1)
	test(1, int64(1), 0)
	test(1, int64(2), 1)
}
