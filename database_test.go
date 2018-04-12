package sqlittle

import (
	"bufio"
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

func TestIOBasic(t *testing.T) {
	f, err := OpenFile("./test/single.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if have, want := f.l.(*filePager).header.PageSize, 4096; have != want {
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
	if have, want := err, ErrInvalidMagic; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

/*
func TestIOWal(t *testing.T) {
	_, err := OpenFile("./test/wal.sqlite")
	if have, want := err, ErrWAL; have != want {
		t.Errorf("have %#v, want %#v", have, want)
	}
}
*/

func TestMasterNoSQL(t *testing.T) {
	// primary key creates an index without SQL
	db, err := OpenFile("./test/primarykey.sqlite")
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
	db, err := OpenFile("./test/index.sqlite")
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
	db, err := OpenFile("./test/single.sqlite")
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
	db, err := OpenFile("./test/words.sqlite")
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
