package sqlittle

import (
	"encoding/binary"
	"errors"
	"math/bits"

	"golang.org/x/exp/mmap"
)

const (
	Magic      = "SQLite format 3\x00"
	headerSize = 100
)

var (
	ErrFileZeroLength        = errors.New("file is 0 bytes")
	ErrFileTooShort          = errors.New("file is too short")
	ErrHeaderInvalidMagic    = errors.New("invalid magic number")
	ErrHeaderInvalidPageSize = errors.New("invalid page size")
	ErrFileTruncated         = errors.New("file truncated")
	ErrNoSuchTable           = errors.New("no such table")
)

type header struct {
	Magic    string
	PageSize int
}

type database struct {
	f      *mmap.ReaderAt
	header header
}

func openFile(f string) (*database, error) {
	r, err := mmap.Open(f)
	if err != nil {
		return nil, err
	}
	if r.Len() == 0 {
		return nil, ErrFileZeroLength
	}

	buf := [headerSize]byte{}
	n, err := r.ReadAt(buf[:], 0)
	if n != headerSize {
		return nil, ErrFileTooShort
	}
	if err != nil {
		return nil, err
	}
	header, err := parseHeader(buf)
	if err != nil {
		return nil, err
	}

	db := &database{
		f:      r,
		header: header,
	}
	return db, nil
}

func (db *database) Close() error {
	return db.f.Close()
}

func (db *database) pageMaster() (TableBtree, error) {
	buf, err := db.page(1)
	if err != nil {
		return nil, err
	}
	return newTableBtree(buf, true)
}

// n starts at 1, sqlite style
func (db *database) page(id int) ([]byte, error) {
	if id < 1 {
		return nil, errors.New("invalid page number")
	}
	buf := make([]byte, db.header.PageSize)
	n, err := db.f.ReadAt(buf[:], int64(id-1)*int64(db.header.PageSize))
	if err != nil {
		return nil, err
	}
	if n != len(buf) {
		return nil, ErrFileTruncated
	}
	return buf, nil
}

func parseHeader(b [headerSize]byte) (header, error) {
	magic := string(b[:16])
	if magic != Magic {
		return header{}, ErrHeaderInvalidMagic
	}

	pageSize := uint(binary.BigEndian.Uint16(b[16:18]))
	if pageSize == 1 {
		pageSize = 1 << 16
	}
	isPower := func(n uint) bool {
		return bits.OnesCount(n) == 1
	}
	if pageSize < 512 || pageSize > 1<<16 || !isPower(pageSize) {
		// TODO: special case for 1
		return header{}, ErrHeaderInvalidPageSize
	}

	h := header{
		Magic:    magic,
		PageSize: int(pageSize),
	}
	return h, nil
}

type table struct {
	name string
	root TableBtree
	// TODO: point to indices, &c.
}

type index struct {
	name string
	root IndexBtree
}

// master records are defined as:
// CREATE TABLE sqlite_master(
//     type text,
//     name text,
//     tbl_name text,
//     rootpage integer,
//     sql text
// );
type Master struct {
	typ, name, tblName string
	rootPage           int
	sql                string
}

func (db *database) master() ([]Master, error) {
	master, err := db.pageMaster()
	if err != nil {
		return nil, err
	}

	var tables []Master
	_, err = master.Iter(db, func(rowid int64, pl Payload) (bool, error) {
		c, err := addOverflow(db, pl)
		if err != nil {
			return false, err
		}

		e, err := parseRecord(c)
		if err != nil {
			return false, err
		}

		m := Master{}
		if s, ok := e[0].(string); !ok {
			return false, errors.New("wrong column type for sqlite_master")
		} else {
			m.typ = s
		}
		if s, ok := e[1].(string); !ok {
			return false, errors.New("wrong column type for sqlite_master")
		} else {
			m.name = s
		}
		if s, ok := e[2].(string); !ok {
			return false, errors.New("wrong column type for sqlite_master")
		} else {
			m.tblName = s
		}
		if n, ok := e[3].(int64); !ok {
			return false, errors.New("wrong column type for sqlite_master")
		} else {
			m.rootPage = int(n)
		}
		if s, ok := e[4].(string); !ok {
			return false, errors.New("wrong column type for sqlite_master")
		} else {
			m.sql = s
		}
		tables = append(tables, m)
		return false, nil
	})
	return tables, err
}

// returns nil if the table isn't found
func (db *database) Table(name string) (*table, error) {
	tables, err := db.master()
	if err != nil {
		return nil, err
	}
	for _, t := range tables {
		if t.typ == "table" && t.name == name {
			root, err := db.openTable(t.rootPage)
			if err != nil {
				return nil, err
			}
			return &table{
				name: t.name,
				root: root,
			}, nil
		}
	}
	return nil, nil
}

// returns nil if the index isn't found
func (db *database) Index(name string) (*index, error) {
	tables, err := db.master()
	if err != nil {
		return nil, err
	}
	for _, t := range tables {
		if t.typ == "index" && t.name == name {
			root, err := db.openIndex(t.rootPage)
			if err != nil {
				return nil, err
			}
			return &index{
				name: t.name,
				root: root,
			}, nil
		}
	}
	return nil, nil
}

func (db *database) openTable(page int) (TableBtree, error) {
	buf, err := db.page(page)
	if err != nil {
		return nil, err
	}
	return newTableBtree(buf, false)
}

func (db *database) openIndex(page int) (IndexBtree, error) {
	buf, err := db.page(page)
	if err != nil {
		return nil, err
	}
	return newIndexBtree(buf)
}

// Call cb() for every row in the table. Will be called in 'database order'.
// Might return ErrNoSuchTable when the table isn't there (or isn't a table),
// or when something's wrong with the DB file.
// There is no way to bail out of the scan halfway.
func (db *database) TableScan(table string, cb func(Record)) error {
	t, err := db.Table(table)
	if err != nil {
		return err
	}
	if t == nil {
		return ErrNoSuchTable
	}
	_, err = t.root.Iter(
		db,
		func(rowid int64, pl Payload) (bool, error) {
			c, err := addOverflow(db, pl)
			if err != nil {
				return false, err
			}

			rec, err := parseRecord(c)
			if err != nil {
				return false, err
			}
			cb(rec)
			return false, nil
		},
	)
	return err
}

// Find a single rowid. Will return nil if it isn't found. The rowid is an
// internal id, but if you have a `primary key(int)` that should be the same.
// Might return ErrNoSuchTable when the table isn't there (or isn't a table),
// or when something's wrong with the DB file.
// Searching is efficient.
func (db *database) TableRowid(table string, rowid int64) (Record, error) {
	t, err := db.Table(table)
	if err != nil {
		return nil, err
	}
	if t == nil {
		return nil, ErrNoSuchTable
	}

	var recPl *Payload
	if _, err := t.root.IterMin(
		db,
		rowid,
		func(k int64, pl Payload) (bool, error) {
			if k == rowid {
				recPl = &pl
			}
			return true, nil
		},
	); err != nil {
		return nil, err
	}
	if recPl == nil {
		return nil, nil
	}

	c, err := addOverflow(db, *recPl)
	if err != nil {
		return nil, err
	}
	return parseRecord(c)
}
