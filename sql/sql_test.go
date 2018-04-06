package sql

import (
	"errors"
	"reflect"
	"testing"
)

type sqlCase struct {
	sql  string
	want interface{}
	err  error
}

func testSQL(t *testing.T, cases []sqlCase) {
	t.Helper()
	for n, c := range cases {
		stmt, err := Parse(c.sql)
		if have, want := err, c.err; !reflect.DeepEqual(have, want) {
			t.Errorf("case %d: have %#v, want %#v", n, have, want)
			continue
		}
		if c.err != nil {
			continue
		}
		if have, want := stmt, c.want; !reflect.DeepEqual(have, want) {
			t.Errorf("case %d: have %#v, want %#v", n, have, want)
		}
	}
}

func TestSQL(t *testing.T) {
	testSQL(
		t,
		[]sqlCase{
			// unknown
			{
				sql: "INSERT INTO FOO",
				err: errors.New("syntax error"),
			},
		},
	)
}

func TestSelect(t *testing.T) {
	testSQL(
		t,
		[]sqlCase{
			// select
			{
				sql:  "SELECT * FROM foo",
				want: SelectStmt{Columns: []string{"*"}, Table: "foo"},
			},
			{
				sql:  "SELECT aap,noot, mies FROM foo2",
				want: SelectStmt{Columns: []string{"aap", "noot", "mies"}, Table: "foo2"},
			},

			// create what?
			{
				sql: "CREATE nothing foo",
				err: errors.New("syntax error"),
			},
		},
	)
}

func TestCreateTable(t *testing.T) {
	testSQL(
		t,
		[]sqlCase{
			{
				sql: "CREATE TABLE foo",
				err: errors.New("syntax error"),
			},
			{
				sql: "CREATE table animals (name varchar not null, age int)",
				want: CreateTableStmt{
					Table: "animals",
					Columns: []ColumnDef{
						{Name: "name", Type: "varchar"}, {Name: "age", Type: "int", Null: true},
					},
				},
			},
		},
	)

	// CREATE TABLE column definition tests.
	// a nil value means we expect an error
	cases := map[string]*ColumnDef{
		"age": &ColumnDef{Name: "age", Null: true},

		// column types
		"age int":        &ColumnDef{Name: "age", Type: "int", Null: true},
		"age integer":    &ColumnDef{Name: "age", Type: "integer", Null: true},
		"age int(1)":     &ColumnDef{Name: "age", Type: "int", Null: true},
		"age int(1,2)":   &ColumnDef{Name: "age", Type: "int", Null: true},
		"age int(1,2,3)": nil,
		"age foo":        &ColumnDef{Name: "age", Type: "foo", Null: true},

		// constraints
		"age int not null":                     &ColumnDef{Name: "age", Type: "int"},
		"i0 integer primary key":               &ColumnDef{Name: "i0", Type: "integer", PrimaryKey: PKAsc, Null: true},
		"i0 integer primary key desc":          &ColumnDef{Name: "i0", Type: "integer", PrimaryKey: PKDesc, Null: true},
		"i0 integer primary key autoincrement": &ColumnDef{Name: "i0", Type: "integer", PrimaryKey: PKAsc, AutoIncrement: true, Null: true},
		"i1 not null unique":                   &ColumnDef{Name: "i1", Unique: true},
	}

	for sql, col := range cases {
		stmt, err := Parse("CREATE TABLE foo (" + sql + ")")
		var werr error
		if col == nil {
			werr = errors.New("syntax error")
		}
		if have, want := err, werr; !reflect.DeepEqual(have, want) {
			t.Errorf("case %q: have %#v, want %#v", sql, have, want)
			continue
		}
		if werr != nil {
			continue
		}
		create, ok := stmt.(CreateTableStmt)
		if !ok {
			t.Errorf("case %q: have %t, want CreateTableStmt", sql, stmt)
		}
		if have, want := len(create.Columns), 1; have != want {
			t.Errorf("case %q: have %#v, want %#v", sql, have, want)
		}
		if have, want := &create.Columns[0], col; !reflect.DeepEqual(have, want) {
			t.Errorf("case %q: have %#v, want %#v", sql, have, want)
		}
	}
}

func TestCreateIndex(t *testing.T) {
	testSQL(
		t,
		[]sqlCase{
			{
				sql: "CREATE INDEX foo",
				err: errors.New("syntax error"),
			},
			{
				sql: "CREATE INDEX foo_index ON foo (name DESC, age)",
				want: CreateIndexStmt{
					Index: "foo_index",
					Table: "foo",
					IndexedColumns: []IndexDef{
						{Column: "name", SortOrder: Desc},
						{Column: "age", SortOrder: Asc},
					},
				},
			},
			{
				sql: "CREATE UNIQUE INDEX foo_index ON foo (name)",
				want: CreateIndexStmt{
					Index:  "foo_index",
					Table:  "foo",
					Unique: true,
					IndexedColumns: []IndexDef{
						{Column: "name", SortOrder: Asc},
					},
				},
			},
		},
	)
}
