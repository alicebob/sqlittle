package db

import (
	"github.com/alicebob/sqlittle"
)

func selectWithRowid(db *sqlittle.Database, s *sqlittle.Schema, cb RowCB, columns []string) error {
	ci, err := toColumnIndexRowid(s, columns)
	if err != nil {
		return err
	}

	t, err := db.Table(s.Table)
	if err != nil {
		return err
	}
	return t.Scan(func(rowid int64, r sqlittle.Record) bool {
		cb(toRow(rowid, ci, r))
		return false
	})
}

func selectWithoutRowid(db *sqlittle.Database, s *sqlittle.Schema, cb RowCB, columns []string) error {
	ci, err := toColumnIndexNonRowid(s, columns)
	if err != nil {
		return err
	}

	t, err := db.Table(s.Table)
	if err != nil {
		return err
	}
	return t.WithoutRowidScan(func(r sqlittle.Record) bool {
		cb(toRow(0, ci, r))
		return false
	})
}

func selectRowid(db *sqlittle.Database, s *sqlittle.Schema, rowid int64, columns []string) (Row, error) {
	ci, err := toColumnIndexRowid(s, columns)
	if err != nil {
		return nil, err
	}

	t, err := db.Table(s.Table)
	if err != nil {
		return nil, err
	}
	r, err := t.Rowid(rowid)
	if err != nil || r == nil {
		return nil, err
	}
	// TODO: decide what to do with shared []byte pointers
	return toRow(rowid, ci, r), nil
}
