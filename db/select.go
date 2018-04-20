package db

import (
	"errors"

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

	t, err := db.WithoutRowidTable(s.Table)
	if err != nil {
		return err
	}
	return t.Scan(func(r sqlittle.Record) bool {
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

func pkSelect(db *sqlittle.Database, s *sqlittle.Schema, key Row, cb RowCB, columns []string) error {
	if s.RowidPK {
		// `integer primary key` table.
		var rowid int64
		if err := key.Scan(&rowid); err != nil {
			return err
		}
		row, err := selectRowid(db, s, rowid, columns)
		if err != nil {
			return err
		}
		if row != nil {
			cb(row)
		}
		return nil
	}
	ind := s.NamedIndex(s.PrimaryKey)
	if ind == nil {
		return errors.New("table has no primary key")
	}
	return indexedSelectEq(
		db,
		s,
		ind,
		key,
		cb,
		columns,
	)
}

func pkSelectNonRowid(db *sqlittle.Database, s *sqlittle.Schema, key Row, cb RowCB, columns []string) error {
	ci, err := toColumnIndexNonRowid(s, columns)
	if err != nil {
		return err
	}
	t, err := db.WithoutRowidTable(s.Table)
	if err != nil {
		return err
	}

	return t.ScanEq(
		sqlittle.Record(key),
		func(r sqlittle.Record) bool {
			cb(toRow(0, ci, r))
			return false
		},
	)
}
