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
%type<columnList> columnList
%type<columnName> columnName
%type<expr> createTableStmt

%token<token> tSelect tFrom tCreate tTable tBare
%token<token> ',' '*'


%%

program
  : selectStmt
  {
    $$ = $1 // needed?
    yylex.(*Lexer).result = $$
  }
  | createTableStmt
  {
    $$ = $1 // needed?
    yylex.(*Lexer).result = $$
  }

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

selectStmt
  : tSelect columnList tFrom tBare
  {
    $$ = SelectStmt{ Columns: $2, Table: $4.s }
  }

createTableStmt
  : tCreate tTable tBare
  {
    $$ = CreatTableStmt{ Table: $3.s }
  }
