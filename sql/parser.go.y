%{
package sql
%}

%union {
	literal string
	identifier string
	signedNumber int64
	expr interface{}
	columnList []string
	columnName string
	columnDefList []ColumnDef
	columnDef ColumnDef
	indexedColumnList []IndexedColumn
	indexedColumn IndexedColumn
	name string
	withoutRowid bool
	unique bool
	bool bool
	collate string
	sortOrder SortOrder
	columnConstraint columnConstraint
	columnConstraintList []columnConstraint
	tableConstraint TableConstraint
	tableConstraintList []TableConstraint
}

%type<expr> program
%type<expr> selectStmt
%type<expr> createTableStmt
%type<expr> createIndexStmt
%type<identifier> identifier
%type<literal> literal
%type<signedNumber> signedNumber
%type<columnList> columnList
%type<columnName> columnName
%type<columnDefList> columnDefList
%type<columnDef> columnDef
%type<indexedColumnList> indexedColumnList
%type<indexedColumn> indexedColumn
%type<name> typeName
%type<unique> unique
%type<withoutRowid> withoutRowid
%type<collate> collate
%type<sortOrder> sortOrder
%type<bool> autoincrement
%type<columnConstraint> columnConstraint
%type<columnConstraintList> columnConstraintList
%type<tableConstraint> tableConstraint
%type<tableConstraintList> tableConstraintList

%token SELECT FROM CREATE TABLE INDEX ON PRIMARY KEY ASC DESC
%token AUTOINCREMENT NOT NULL UNIQUE COLLATE WITHOUT ROWID DEFAULT
%token<identifier> tBare tLiteral tIdentifier
%token<signedNumber> tSignedNumber

%%

program:
	selectStmt |
	createTableStmt |
	createIndexStmt

literal:
	tBare {
		$$ = $1
	} |
	tLiteral {
		$$ = $1
	}

identifier:
	tBare {
		$$ = $1
	} |
	tIdentifier {
		$$ = $1
	}

signedNumber:
	tSignedNumber {
		$$ = $1
	}

columnName:
	identifier {
		$$ = $1
	} |
	'*' {
		$$ = "*"
	}

columnList:
	columnName {
		$$ = []string{$1}
	} |
	columnList ',' columnName {
		$$ = append($1, $3)
	}

columnConstraint:
	PRIMARY KEY sortOrder autoincrement {
		$$ = ccPrimaryKey{$3, $4}
	} |
	UNIQUE {
		$$ = ccUnique(true)
	} |
	NULL {
		$$ = ccNull(true)
	} |
	NOT NULL {
		$$ = ccNull(false)
	} |
	COLLATE identifier {
		$$ = ccCollate($2)
	} |
	DEFAULT signedNumber {
		$$ = ccDefault($2)
	} |
	DEFAULT literal {
		$$ = ccDefault($2)
	}

columnConstraintList:
	{
	} |
	columnConstraint {
		$$ = []columnConstraint{$1}
	} |
	columnConstraintList columnConstraint {
		$$ = append($1, $2)
	}

tableConstraint:
	PRIMARY KEY '(' indexedColumnList ')' {
		$$ = TablePrimaryKey{$4}
	} |
	UNIQUE '(' indexedColumnList ')' {
		$$ = TableUnique{$3}
	}

tableConstraintList:
	{ } |
	',' tableConstraint {
		$$ = []TableConstraint{$2}
	} |
	tableConstraintList ',' tableConstraint {
		$$ = append($1, $3)
	}


autoincrement:
	{ } |
	AUTOINCREMENT {
		$$ = true
	}

columnDefList:
	columnDef {
		$$ = []ColumnDef{$1}
	} |
	columnDefList ',' columnDef {
		$$ = append($1, $3)
	}

columnDef:
	identifier typeName columnConstraintList {
		$$ = makeColumnDef($1, $2, $3)
	}

typeName:
	{
		$$ = ""
	} |
	identifier {
		$$ = $1
	} |
	identifier '(' signedNumber ')' {
		$$ = $1
	} |
	identifier '(' signedNumber ',' signedNumber ')' {
		$$ = $1
	}

collate:
	{ } |
	COLLATE literal {
		$$ = $2
	}

sortOrder:
	{
		$$ = Asc
	} |
	ASC {
		$$ = Asc
	} |
	DESC {
		$$ = Desc
	}

withoutRowid:
	{
		$$ = false
	} |
	WITHOUT ROWID {
		$$ = true
	}

unique:
	{
		$$ = false
	} |
	UNIQUE {
		$$ = true
	}

indexedColumnList:
	indexedColumn {
		$$ = []IndexedColumn{$1}
	} |
	indexedColumnList ',' indexedColumn {
		$$ = append($1, $3)
	}

indexedColumn:
	identifier collate sortOrder {
		$$ = IndexedColumn{
			Column: $1,
			Collate: $2,
			SortOrder: $3,
		}
	}

selectStmt:
	SELECT columnList FROM identifier {
		yylex.(*lexer).result = SelectStmt{ Columns: $2, Table: $4 }
	}

createTableStmt:
	CREATE TABLE identifier '(' columnDefList tableConstraintList ')' withoutRowid {
		yylex.(*lexer).result = CreateTableStmt{
			Table: $3,
			Columns: $5,
			Constraints: $6,
			WithoutRowid: $8,
		}
	}

createIndexStmt:
	CREATE unique INDEX identifier ON identifier '(' indexedColumnList ')' {
		yylex.(*lexer).result = CreateIndexStmt{
			Index: $4,
			Table: $6,
			Unique: $2,
			IndexedColumns: $8,
		}
	}
