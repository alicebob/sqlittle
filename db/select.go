package db

import (
	"fmt"

	"github.com/alicebob/sqlittle"
)

func selectWithRowid(db *sqlittle.Database, s *sqlittle.SchemaTable, cb RowCB, columns []string) error {
	ci, err := toColumnIndex(s, columns, true)
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

func selectWithoutRowid(db *sqlittle.Database, s *sqlittle.SchemaTable, cb RowCB, columns []string) error {
	ci, err := toColumnIndex(s, columns, false)
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

func toRow(rowid int64, cis []columIndex, r sqlittle.Record) Row {
	row := make(Row, len(cis))
	for i, c := range cis {
		if c.rowid {
			row[i] = rowid
			continue
		}
		if len(r) <= c.rowIndex {
			// use 'DEFAULT' when the record is too short
			row[i] = c.col.Default
		} else {
			row[i] = r[c.rowIndex]
		}
	}
	return row
}

type columIndex struct {
	col      *sqlittle.TableColumn
	rowIndex int
	rowid    bool
}

// given column names returns the index in a Row this column is expected, and
// the column definition
func toColumnIndex(s *sqlittle.SchemaTable, columns []string, allowRowid bool) ([]columIndex, error) {
	res := make([]columIndex, 0, len(columns))
	for _, c := range columns {
		n := s.Column(c)
		if n < 0 {
			if allowRowid && (c == "rowid" || c == "oid" || c == "_rowid_") {
				res = append(res, columIndex{nil, n, true})
				continue
			} else {
				return nil, fmt.Errorf("no such column: %q", c)
			}
		}
		res = append(res, columIndex{&s.Columns[n], n, false})
	}
	return res, nil
}
