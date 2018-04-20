package db

import (
	"github.com/alicebob/sqlittle"
)

func indexedSelect(
	db *sqlittle.Database,
	schema *sqlittle.Schema,
	index *sqlittle.SchemaIndex,
	cb RowCB,
	columns []string,
) error {
	ci, err := toColumnIndexRowid(schema, columns)
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

func indexedSelectEq(
	db *sqlittle.Database,
	schema *sqlittle.Schema,
	index *sqlittle.SchemaIndex,
	key Row,
	cb RowCB,
	columns []string,
) error {
	ci, err := toColumnIndexRowid(schema, columns)
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

	return ind.ScanEq(
		sqlittle.Record(key),
		func(r sqlittle.Record) bool {
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

func indexedSelectNonRowid(
	db *sqlittle.Database,
	schema *sqlittle.Schema,
	index *sqlittle.SchemaIndex,
	cb RowCB,
	columns []string,
) error {
	ci, err := toColumnIndexNonRowid(schema, columns)
	if err != nil {
		return err
	}

	tab, err := db.WithoutRowidTable(schema.Table)
	if err != nil {
		return err
	}

	ind, err := db.Index(index.Index)
	if err != nil {
		return err
	}

	cols := pkColumns(schema, index)
	return ind.Scan(func(r sqlittle.Record) bool {
		pk := reRecord(r, cols)

		var found sqlittle.Record
		err := tab.ScanEq(pk, func(row sqlittle.Record) bool { found = row; return true })
		if err != nil || found == nil {
			// should never be nil
			return false
		}
		cb(toRow(0, ci, found))
		return false
	})
}

func indexedSelectEqNonRowid(
	db *sqlittle.Database,
	schema *sqlittle.Schema,
	index *sqlittle.SchemaIndex,
	key Row,
	cb RowCB,
	columns []string,
) error {
	ci, err := toColumnIndexNonRowid(schema, columns)
	if err != nil {
		return err
	}

	tab, err := db.WithoutRowidTable(schema.Table)
	if err != nil {
		return err
	}

	ind, err := db.Index(index.Index)
	if err != nil {
		return err
	}

	cols := pkColumns(schema, index)
	return ind.ScanEq(
		sqlittle.Record(key),
		func(r sqlittle.Record) bool {
			pk := reRecord(r, cols)

			var found sqlittle.Record
			err := tab.ScanEq(pk, func(row sqlittle.Record) bool { found = row; return true })
			if err != nil || found == nil {
				// should never be nil
				return false
			}
			cb(toRow(0, ci, found))
			return false
		},
	)
}
