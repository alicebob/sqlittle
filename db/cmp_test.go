package db

import (
	"testing"
)

func TestSearch(t *testing.T) {
	test := func(a Key, b Record, equals bool, search bool) {
		t.Helper()
		if have, want := Equals(a, b), equals; have != want {
			t.Errorf("equals: have %t, want %t", have, want)
		}

		if have, want := Search(a, b), search; have != want {
			t.Errorf("search; have %t, want %t", have, want)
		}
	}
	test(
		Key{{V: int64(1)}},
		Record{int64(42)},
		false,
		true,
	)
	test(
		Key{{V: int64(42)}},
		Record{int64(42)},
		true,
		true,
	)
	test(
		Key{{V: int64(82)}},
		Record{int64(42)},
		false,
		false,
	)
	test(
		Key{{V: int64(82)}},
		Record{int64(82), int64(14)},
		true,
		true,
	)
	test(
		Key{{V: int64(82)}, {V: int64(22)}},
		Record{int64(82), int64(14)},
		false,
		false,
	)
	test(
		Key{{V: int64(82)}, {V: int64(14)}},
		Record{int64(82), int64(14)},
		true,
		true,
	)
	test(
		Key{{V: int64(82)}, {V: int64(12)}},
		Record{int64(82), int64(14)},
		false,
		true,
	)

	test(
		Key{{V: int64(82)}, {V: int64(12), Desc: true}},
		Record{int64(82), int64(14)},
		false,
		false,
	)
	test(
		Key{{V: int64(82)}, {V: int64(12), Desc: true}},
		Record{int64(82), int64(12)},
		true,
		true,
	)
	test(
		Key{{V: int64(82)}, {V: int64(14), Desc: true}},
		Record{int64(82), int64(12)},
		false,
		true,
	)

	test( // invalid, a shouldn't have more records than b
		Key{{V: int64(82)}, {V: int64(14)}},
		Record{int64(82)},
		false,
		false,
	)

	test(
		Key{{V: "foo  ", Collate: "rtrim"}},
		Record{"foo"},
		true,
		true,
	)
	test(
		Key{{V: "fOo", Collate: "nocase"}},
		Record{"foO"},
		true,
		true,
	)
}

func Testcompare(t *testing.T) {
	test := func(a interface{}, b interface{}, want int) {
		t.Helper()

		if have, want := compare(a, b, CollateFuncs[""]), want; have != want {
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

	/*
		test(CollateRtrim("aap"), "noot  ", -1)
		test(CollateRtrim("aap   "), "aap", 0)
		test(CollateRtrim("aap"), "aap   ", 0)
		test(CollateRtrim("aap"), "mies   ", -1)
	*/
}
