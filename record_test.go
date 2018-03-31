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
