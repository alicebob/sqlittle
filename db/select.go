package db

import (
	"fmt"

	"github.com/alicebob/sqlittle"
)

func selectWithRowid(db *sqlittle.Database, s *sqlittle.SchemaTable, cb RowCB, columns []string) error {
	ci, err := toColumnIndex(s, columns)
	if err != nil {
		return err
	}

	t, err := db.Table(s.Table)
	if err != nil {
		return err
	}
	return t.Scan(func(_ int64, r sqlittle.Record) bool {
		row := make(Row, len(ci))
		for i, c := range ci {
			// TODO: use 'DEFAULT' when the record is too short
			row[i] = r[c]
		}
		cb(row)
		return false
	})
}

func selectWithoutRowid(db *sqlittle.Database, s *sqlittle.SchemaTable, cb RowCB, columns []string) error {
	ci, err := toColumnIndex(s, columns)
	if err != nil {
		return err
	}

	t, err := db.Table(s.Table)
	if err != nil {
		return err
	}
	return t.WithoutRowidScan(func(r sqlittle.Record) bool {
		row := make(Row, len(ci))
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
