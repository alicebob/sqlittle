package sqlit

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestHeader(t *testing.T) {
	// hexdump -v -e '/1 "0x%02x, "' -n 100 test/single.sqlite
	base := [headerSize]byte{
		0x53, 0x51, 0x4c, 0x69, 0x74, 0x65, 0x20, 0x66,
		0x6f, 0x72, 0x6d, 0x61, 0x74, 0x20, 0x33, 0x00,
		0x10, 0x00, 0x01, 0x01, 0x00, 0x40, 0x20, 0x20,
		0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x04,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04,
		0x00, 0x2e, 0x1c, 0xb0,
	}
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
	f, err := openFile("./test/single.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if have, want := f.header.PageSize, 4096; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIONoSuch(t *testing.T) {
	_, err := openFile("./test/nosuch.sqlite")
	if have, want := err.Error(), "open ./test/nosuch.sqlite: no such file or directory"; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIOZero(t *testing.T) {
	_, err := openFile("./test/zerolength.sqlite")
	if have, want := err, ErrFileZeroLength; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIOTruncated(t *testing.T) {
	_, err := openFile("./test/truncated.sqlite")
	if have, want := err, ErrFileTooShort; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIOInvalidMagic(t *testing.T) {
	_, err := openFile("./test/magic.sqlite")
	if have, want := err, ErrHeaderInvalidMagic; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIOTableRowidSingle(t *testing.T) {
	db, err := openFile("./test/single.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	{
		_, err := db.TableRowid("nosuch", 999)
		if have, want := err, ErrNoSuchTable; have != want {
			t.Errorf("have %#v, want %#v", have, want)
		}
	}

	for _, c := range []struct {
		rowid int64
		want  Row
	}{
		{-1, nil},
		{0, nil},
		{1, Row{"world"}},
		{2, Row{"universe"}},
		{3, Row{"town"}},
		{4, nil},
		{4000, nil},
	} {
		row, err := db.TableRowid("hello", c.rowid)
		if err != nil {
			t.Fatal(err)
		}
		if have, want := row, c.want; !reflect.DeepEqual(have, want) {
			t.Errorf("have %#v, want %#v", have, want)
		}
	}
}

func TestIOTableRowidLong(t *testing.T) {
	db, err := openFile("./test/long.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	type cas struct {
		rowid int64
		want  Row
	}

	var cases = []cas{
		{-1, nil},
		{0, nil},
		{1, Row{"bottles of beer on the wall 1"}},
		{1000, Row{"bottles of beer on the wall 1000"}},
		{1001, nil},
		{4000, nil},
	}
	for i := int64(1); i <= 1000; i++ {
		cases = append(cases, cas{
			i,
			Row{fmt.Sprintf("bottles of beer on the wall %d", i)},
		})
	}
	rand.Shuffle(len(cases), func(i, j int) {
		cases[i], cases[j] = cases[j], cases[i]
	})

	for _, c := range cases {
		row, err := db.TableRowid("bottles", c.rowid)
		if err != nil {
			t.Fatal(err)
		}
		if have, want := row, c.want; !reflect.DeepEqual(have, want) {
			t.Errorf("have %#v, want %#v", have, want)
		}
	}
}
