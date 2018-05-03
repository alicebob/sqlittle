package sql

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
	Constraints  []TableConstraint
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
	Default       interface{}
	Collate       string
	// Check
	// foreign key
}

// column constraints, used while parsing a constraint list
type columnConstraint interface{}
type ccPrimaryKey struct {
	sort          SortOrder
	autoincrement bool
}
type ccUnique bool
type ccNull bool
type ccAutoincrement bool
type ccCollate string
type ccDefault interface{}

func makeColumnDef(name string, typ string, cs []columnConstraint) ColumnDef {
	cd := ColumnDef{
		Name: name,
		Type: typ,
		Null: true,
	}
	for _, c := range cs {
		switch v := c.(type) {
		case ccNull:
			cd.Null = bool(v)
		case ccPrimaryKey:
			cd.PrimaryKey = true
			cd.PrimaryKeyDir = SortOrder(v.sort)
			cd.AutoIncrement = v.autoincrement
		case ccUnique:
			cd.Unique = bool(v)
		case ccCollate:
			cd.Collate = string(v)
		case ccDefault:
			cd.Default = interface{}(v)
		default:
			panic("unhandled constraint")
		}
	}
	return cd
}

// CREATE TABLE constraint (primary key, index)
type TableConstraint interface{}
type TablePrimaryKey struct {
	IndexedColumns []IndexedColumn
}
type TableUnique struct {
	IndexedColumns []IndexedColumn
}
type TableForeignKey struct {
	Columns        []string
	ForeignTable   string
	ForeignColumns []string
	Triggers       []Trigger
}
type Trigger interface{}
type TriggerOnDelete TriggerAction
type TriggerOnUpdate TriggerAction

type TriggerAction int

const (
	ActionSetNull TriggerAction = iota
	ActionSetDefault
	ActionCascade
	ActionRestrict
	ActionNoAction
)

// TriggerMatch string

// A `CREATE INDEX` statement
type CreateIndexStmt struct {
	Index          string
	Table          string
	Unique         bool
	IndexedColumns []IndexedColumn
	Where          Expression
}

// Indexed column, for CreateIndexStmt, and index table constraints
type IndexedColumn struct {
	Column    string
	Collate   string
	SortOrder SortOrder
}

type Expression interface{}

type ExBinaryOp struct {
	Op          string
	Left, Right Expression
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
