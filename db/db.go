package db

import (
	"errors"
	"fmt"

	"github.com/alicebob/sqlittle"
)

type DB struct {
	db *sqlittle.Database
}

func Open(filename string) (*DB, error) {
	db, err := sqlittle.OpenFile(filename)
	if err != nil {
		return nil, err
	}
	return &DB{
		db: db,
	}, nil
}

func (db *DB) Close() error {
	return db.db.Close()
}

type RowCB func(Row)

// Select the columns from every row from the given table. Order is the rowid
// order for rowid tables, and the orderd primary key for non-rowid tables.
//
// For rowid tables the special values "rowid", "oid", and "_rowid_" will load
// the rowid (unless there is a column with that name).
func (db *DB) Select(table string, cb RowCB, columns ...string) error {
	if err := db.db.RLock(); err != nil {
		return err
	}
	defer db.db.RUnlock()

	s, err := db.db.Schema(table)
	if err != nil {
		return err
	}

	if s.WithoutRowid {
		return selectWithoutRowid(db.db, s, cb, columns)
	} else {
		return selectWithRowid(db.db, s, cb, columns)
	}
}

// Select all rows from the given table via the index. The order will be the
// index order.
func (db *DB) IndexedSelect(table, index string, cb RowCB, columns ...string) error {
	if err := db.db.RLock(); err != nil {
		return err
	}
	defer db.db.RUnlock()

	s, err := db.db.Schema(table)
	if err != nil {
		return err
	}

	ind := s.NamedIndex(index)
	if ind == nil {
		return fmt.Errorf("no such index: %q", index)
	}

	if s.WithoutRowid {
		return errors.New("WITHOUT ROWID not supported yet")
		// return selectWithoutRowid(db.db, s, cb, columns)
	} else {
		return indexedSelectRowid(db.db, s, ind, cb, columns)
	}
}

// func (db *DB) SelectWhere(table string, cb RowCB, columns ...string) error

// func (db *DB) IndexedSelectWhere(table string, cb RowCB, columns ...string) error
