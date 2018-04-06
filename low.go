// all public low level methods

package sqlittle

import (
	"errors"

	"github.com/alicebob/sqlittle/sql"
)

type Table struct {
	db   *Database
	root int
	sql  string
}

// Def returns the table definition. Not everything SQLite supports is
// supported (yet).
func (t *Table) Def() (*sql.CreateTableStmt, error) {
	c, err := sql.Parse(t.sql)
	if err != nil {
		return nil, err
	}
	stmt, ok := c.(sql.CreateTableStmt)
	if !ok {
		return nil, errors.New("no CREATE TABLE attached")
	}
	return &stmt, nil
}

type Index struct {
	db   *Database
	root int
}

// TableScanCB is the callback for Table.Scan(). It gets the rowid (usually an
// internal number), and the data of the row. It should return true when the
// scan should be terminated.
type TableScanCB func(int64, Record) bool

// Scan calls cb() for every row in the table. Will be called in 'database
// order'.
// The record is given as sqlite stores it; this means:
//  - float64 columns might be stored as int64
//  - after an alter table which adds columns a row might miss the new columns
//  - an "integer primary key" column will be always be nil, and the rowid is
//  the value
// If the callback returns true (done) the scan will be stopped.
func (t *Table) Scan(cb TableScanCB) error {
	root, err := t.db.openTable(t.root)
	if err != nil {
		return err
	}
	_, err = root.Iter(
		maxRecursion,
		t.db,
		func(rowid int64, pl cellPayload) (bool, error) {
			c, err := addOverflow(t.db, pl)
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

// Rowid finds a single row by rowid. Will return nil if it isn't found.
// The rowid is an internal id, but if you have an `integer primary key` column
// that should be the same.
// See Table.Scan comments about the Record
func (t *Table) Rowid(rowid int64) (Record, error) {
	root, err := t.db.openTable(t.root)
	if err != nil {
		return nil, err
	}

	var recPl *cellPayload
	if _, err := root.IterMin(
		maxRecursion,
		t.db,
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

	c, err := addOverflow(t.db, *recPl)
	if err != nil {
		return nil, err
	}
	return parseRecord(c)
}

// IndexScanCB is passed to Index.Scan() and Index.ScanMin(). It gets the rowid
// and the values from the index. It should return true when the scan should be
// stopped.
type IndexScanCB func(int64, Record) bool

// Scan calls cb() for every row in the index. These will be called in the
// index order.
// The callback gets the rowid the row is about (use Table.Rowid() to load the
// row, if you need it), and all the columns present in the index.
// If the callback returns true (done) the scan will be stopped.
func (in *Index) Scan(cb IndexScanCB) error {
	root, err := in.db.openIndex(in.root)
	if err != nil {
		return err
	}

	_, err = root.Iter(
		maxRecursion,
		in.db,
		func(pl cellPayload) (bool, error) {
			full, err := addOverflow(in.db, pl)
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

// ScanMin calls cb() for every row in the index, starting from the first
// record equal or bigger than the given record. If the type of columns in the given
// record don't match those in the index an error will be returned.
// If the callback returns true (done) the scan will be stopped.
// All comments from Index.Scan are valid here as well.
func (in *Index) ScanMin(from Record, cb IndexScanCB) error {
	root, err := in.db.openIndex(in.root)
	if err != nil {
		return err
	}

	_, err = root.IterMin(
		maxRecursion,
		in.db,
		from,
		func(rowid int64, rec Record) (bool, error) {
			return cb(rowid, rec), nil
		},
	)
	return err
}
