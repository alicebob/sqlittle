%{
package sql
%}

%union {
	identifier string
	signedNumber string
	expr interface{}
	columnList []string
	columnName string
	columnDefList []ColumnDef
	columnDef ColumnDef
	indexedColumnDefList []IndexDef
	indexedColumnDef IndexDef
	name string
	withoutRowid bool
	unique bool
	sortOrder SortOrder
	iface interface{}
	ifaceList []interface{}
}

%type<expr> program
%type<expr> selectStmt
%type<expr> createTableStmt
%type<expr> createIndexStmt
%type<identifier> identifier
%type<signedNumber> signedNumber
%type<columnList> columnList
%type<columnName> columnName
%type<columnDefList> columnDefList
%type<columnDef> columnDef
%type<indexedColumnDefList> indexedColumnDefList
%type<indexedColumnDef> indexedColumnDef
%type<name> typeName
%type<unique> unique
%type<withoutRowid> withoutRowid
%type<sortOrder> sortOrder
%type<iface> columnConstraint
%type<ifaceList> columnConstraintList

%token SELECT FROM CREATE TABLE INDEX ON PRIMARY KEY ASC DESC
%token AUTOINCREMENT NOT NULL UNIQUE WITHOUT ROWID
%token<identifier> tBare tLiteral tIdentifier
%token<signedNumber> tSignedNumber

%%

program:
	selectStmt |
	createTableStmt |
	createIndexStmt

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
	PRIMARY KEY ASC {
		$$ = primaryKey(Asc)
	} |
	PRIMARY KEY DESC {
		$$ = primaryKey(Desc)
	} |
	PRIMARY KEY {
		$$ = primaryKey(Asc)
	} |
	UNIQUE {
		$$ = unique(true)
	} |
	NULL {
		$$ = null(true)
	} |
	NOT NULL {
		$$ = null(false)
	} |
	AUTOINCREMENT {
		$$ = autoincrement(true)
	}

columnConstraintList:
	{
		$$ = nil
	} |
	columnConstraint {
		$$ = []interface{}{$1}
	} |
	columnConstraintList columnConstraint {
		$$ = append($1, $2)
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
		$$ = makeDef($1, $2, $3)
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

indexedColumnDefList:
	indexedColumnDef {
		$$ = []IndexDef{$1}
	} |
	indexedColumnDefList ',' indexedColumnDef {
		$$ = append($1, $3)
	}

indexedColumnDef:
	identifier sortOrder {
		$$ = IndexDef{
			Column: $1,
			SortOrder: $2,
		}
	}

selectStmt:
	SELECT columnList FROM identifier {
		yylex.(*lexer).result = SelectStmt{ Columns: $2, Table: $4 }
	}

createTableStmt:
	CREATE TABLE identifier '(' columnDefList ')' withoutRowid {
		yylex.(*lexer).result = CreateTableStmt{
			Table: $3,
			Columns: $5,
			WithoutRowid: $7,
		}
	}

createIndexStmt:
	CREATE unique INDEX identifier ON identifier '(' indexedColumnDefList ')' {
		yylex.(*lexer).result = CreateIndexStmt{
			Index: $4,
			Table: $6,
			Unique: $2,
			IndexedColumns: $8,
		}
	}
