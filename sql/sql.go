package sql

type SelectStmt struct {
	Columns []string
	Table   string
}

type PrimaryKey int

const (
	PKNone PrimaryKey = iota
	PKAsc
	PKDesc
)

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

type CreateTableStmt struct {
	Table   string
	Columns []ColumnDef
}
