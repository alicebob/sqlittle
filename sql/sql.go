package sql

type PrimaryKey int

const (
	PKNone PrimaryKey = iota
	PKAsc
	PKDesc
)

// Type of a column, as found by CreateTableStmt
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
