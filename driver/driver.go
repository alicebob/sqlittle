package driver

import (
	"database/sql"
	"database/sql/driver"
	"io"

	"github.com/alicebob/sqlittle/db"
)

func init() {
	sql.Register("sqlittle", &Driver{})
}

// Driver is the sqlittle database driver.
// It implements the driver.Driver interface
type Driver struct{}

func (d *Driver) Open(name string) (driver.Conn, error) {
	return Open(name)
}

func Open(dsn string) (driver.Conn, error) {
	return &Connection{
		File: dsn,
	}, nil
}

// Connection is a single sqlite file
// It implements the driver.Conn and driver.Tx interfaces
type Connection struct {
	File string
}

func (c *Connection) Begin() (driver.Tx, error) {
	return &Tx{}, nil
}

func (c *Connection) Close() error {
	return nil
}

func (c *Connection) Prepare(q string) (driver.Stmt, error) {
	dbh, err := db.OpenFile(c.File)
	if err != nil {
		return nil, err
	}
	return &Statement{
		dbh: dbh,
		SQL: q,
	}, nil
}

type Tx struct{}

func (*Tx) Rollback() error {
	return nil
}

func (*Tx) Commit() error {
	return nil
}

// Statement is a single statement, belonging to a particular Connection.
// It implements the driver.Stmt interface.
type Statement struct {
	dbh *db.Database
	SQL string
}

func (st *Statement) Close() error {
	return st.dbh.Close()
}

// Exec is not relevant and is a NOOP
func (st Statement) Exec(v []driver.Value) (driver.Result, error) {
	return driver.RowsAffected(0), nil
}

func (st Statement) Query(v []driver.Value) (driver.Rows, error) {
	s, err := st.dbh.Schema("tracks")
	if err != nil {
		return nil, err
	}

	var cols []string
	for _, c := range s.Columns {
		cols = append(cols, c.Column)
	}

	var rows chan db.Record
	if s.WithoutRowid {
		t, err := st.dbh.NonRowidTable(s.Table)
		if err != nil {
			return nil, err
		}
		rows = indexScan(t)
	} else {
		t, err := st.dbh.Table(s.Table)
		if err != nil {
			return nil, err
		}
		rows = tableScan(t)
	}
	return &Rows{
		columns: cols,
		rows:    rows,
	}, nil
}

func (st Statement) NumInput() int {
	return 0
}

// Rows is the result set. It implements the driver.Rows interface.
type Rows struct {
	columns []string
	rows    chan db.Record
}

func (*Rows) Close() error {
	return nil
}

func (r *Rows) Columns() []string {
	return r.columns
}

func (r *Rows) Next(dest []driver.Value) error {
	row, ok := <-r.rows
	if !ok {
		return io.EOF
	}

	for i, c := range row {
		if len(row) <= i {
			dest[i] = ""
			continue
		}
		dest[i] = c
	}
	return nil
}

// runs a table scan in a Go routine.
// Closes the channel only when the whole table has been scanned.
func tableScan(t *db.Table) chan db.Record {
	// This needs to deal with errors much nicer
	rows := make(chan db.Record)
	go func() {
		defer close(rows)
		err := t.Scan(func(rowID int64, rec db.Record) bool {
			rows <- rec
			return false
		})
		if err != nil {
			panic(err) // FIXME :)
		}
	}()
	return rows
}

// see tableScan
func indexScan(ind *db.Index) chan db.Record {
	// This needs to deal with errors much nicer
	rows := make(chan db.Record)
	go func() {
		defer close(rows)
		err := ind.Scan(func(rec db.Record) bool {
			rows <- rec
			return false
		})
		if err != nil {
			panic(err) // FIXME :)
		}
	}()
	return rows
}
