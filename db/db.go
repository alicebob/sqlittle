package db

import (
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

func (db *DB) Select(table string, cb RowCB, columns ...string) error {
	if err := db.db.RLock(); err != nil {
		return err
	}
	defer db.db.RUnlock()

	s, err := db.db.Schema(table)
	if err != nil {
		return err
	}
	ci, err := toColumnIndex(s, columns)
	if err != nil {
		return err
	}
	// TODO: withoutrowid

	t, err := db.db.Table(table)
	if err != nil {
		return err
	}
	return t.Scan(func(_ int64, r sqlittle.Record) bool {
		row := make([]interface{}, len(ci))
		for i, c := range ci {
			// TODO: use 'DEFAULT' when the record is too short
			row[i] = r[c]
		}
		cb(row)
		return false
	})
}

func toColumnIndex(s *sqlittle.SchemaTable, columns []string) ([]int, error) {
	res := make([]int, 0, len(columns))
	for _, c := range columns {
		i := s.Column(c)
		// TODO: accept 'rowid' and friends
		if i < 0 {
			return nil, fmt.Errorf("no such column: %q", c)
		}
		res = append(res, i)
	}
	return res, nil
}
