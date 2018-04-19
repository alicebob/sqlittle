package sqlittle

import (
	"reflect"
	"testing"
)

func TestRecord(t *testing.T) {
	test := func(e string, want []interface{}, wantErr error) {
		t.Helper()
		parsed, err := parseRecord([]byte(e))
		if have, want := err, wantErr; !reflect.DeepEqual(have, want) {
			t.Fatalf("have %v, want %v", have, want)
		}
		if have, want := parsed, want; !reflect.DeepEqual(have, want) {
			t.Errorf("have\n-[[%#v]], want\n-[[%#v]]", have, want)
		}
	}
	test(
		"\x06\x17\x17\x17\x01Wtablehellohello\x02CREATE TABLE hello (who varchar(255))",
		[]interface{}{
			"table",
			"hello",
			"hello",
			int64(2),
			"CREATE TABLE hello (who varchar(255))",
		},
		nil,
	)
	test(
		// type 1: 8 bit
		"\x02\x01P",
		[]interface{}{int64(80)},
		nil,
	)
	test(
		// type 1: 8 bit
		"\x02\x01\xb0",
		[]interface{}{-int64(80)},
		nil,
	)
	test(
		// type 2: 16 bit
		"\x02\x02@\x00",
		[]interface{}{int64(1 << 14)},
		nil,
	)
	test(
		// type 2: 16 bit
		"\x02\x02\xc0\x00",
		[]interface{}{-int64(1 << 14)},
		nil,
	)
	test(
		// type 3: 24 bit
		"\x02\x03\x7f\x00\x00",
		[]interface{}{int64(0x7f0000)},
		nil,
	)
	test(
		// type 3: 24 bit
		"\x02\x03\xff\xff\xff",
		[]interface{}{-int64(1)},
		nil,
	)
	test(
		// type 4: 32 bit
		"\x02\x04\x7f\x00\x00\x00",
		[]interface{}{int64(0x7f000000)},
		nil,
	)
	test(
		// type 4: 32 bit
		"\x02\x04\xff\xff\xff\xff",
		[]interface{}{-int64(1)},
		nil,
	)
	test(
		// type 5: 48 bit
		"\x02\x05\x7f\x00\x00\x00\x00\x00",
		[]interface{}{int64(0x7f0000000000)},
		nil,
	)
	test(
		// type 5: 48 bit
		"\x02\x05\xff\xff\xff\xff\xff\xff",
		[]interface{}{-int64(1)},
		nil,
	)
	test(
		// type 6: 64 bit
		"\x02\x06\x7f\x00\x00\x00\x00\x00\x00\x00",
		[]interface{}{int64(0x7f00000000000000)},
		nil,
	)
	test(
		// type 6: 64 bit
		"\x02\x06\xff\xff\xff\xff\xff\xff\xff\xff",
		[]interface{}{-int64(1)},
		nil,
	)
	test(
		// type 7: float
		"\x02\x07\x00\x00\x00\x00\x00\x00\x00\x00",
		[]interface{}{0.0},
		nil,
	)
	test(
		// type 7: float
		"\x02\x07\x40\x09\x21\xfb\x54\x44\x2d\x18",
		[]interface{}{3.141592653589793},
		nil,
	)
	test(
		// type 8: int 0
		"\x02\b",
		[]interface{}{int64(0)},
		nil,
	)
	test(
		// type 9: int 1
		"\x02\t",
		[]interface{}{int64(1)},
		nil,
	)
	test(
		// type 10: internal
		"\x02\x0a",
		nil,
		errInternal,
	)
	test(
		// type 11: internal
		"\x02\x0b",
		nil,
		errInternal,
	)
	test(
		// type > 11, even: bytes
		"\x02VCREATE TABLE hello (who varchar(255))",
		[]interface{}{
			[]byte("CREATE TABLE hello (who varchar(255))"),
		},
		nil,
	)
	test(
		// type > 11, odd: string
		"\x02WCREATE TABLE hello (who varchar(255))",
		[]interface{}{
			"CREATE TABLE hello (who varchar(255))",
		},
		nil,
	)

	// Error cases
	test(
		// truncated record
		"\x02",
		nil,
		ErrCorrupted,
	)
	test(
		// truncated 8 bit
		"\x02\x01",
		nil,
		ErrCorrupted,
	)
	test(
		// truncated 16 bit
		"\x02\x02@",
		nil,
		ErrCorrupted,
	)
	test(
		// truncated 32 bit
		"\x02\x04\x7f\x00\x00",
		nil,
		ErrCorrupted,
	)
	test(
		// truncated 48 bit
		"\x02\x05\x7f\x00\x00",
		nil,
		ErrCorrupted,
	)
	test(
		// truncated 64 bit
		"\x02\x06\x7f\x00\x00\x00\x00\x00\x00",
		nil,
		ErrCorrupted,
	)
	test(
		// truncated float
		"\x02\x07\x40\x09\x21\xfb\x54\x44\x2d",
		nil,
		ErrCorrupted,
	)
	test(
		// truncated bytes
		"\x02VCREATE TABLE hello (who varchar(255)",
		nil,
		ErrCorrupted,
	)
	test(
		// truncated string
		"\x02WCREATE TABLE hello (who varchar(255)",
		nil,
		ErrCorrupted,
	)
	test(
		// truncated multi field record
		"\x06\x17\x17\x17\x01Wtablehellohello\x02",
		[]interface{}{
			"table",
			"hello",
			"hello",
			int64(2),
		},
		ErrCorrupted,
	)
}

func TestColumnCompare(t *testing.T) {
	test := func(a, b interface{}, want int, wantErr error) {
		t.Helper()
		o, err := columnCmp(a, b)
		if have, want := err, wantErr; !reflect.DeepEqual(have, want) {
			t.Fatalf("have %q, want %q", have, want)
		}
		if have, want := o, want; have != want {
			t.Errorf("have %d, want %d", have, want)
		}
	}
	// nil
	test(nil, nil, 0, nil)

	// int64
	test(nil, int64(1), -1, nil)
	test(int64(1), nil, 1, nil)
	test(int64(42), int64(42), 0, nil)
	test(int64(-42), int64(42), -1, nil)
	test(int64(42), int64(-42), 1, nil)

	// float64
	test(nil, 3.14, -1, nil)
	test(3.14, nil, 1, nil)
	test(1.12, 3.14, -1, nil)
	test(3.14, 3.14, 0, nil)
	test(3.14, -3.14, 1, nil)

	// float64 'n int
	test(-int64(12), -3.14, -1, nil)
	test(int64(3), 3.0, 0, nil)
	test(int64(3), -3.14, 1, nil)

	// strings
	test(nil, "bar", -1, nil)
	test("bar", nil, 1, nil)
	test("foo", "bar", 1, nil)
	test("foo", "foo", 0, nil)
	test("bar", "foo", -1, nil)
	test("foofoo", "foo", 1, nil)
	test("foo", "foofoo", -1, nil)
	test("foo", "Foo", 1, nil)

	// bytes
	test(nil, []byte("bar"), -1, nil)
	test([]byte("bar"), nil, 1, nil)
	test([]byte("foo"), []byte("bar"), 1, nil)
	test([]byte("bar"), []byte("bar"), 0, nil)
	test([]byte("bar"), []byte("foo"), -1, nil)

	// combos
	test(int64(42), "string", 0, errCmp)
	test("string", int64(42), 0, errCmp)
	test(int64(42), []byte("byte"), 0, errCmp)
	test([]byte("byte"), int64(42), 0, errCmp)
	test(3.14, "string", 0, errCmp)
	test("string", 3.14, 0, errCmp)
	test(3.14, []byte("byte"), 0, errCmp)
	test([]byte("byte"), 3.14, 0, errCmp)
	test([]byte("byte"), "string", 0, errCmp)
	test("string", []byte("byte"), 0, errCmp)
}

func TestRecordCompare(t *testing.T) {
	test := func(a, b Record, want int) {
		t.Helper()
		o, err := cmp(a, b)
		if err != nil {
			t.Fatal(err)
		}
		if have, want := o, want; have != want {
			t.Errorf("have %d, want %d", have, want)
		}
	}
	test(
		Record{int64(1)},
		Record{int64(42)},
		-1,
	)
	test(
		Record{int64(42)},
		Record{int64(42)},
		0,
	)
	test(
		Record{int64(42)},
		Record{int64(1)},
		1,
	)
	test(
		Record{int64(42), int64(43)},
		Record{int64(42)},
		0,
	)
	test(
		Record{int64(42)},
		Record{int64(42), int64(43)},
		0,
	)
}
