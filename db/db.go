package db

import (
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
