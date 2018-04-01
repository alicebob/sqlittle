package sqlittle

import (
	"encoding/binary"
	"errors"
	"math/bits"
)

const (
	headerMagic = "SQLite format 3\x00"
	headerSize  = 100
)

var (
	ErrHeaderInvalidMagic    = errors.New("invalid magic number")
	ErrHeaderInvalidPageSize = errors.New("invalid page size")
	ErrNoSuchTable           = errors.New("no such table")
	ErrNoSuchIndex           = errors.New("no such index")
	ErrCorrupted             = errors.New("database corrupted")
)

type header struct {
	Magic    string
	PageSize int
}

type Database struct {
	l      loader
	header header
}

// OpenFile opens a .sqlite file. This is the main entry point.
// Use database.Close() when done.
func OpenFile(f string) (*Database, error) {
	l, err := newMMapLoader(f)
	if err != nil {
		return nil, err
	}
	return newDatabase(l)
}

func newDatabase(l loader) (*Database, error) {
	buf, err := l.header()
	if err != nil {
		return nil, err
	}
	header, err := parseHeader(buf)
	if err != nil {
		return nil, err
	}

	db := &Database{
		l:      l,
		header: header,
	}
	return db, nil
}

// Close the database.
func (db *Database) Close() error {
	return db.l.Close()
}

func (db *Database) pageMaster() (tableBtree, error) {
	buf, err := db.page(1)
	if err != nil {
		return nil, err
	}
	return newTableBtree(buf, true)
}

// n starts at 1, sqlite style
func (db *Database) page(id int) ([]byte, error) {
	if id < 1 {
		return nil, errors.New("invalid page number")
	}
	return db.l.page(id, db.header.PageSize)
}

func parseHeader(b [headerSize]byte) (header, error) {
	magic := string(b[:16])
	if magic != headerMagic {
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
	root tableBtree
	// TODO: point to indices, &c.
}

type index struct {
	name string
	root indexBtree
}

// master records are defined as:
// CREATE TABLE sqlite_master(
//     type text,
//     name text,
//     tbl_name text,
//     rootpage integer,
//     sql text
// );
type sqliteMaster struct {
	typ, name, tblName string
	rootPage           int
	sql                string
}

func (db *Database) master() ([]sqliteMaster, error) {
	master, err := db.pageMaster()
	if err != nil {
		return nil, err
	}

	var tables []sqliteMaster
	_, err = master.Iter(db, func(rowid int64, pl cellPayload) (bool, error) {
		c, err := addOverflow(db, pl)
		if err != nil {
			return false, err
		}

		e, err := parseRecord(c)
		if err != nil {
			return false, err
		}

		m := sqliteMaster{}
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
func (db *Database) table(name string) (*table, error) {
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
func (db *Database) index(name string) (*index, error) {
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

func (db *Database) openTable(page int) (tableBtree, error) {
	buf, err := db.page(page)
	if err != nil {
		return nil, err
	}
	return newTableBtree(buf, false)
}

func (db *Database) openIndex(page int) (indexBtree, error) {
	buf, err := db.page(page)
	if err != nil {
		return nil, err
	}
	return newIndexBtree(buf)
}

// TablScanCB is the callback for TableScan(). It gets the rowid (usually an
// internal number), and the data of a row. It should return true when the scan
// should be terminated.
type TableScanCB func(int64, Record) bool

// TableScan calls cb() for every row in the table. Will be called in 'database order'.
// Will return ErrNoSuchTable when the table isn't there (or isn't a table).
// The record is given as sqlite stores it; this means:
//  - float64 columns might be stored as int64
//  - after an alter table which adds columns a row might miss those columns
//  - "integer primary key" column will be always be nil, and the rowid is the
//  value
// If the callback returns true (done) the scan will be stopped.
func (db *Database) TableScan(table string, cb TableScanCB) error {
	t, err := db.table(table)
	if err != nil {
		return err
	}
	if t == nil {
		return ErrNoSuchTable
	}
	_, err = t.root.Iter(
		db,
		func(rowid int64, pl cellPayload) (bool, error) {
			c, err := addOverflow(db, pl)
			if err != nil {
				return false, err
			}

			rec, err := parseRecord(c)
			if err != nil {
				return false, err
			}
			return cb(rowid, rec), nil
		},
	)
	return err
}

// TableRowid finds a single row by rowid. Will return nil if it isn't found.
// The rowid is an internal id, but if you have an `integer primary key` column
// that should be the same.
// Will return ErrNoSuchTable when the table isn't there (or isn't a table).
func (db *Database) TableRowid(table string, rowid int64) (Record, error) {
	t, err := db.table(table)
	if err != nil {
		return nil, err
	}
	if t == nil {
		return nil, ErrNoSuchTable
	}

	var recPl *cellPayload
	if _, err := t.root.IterMin(
		db,
		rowid,
		func(k int64, pl cellPayload) (bool, error) {
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

// IndexScanCB is passed to IndexScan() and IndexScanMin(). It gets the rowid
// and the values from the index. It should return true when the scan should be
// stopped.
type IndexScanCB func(int64, Record) bool

// IndexScan calls cb() for every row in the index. These will be called in the
// index order.
// The callback gets the rowid the row is about (use TableRowid() to load the
// row, if you need it), and all the columns present in the index.
// If the callback returns true (done) the scan will be stopped.
func (db *Database) IndexScan(index string, cb IndexScanCB) error {
	ind, err := db.index(index)
	if err != nil {
		return err
	}
	if ind == nil {
		return ErrNoSuchIndex
	}
	_, err = ind.root.Iter(
		db,
		func(pl cellPayload) (bool, error) {
			full, err := addOverflow(db, pl)
			if err != nil {
				return false, err
			}
			rec, err := parseRecord(full)
			if err != nil {
				return false, err
			}
			rowid, rec, err := chompRowid(rec)
			if err != nil {
				return false, err
			}
			return cb(rowid, rec), nil
		},
	)
	return err
}

// IndexScanMin calls cb() for every row in the index, starting from the first
// record equal or bigger then the given record. If the type of columns in the given
// record don't match those in the index a error will be returned.
// If the callback returns true (done) the scan will be stopped.
// All comments from IndexScan are valid here as well.
func (db *Database) IndexScanMin(index string, from Record, cb IndexScanCB) error {
	ind, err := db.index(index)
	if err != nil {
		return err
	}
	if ind == nil {
		return ErrNoSuchIndex
	}
	_, err = ind.root.IterMin(
		db,
		from,
		func(rowid int64, rec Record) (bool, error) {
			return cb(rowid, rec), nil
		},
	)
	return err
}
