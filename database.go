package sqlittle

import (
	"errors"
)

const (
	headerMagic = "SQLite format 3\x00"
	headerSize  = 100
	// CachePages is the number of pages to keep in memory. Default size per
	// page is 4K (1K on older databases).
	CachePages = 100

	ModeJournal = 1
	ModeWal     = 2
)

var (
	// Various error messages returned when the database is corrupted
	ErrInvalidMagic    = errors.New("invalid magic number")
	ErrInvalidPageSize = errors.New("invalid page size")
	ErrReservedSpace   = errors.New("unsupported database (encrypted?)")
	ErrCorrupted       = errors.New("database corrupted")
	ErrInvalidDef      = errors.New("invalid object definition")
	ErrRecursion       = errors.New("tree is too deep")
	ErrFileTruncated   = errors.New("file truncated")

	// Various error messages returned when the database uses features sqlittle
	// doesn't support.
	ErrIncompatible = errors.New("incompatible database version")
	// There is a stale `-journal` file present with an unfinished transaction.
	// Open the database in sqlite3 to repair the database.
	ErrHotJournal = errors.New("crashed transaction present")

	ErrNoSuchTable = errors.New("no such table")
	ErrNoSuchIndex = errors.New("no such index")
)

// type objectCache struct {
// objects []sqliteMaster
// err     error
// }

type Database struct {
	l pager
	f format
	// header      *header
	// objectCache *objectCache
}

// OpenFile opens a .sqlite file. This is the main entry point.
// Use database.Close() when done.
func OpenFile(f string) (*Database, error) {
	l, err := newFilePager(f)
	if err != nil {
		return nil, err
	}
	db := &Database{
		l: l,
	}
	h, err := getHeader(l)
	var fm format
	switch h.Mode {
	case ModeJournal:
		fm, err = newFJournal(l, f+"-journal")
		if err != nil {
			return nil, err
		}
	case ModeWal:
		fm, err = newFormatWal(l, f)
		if err != nil {
			return nil, err
		}
	default:
		panic("impossible")
	}
	db.f = fm
	return db, nil
}

// Close the database.
func (db *Database) Close() error {
	if err := db.l.Close(); err != nil {
		return err
	}
	if err := db.f.Close(); err != nil {
		return err
	}
	return nil
}

// Lock database for reading. Blocks. Don't nest RLock() calls.
func (db *Database) RLock() error {
	return db.f.RLock()
}

// Unlock a read lock. Use a single RUnlock() for every RLock().
func (db *Database) RUnlock() error {
	return db.f.RUnlock()
}

// n starts at 1, sqlite style
func (db *Database) page(id int) ([]byte, error) {
	if id < 1 {
		return nil, errors.New("invalid page number")
	}
	return db.l.page(id)
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
	// if o := db.objectCache; o != nil {
	// return o.objects, o.err
	// }

	master, err := db.openTable(1)
	if err != nil {
		return nil, err
	}

	var objects []sqliteMaster
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
		switch s := e[4].(type) {
		case string:
			m.sql = s
		case nil:
		default:
			return false, ErrInvalidDef
		}
		objects = append(objects, m)
		return false, nil
	})

	// db.objectCache = &objectCache{
	// objects: objects,
	// err:     err,
	// }

	return objects, err
}

func (db *Database) openTable(page int) (tableBtree, error) {
	p, err := db.f.Page(page)
	if err != nil {
		return nil, err
	}
	tb, ok := p.(tableBtree)
	if !ok {
		return nil, errors.New("found an index, expected a table")
	}
	return tb, nil
}

func (db *Database) openIndex(page int) (indexBtree, error) {
	p, err := db.f.Page(page)
	if err != nil {
		return nil, err
	}
	tb, ok := p.(indexBtree)
	if !ok {
		return nil, errors.New("found a table, expected an index")
	}
	return tb, nil
}

// Tables lists all table names. Also sqlite internal ones.
func (db *Database) Tables() ([]string, error) {
	return db.objectNames("table")
}

// Indexes lists all index names.
func (db *Database) Indexes() ([]string, error) {
	return db.objectNames("index")
}

func (db *Database) objectNames(typ string) ([]string, error) {
	objects, err := db.master()
	if err != nil {
		return nil, err
	}
	var names []string
	for _, o := range objects {
		if o.typ == typ {
			names = append(names, o.name)
		}
	}
	return names, nil
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
			return &Table{db: db, root: o.rootPage, sql: o.sql}, nil
		}
	}
	return nil, ErrNoSuchTable
}

// Index opens the named index.
// Will return ErrNoSuchIndex when the index isn't there (or isn't an index).
// Index pointer is always valid if err == nil.
func (db *Database) Index(name string) (*Index, error) {
	objects, err := db.master()
	if err != nil {
		return nil, err
	}
	for _, o := range objects {
		if o.typ == "index" && o.name == name {
			return &Index{db: db, root: o.rootPage, sql: o.sql}, nil
		}
	}
	return nil, ErrNoSuchIndex
}
