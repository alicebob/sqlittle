package sql

import (
	"errors"
	"reflect"
	"testing"

	"github.com/andreyvit/diff"
	"github.com/davecgh/go-spew/spew"
)

func init() {
	spew.Config.DisablePointerAddresses = true
	spew.Config.DisableCapacities = true
	spew.Config.SortKeys = true
}

func TestExpr(t *testing.T) {
	test := func(src Expression, want string) {
		t.Helper()
		s := AsString(src)
		if have := s; have != want {
			t.Errorf("have %q, want %q", have, want)
		}
		// try to parse our serialized expression
		sqlOK(t,
			"CREATE INDEX foo_index ON foo (name) WHERE "+s,
			CreateIndexStmt{
				Index: "foo_index",
				Table: "foo",
				IndexedColumns: []IndexedColumn{
					{Column: "name"},
				},
				Where: src,
			},
		)
	}

	test(Expression(int64(123)), "123")
	test(Expression("123"), "'123'")
	test(Expression(12.3), "12.3")
	test(Expression(ExColumn("foo")), `"foo"`)
	test(Expression(ExFunction{"foo", []Expression{int64(123)}}), `"foo"(123)`)
	test(Expression(ExBinaryOp{"+", int64(1), int64(2)}), "1+2")
	test(Expression(ExBinaryOp{">", int64(1), int64(2)}), "1>2")
	test(Expression(ExFunction{"foo", []Expression{ExBinaryOp{"*", int64(1), int64(2)}}}), `"foo"(1*2)`)
}

func sqlOK(t *testing.T, sql string, want interface{}) {
	t.Helper()
	stmt, err := Parse(sql)
	if err != nil {
		t.Error(err)
		return
	}
	if have := stmt; !reflect.DeepEqual(have, want) {
		t.Errorf("diff:\n%s", diff.LineDiff(spew.Sdump(want), spew.Sdump(have)))
	}
}

func sqlError(t *testing.T, sql string, want error) {
	t.Helper()
	_, err := Parse(sql)
	if have := err; !reflect.DeepEqual(have, want) {
		t.Errorf("have %#v, want %#v", have, want)
	}
}

func TestSQL(t *testing.T) {
	sqlError(t, "INSERT INTO FOO", errors.New("syntax error"))
}

func TestSelect(t *testing.T) {
	sqlOK(t,
		"SELECT aap,noot, mies FROM foo2",
		SelectStmt{Columns: []string{"aap", "noot", "mies"}, Table: "foo2"},
	)

	// create what?
	sqlError(t, "CREATE nothing foo", errors.New("syntax error"))
}

func TestCreateTable(t *testing.T) {
	sqlError(t,
		"CREATE TABLE foo",
		errors.New("syntax error"),
	)
	sqlOK(t,
		"CREATE table animals (name varchar not null, age int)",
		CreateTableStmt{
			Table: "animals",
			Columns: []ColumnDef{
				{Name: "name", Type: "varchar"}, {Name: "age", Type: "int", Null: true},
			},
		},
	)
	sqlOK(t,
		`create table "table" ([table] "table")`,
		CreateTableStmt{
			Table: "table",
			Columns: []ColumnDef{
				{Name: "table", Type: "table", Null: true},
			},
		},
	)
	// (you can sort-of use a string literal as a name in sqlite)
	sqlError(t,
		`create table "table" ([table] 'table')`,
		errors.New("syntax error"),
	)

	// WITHOUT ROWID
	sqlOK(t,
		"CREATE TABLE foo (name NOT NULL PRIMARY KEY) WITHOUT ROWID",
		CreateTableStmt{
			Table: "foo",
			Columns: []ColumnDef{
				{Name: "name", Type: "", PrimaryKey: true},
			},
			WithoutRowid: true,
		},
	)

	sqlOK(t,
		"CREATE TABLE aap (noot, mies, PRIMARY KEY (noot, mies DESC))",
		CreateTableStmt{
			Table: "aap",
			Columns: []ColumnDef{
				{Name: "noot", Null: true},
				{Name: "mies", Null: true},
			},
			Constraints: []TableConstraint{
				TablePrimaryKey{
					IndexedColumns: []IndexedColumn{
						{Column: "noot", SortOrder: Asc},
						{Column: "mies", SortOrder: Desc},
					},
				},
			},
		},
	)

	// constraint names (ignored)
	sqlOK(t,
		"CREATE TABLE aap (noot, mies, CONSTRAINT foo PRIMARY KEY (noot), CONSTRAINT bar UNIQUE (mies))",
		CreateTableStmt{
			Table: "aap",
			Columns: []ColumnDef{
				{Name: "noot", Null: true},
				{Name: "mies", Null: true},
			},
			Constraints: []TableConstraint{
				TablePrimaryKey{IndexedColumns: []IndexedColumn{
					{Column: "noot"},
				}},
				TableUnique{IndexedColumns: []IndexedColumn{
					{Column: "mies"},
				}},
			},
		},
	)

	sqlOK(t,
		"CREATE TABLE aap (noot, mies, UNIQUE (mies DESC), UNIQUE (noot))",
		CreateTableStmt{
			Table: "aap",
			Columns: []ColumnDef{
				{Name: "noot", Null: true},
				{Name: "mies", Null: true},
			},
			Constraints: []TableConstraint{
				TableUnique{
					IndexedColumns: []IndexedColumn{
						{Column: "mies", SortOrder: Desc},
					},
				},
				TableUnique{
					IndexedColumns: []IndexedColumn{
						{Column: "noot", SortOrder: Asc},
					},
				},
			},
		},
	)

	sqlOK(t,
		"CREATE TABLE aap (noot, FOREIGN KEY (noot) REFERENCES something (fnoot) ON DELETE NO ACTION ON UPDATE CASCADE ON UPDATE RESTRICT ON DELETE SET NULL ON UPDATE SET DEFAULT)",
		CreateTableStmt{
			Table: "aap",
			Columns: []ColumnDef{
				{Name: "noot", Null: true},
			},
			Constraints: []TableConstraint{
				TableForeignKey{
					Columns: []string{"noot"},
					Clause: ForeignKeyClause{
						ForeignTable:   "something",
						ForeignColumns: []string{"fnoot"},
						Triggers: []Trigger{
							TriggerOnDelete(ActionNoAction),
							TriggerOnUpdate(ActionCascade),
							TriggerOnUpdate(ActionRestrict),
							TriggerOnDelete(ActionSetNull),
							TriggerOnUpdate(ActionSetDefault),
						},
					},
				},
			},
		},
	)

	sqlOK(t,
		"CREATE TABLE aap (noot DEFAULT -12)",
		CreateTableStmt{
			Table: "aap",
			Columns: []ColumnDef{
				{Name: "noot", Null: true, Default: int64(-12)},
			},
		},
	)

	sqlOK(t,
		"create table foo3 (a unique, b PRIMARY KEY, c, unique(a), unique(c))",
		CreateTableStmt{
			Table: "foo3",
			Columns: []ColumnDef{
				{Name: "a", Null: true, Unique: true},
				{Name: "b", Null: true, PrimaryKey: true},
				{Name: "c", Null: true},
			},
			Constraints: []TableConstraint{
				TableUnique{
					IndexedColumns: []IndexedColumn{
						{Column: "a", SortOrder: Asc},
					},
				},
				TableUnique{
					IndexedColumns: []IndexedColumn{
						{Column: "c", SortOrder: Asc},
					},
				},
			},
		},
	)

	sqlOK(t,
		"CREATE TABLE table1 ( data INT, UNIQUE (data) ON CONFLICT REPLACE )",
		CreateTableStmt{
			Table: "table1",
			Columns: []ColumnDef{
				{Name: "data", Null: true, Type: "INT"},
			},
			Constraints: []TableConstraint{
				TableUnique{
					IndexedColumns: []IndexedColumn{
						{Column: "data", SortOrder: Asc},
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
		"i0 STRING DEFAULT NULL":               &ColumnDef{Name: "i0", Type: "STRING", Null: true, Default: nil},

		"integer integer primary key":                      &ColumnDef{Name: "integer", Type: "integer", PrimaryKey: true, PrimaryKeyDir: Asc, Null: true},
		"ROWID INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE":   &ColumnDef{Name: "ROWID", Type: "INTEGER", PrimaryKey: true, Unique: true, AutoIncrement: true, Null: true},
		"REPLACE INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE": &ColumnDef{Name: "REPLACE", Type: "INTEGER", PrimaryKey: true, Unique: true, AutoIncrement: true, Null: true},
		"select integer primary key":                       nil,

		"message_id INTEGER REFERENCES message (foo)": &ColumnDef{
			Name: "message_id",
			Type: "INTEGER",
			Null: true,
			References: &ForeignKeyClause{
				ForeignTable:   "message",
				ForeignColumns: []string{"foo"},
				// Triggers:
			},
		},
		"message_id INTEGER REFERENCES message (foo) ON DELETE CASCADE": &ColumnDef{
			Name: "message_id",
			Type: "INTEGER",
			Null: true,
			References: &ForeignKeyClause{
				ForeignTable:   "message",
				ForeignColumns: []string{"foo"},
				Triggers:       []Trigger{TriggerOnDelete(ActionCascade)},
			},
		},
		"message_id INTEGER REFERENCES message (ROWID) ON DELETE CASCADE": &ColumnDef{
			Name: "message_id",
			Type: "INTEGER",
			Null: true,
			References: &ForeignKeyClause{
				ForeignTable:   "message",
				ForeignColumns: []string{"ROWID"},
				Triggers:       []Trigger{TriggerOnDelete(ActionCascade)},
			},
		},

		"foo integer CHECK ( 1 > 2 ) NOT NULL": &ColumnDef{
			Name: "foo",
			Type: "integer",
			Checks: []Expression{
				ExBinaryOp{Op: ">", Left: int64(1), Right: int64(2)},
			},
		},
		"foo integer CHECK (foo = 4) NOT NULL": &ColumnDef{
			Name: "foo",
			Type: "integer",
			Checks: []Expression{
				ExBinaryOp{Op: "=", Left: ExColumn("foo"), Right: int64(4)},
			},
		},
		"foo integer CHECK 12 NOT NULL": nil,
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
	sqlError(t,
		"CREATE INDEX foo",
		errors.New("syntax error"),
	)

	sqlOK(t,
		"CREATE INDEX foo_index ON foo (name DESC, age)",
		CreateIndexStmt{
			Index: "foo_index",
			Table: "foo",
			IndexedColumns: []IndexedColumn{
				{Column: "name", SortOrder: Desc},
				{Column: "age", SortOrder: Asc},
			},
		},
	)

	sqlOK(t,
		"CREATE UNIQUE INDEX foo_index ON foo (name)",
		CreateIndexStmt{
			Index:  "foo_index",
			Table:  "foo",
			Unique: true,
			IndexedColumns: []IndexedColumn{
				{Column: "name", SortOrder: Asc},
			},
		},
	)

	sqlOK(t,
		"CREATE UNIQUE INDEX foo_index ON foo (name COLLATE RTRIM DESC)",
		CreateIndexStmt{
			Index:  "foo_index",
			Table:  "foo",
			Unique: true,
			IndexedColumns: []IndexedColumn{
				{Column: "name", Collate: "RTRIM", SortOrder: Desc},
			},
		},
	)

	// WHERE expressions
	for sql, where := range map[string]interface{}{
		"1":                  int64(1),
		"- 1":                int64(-1),
		"- - 1":              int64(1),
		"2>3":                ExBinaryOp{">", int64(2), int64(3)},
		"2>=3":               ExBinaryOp{">=", int64(2), int64(3)},
		"(2>3)":              ExBinaryOp{">", int64(2), int64(3)},
		"1>(2>3)":            ExBinaryOp{">", int64(1), ExBinaryOp{">", int64(2), int64(3)}},
		"3.14":               float64(3.14),
		"+3":                 int64(3),
		"-3.14":              -3.14,
		"2+3":                ExBinaryOp{"+", int64(2), int64(3)},
		"'abc'":              "abc",
		"foo":                ExColumn("foo"),
		"[foo]":              ExColumn("foo"),
		"NULL":               nil,
		`length('abcdef')`:   ExFunction{"length", []Expression{"abcdef"}},
		`length()`:           ExFunction{"length", nil},
		`"length"('abcdef')`: ExFunction{"length", []Expression{"abcdef"}},
		`"length"([abcdef])`: ExFunction{"length", []Expression{ExColumn("abcdef")}},
		`f(g(), 123)`: ExFunction{"f", []Expression{
			ExFunction{"g", nil},
			int64(123),
		}},
	} {
		sqlOK(t,
			"CREATE INDEX foo_index ON foo (name) WHERE "+sql,
			CreateIndexStmt{
				Index: "foo_index",
				Table: "foo",
				IndexedColumns: []IndexedColumn{
					{Column: "name"},
				},
				Where: where,
			},
		)
	}

	sqlOK(t,
		"CREATE INDEX foo_index ON foo (length(name) + 12)",
		CreateIndexStmt{
			Index: "foo_index",
			Table: "foo",
			IndexedColumns: []IndexedColumn{
				{Expression: `"length"("name")+12`},
			},
		},
	)

	sqlOK(t,
		"CREATE INDEX foo_index ON foo (name DESC, length(name) DESC)",
		CreateIndexStmt{
			Index: "foo_index",
			Table: "foo",
			IndexedColumns: []IndexedColumn{
				{Column: "name", SortOrder: Desc},
				{Expression: `"length"("name")`, SortOrder: Desc},
			},
		},
	)
}
