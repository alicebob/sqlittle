// all public low level methods

package sqlittle

type Table struct {
	db   *Database
	root int
}

// Scan calls cb() for every row in the table. Will be called in 'database
// order'.
// The record is given as sqlite stores it; this means:
//  - float64 columns might be stored as int64
//  - after an alter table which adds columns a row might miss the new columns
//  - an "integer primary key" column will be always be nil, and the rowid is
//  the value
// If the callback returns true (done) the scan will be stopped.
func (t *Table) Scan(cb TableScanCB) error {
	td, err := t.db.openTable(t.root)
	if err != nil {
		return err
	}
	_, err = td.Iter(
		t.db,
		func(rowid int64, pl cellPayload) (bool, error) {
			c, err := addOverflow(t.db, pl)
			if err != nil {
				return false, err
			}

			rec, err := parseRecord(c)
			if err != nil {
				return false, err
			}
			return cb(rowid, rec), nil
		},
	)
	return err
}

// Rowid finds a single row by rowid. Will return nil if it isn't found.
// The rowid is an internal id, but if you have an `integer primary key` column
// that should be the same.
// See Table.Scan comments about the Record
func (t *Table) Rowid(rowid int64) (Record, error) {
	td, err := t.db.openTable(t.root)
	if err != nil {
		return nil, err
	}

	var recPl *cellPayload
	if _, err := td.IterMin(
		t.db,
		rowid,
		func(k int64, pl cellPayload) (bool, error) {
			if k == rowid {
				recPl = &pl
			}
			return true, nil
		},
	); err != nil {
		return nil, err
	}
	if recPl == nil {
		return nil, nil
	}

	c, err := addOverflow(t.db, *recPl)
	if err != nil {
		return nil, err
	}
	return parseRecord(c)
}
