package db

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
	// hexdump -v -e '/1 "0x%02x, "' -n 100 ../testdata/single.sqlite
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
	test := func(f change, expect *header, expectErr error) {
		t.Helper()
		hb := f(base)
		h, err := parseHeader(hb[:])
		if have, want := err, expectErr; !reflect.DeepEqual(have, want) {
			t.Fatalf(" have %v, want %v", have, want)
		}
		if expectErr != nil {
			return
		}
		if have, want := h, *expect; have != want {
			t.Fatalf("have %#v, want %#v", have, want)
		}
	}

	// All fine
	test(
		func(h [headerSize]byte) [headerSize]byte {
			return h
		},
		&header{
			PageSize:      4096,
			ChangeCounter: 4,
			SchemaCookie:  1,
		},
		nil,
	)

	// Magic number
	test(
		// invalid magic numner
		func(h [headerSize]byte) [headerSize]byte {
			h[0] = 's'
			return h
		},
		nil,
		ErrInvalidMagic,
	)

	// PageSize
	test(
		// page size 4
		func(h [headerSize]byte) [headerSize]byte {
			h[16], h[17] = 0, 4
			return h
		},
		nil,
		errors.New("invalid page size"),
	)
	test(
		// page size not a power of two
		func(h [headerSize]byte) [headerSize]byte {
			h[17] = 0x12
			return h
		},
		nil,
		ErrInvalidPageSize,
	)
	test(
		// page size 0xffff
		func(h [headerSize]byte) [headerSize]byte {
			h[16], h[17] = 0xFF, 0xFF
			return h
		},
		nil,
		ErrInvalidPageSize,
	)
	test(
		// page size 1 is special case, according to the docs
		func(h [headerSize]byte) [headerSize]byte {
			h[16], h[17] = 0, 1
			return h
		},
		&header{
			PageSize:      0x010000,
			ChangeCounter: 4,
			SchemaCookie:  1,
		},
		nil,
	)

	// read version
	test(
		// read version > 2
		func(h [headerSize]byte) [headerSize]byte {
			h[19] = 3
			return h
		},
		nil,
		ErrIncompatible,
	)

	// reserved space
	test(
		// test #7
		func(h [headerSize]byte) [headerSize]byte {
			h[20] = 0x10
			return h
		},
		nil,
		ErrReservedSpace,
	)

	// constants
	test(
		// maximum fraction
		func(h [headerSize]byte) [headerSize]byte {
			h[21] = 123
			return h
		},
		nil,
		ErrIncompatible,
	)
	test(
		// minimum fraction
		func(h [headerSize]byte) [headerSize]byte {
			h[22] = 123
			return h
		},
		nil,
		ErrIncompatible,
	)
	test(
		// leaf fraction
		func(h [headerSize]byte) [headerSize]byte {
			h[23] = 123
			return h
		},
		nil,
		ErrIncompatible,
	)

	// Schema format numner
	test(
		// we do support version 1
		func(h [headerSize]byte) [headerSize]byte {
			h[44+3] = 1
			return h
		},
		&header{
			PageSize:      4096,
			ChangeCounter: 4,
			SchemaCookie:  1,
		},
		nil,
	)
	test(
		// invalid value
		func(h [headerSize]byte) [headerSize]byte {
			h[44+3] = 0
			return h
		},
		nil,
		ErrIncompatible,
	)
	test(
		// invalid value
		func(h [headerSize]byte) [headerSize]byte {
			h[44+3] = 5
			return h
		},
		nil,
		ErrIncompatible,
	)

	// Text Encoding
	test(
		// invalid value
		func(h [headerSize]byte) [headerSize]byte {
			h[56+3] = 0
			return h
		},
		nil,
		ErrIncompatible,
	)
	test(
		// invalid value
		func(h [headerSize]byte) [headerSize]byte {
			h[56+3] = 2
			return h
		},
		nil,
		ErrEncoding,
	)
	test(
		func(h [headerSize]byte) [headerSize]byte {
			h[56+3] = 1
			return h
		},
		&header{
			PageSize:      0x1000,
			ChangeCounter: 4,
			SchemaCookie:  1,
		},
		nil,
	)

	// empty
	test(
		// 'Reserved for expansion'. Must be 0s.
		func(h [headerSize]byte) [headerSize]byte {
			h[78] = 1
			return h
		},
		nil,
		ErrIncompatible,
	)
}

func TestIOBasic(t *testing.T) {
	db, err := OpenFile("./../testdata/single.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if have, want := db.header.PageSize, 4096; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIONoSuch(t *testing.T) {
	_, err := OpenFile("./../testdata/nosuch.sqlite")
	if have, want := err.Error(), "open ./../testdata/nosuch.sqlite: no such file or directory"; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIOZero(t *testing.T) {
	_, err := OpenFile("./../testdata/zerolength.sqlite")
	if have, want := err, errors.New("mmap: closed"); !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIOTruncated(t *testing.T) {
	_, err := OpenFile("./../testdata/truncated.sqlite")
	if have, want := err, io.EOF; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIOInvalidMagic(t *testing.T) {
	_, err := OpenFile("./../testdata/magic.sqlite")
	if have, want := err, ErrInvalidMagic; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestIOWal(t *testing.T) {
	_, err := OpenFile("./../testdata/wal.sqlite")
	if have, want := err, ErrWAL; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestMasterNoSQL(t *testing.T) {
	// primary key creates an index without SQL statement
	db, err := OpenFile("./../testdata/primarykey.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	tables, err := db.Tables()
	if err != nil {
		t.Fatal(err)
	}
	for _, tname := range tables {
		_, err := db.Table(tname)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestDatabaseTable(t *testing.T) {
	db, err := OpenFile("./../testdata/index.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	{
		_, err := db.Table("nosuch")
		if have, want := err, ErrNoSuchTable; have != want {
			t.Errorf("have %#v, want %#v", have, want)
		}
	}

	{
		_, err := db.Table("hello_index")
		if have, want := err, ErrNoSuchTable; have != want {
			t.Errorf("have %#v, want %#v", have, want)
		}
	}

	{
		_, err := db.Table("hello")
		if err != nil {
			t.Errorf("have err: %s", err)
		}
	}
}

func TestIOTableRowidSingle(t *testing.T) {
	db, err := OpenFile("./../testdata/single.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	table, err := db.Table("hello")
	if err != nil {
		t.Fatal(err)
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
		row, err := table.Rowid(c.rowid)
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
	f, err := os.Open("./../testdata/words.txt")
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

func shuffle(n int, swap func(i, j int)) {
	if n < 0 {
		panic("invalid argument to shuffle")
	}
	i := n - 1
	for ; i > 1<<31-1-1; i-- {
		j := int(rand.Int63n(int64(i + 1)))
		swap(i, j)
	}
	for ; i > 0; i-- {
		j := int(rand.Int31n(int32(i + 1)))
		swap(i, j)
	}
}

func TestIOTableRowidLong(t *testing.T) {
	db, err := OpenFile("./../testdata/words.sqlite")
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
	shuffle(len(cases), func(i, j int) {
		cases[i], cases[j] = cases[j], cases[i]
	})

	table, err := db.Table("words")
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range cases {
		row, err := table.Rowid(c.rowid)
		if err != nil {
			t.Fatal(err)
		}
		if have, want := row, c.want; !reflect.DeepEqual(have, want) {
			t.Errorf("have %#v, want %#v", have, want)
		}
	}
}

func TestDatabaseLock(t *testing.T) {
	// can we lock at all?
	db, err := OpenFile("./../testdata/words.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if err := db.RLock(); err != nil {
		t.Fatal(err)
	}

	table, err := db.Table("words")
	if err != nil {
		t.Fatal(err)
	}
	row, err := table.Rowid(42)
	if err != nil {
		t.Fatal(err)
	}
	if have, want := row, (Record{"aniseed", int64(7)}); !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, want %#v", have, want)
	}

	if err := db.RUnlock(); err != nil {
		t.Fatal(err)
	}
}

func TestDatabaseSchema(t *testing.T) {
	db, err := OpenFile("./../testdata/words.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	s, err := db.Schema("words")
	if err != nil {
		t.Fatal(err)
	}
	if have, want := len(s.Columns), 2; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
	if have, want := len(s.Indexes), 2; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}
