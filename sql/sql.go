package sql

type PrimaryKey int

const (
	PKNone PrimaryKey = iota
	PKAsc
	PKDesc
)

type SortOrder int

const (
	Asc SortOrder = iota
	Desc
)

// A `SELECT` statement
type SelectStmt struct {
	Columns []string
	Table   string
}

// A `CREATE TABLE` statement
type CreateTableStmt struct {
	Table   string
	Columns []ColumnDef
}

// Definition of a column, as found by CreateTableStmt
type ColumnDef struct {
	Name          string
	Type          string
	PrimaryKey    PrimaryKey
	AutoIncrement bool
	Null          bool
	Unique        bool
	// Check
	// Default
	// Collate
	// foreign key
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
	l := &Lexer{tokens: ts}
	yyParse(l)
	return l.result, l.err
}
