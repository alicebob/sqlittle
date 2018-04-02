package sqlittle

import (
	"bufio"
	"errors"
	"io"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
	"unicode/utf8"
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
				PageSize:      4096,
				ReservedSpace: 0,
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

		// PageSize
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
				PageSize:      0x010000,
				ReservedSpace: 0,
			},
		},

		// read version
		{
			// read version > 2
			change: func(h [headerSize]byte) [headerSize]byte {
				h[19] = 3
				return h
			},
			err: ErrIncompatible,
		},

		// reserved space
		{
			// test #7
			change: func(h [headerSize]byte) [headerSize]byte {
				h[20] = 0x10
				return h
			},
			want: header{
				PageSize:      0x1000,
				ReservedSpace: 0x10,
			},
		},

		// constants
		{
			// maximum fraction
			change: func(h [headerSize]byte) [headerSize]byte {
				h[21] = 123
				return h
			},
			err: ErrIncompatible,
		},
		{
			// minimum fraction
			change: func(h [headerSize]byte) [headerSize]byte {
				h[22] = 123
				return h
			},
			err: ErrIncompatible,
		},
		{
			// leaf fraction
			change: func(h [headerSize]byte) [headerSize]byte {
				h[23] = 123
				return h
			},
			err: ErrIncompatible,
		},

		// Schema format numner
		{
			// no support for version 1
			change: func(h [headerSize]byte) [headerSize]byte {
				h[44+3] = 1
				return h
			},
			err: ErrIncompatible,
		},
		{
			// invalid value
			change: func(h [headerSize]byte) [headerSize]byte {
				h[44+3] = 0
				return h
			},
			err: ErrIncompatible,
		},
		{
			// invalid value
			change: func(h [headerSize]byte) [headerSize]byte {
				h[44+3] = 5
				return h
			},
			err: ErrIncompatible,
		},

		// Text Encoding
		{
			// invalid value
			change: func(h [headerSize]byte) [headerSize]byte {
				h[56+3] = 0
				return h
			},
			err: ErrIncompatible,
		},
		{
			// invalid value
			change: func(h [headerSize]byte) [headerSize]byte {
				h[56+3] = 2
				return h
			},
			err: ErrEncoding,
		},
		{
			change: func(h [headerSize]byte) [headerSize]byte {
				h[56+3] = 1
				return h
			},
			want: header{
				PageSize:      0x1000,
				ReservedSpace: 0,
			},
		},

		// empty
		{
			// 'Reserved for expansion'. Must be 0s.
			change: func(h [headerSize]byte) [headerSize]byte {
				h[78] = 1
				return h
			},
			err: ErrIncompatible,
		},
	} {
		h, err := parseHeader(c.change(base))
		if have, want := err, c.err; !reflect.DeepEqual(have, want) {
			t.Fatalf("case %d: have %v, want %v", n, have, want)
		}
		if c.err != nil {
			continue
		}
		if have, want := h, c.want; have != want {
			t.Fatalf("case %d: have %#v, want %#v", n, have, want)
		}
	}
}

func TestIOBasic(t *testing.T) {
	f, err := OpenFile("./test/single.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if have, want := f.header.PageSize, 4096; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIONoSuch(t *testing.T) {
	_, err := OpenFile("./test/nosuch.sqlite")
	if have, want := err.Error(), "open ./test/nosuch.sqlite: no such file or directory"; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIOZero(t *testing.T) {
	_, err := OpenFile("./test/zerolength.sqlite")
	if have, want := err, ErrFileTruncated; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIOTruncated(t *testing.T) {
	_, err := OpenFile("./test/truncated.sqlite")
	if have, want := err, ErrFileTruncated; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIOInvalidMagic(t *testing.T) {
	_, err := OpenFile("./test/magic.sqlite")
	if have, want := err, ErrHeaderInvalidMagic; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIOTableRowidSingle(t *testing.T) {
	db, err := OpenFile("./test/single.sqlite")
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
		want  Record
	}{
		{-1, nil},
		{0, nil},
		{1, Record{"world"}},
		{2, Record{"universe"}},
		{3, Record{"town"}},
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

// wordList gives the contents of words.txt
func wordList(t *testing.T) []string {
	f, err := os.Open("./test/words.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	var words []string
	b := bufio.NewReader(f)
	for {
		w, err := b.ReadString('\n')
		if err == io.EOF {
			return words
		}
		if err != nil {
			t.Fatal(err)
		}
		words = append(words, strings.TrimRight(w, "\n"))
	}
}

func TestIOTableRowidLong(t *testing.T) {
	db, err := OpenFile("./test/words.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	type cas struct {
		rowid int64
		want  Record
	}

	var cases = []cas{
		{-1, nil},
		{0, nil},
		{1, Record{"hangdog", int64(7)}},
		{1000, Record{"ideologist", int64(10)}},
		{1001, nil},
		{4000, nil},
	}
	for i, w := range wordList(t) {
		cases = append(cases, cas{
			rowid: int64(i) + 1,
			want: Record{
				w,
				int64(utf8.RuneCountInString(w)),
			},
		})
	}
	rand.Shuffle(len(cases), func(i, j int) {
		cases[i], cases[j] = cases[j], cases[i]
	})

	for _, c := range cases {
		row, err := db.TableRowid("words", c.rowid)
		if err != nil {
			t.Fatal(err)
		}
		if have, want := row, c.want; !reflect.DeepEqual(have, want) {
			t.Errorf("have %#v, want %#v", have, want)
		}
	}
}
