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
	name string
	primaryKey PrimaryKey
	autoIncrement bool
	unique bool
	null bool
}

%type<expr> program
%type<expr> selectStmt
%type<expr> createTableStmt
%type<identifier> identifier
%type<signedNumber> signedNumber
%type<columnList> columnList
%type<columnName> columnName
%type<columnDefList> columnDefList
%type<columnDef> columnDef
%type<name> typeName
%type<primaryKey> primaryKey
%type<null> null
%type<unique> unique
%type<autoIncrement> autoIncrement

%token SELECT FROM CREATE TABLE PRIMARY KEY ASC DESC AUTOINCREMENT NOT NULL UNIQUE
%token<identifier> tBare
%token<signedNumber> tSignedNumber

%%

program:
	selectStmt |
	createTableStmt

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

selectStmt:
	SELECT columnList FROM identifier {
		yylex.(*Lexer).result = SelectStmt{ Columns: $2, Table: $4 }
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

createTableStmt:
	CREATE TABLE identifier '(' columnDefList ')' {
		yylex.(*Lexer).result = CreateTableStmt{ Table: $3, Columns: $5 }
	}
