// tests with corrupted data files

package sqlittle

import (
	"errors"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func twiddleAByte(b []byte) {
	new := uint8(0)
	if rand.Intn(2) == 1 {
		new = uint8(rand.Uint32())
	}
	b[rand.Intn(len(b))] = new
}

type corrupter mmapLoader

func corruptDatabase(f string) (*database, error) {
	l, err := newMMapLoader(f)
	if err != nil {
		return nil, err
	}
	return newDatabase((*corrupter)(l))
}

func (c *corrupter) header() ([headerSize]byte, error) {
	b, err := (*mmapLoader)(c).header()
	// don't mess too often with the header bytes. It mostly messes up the
	// magic number, which is a boring test.
	if err == nil {
		if rand.Intn(100) < 3 {
			for i := 0; i < rand.Intn(10); i++ {
				twiddleAByte(b[:])
			}
		}
		if rand.Intn(100) < 1 {
			err = errors.New("header corrupter strikes")
		}
	}
	return b, err
}

func (c *corrupter) page(n int, pagesize int) ([]byte, error) {
	p, err := (*mmapLoader)(c).page(n, pagesize)
	if err == nil {
		if rand.Intn(100) < 40 {
			for i := 0; i < rand.Intn(10); i++ {
				twiddleAByte(p[:])
			}
		}
		if rand.Intn(100) < 1 {
			err = errors.New("page corrupter strikes again")
		}
	}
	return p, err
}
func (c *corrupter) Close() error {
	return (*mmapLoader)(c).Close()
}

func readTableWords() ([]string, error) {
	db, err := corruptDatabase("test/words.sqlite")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	tab, err := db.Table("words")
	if err != nil {
		return nil, err
	}
	if tab == nil {
		return nil, errors.New("no table found")
	}

	if _, err := tab.root.Count(db); err != nil {
		return nil, err
	}

	var rows []string
	_, err = tab.root.Iter(
		db,
		func(rowid int64, pl Payload) (bool, error) {
			c, err := addOverflow(db, pl)
			if err != nil {
				return false, err
			}
			e, err := parseRecord(c)
			if err != nil {
				return false, err
			}
			// could be wrong length with some flipped bits
			if len(e) != 2 {
				return false, errors.New("wrong field count")
			}
			// could be broken with some flipped bits
			word, ok := e[0].(string)
			if !ok {
				return false, errors.New("not a string")
			}
			rows = append(rows, word)
			return false, nil
		})
	return rows, err
}

func TestCorrupted(t *testing.T) {
	for i := 0; i < 100; i++ {
		_, err := readTableWords()
		if err != nil {
			// t.Log(err)
			// t.Fatal(err)
		}
		/*
			if have, want := len(words), 1000; have != want {
				t.Errorf("have %#v, want %#v", have, want)
			}
		*/
	}
}
