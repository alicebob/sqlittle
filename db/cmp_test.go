package db

import (
	"testing"
)

func TestCompare(t *testing.T) {
	test := func(a interface{}, b interface{}, want int) {
		t.Helper()

		if have, want := Compare(a, b), want; have != want {
			t.Errorf("have %d, want %d", have, want)
		}
	}
	test(nil, nil, 0)
	test(nil, int64(42), -1)
	test(nil, 3.14, -1)
	test(nil, "bar", -1)
	test(nil, []byte("bar"), -1)
	test(int64(42), nil, 1)
	test(int64(42), int64(41), 1)
	test(int64(42), int64(42), 0)
	test(int64(42), int64(43), -1)
	test(int64(42), 41.00, 1)
	test(int64(42), 42.00, 0)
	test(int64(42), 43.00, -1)
	test(int64(42), "bar", -1)
	test(int64(42), []byte("bar"), -1)
	test(3.14, nil, 1)
	test(3.14, int64(2), 1)
	test(3.00, int64(3), 0)
	test(3.14, int64(3), 1)
	test(3.14, int64(4), -1)
	test(3.14, 2.14, 1)
	test(3.14, 3.14, 0)
	test(3.14, 4.14, -1)
	test(3.14, "bar", -1)
	test(3.14, []byte("bar"), -1)
	test("foo", "bar", 1)
	test("aap", nil, 1)
	test("aap", int64(42), 1)
	test("aap", 3.14, 1)
	test("aap", "aaap", 1)
	test("aap", "aap", 0)
	test("aap", "noot", -1)
	test("aap", []byte("aaap"), -1)
	test([]byte("aap"), nil, 1)
	test([]byte("aap"), int64(42), 1)
	test([]byte("aap"), 3.14, 1)
	test([]byte("aap"), "noot", 1)
	test([]byte("aap"), []byte("aaap"), 1)
	test([]byte("aap"), []byte("aap"), 0)
	test([]byte("aap"), []byte("ap"), -1)

	test(CmpPrefix("aap"), nil, 1)
	test(CmpPrefix("aap"), "aaaal", 1)
	test(CmpPrefix("aap"), "aap", 0)
	test(CmpPrefix("aap"), "aapnootmies", 0)
	test(CmpPrefix("aap"), "zapzap", -1)
	test(CmpPrefix("aap"), []byte("foo"), -1)

	test(NewCmpDesc("aap"), "noot", 1)
	test(NewCmpDesc("aap"), "aap", 0)

	test(CollateRtrim("aap"), "noot  ", -1)
	test(CollateRtrim("aap   "), "aap", 0)
	test(CollateRtrim("aap"), "aap   ", 0)
	test(CollateRtrim("aap"), "mies   ", -1)
	test(NewCmpDesc(CollateRtrim("aap")), "noot  ", 1)
}
