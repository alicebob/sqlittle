%{
package sql
%}

%union {
  	token token
  	expr interface{}
  	columnList []string
	columnName string
}

%type<expr> program
%type<expr> selectStmt
%type<expr> createTableStmt
%type<columnList> columnList
%type<columnName> columnName

%token<token> SELECT FROM CREATE TABLE
%token<token> tBare

%%

program
  : selectStmt
  | createTableStmt

columnName:
  	tBare {
		$$ = $1.s
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
    SELECT columnList FROM tBare {
        yylex.(*Lexer).result = SelectStmt{ Columns: $2, Table: $4.s }
    }

createTableStmt:
    CREATE TABLE tBare {
        yylex.(*Lexer).result = CreatTableStmt{ Table: $3.s }
    }
