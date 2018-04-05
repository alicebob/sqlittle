%{
package sql
%}

%union {
	identifier string
	expr interface{}
	columnList []string
	columnName string
	columnDefList []ColumnDef
	columnDef ColumnDef
	null bool
}

%type<expr> program
%type<expr> selectStmt
%type<expr> createTableStmt
%type<identifier> identifier
%type<columnList> columnList
%type<columnName> columnName
%type<columnDefList> columnDefList
%type<columnDef> columnDef
%type<null> null

%token SELECT FROM CREATE TABLE NOT NULL
%token<identifier> tBare

%%

program:
	selectStmt |
	createTableStmt

identifier:
	tBare {
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
	identifier null {
		$$ = ColumnDef{Name: $1, Null: $2}
	} |
	identifier tBare null {
		$$ = ColumnDef{Name: $1, Type: $2, Null: $3}
	}

null:
	/* nothing */ {
		$$ = true
	} |
	NOT NULL {
		$$ = false
	} |
	NULL {
		$$ = true
	}

createTableStmt:
	CREATE TABLE identifier '(' columnDefList ')' {
		yylex.(*Lexer).result = CreateTableStmt{ Table: $3, Columns: $5 }
	}
