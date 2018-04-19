package sqlittle

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alicebob/sqlittle/sql"
	"github.com/andreyvit/diff"
	"github.com/davecgh/go-spew/spew"
)

func TestIsRowid(t *testing.T) {
	for i, c := range [][2]bool{
		[...]bool{isRowid(true, "Integer", sql.Asc), true},
		[...]bool{isRowid(true, "Int", sql.Asc), false},
		[...]bool{isRowid(true, "Integer", sql.Desc), true},
		[...]bool{isRowid(false, "Integer", sql.Desc), false},
		[...]bool{isRowid(false, "Integer", sql.Asc), true},
	} {
		if have, want := c[0], c[1]; have != want {
			t.Errorf("%#v: have %t, want %t", i, have, want)
		}
	}
}

func testSchema(
	t *testing.T,
	table string,
	master []sqliteMaster,
	want *SchemaTable,
	wantErr error,
) {
	t.Helper()
	have, err := newSchema(table, master)
	if !reflect.DeepEqual(err, wantErr) {
		t.Errorf("%s error: have %#v, want %#v", table, err, wantErr)
		return
	}
	if wantErr != nil {
		return
	}
	if !reflect.DeepEqual(have, want) {
		spew.Config.DisablePointerAddresses = true
		spew.Config.DisableCapacities = true
		spew.Config.SortKeys = true
		t.Errorf("%s: diff:\n%s", table, diff.LineDiff(spew.Sdump(want), spew.Sdump(have)))
	}
}

func TestSchemaNosuch(t *testing.T) {
	testSchema(
		t,
		"bar",
		[]sqliteMaster{
			{"table", "foo", "foo", 42, "CREATE TABLE foo (a, b not null)"},
		},
		nil,
		errors.New(`no such table: "bar"`),
	)
}

func TestSchemaSimple(t *testing.T) {
	testSchema(
		t,
		"foo",
		[]sqliteMaster{
			{"table", "foo", "foo", 42, "CREATE TABLE foo (a, b not null)"},
		},
		&SchemaTable{
			Table: "foo",
			Columns: []TableColumn{
				{Column: "a", Null: true},
				{Column: "b"},
			},
		},
		nil,
	)
}

func TestSchemaConstrPK(t *testing.T) {
	testSchema(
		t,
		"foo",
		[]sqliteMaster{
			{"table", "foo", "foo", 42, "CREATE TABLE foo (a, PRIMARY KEY(a))"},
			{"index", "sqlite_autoindex_foo_1", "foo", 42, ""},
		},
		&SchemaTable{
			Table: "foo",
			Columns: []TableColumn{
				{Column: "a", Null: true},
			},
			Indexes: []SchemaIndex{
				{
					Index: "sqlite_autoindex_foo_1",
					Columns: []IndexColumn{
						{
							Column: "a",
						},
						{
							Rowid: true,
						},
					},
				},
			},
		},
		nil,
	)
}

func TestSchemaUnique(t *testing.T) {
	testSchema(
		t,
		"foo3",
		[]sqliteMaster{
			{"table", "foo3", "foo3", 42, "create table foo3 (a unique, b PRIMARY KEY, c, unique(a), unique(c))"},
			{"index", "sqlite_autoindex_foo_1", "foo3", 42, ""},
		},
		&SchemaTable{
			Table: "foo3",
			Columns: []TableColumn{
				{Column: "a", Null: true},
				{Column: "b", Null: true},
				{Column: "c", Null: true},
			},
			Indexes: []SchemaIndex{
				{
					Index:   "sqlite_autoindex_foo3_1",
					Columns: []IndexColumn{{Column: "a"}, {Rowid: true}},
				},
				{
					Index:   "sqlite_autoindex_foo3_2",
					Columns: []IndexColumn{{Column: "b"}, {Rowid: true}},
				},
				{
					Index:   "sqlite_autoindex_foo3_3",
					Columns: []IndexColumn{{Column: "c"}, {Rowid: true}},
				},
			},
		},
		nil,
	)
}

func TestSchemaRowid(t *testing.T) {
	testSchema(
		t,
		"foo",
		[]sqliteMaster{
			{"table", "foo", "foo", 42, `create table foo(a integer primary key, b unique, unique(b collate 'rtrim' desc))`},
			{"index", "sqlite_autoindex_foo_1", "foo", 42, ""},
			{"index", "sqlite_autoindex_foo_2", "foo", 42, ""},
		},
		&SchemaTable{
			Table: "foo",
			Columns: []TableColumn{
				{Column: "a", Type: "integer", Null: true, Rowid: true},
				{Column: "b", Null: true},
			},
			Indexes: []SchemaIndex{
				{
					Index:   "sqlite_autoindex_foo_1",
					Columns: []IndexColumn{{Column: "b"}, {Rowid: true}},
				},
				{
					Index: "sqlite_autoindex_foo_2",
					Columns: []IndexColumn{
						{Column: "b", Collate: "rtrim", SortOrder: sql.Desc},
						{Rowid: true},
					},
				},
			},
		},
		nil,
	)
}

func TestSchemaRowid2(t *testing.T) {
	testSchema(
		t,
		"foo",
		[]sqliteMaster{
			{"table", "foo", "foo", 42, `create table foo(a integer, primary key(a desc))`},
		},
		&SchemaTable{
			Table: "foo",
			Columns: []TableColumn{
				{Column: "a", Type: "integer", Null: true, Rowid: true},
			},
		},
		nil,
	)
}

func TestSchemaWithoutRowid(t *testing.T) {
	testSchema(
		t,
		"foo4",
		[]sqliteMaster{
			{"table", "foo4", "foo4", 42, `create table foo4(a varchar, b, primary key(a), unique(b)) without rowid`},
		},
		&SchemaTable{
			Table:        "foo4",
			WithoutRowid: true,
			Columns: []TableColumn{
				{Column: "a", Type: "varchar"},
				{Column: "b", Type: "", Null: true},
			},
			Indexes: []SchemaIndex{
				{
					// "without rowid" primary key
					Index:     "",
					Columns:   []IndexColumn{{Column: "a"}, {Column: "b"}},
					PKColumns: []int{0},
				},
				{
					Index:     "sqlite_autoindex_foo4_2", // _1 is reserved
					Columns:   []IndexColumn{{Column: "b"}, {Column: "a"}},
					PKColumns: []int{1},
				},
			},
		},
		nil,
	)
}

func TestSchemaWithoutRowid2(t *testing.T) {
	// w/o rowid table: first a unique, then a primary key
	testSchema(
		t,
		"foo7",
		[]sqliteMaster{
			{"table", "foo7", "foo7", 42, `CREATE TABLE foo7(a, unique(a), primary key(a)) without rowid`},
		},
		&SchemaTable{
			Table:        "foo7",
			WithoutRowid: true,
			Columns: []TableColumn{
				{Column: "a", Null: false},
			},
			Indexes: []SchemaIndex{
				{
					Index:     "",
					Columns:   []IndexColumn{{Column: "a"}},
					PKColumns: []int{0},
				},
			},
		},
		nil,
	)
}

func TestSchemaWithoutRowid3(t *testing.T) {
	testSchema(
		t,
		"foo6",
		[]sqliteMaster{
			{"table", "foo6", "foo6", 42, `create table foo6(a integer unique primary key) without rowid`},
		},
		&SchemaTable{
			Table:        "foo6",
			WithoutRowid: true,
			Columns: []TableColumn{
				{Column: "a", Type: "integer", Null: false}, // forced by w/o rowid PK
			},
			Indexes: []SchemaIndex{
				{
					Index:     "",
					Columns:   []IndexColumn{{Column: "a"}},
					PKColumns: []int{0},
				},
			},
		},
		nil,
	)
}

func TestSchemaIndex(t *testing.T) {
	testSchema(
		t,
		"foo9",
		[]sqliteMaster{
			{"table", "foo9", "foo9", 42, `CREATE TABLE foo9 (a, b, c, unique(c, b))`},
			{"index", "fooi", "foo9", 42, `CREATE INDEX fooi ON foo9 (c, b)`},
			{"index", "fooi2", "foo9", 42, `CREATE INDEX fooi2 ON foo9 (c, b)`},
		},
		&SchemaTable{
			Table: "foo9",
			Columns: []TableColumn{
				{Column: "a", Null: true},
				{Column: "b", Null: true},
				{Column: "c", Null: true},
			},
			Indexes: []SchemaIndex{
				{
					Index:   "sqlite_autoindex_foo9_1",
					Columns: []IndexColumn{{Column: "c"}, {Column: "b"}, {Rowid: true}},
				},
				{
					Index:   "fooi",
					Columns: []IndexColumn{{Column: "c"}, {Column: "b"}, {Rowid: true}},
				},
				{
					Index:   "fooi2",
					Columns: []IndexColumn{{Column: "c"}, {Column: "b"}, {Rowid: true}},
				},
			},
		},
		nil,
	)
}

func TestSchemaIndexNonRowid(t *testing.T) {
	testSchema(
		t,
		"foo9",
		[]sqliteMaster{
			{"table", "foo9", "foo9", 42, `CREATE TABLE foo9 (a, b, c, primary key (c, b)) WITHOUT ROWID`},
			{"index", "fooi", "foo9", 42, `CREATE INDEX fooi ON foo9 (b, c)`},
			{"index", "fooj", "foo9", 42, `CREATE INDEX fooj ON foo9 (a, c)`},
		},
		&SchemaTable{
			Table:        "foo9",
			WithoutRowid: true,
			Columns: []TableColumn{
				{Column: "a", Null: true},
				{Column: "b", Null: false},
				{Column: "c", Null: false},
			},
			Indexes: []SchemaIndex{
				{
					Index: "",
					// 'a' is not in the primary key, but this describes the
					// column order in the database file
					Columns:   []IndexColumn{{Column: "c"}, {Column: "b"}, {Column: "a"}},
					PKColumns: []int{0, 1},
				},
				{
					Index:     "fooi",
					Columns:   []IndexColumn{{Column: "b"}, {Column: "c"}},
					PKColumns: []int{1, 0},
				},
				{
					Index:     "fooj",
					Columns:   []IndexColumn{{Column: "a"}, {Column: "c"}, {Column: "b"}},
					PKColumns: []int{1, 2},
				},
			},
		},
		nil,
	)
}
