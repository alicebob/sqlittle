package db

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alicebob/sqlittle/sql"
	"github.com/andreyvit/diff"
	"github.com/davecgh/go-spew/spew"
)

func init() {
	spew.Config.DisablePointerAddresses = true
	spew.Config.DisableCapacities = true
	spew.Config.SortKeys = true
}

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
	want *Schema,
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
		&Schema{
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
		&Schema{
			Table:      "foo",
			PrimaryKey: "sqlite_autoindex_foo_1",
			Columns: []TableColumn{
				{Column: "a", Null: true},
			},
			Indexes: []SchemaIndex{
				{
					Index:   "sqlite_autoindex_foo_1",
					Columns: []IndexColumn{{Column: "a"}},
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
		&Schema{
			Table:      "foo3",
			PrimaryKey: "sqlite_autoindex_foo3_2",
			Columns: []TableColumn{
				{Column: "a", Null: true},
				{Column: "b", Null: true},
				{Column: "c", Null: true},
			},
			Indexes: []SchemaIndex{
				{
					Index:   "sqlite_autoindex_foo3_1",
					Columns: []IndexColumn{{Column: "a"}},
				},
				{
					Index:   "sqlite_autoindex_foo3_2",
					Columns: []IndexColumn{{Column: "b"}},
				},
				{
					Index:   "sqlite_autoindex_foo3_3",
					Columns: []IndexColumn{{Column: "c"}},
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
		&Schema{
			Table: "foo",
			Columns: []TableColumn{
				{Column: "a", Type: "integer", Null: true, Rowid: true},
				{Column: "b", Null: true},
			},
			RowidPK: true,
			Indexes: []SchemaIndex{
				{
					Index:   "sqlite_autoindex_foo_1",
					Columns: []IndexColumn{{Column: "b"}},
				},
				{
					Index: "sqlite_autoindex_foo_2",
					Columns: []IndexColumn{
						{Column: "b", Collate: "rtrim", SortOrder: sql.Desc},
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
		&Schema{
			Table:   "foo",
			RowidPK: true,
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
		&Schema{
			Table:        "foo4",
			WithoutRowid: true,
			Columns: []TableColumn{
				{Column: "a", Type: "varchar"},
				{Column: "b", Type: "", Null: true},
			},
			Indexes: []SchemaIndex{
				{
					Index:   "sqlite_autoindex_foo4_2", // _1 is reserved
					Columns: []IndexColumn{{Column: "b"}},
				},
			},
			PK: []IndexColumn{{Column: "a"}},
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
		&Schema{
			Table:        "foo7",
			WithoutRowid: true,
			Columns: []TableColumn{
				{Column: "a", Null: false},
			},
			PK: []IndexColumn{{Column: "a"}},
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
		&Schema{
			Table:        "foo6",
			WithoutRowid: true,
			Columns: []TableColumn{
				{Column: "a", Type: "integer", Null: false}, // forced not null by w/o rowid PK
			},
			PK: []IndexColumn{{Column: "a"}},
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
		&Schema{
			Table: "foo9",
			Columns: []TableColumn{
				{Column: "a", Null: true},
				{Column: "b", Null: true},
				{Column: "c", Null: true},
			},
			Indexes: []SchemaIndex{
				{
					Index:   "sqlite_autoindex_foo9_1",
					Columns: []IndexColumn{{Column: "c"}, {Column: "b"}},
				},
				{
					Index:   "fooi",
					Columns: []IndexColumn{{Column: "c"}, {Column: "b"}},
				},
				{
					Index:   "fooi2",
					Columns: []IndexColumn{{Column: "c"}, {Column: "b"}},
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
		&Schema{
			Table:        "foo9",
			WithoutRowid: true,
			Columns: []TableColumn{
				{Column: "a", Null: true},
				{Column: "b", Null: false},
				{Column: "c", Null: false},
			},
			Indexes: []SchemaIndex{
				{
					Index:   "fooi",
					Columns: []IndexColumn{{Column: "b"}, {Column: "c"}},
				},
				{
					Index:   "fooj",
					Columns: []IndexColumn{{Column: "a"}, {Column: "c"}},
				},
			},
			PK: []IndexColumn{{Column: "c"}, {Column: "b"}},
		},
		nil,
	)
}

func TestSchemaUnique2(t *testing.T) {
	// ignore duplicate uniques
	testSchema(
		t,
		"foo",
		[]sqliteMaster{
			{"table", "foo", "foo", 42, `CREATE TABLE foo (a,b,c, unique(a,c,b), unique(a,c), unique(a,c,b), primary key(a,c,b)) without rowid`},
		},
		&Schema{
			Table:        "foo",
			WithoutRowid: true,
			Columns: []TableColumn{
				{Column: "a", Null: false},
				{Column: "b", Null: false},
				{Column: "c", Null: false},
			},
			Indexes: []SchemaIndex{
				{
					Index:   "sqlite_autoindex_foo_2",
					Columns: []IndexColumn{{Column: "a"}, {Column: "c"}},
				},
			},
			PK: []IndexColumn{{Column: "a"}, {Column: "c"}, {Column: "b"}},
		},
		nil,
	)
}
