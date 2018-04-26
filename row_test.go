package sqlittle

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestScanNil(t *testing.T) {
	// nil skips the column
	if err := (Row{123}).Scan(nil); err != nil {
		t.Fatal(err)
	}
}

func TestScanString(t *testing.T) {
	test := func(v interface{}, want string) {
		t.Helper()
		n := ""
		if err := (Row{v}).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if have, want := n, want; have != want {
			t.Errorf("have %v, want %v", have, want)
		}
	}

	test(nil, "")
	test(int64(42), "42")
	test(float64(3.14), "3.14")
	test("world", "world")
	test([]byte("hello"), "hello")
}

func TestScanBytes(t *testing.T) {
	test := func(v interface{}, want []byte) {
		t.Helper()
		var n []byte
		if err := (Row{v}).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if have, want := n, want; !reflect.DeepEqual(have, want) {
			t.Errorf("have %v, want %v", have, want)
		}
	}

	test(nil, []byte(nil))
	test(int64(42), []byte("42"))
	test(float64(3.14), []byte("3.14"))
	test("world", []byte("world"))
	test([]byte("hello"), []byte("hello"))
}

func TestScanFloat64(t *testing.T) {
	test := func(v interface{}, want float64) {
		t.Helper()
		var f float64
		if err := (Row{v}).Scan(&f); err != nil {
			t.Fatal(err)
		}
		if have, want := f, want; have != want {
			t.Errorf("have %v, want %v", have, want)
		}
	}
	fail := func(v interface{}) {
		t.Helper()
		var f float64
		if have, want := (Row{v}).Scan(&f), fmt.Errorf("invalid number: %q", v); !reflect.DeepEqual(have, want) {
			t.Errorf("have %v, want %v", have, want)
		}
	}
	test(nil, 0)
	test(int64(42), 42)
	test(float64(3.14), 3.14)
	test("-3.14", -3.14)
	test([]byte("2.71828"), 2.71828)
	fail("hi")
	fail([]byte("bye"))
	fail("123world")
}

func TestScanInt64(t *testing.T) {
	test := func(v interface{}, want int64) {
		t.Helper()
		var n int64
		if err := (Row{v}).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if have, want := n, want; have != want {
			t.Errorf("have %v, want %v", have, want)
		}
	}
	fail := func(v interface{}) {
		t.Helper()
		var n int64
		if have, want := (Row{v}).Scan(&n), fmt.Errorf("invalid number: %q", v); !reflect.DeepEqual(have, want) {
			t.Errorf("have %v, want %v", have, want)
		}
	}
	test(nil, 0)
	test(int64(42), 42)
	test(float64(3.14), 3)
	test("-3.14", -3)
	test([]byte("2.71828"), 2)
	fail("hi")
	fail([]byte("bye"))
	fail("123world")
}

func TestScanInt32(t *testing.T) {
	test := func(v interface{}, want int32) {
		t.Helper()
		var n int32
		if err := (Row{v}).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if have, want := n, want; have != want {
			t.Errorf("have %v, want %v", have, want)
		}
	}
	test(nil, 0)
	test(int64(42), 42)
	test(int64(1<<42+1<<4), 1<<4)
	test(float64(3.14), 3)
	test("-3.14", -3)
	test([]byte("2.71828"), 2)
}

func TestScanInt(t *testing.T) {
	test := func(v interface{}, want int) {
		t.Helper()
		var n int
		if err := (Row{v}).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if have, want := n, want; have != want {
			t.Errorf("have %v, want %v", have, want)
		}
	}
	test(nil, 0)
	test(int64(42), 42)
	test(float64(3.14), 3)
	test("3.1415", 3)
	test([]byte("2.71828"), 2)
}

func TestScanBool(t *testing.T) {
	test := func(v interface{}, want bool) {
		t.Helper()
		var n bool
		if err := (Row{v}).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if have, want := n, want; have != want {
			t.Errorf("have %v, want %v", have, want)
		}
	}
	fail := func(v interface{}) {
		t.Helper()
		var n bool
		if have, want := (Row{v}).Scan(&n), fmt.Errorf("invalid boolean: %q", v); !reflect.DeepEqual(have, want) {
			t.Errorf("have %v, want %v", have, want)
		}
	}
	test(nil, false)
	test(int64(42), true)
	test(int64(0), false)
	test(float64(3.14), true)
	test(float64(0.0), false)
	test(float64(-0.0), false)
	test("3.1414", true)
	test("0.0", false)
	test([]byte("2.71828"), true)
	test([]byte("0.0"), false)
	fail("hi")
	fail("0hi")
}

func TestScanTime(t *testing.T) {
	test := func(v interface{}, want time.Time) {
		t.Helper()
		var tim time.Time
		if err := (Row{v}).Scan(&tim); err != nil {
			t.Fatal(err)
		}
		if have, want := tim, want; !have.Equal(want) {
			t.Errorf("have %v, want %v", have, want)
		}
	}
	fail := func(v interface{}, err error) {
		t.Helper()
		var tim time.Time
		if have, want := (Row{v}).Scan(&tim), err; !reflect.DeepEqual(have, want) {
			t.Errorf("have %v, want %v", have, want)
		}
	}
	test(nil, time.Time{})
	test(int64(42), time.Unix(42, 0))
	test("1999-02-22 23:45:34.333", time.Date(1999, 2, 22, 23, 45, 34, 333000000, time.UTC))
	fail([]byte("aaa"), errors.New("BLOB timestamps are invalid"))
	fail("aaa", errors.New(`invalid time: "aaa"`))
	fail(float64(3.14), errors.New("float timestamps not supported")) // TODO
}

func TestScanStrings(t *testing.T) {
	test := func(row Row, want []string) {
		t.Helper()
		s := row.ScanStrings()
		if have, want := s, want; !reflect.DeepEqual(have, want) {
			t.Errorf("have %v, want %v", have, want)
		}
	}

	test(Row{}, []string{})
	test(Row{"foo", "bar"}, []string{"foo", "bar"})
	test(Row{"foo", "bar", int64(12)}, []string{"foo", "bar", "12"})
}
