package sqlittle

import (
	sdb "github.com/alicebob/sqlittle/db"
)

func indexedSelect(
	db *sdb.Database,
	schema *sdb.Schema,
	index *sdb.SchemaIndex,
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

	return ind.Scan(func(r sdb.Record) bool {
		rowid, _, err := sdb.ChompRowid(r)
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
	db *sdb.Database,
	schema *sdb.Schema,
	index *sdb.SchemaIndex,
	key sdb.Key,
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
		key,
		func(r sdb.Record) bool {
			rowid, _, err := sdb.ChompRowid(r)
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
	db *sdb.Database,
	schema *sdb.Schema,
	index *sdb.SchemaIndex,
	cb RowCB,
	columns []string,
) error {
	ci, err := toColumnIndexNonRowid(schema, columns)
	if err != nil {
		return err
	}

	tab, err := db.NonRowidTable(schema.Table)
	if err != nil {
		return err
	}

	ind, err := db.Index(index.Index)
	if err != nil {
		return err
	}

	cols := pkColumns(schema, index)
	return ind.Scan(func(r sdb.Record) bool {
		pk := asKey(r, cols)

		var found sdb.Record
		err := tab.ScanEq(pk, func(row sdb.Record) bool { found = row; return true })
		if err != nil || found == nil {
			// should never be nil
			return false
		}
		cb(toRow(0, ci, found))
		return false
	})
}

func indexedSelectEqNonRowid(
	db *sdb.Database,
	schema *sdb.Schema,
	index *sdb.SchemaIndex,
	key sdb.Key,
	cb RowCB,
	columns []string,
) error {
	ci, err := toColumnIndexNonRowid(schema, columns)
	if err != nil {
		return err
	}

	tab, err := db.NonRowidTable(schema.Table)
	if err != nil {
		return err
	}

	ind, err := db.Index(index.Index)
	if err != nil {
		return err
	}

	cols := pkColumns(schema, index)
	return ind.ScanEq(
		key,
		func(r sdb.Record) bool {
			pk := asKey(r, cols)

			var found sdb.Record
			err := tab.ScanEq(pk, func(row sdb.Record) bool { found = row; return true })
			if err != nil || found == nil {
				// should never be nil
				return false
			}
			cb(toRow(0, ci, found))
			return false
		},
	)
}
