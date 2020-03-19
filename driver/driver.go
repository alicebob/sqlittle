package driver

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"io"

	"github.com/alicebob/sqlittle"
	sqsql "github.com/alicebob/sqlittle/sql"
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
	dbh, err := sqlittle.Open(c.File)
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
	dbh *sqlittle.DB
	SQL string
}

var (
	_ driver.Stmt = (*Statement)(nil)
)

func (st *Statement) Close() error {
	return st.dbh.Close()
}

// Exec is not relevant and is a NOOP
func (st Statement) Exec(v []driver.Value) (driver.Result, error) {
	return driver.RowsAffected(0), nil
}

func (st Statement) Query(v []driver.Value) (driver.Rows, error) {
	stmt, err := sqsql.Parse(st.SQL)
	if err != nil {
		return nil, err
	}
	sel, ok := stmt.(sqsql.SelectStmt)
	if !ok {
		return nil, fmt.Errorf("only SELECT is supported (we got a %T)", stmt)
	}
	table := sel.Table

	// ignore SELECT columns for now
	cols, err := st.dbh.Columns(sel.Table)
	if err != nil {
		return nil, err
	}

	rows := make(chan sqlittle.Row)
	go func() {
		// fix error reporting and Close()
		defer close(rows)
		cb := func(row sqlittle.Row) {
			rows <- row
		}
		st.dbh.Select(table, cb, cols...)
	}()

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
	rows    chan sqlittle.Row
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
		dest[i] = c
	}
	return nil
}
