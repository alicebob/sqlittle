package sqlittle

import (
	"errors"
	"math/bits"
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
	journal string
	// dirty       bool // reload header if true
	l pager
	// header      *header
	// btreeCache *btreeCache // table and index page cache
	// objectCache *objectCache
}

// OpenFile opens a .sqlite file. This is the main entry point.
// Use database.Close() when done.
func OpenFile(f string) (*Database, error) {
	l, err := newFilePager(f)
	if err != nil {
		return nil, err
	}
	h, err := getHeader(l)
	switch h.Mode {
	case ModeJournal:
		return newJournalDB(l, f+"-journal")
	case ModeWal:
		return newWalDB(l, f)
	default:
		panic("impossible")
	}
}

func newJournalDB(l pager, journal string) (*Database, error) {
	db := &Database{
		journal: journal,
		// dirty:      true,
		l: l,
		// btreeCache: newBtreeCache(CachePages),
		// header: &h,
	}
	return db, db.checkJournal()
}

func (db *Database) checkJournal() error {
	// this should go elsewhere
	if db.journal != "" {
		hot, err := validJournal(db.journal)
		if err != nil {
			return err
		}
		if hot {
			// If something is using the transaction the db will have a RESERVED
			// lock.
			locked, err := db.l.CheckReservedLock()
			if err != nil {
				return err
			}
			if !locked {
				return ErrHotJournal
			}
		}
	}
	return nil
}

// Close the database.
func (db *Database) Close() error {
	return db.l.Close()
}

// Lock database for reading. Blocks. Don't nest RLock() calls.
func (db *Database) RLock() error {
	if err := db.checkJournal(); err != nil {
		return err
	}
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
	return db.l.page(id)
}

func (db *Database) resolveDirty() error {
	return nil
	/*
		if !db.dirty {
			return nil
		}

		if db.journal != "" {
			hot, err := validJournal(db.journal)
			if err != nil {
				return err
			}
			if hot {
				// If something is using the transaction the db will have a RESERVED
				// lock.
				locked, err := db.l.CheckReservedLock()
				if err != nil {
					return err
				}
				if !locked {
					return ErrHotJournal
				}
			}
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
			db.btreeCache.clear()
		}
		if db.header != nil && db.header.SchemaCookie != newHeader.SchemaCookie {
			db.objectCache = nil
		}
		db.dirty = false
		db.header = &newHeader
		return nil
	*/
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
	if err := db.resolveDirty(); err != nil {
		return nil, err
	}

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

// openPage returns a tableBtree or indexBtree
func (db *Database) openPage(page int) (interface{}, error) {
	if err := db.resolveDirty(); err != nil {
		return nil, err
	}

	// if db.btreeCache != nil {
	// if p := db.btreeCache.get(page); p != nil {
	// return p, nil
	// }
	// }

	buf, err := db.page(page)
	if err != nil {
		return nil, err
	}
	p, err := newBtree(buf, page == 1)
	// if err == nil && db.btreeCache != nil {
	// db.btreeCache.set(page, p)
	// }
	return p, err
}

func (db *Database) openTable(page int) (tableBtree, error) {
	p, err := db.openPage(page)
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
	p, err := db.openPage(page)
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

func validPageSize(s uint) bool {
	return s >= 512 && s <= 1<<16 && bits.OnesCount(s) == 1
}
