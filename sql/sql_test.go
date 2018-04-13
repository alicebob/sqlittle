package sql

import (
	"errors"
	"reflect"
	"testing"
)

func TestColumnIsRowid(t *testing.T) {
	for d, want := range map[ColumnDef]bool{
		ColumnDef{
			Name:          "c 1",
			Type:          "Integer",
			PrimaryKey:    true,
			PrimaryKeyDir: Asc,
		}: true,
		ColumnDef{
			Name:          "c 1",
			Type:          "Integer",
			PrimaryKey:    true,
			PrimaryKeyDir: Desc,
		}: false,
		ColumnDef{
			Name: "c 1",
			Type: "Integer",
		}: false,
		ColumnDef{
			Name:          "c 1",
			Type:          "Int",
			PrimaryKey:    true,
			PrimaryKeyDir: Asc,
		}: false,
	} {
		if have := d.IsRowid(); have != want {
			t.Errorf("%#v: have %t, want %t", d, have, want)
		}
	}
}

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
			{
				sql: `create table "table" ([table] "table")`,
				want: CreateTableStmt{
					Table: "table",
					Columns: []ColumnDef{
						{Name: "table", Type: "table", Null: true},
					},
				},
			},
			{
				// You can sort-of use a string literal as a name
				sql: `create table "table" ([table] 'table')`,
				err: errors.New("syntax error"),
			},
			{
				// WITHOUT ROWID
				sql: "CREATE TABLE foo (name NOT NULL PRIMARY KEY) WITHOUT ROWID",
				want: CreateTableStmt{
					Table: "foo",
					Columns: []ColumnDef{
						{Name: "name", Type: "", PrimaryKey: true},
					},
					WithoutRowid: true,
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
		"age int not null not null null":       &ColumnDef{Name: "age", Type: "int", Null: true},
		"i0 integer primary key":               &ColumnDef{Name: "i0", Type: "integer", PrimaryKey: true, PrimaryKeyDir: Asc, Null: true},
		"i0 integer primary key desc":          &ColumnDef{Name: "i0", Type: "integer", PrimaryKey: true, PrimaryKeyDir: Desc, Null: true},
		"i0 integer primary key autoincrement": &ColumnDef{Name: "i0", Type: "integer", PrimaryKey: true, PrimaryKeyDir: Asc, AutoIncrement: true, Null: true},
		"i0 integer autoincrement":             nil,
		"i1 not null unique":                   &ColumnDef{Name: "i1", Unique: true},
		"i1 unique not null":                   &ColumnDef{Name: "i1", Unique: true},
		"i0 NOT NULL primary key":              &ColumnDef{Name: "i0", Type: "", PrimaryKey: true, PrimaryKeyDir: Asc},
		"i0 not null collate rtrim":            &ColumnDef{Name: "i0", Collate: "rtrim"},
		"i0 not null collate rtrim rtrim":      nil,
		"i0 not null default 1":                &ColumnDef{Name: "i0", Default: int64(1)},
		"i0 not null default foo":              &ColumnDef{Name: "i0", Default: "foo"},
		"i0 not null default 'foo'":            &ColumnDef{Name: "i0", Default: "foo"},
		"i0 not null default [foo]":            nil,
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
