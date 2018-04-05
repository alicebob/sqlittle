package sql

type SelectStmt struct {
	Columns []string
	Table   string
}

type CreatTableStmt struct {
	Table string
}
