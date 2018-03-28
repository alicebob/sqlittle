package sqlit

import (
	"errors"
	"reflect"
	"testing"
)

func TestHeader(t *testing.T) {
	// hexdump -v -e '/1 "0x%02x, "' -n 100 test/single.sql
	base := [headerSize]byte{0x53, 0x51, 0x4c, 0x69, 0x74, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x20, 0x33, 0x00, 0x10, 0x00, 0x01, 0x01, 0x00, 0x40, 0x20, 0x20, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x2e, 0x1c, 0xb0}
	type change func([headerSize]byte) [headerSize]byte
	type cas struct {
		change change
		want   header
		err    error
	}
	for n, c := range []cas{
		// All fine
		{
			change: func(h [headerSize]byte) [headerSize]byte {
				return h
			},
			want: header{
				Magic:    "SQLite format 3\x00",
				PageSize: 4096,
			},
		},

		// Magic number
		{
			// invalid magic numner
			change: func(h [headerSize]byte) [headerSize]byte {
				h[0] = 's'
				return h
			},
			err: ErrHeaderInvalidMagic,
		},

		// PageSize tests
		{
			// page size 4
			change: func(h [headerSize]byte) [headerSize]byte {
				h[16], h[17] = 0, 4
				return h
			},
			err: errors.New("invalid page size"),
		},
		{
			// page size not a power of two
			change: func(h [headerSize]byte) [headerSize]byte {
				h[17] = 0x12
				return h
			},
			err: ErrHeaderInvalidPageSize,
		},
		{
			// page size 0xffff
			change: func(h [headerSize]byte) [headerSize]byte {
				h[16], h[17] = 0xFF, 0xFF
				return h
			},
			err: ErrHeaderInvalidPageSize,
		},
		{
			// page size 1 is special case, according to the docs
			change: func(h [headerSize]byte) [headerSize]byte {
				h[16], h[17] = 0, 1
				return h
			},
			want: header{
				Magic:    "SQLite format 3\x00",
				PageSize: 65536,
			},
		},
	} {
		h, err := parseHeader(c.change(base))
		if have, want := err, c.err; !reflect.DeepEqual(have, want) {
			t.Fatalf("case %d: have %v, want %v", n, have, want)
		}
		if have, want := h, c.want; have != want {
			t.Fatalf("case %d: have %#v, want %#v", n, have, want)
		}
	}
}

func TestIOBasic(t *testing.T) {
	f, err := openFile("./test/single.sql")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if have, want := f.header.PageSize, 4096; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIONoSuch(t *testing.T) {
	_, err := openFile("./test/nosuch.sql")
	if have, want := err.Error(), "open ./test/nosuch.sql: no such file or directory"; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIOZero(t *testing.T) {
	_, err := openFile("./test/zerolength.sql")
	if have, want := err, ErrFileZeroLength; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIOTruncated(t *testing.T) {
	_, err := openFile("./test/truncated.sql")
	if have, want := err, ErrFileTooShort; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIOInvalidMagic(t *testing.T) {
	_, err := openFile("./test/magic.sql")
	if have, want := err, ErrHeaderInvalidMagic; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}
