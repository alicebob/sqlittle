package driver

import (
	"database/sql"
	"database/sql/driver"
	"io"
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
		File: "./test.sqlite",
	}, nil
}

// Connection is a single sqlite file
// It implements the driver.Conn and driver.Tx interfaces
type Connection struct {
	File string
}

func (c *Connection) Begin() (driver.Tx, error) {
	return c, nil
}

func (c *Connection) Rollback() error {
	return nil
}

func (c *Connection) Commit() error {
	return nil
}

func (c *Connection) Close() error {
	return nil
}

func (c *Connection) Prepare(q string) (driver.Stmt, error) {
	return Statement{
		SQL: q,
	}, nil
}

// Statement is a single statement, belonging to a particular Connection.
// It implements the driver.Stmt interface.
type Statement struct {
	SQL string
}

func (Statement) Close() error {
	return nil
}

// Exec is not relevant and is a NOOP
func (st Statement) Exec(v []driver.Value) (driver.Result, error) {
	return driver.RowsAffected(0), nil
}

func (st Statement) Query(v []driver.Value) (driver.Rows, error) {
	return &Rows{
		columns: []string{"test", "columns"},
		rows: [][]string{
			{"aap", "noot"},
			{"mies", "wim"},
			{"vuur", "eekhoorn"},
		},
	}, nil
}

func (st Statement) NumInput() int {
	return 0
}

// Rows is the result set. It implements the driver.Rows interface.
type Rows struct {
	columns []string
	rows    [][]string
}

func (*Rows) Close() error {
	return nil
}

func (r *Rows) Columns() []string {
	return r.columns
}

func (r *Rows) Next(dest []driver.Value) error {
	if len(r.rows) == 0 {
		return io.EOF
	}
	row := r.rows[0]
	r.rows = r.rows[1:]

	for i, c := range row {
		dest[i] = c
	}
	return nil
}
