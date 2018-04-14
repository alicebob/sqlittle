%{
package sql
%}

%union {
	literal string
	identifier string
	signedNumber int64
	expr interface{}
	columnNameList []string
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
	triggerAction TriggerAction
	trigger Trigger
	triggerList []Trigger
}

%type<expr> program
%type<expr> selectStmt
%type<expr> createTableStmt
%type<expr> createIndexStmt
%type<identifier> identifier
%type<literal> literal
%type<signedNumber> signedNumber
%type<columnName> columnName resultColumn
%type<columnNameList> columnNameList optColumnNameList resultColumnList
%type<columnDefList> columnDefList
%type<columnDef> columnDef
%type<indexedColumnList> indexedColumnList
%type<indexedColumn> indexedColumn
%type<name> typeName constraintName
%type<unique> unique
%type<withoutRowid> withoutRowid
%type<collate> collate
%type<sortOrder> sortOrder
%type<bool> autoincrement
%type<columnConstraint> columnConstraint
%type<columnConstraintList> columnConstraintList
%type<tableConstraint> tableConstraint
%type<tableConstraintList> tableConstraintList
%type<triggerAction> triggerAction
%type<trigger> trigger
%type<triggerList> triggerList

%token ACTION
%token ASC
%token AUTOINCREMENT
%token CASCADE
%token COLLATE
%token CONSTRAINT
%token CREATE
%token DEFAULT
%token DELETE
%token DESC
%token FOREIGN
%token FROM
%token INDEX
%token KEY
%token NO
%token NOT
%token NULL
%token ON
%token PRIMARY
%token REFERENCES
%token RESTRICT
%token ROWID
%token SELECT
%token SET
%token TABLE
%token UNIQUE
%token UPDATE
%token WITHOUT
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
	}

columnNameList:
	columnName {
		$$ = []string{$1}
	} |
	columnNameList ',' columnName {
		$$ = append($1, $3)
	}

optColumnNameList:
	'(' columnNameList ')' {
		$$ = $2
	}

resultColumn:
	columnName {
		$$ = $1
	} |
	'*' {
		$$ = "*"
	}

resultColumnList:
	resultColumn {
		$$ = []string{$1}
	} |
	resultColumnList ',' resultColumn {
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
	} |
	FOREIGN KEY '(' columnNameList ')' REFERENCES identifier optColumnNameList triggerList {
		$$ = TableForeignKey{
			Columns: $4,
			ForeignTable: $7,
			ForeignColumns: $8,
			Triggers: $9,
		}
	}

constraintName:
	{ } |
	CONSTRAINT identifier {
	}

tableConstraintList:
	{ } |
	',' constraintName tableConstraint {
		$$ = []TableConstraint{$3}
	} |
	tableConstraintList ',' constraintName tableConstraint {
		$$ = append($1, $4)
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

triggerAction:
	SET NULL {
		$$ = ActionSetNull
	} |
	SET DEFAULT {
		$$ = ActionSetDefault
	} |
	CASCADE {
		$$ = ActionCascade
	} |
	RESTRICT {
		$$ = ActionRestrict
	} |
	NO ACTION {
		$$ = ActionNoAction
	}

trigger:
	ON DELETE triggerAction {
		$$ = TriggerOnDelete($3)
	} |
	ON UPDATE triggerAction {
		$$ = TriggerOnUpdate($3)
	}

triggerList:
	{ } |
	triggerList trigger {
		$$ = append($1, $2)
	}

selectStmt:
	SELECT resultColumnList FROM identifier {
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
