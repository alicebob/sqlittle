package sql

type SelectStmt struct {
	Columns []string
	Table   string
}

type ColumnDef struct {
	Name string
	Type string
	Null bool
}
type CreateTableStmt struct {
	Table   string
	Columns []ColumnDef
}
