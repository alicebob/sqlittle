package sqlit

import (
	"reflect"
	"testing"
)

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
		{
			// type 8: int 0
			e:    "\x02\b",
			want: []interface{}{int64(0)},
		},
		{
			// type 9: int 1
			e:    "\x02\t",
			want: []interface{}{int64(1)},
		},
		{
			// type 1: 8 bit
			e:    "\x02\x01P",
			want: []interface{}{int64(80)},
		},
		{
			// type 1: 8 bit
			e:    "\x02\x01\xb0",
			want: []interface{}{-int64(80)},
		},
		{
			// type 2: 16 bit
			e:    "\x02\x02@\x00",
			want: []interface{}{int64(1 << 14)},
		},
		{
			// type 2: 16 bit
			e:    "\x02\x02\xc0\x00",
			want: []interface{}{-int64(1 << 14)},
		},
		{
			// type 3: 24 bit
			e:    "\x02\x03\x7f\x00\x00",
			want: []interface{}{int64(0x7f0000)},
		},
		{
			// type 3: 24 bit
			e:    "\x02\x03\xff\xff\xff",
			want: []interface{}{-int64(1)},
		},
		{
			// type 4: 32 bit
			e:    "\x02\x04\x7f\x00\x00\x00",
			want: []interface{}{int64(0x7f000000)},
		},
		{
			// type 4: 32 bit
			e:    "\x02\x04\xff\xff\xff\xff",
			want: []interface{}{-int64(1)},
		},
		{
			// type 5: 48 bit
			e:    "\x02\x05\x7f\x00\x00\x00\x00\x00",
			want: []interface{}{int64(0x7f0000000000)},
		},
		{
			// type 5: 48 bit
			e:    "\x02\x05\xff\xff\xff\xff\xff\xff",
			want: []interface{}{-int64(1)},
		},
		{
			// type 6: 64 bit
			e:    "\x02\x06\x7f\x00\x00\x00\x00\x00\x00\x00",
			want: []interface{}{int64(0x7f00000000000000)},
		},
		{
			// type 6: 64 bit
			e:    "\x02\x06\xff\xff\xff\xff\xff\xff\xff\xff",
			want: []interface{}{-int64(1)},
		},
		{
			// type 7: float
			e:    "\x02\x07\x00\x00\x00\x00\x00\x00\x00\x00",
			want: []interface{}{0.0},
		},
		{
			// type 7: float
			e:    "\x02\x07\x40\x09\x21\xfb\x54\x44\x2d\x18",
			want: []interface{}{3.141592653589793},
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

func TestColumnCompare(t *testing.T) {
	type cas struct {
		a, b interface{}
		want int
		err  error
	}
	for i, c := range []cas{
		// nil
		{nil, nil, 0, nil},

		// int64
		{nil, int64(1), -1, nil},
		{int64(1), nil, 1, nil},
		{int64(42), int64(42), 0, nil},
		{int64(-42), int64(42), -1, nil},
		{int64(42), int64(-42), 1, nil},

		// float64
		{nil, 3.14, -1, nil},
		{3.14, nil, 1, nil},
		{1.12, 3.14, -1, nil},
		{3.14, 3.14, 0, nil},
		{3.14, -3.14, 1, nil},

		// float64 'n int
		{-int64(12), -3.14, -1, nil},
		{int64(3), 3.0, 0, nil},
		{int64(3), -3.14, 1, nil},

		// strings
		{nil, "bar", -1, nil},
		{"bar", nil, 1, nil},
		{"foo", "bar", 1, nil},
		{"foo", "foo", 0, nil},
		{"bar", "foo", -1, nil},
		{"foofoo", "foo", 1, nil},
		{"foo", "foofoo", -1, nil},
		{"foo", "Foo", 1, nil},

		// bytes
		{nil, []byte("bar"), -1, nil},
		{[]byte("bar"), nil, 1, nil},
		{[]byte("foo"), []byte("bar"), 1, nil},
		{[]byte("bar"), []byte("bar"), 0, nil},
		{[]byte("bar"), []byte("foo"), -1, nil},

		// combos
		{int64(42), "string", 0, errCmp},
		{"string", int64(42), 0, errCmp},
		{int64(42), []byte("byte"), 0, errCmp},
		{[]byte("byte"), int64(42), 0, errCmp},
		{3.14, "string", 0, errCmp},
		{"string", 3.14, 0, errCmp},
		{3.14, []byte("byte"), 0, errCmp},
		{[]byte("byte"), 3.14, 0, errCmp},
		{[]byte("byte"), "string", 0, errCmp},
		{"string", []byte("byte"), 0, errCmp},
	} {
		o, err := columnCmp(c.a, c.b)
		if have, want := err, c.err; !reflect.DeepEqual(have, want) {
			t.Fatalf("case %d: have %q, want %q", i, have, want)
		}
		if have, want := o, c.want; have != want {
			t.Errorf("case %d-(%v,%v): have %d, want %d", i, c.a, c.b, have, want)
		}
	}
}

func TestRecordCompare(t *testing.T) {
	type cas struct {
		a, b Row
		want int
	}
	for i, c := range []cas{
		{
			a:    Row{int64(1)},
			b:    Row{int64(42)},
			want: -1,
		},
		{
			a:    Row{int64(42)},
			b:    Row{int64(42)},
			want: 0,
		},
		{
			a:    Row{int64(42)},
			b:    Row{int64(1)},
			want: 1,
		},
		{
			a:    Row{int64(42), int64(43)},
			b:    Row{int64(42)},
			want: 1,
		},
		{
			a:    Row{int64(42)},
			b:    Row{int64(42), int64(43)},
			want: -1,
		},
	} {
		o, err := cmp(c.a, c.b)
		if err != nil {
			t.Fatal(err)
		}
		if have, want := o, c.want; have != want {
			t.Errorf("case %d: have %d, want %d", i, have, want)
		}
	}
}
