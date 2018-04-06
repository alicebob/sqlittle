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
	primaryKey PrimaryKey
	autoIncrement bool
	unique bool
	null bool
	sortOrder SortOrder
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
%type<primaryKey> primaryKey
%type<null> null
%type<unique> unique
%type<autoIncrement> autoIncrement
%type<sortOrder> sortOrder

%token SELECT FROM CREATE TABLE INDEX ON PRIMARY KEY ASC DESC AUTOINCREMENT NOT NULL UNIQUE
%token<identifier> tBare
%token<signedNumber> tSignedNumber

%%

program:
	selectStmt |
	createTableStmt |
	createIndexStmt

identifier:
	tBare {
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

columnDefList:
	columnDef {
		$$ = []ColumnDef{$1}
	} |
	columnDefList ',' columnDef {
		$$ = append($1, $3)
	}

columnDef:
	identifier typeName primaryKey autoIncrement null unique {
		$$ = ColumnDef{
			Name: $1,
			Type: $2,
			PrimaryKey: $3,
			AutoIncrement: $4,
			Null: $5,
			Unique: $6,
		}
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

primaryKey:
	{
		$$ = PKNone
	} |
	PRIMARY KEY {
		$$ = PKAsc
	} |
	PRIMARY KEY ASC {
		$$ = PKAsc
	} |
	PRIMARY KEY DESC {
		$$ = PKDesc
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

autoIncrement:
	{
		$$ = false
	} |
	AUTOINCREMENT {
		$$ = true
	}

null:
	{
		$$ = true
	} |
	NOT NULL {
		$$ = false
	} |
	NULL {
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
		yylex.(*Lexer).result = SelectStmt{ Columns: $2, Table: $4 }
	}

createTableStmt:
	CREATE TABLE identifier '(' columnDefList ')' {
		yylex.(*Lexer).result = CreateTableStmt{ Table: $3, Columns: $5 }
	}

createIndexStmt:
	CREATE unique INDEX identifier ON identifier '(' indexedColumnDefList ')' {
		yylex.(*Lexer).result = CreateIndexStmt{
			Index: $4,
			Table: $6,
			Unique: $2,
			IndexedColumns: $8,
		}
	}
