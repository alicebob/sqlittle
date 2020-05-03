package driver

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"sync"

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

// Exec is not relevant and always returns an error
func (st *Statement) Exec(v []driver.Value) (driver.Result, error) {
	return nil, errors.New("Exec() is not supported")
	// return driver.ResultNoRows, nil
}

func (st *Statement) Query(v []driver.Value) (driver.Rows, error) {
	stmt, err := sqsql.Parse(st.SQL)
	if err != nil {
		return nil, err
	}
	sel, ok := stmt.(sqsql.SelectStmt)
	if !ok {
		return nil, fmt.Errorf("only SELECT is supported (we got a %T)", stmt)
	}
	table := sel.Table

	cols, err := st.expandSelectColumns(sel)
	if err != nil {
		return nil, err
	}

	rows := make(chan sqlittle.Row)
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer close(rows)
		cb := func(row sqlittle.Row) bool {
			select {
			case <-ctx.Done():
				return true
			case rows <- row:
				return false
			}
		}
		// FIXME: the error is currently lost
		st.dbh.SelectDone(table, cb, cols...)
		wg.Done()
	}()

	return &Rows{
		columns: cols,
		rows:    rows,
		wg:      wg,
		cancel:  cancel,
	}, nil
}

func (st Statement) NumInput() int {
	return 0
}

// expand the '*' from SELECT statements. Doesn't check column names.
func (st *Statement) expandSelectColumns(sel sqsql.SelectStmt) ([]string, error) {
	allCols, err := st.dbh.Columns(sel.Table)
	if err != nil {
		return nil, err
	}

	var cols []string
	for _, c := range sel.Columns {
		if c == "*" {
			cols = append(cols, allCols...)
			continue
		}
		cols = append(cols, c)
	}
	return cols, nil
}

// Rows is the result set. It implements the driver.Rows interface.
type Rows struct {
	columns []string
	rows    chan sqlittle.Row
	wg      *sync.WaitGroup
	cancel  func()
}

func (r *Rows) Close() error {
	r.cancel()
	r.wg.Wait()
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
