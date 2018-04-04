package sqlittle

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math/bits"
)

const (
	headerMagic = "SQLite format 3\x00"
	headerSize  = 100
	// CacheTables is the number of tables pages to keep in memory. Default
	// size per page is about 4K.
	CacheTables = 100
)

var (
	ErrHeaderInvalidMagic    = errors.New("invalid magic number")
	ErrHeaderInvalidPageSize = errors.New("invalid page size")
	ErrNoSuchTable           = errors.New("no such table")
	ErrNoSuchIndex           = errors.New("no such index")
	ErrCorrupted             = errors.New("database corrupted")
	ErrIncompatible          = errors.New("incompatible database version")
	ErrEncoding              = errors.New("unsupported encoding")
	ErrInvalidDef            = errors.New("invalid object definition")
	ErrRecursion             = errors.New("tree is too deep")
)

type header struct {
	// The database page size in bytes.
	PageSize int
	// Bytes of unused "reserved" space at the end of each page. Usually 0.
	ReservedSpace int
	// Updated when anything changes (only for non-WAL files).
	ChangeCounter uint32
}

type Database struct {
	dirty  bool // reload header if true
	l      pager
	header *header
	tables *tableCache
}

// OpenFile opens a .sqlite file. This is the main entry point.
// Use database.Close() when done.
func OpenFile(f string) (*Database, error) {
	l, err := newFilePager(f)
	if err != nil {
		return nil, err
	}
	return newDatabase(l)
}

func newDatabase(l pager) (*Database, error) {
	d := &Database{
		dirty:  true,
		l:      l,
		tables: newTableCache(CacheTables),
	}
	return d, d.resolveDirty()
}

// Close the database.
func (db *Database) Close() error {
	return db.l.Close()
}

// Lock database for reading. Blocks. Don't nest RLock() calls.
func (db *Database) RLock() error {
	db.dirty = true
	return db.l.RLock()
}

// Unlock a read lock. Use a single RUnlock() for every RLock().
func (db *Database) RUnlock() error {
	return db.l.RUnlock()
}

// n starts at 1, sqlite style
func (db *Database) page(id int) ([]byte, error) {
	if id < 1 {
		return nil, errors.New("invalid page number")
	}
	return db.l.page(id, db.header.PageSize)
}

// the file header, as described in "1.2. The Database Header"
func parseHeader(b [headerSize]byte) (header, error) {
	hs := struct {
		Magic                [16]byte
		PageSize             uint16
		WriteVersion         uint8
		ReadVersion          uint8
		ReservedSpace        uint8
		MaxFraction          uint8
		MinFraction          uint8
		LeafFraction         uint8
		ChangeCounter        uint32
		_                    uint32
		_                    uint32
		_                    uint32
		_                    uint32 // SchemaCookie
		SchemaFormat         uint32
		_                    uint32
		_                    uint32
		TextEncoding         uint32
		_                    uint32
		_                    uint32
		_                    uint32
		ReservedForExpansion [20]byte
		_                    uint32
		_                    uint32
	}{}
	if err := binary.Read(bytes.NewBuffer(b[:]), binary.BigEndian, &hs); err != nil {
		return header{}, err
	}

	h := header{}

	if string(hs.Magic[:]) != headerMagic {
		return h, ErrHeaderInvalidMagic
	}

	{
		s := uint(hs.PageSize)
		if s == 1 {
			s = 1 << 16
		}
		isPower := func(n uint) bool {
			return bits.OnesCount(n) == 1
		}
		if s < 512 || s > 1<<16 || !isPower(s) {
			return header{}, ErrHeaderInvalidPageSize
		}
		h.PageSize = int(s)
	}

	if hs.ReadVersion > 2 {
		return h, ErrIncompatible
	}

	h.ReservedSpace = int(hs.ReservedSpace)

	if hs.MaxFraction != 64 ||
		hs.MinFraction != 32 ||
		hs.LeafFraction != 32 {
		return h, ErrIncompatible
	}

	h.ChangeCounter = hs.ChangeCounter

	// 1,2,3,4 are the only valid values. Version 1 ignores 'DESC' on indexes,
	// so we could support that as long as we ignore any 'DESC' index, but...
	switch hs.SchemaFormat {
	case 2, 3, 4:
	default:
		return h, ErrIncompatible
	}

	switch hs.TextEncoding {
	case 1:
		// UTF8. It's the only thing we currently support
	case 2, 3:
		// UTF16le and UTF16be
		return h, ErrEncoding
	default:
		return h, ErrIncompatible
	}

	for _, v := range hs.ReservedForExpansion {
		if v != 0 {
			return h, ErrIncompatible
		}
	}

	return h, nil
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

func (db *Database) resolveDirty() error {
	if !db.dirty {
		return nil
	}

	buf, err := db.l.header()
	if err != nil {
		return err
	}
	newHeader, err := parseHeader(buf)
	if err != nil {
		return err
	}
	if db.header != nil && db.header.ChangeCounter != newHeader.ChangeCounter {
		db.tables.clear()
	}
	db.dirty = false
	db.header = &newHeader
	return nil
}

func (db *Database) master() ([]sqliteMaster, error) {
	if err := db.resolveDirty(); err != nil {
		return nil, err
	}

	master, err := db.openTable(1)
	if err != nil {
		return nil, err
	}

	var tables []sqliteMaster
	_, err = master.Iter(maxRecursion, db, func(rowid int64, pl cellPayload) (bool, error) {
		c, err := addOverflow(db, pl)
		if err != nil {
			return false, err
		}

		e, err := parseRecord(c)
		if err != nil {
			return false, err
		}
		if len(e) != 5 {
			return false, ErrInvalidDef
		}

		m := sqliteMaster{}
		if s, ok := e[0].(string); !ok {
			return false, ErrInvalidDef
		} else {
			m.typ = s
		}
		if s, ok := e[1].(string); !ok {
			return false, ErrInvalidDef
		} else {
			m.name = s
		}
		if s, ok := e[2].(string); !ok {
			return false, ErrInvalidDef
		} else {
			m.tblName = s
		}
		if n, ok := e[3].(int64); !ok {
			return false, ErrInvalidDef
		} else {
			m.rootPage = int(n)
		}
		if s, ok := e[4].(string); !ok {
			return false, ErrInvalidDef
		} else {
			m.sql = s
		}
		tables = append(tables, m)
		return false, nil
	})
	return tables, err
}

func (db *Database) openTable(page int) (tableBtree, error) {
	if err := db.resolveDirty(); err != nil {
		return nil, err
	}

	if p := db.tables.get(page); p != nil {
		return p, nil
	}

	buf, err := db.page(page)
	if err != nil {
		return nil, err
	}
	p, err := newTableBtree(buf, page == 1)
	if err == nil {
		db.tables.set(page, p)
	}
	return p, err
}

func (db *Database) openIndex(page int) (indexBtree, error) {
	if err := db.resolveDirty(); err != nil {
		return nil, err
	}

	buf, err := db.page(page)
	if err != nil {
		return nil, err
	}
	return newIndexBtree(buf)
}

// Tables lists the table names.
func (db *Database) Tables() ([]string, error) {
	objects, err := db.master()
	if err != nil {
		return nil, err
	}
	var tables []string
	for _, o := range objects {
		if o.typ == "table" {
			tables = append(tables, o.name)
		}
	}
	return tables, nil
}

// Table opens the named table.
// Will return ErrNoSuchTable when the table isn't there (or isn't a table).
// Table pointer is always valid if err == nil.
func (db *Database) Table(name string) (*Table, error) {
	objects, err := db.master()
	if err != nil {
		return nil, err
	}
	for _, o := range objects {
		if o.typ == "table" && o.name == name {
			return &Table{db: db, root: o.rootPage}, nil
		}
	}
	return nil, ErrNoSuchTable
}

// Index opens the named index.
// Will return ErrNoSuchIndex when the index isn't there (or isn't an index).
// Index pointer is always valid if err == nil.
func (db *Database) Index(name string) (*Index, error) {
	tables, err := db.master()
	if err != nil {
		return nil, err
	}
	for _, t := range tables {
		if t.typ == "index" && t.name == name {
			return &Index{db: db, root: t.rootPage}, nil
		}
	}
	return nil, ErrNoSuchIndex
}
