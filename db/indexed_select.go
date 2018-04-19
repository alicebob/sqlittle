package db

import (
	"github.com/alicebob/sqlittle"
)

func indexedSelectRowid(
	db *sqlittle.Database,
	schema *sqlittle.SchemaTable,
	index *sqlittle.SchemaIndex,
	cb RowCB,
	columns []string,
) error {
	ci, err := toColumnIndex(schema, columns, true)
	if err != nil {
		return err
	}

	tab, err := db.Table(schema.Table)
	if err != nil {
		return err
	}

	ind, err := db.Index(index.Index)
	if err != nil {
		return err
	}

	return ind.Scan(func(r sqlittle.Record) bool {
		rowid, _, err := sqlittle.ChompRowid(r)
		if err != nil {
			return false
		}
		row, err := tab.Rowid(rowid)
		if err != nil || row == nil {
			// should never be nil
			return false
		}
		cb(toRow(rowid, ci, row))
		return false
	})
}

func indexedSelectWithoutRowid(
	db *sqlittle.Database,
	schema *sqlittle.SchemaTable,
	index *sqlittle.SchemaIndex,
	cb RowCB,
	columns []string,
) error {
	ci, err := toColumnIndex(schema, columns, false)
	if err != nil {
		return err
	}

	tab, err := db.Table(schema.Table)
	if err != nil {
		return err
	}

	ind, err := db.Index(index.Index)
	if err != nil {
		return err
	}

	pkIndexes := []int{1} // FIXME :)

	return ind.Scan(func(r sqlittle.Record) bool {
		pk := reRecord(r, pkIndexes)

		row, err := tab.WithoutRowidPK(pk)
		if err != nil || row == nil {
			// should never be nil
			return false
		}
		cb(toRow(0, ci, row))
		return false
	})
}
