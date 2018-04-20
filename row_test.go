package sqlittle

import (
	"fmt"
	"reflect"
	"testing"
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
