package sql

import (
	"strings"
)

type SortOrder int

const (
	Asc SortOrder = iota
	Desc
)

func (so SortOrder) String() string {
	switch so {
	case Asc:
		return "ASC"
	case Desc:
		return "DESC"
	default:
		return "???"
	}
}

// A `SELECT` statement
type SelectStmt struct {
	Columns []string
	Table   string
}

// A `CREATE TABLE` statement
type CreateTableStmt struct {
	Table        string
	Columns      []ColumnDef
	WithoutRowid bool
}

// Definition of a column, as found by CreateTableStmt
type ColumnDef struct {
	Name          string
	Type          string
	PrimaryKey    bool
	PrimaryKeyDir SortOrder
	AutoIncrement bool
	Null          bool
	Unique        bool
	// Check
	// Default
	// Collate
	// foreign key
}

// Column constraints
type primaryKey SortOrder
type unique bool
type null bool
type autoincrement bool

// make a ColumnDef with a list of constraints
func makeDef(name string, typ string, cs []interface{}) ColumnDef {
	cd := ColumnDef{
		Name: name,
		Type: typ,
		Null: true,
	}
	for _, c := range cs {
		switch v := c.(type) {
		case null:
			cd.Null = bool(v)
		case primaryKey:
			cd.PrimaryKey = true
			cd.PrimaryKeyDir = SortOrder(v)
		case unique:
			cd.Unique = bool(v)
		case autoincrement:
			cd.AutoIncrement = bool(v)
		default:
			panic("unhandled constraint")
		}
	}
	return cd
}

// The column is an alias for the rowid, and not stored in a row.
// https://sqlite.org/lang_createtable.html#rowid
func (c ColumnDef) IsRowid() bool {
	// supported:
	// CREATE TABLE t(x INTEGER PRIMARY KEY ASC, y, z);
	// TODO:
	// CREATE TABLE t(x INTEGER, y, z, PRIMARY KEY(x ASC));
	// CREATE TABLE t(x INTEGER, y, z, PRIMARY KEY(x DESC));
	return c.PrimaryKey && c.PrimaryKeyDir == Asc && strings.ToUpper(c.Type) == "INTEGER"
}

// A `CREATE INDEX` statement
type CreateIndexStmt struct {
	Index          string
	Table          string
	Unique         bool
	IndexedColumns []IndexDef
	// Where
}

// Indexed column, for CreateIndexStmt
type IndexDef struct {
	Column    string
	SortOrder SortOrder
	// Collate
}

// Parse is the main function. It will return either an error or a *Stmt
// struct.
func Parse(sql string) (interface{}, error) {
	ts, err := tokenize(sql)
	if err != nil {
		return nil, err
	}
	l := &lexer{tokens: ts}
	yyParse(l)
	return l.result, l.err
}
