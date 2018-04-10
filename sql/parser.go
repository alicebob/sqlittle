//line parser.go.y:2
package sql

import __yyfmt__ "fmt"

//line parser.go.y:2
//line parser.go.y:5
type yySymType struct {
	yys                  int
	identifier           string
	signedNumber         string
	expr                 interface{}
	columnList           []string
	columnName           string
	columnDefList        []ColumnDef
	columnDef            ColumnDef
	indexedColumnDefList []IndexDef
	indexedColumnDef     IndexDef
	name                 string
	primaryKey           PrimaryKey
	autoIncrement        bool
	unique               bool
	null                 bool
	sortOrder            SortOrder
}

const SELECT = 57346
const FROM = 57347
const CREATE = 57348
const TABLE = 57349
const INDEX = 57350
const ON = 57351
const PRIMARY = 57352
const KEY = 57353
const ASC = 57354
const DESC = 57355
const AUTOINCREMENT = 57356
const NOT = 57357
const NULL = 57358
const UNIQUE = 57359
const tBare = 57360
const tLiteral = 57361
const tIdentifier = 57362
const tSignedNumber = 57363

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"SELECT",
	"FROM",
	"CREATE",
	"TABLE",
	"INDEX",
	"ON",
	"PRIMARY",
	"KEY",
	"ASC",
	"DESC",
	"AUTOINCREMENT",
	"NOT",
	"NULL",
	"UNIQUE",
	"tBare",
	"tLiteral",
	"tIdentifier",
	"tSignedNumber",
	"'*'",
	"','",
	"'('",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 67

var yyAct = [...]int{

	9, 44, 14, 41, 25, 53, 52, 54, 51, 28,
	62, 29, 37, 36, 18, 22, 17, 11, 21, 12,
	23, 10, 11, 26, 12, 13, 42, 31, 32, 26,
	15, 47, 48, 33, 16, 15, 8, 59, 45, 56,
	57, 49, 50, 40, 39, 35, 27, 19, 5, 58,
	6, 55, 38, 20, 45, 61, 60, 46, 34, 30,
	43, 24, 7, 4, 3, 2, 1,
}
var yyPact = [...]int{

	44, -1000, -1000, -1000, -1000, -1, 18, 11, -1000, -1000,
	-1000, -1000, -1000, 4, 39, -1000, -1, 4, -9, 4,
	-1000, -1000, 4, 37, -14, -1000, 4, 4, 4, -1000,
	35, -11, -12, -1000, 30, 32, 5, 4, 16, -1000,
	29, -17, -1000, -18, -1000, 27, 13, 21, -1000, -1000,
	-1000, -1000, 5, 4, -1000, -1000, -1000, -1000, -1000, -1000,
	-15, -1000, -1000,
}
var yyPgo = [...]int{

	0, 66, 65, 64, 63, 0, 3, 62, 36, 61,
	4, 60, 1, 59, 58, 57, 2, 52, 51,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 5, 5, 6, 8, 8, 7,
	7, 9, 9, 10, 13, 13, 13, 13, 14, 14,
	14, 14, 18, 18, 18, 17, 17, 15, 15, 15,
	16, 16, 11, 11, 12, 2, 3, 4,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 1, 3, 6, 0, 1, 4, 6, 0, 2,
	3, 3, 0, 1, 1, 0, 1, 0, 2, 1,
	0, 1, 1, 3, 2, 4, 6, 9,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, 4, 6, -7, -8, -5,
	22, 18, 20, 7, -16, 17, 23, 5, -5, 8,
	-8, -5, 24, -5, -9, -10, -5, 9, 23, 25,
	-13, -5, -5, -10, -14, 10, 24, 24, -17, 14,
	11, -6, 21, -11, -12, -5, -15, 15, 16, 12,
	13, 25, 23, 23, 25, -18, 12, 13, -16, 16,
	-6, -12, 25,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 0, 30, 0, 9, 7,
	8, 4, 5, 0, 0, 31, 0, 0, 0, 0,
	10, 35, 0, 0, 0, 11, 14, 0, 0, 36,
	18, 15, 0, 12, 25, 0, 0, 0, 27, 26,
	19, 0, 6, 0, 32, 22, 30, 0, 29, 20,
	21, 16, 0, 0, 37, 34, 23, 24, 13, 28,
	0, 33, 17,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	24, 25, 22, 3, 23,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:54
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:57
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:62
		{
			yyVAL.signedNumber = yyDollar[1].signedNumber
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:67
		{
			yyVAL.columnName = yyDollar[1].identifier
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:70
		{
			yyVAL.columnName = "*"
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:75
		{
			yyVAL.columnList = []string{yyDollar[1].columnName}
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:78
		{
			yyVAL.columnList = append(yyDollar[1].columnList, yyDollar[3].columnName)
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:83
		{
			yyVAL.columnDefList = []ColumnDef{yyDollar[1].columnDef}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:86
		{
			yyVAL.columnDefList = append(yyDollar[1].columnDefList, yyDollar[3].columnDef)
		}
	case 13:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:91
		{
			yyVAL.columnDef = ColumnDef{
				Name:          yyDollar[1].identifier,
				Type:          yyDollar[2].name,
				PrimaryKey:    yyDollar[3].primaryKey,
				AutoIncrement: yyDollar[4].autoIncrement,
				Null:          yyDollar[5].null,
				Unique:        yyDollar[6].unique,
			}
		}
	case 14:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:103
		{
			yyVAL.name = ""
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:106
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:109
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 17:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:112
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 18:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:117
		{
			yyVAL.primaryKey = PKNone
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:120
		{
			yyVAL.primaryKey = PKAsc
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:123
		{
			yyVAL.primaryKey = PKAsc
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:126
		{
			yyVAL.primaryKey = PKDesc
		}
	case 22:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:131
		{
			yyVAL.sortOrder = Asc
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:134
		{
			yyVAL.sortOrder = Asc
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:137
		{
			yyVAL.sortOrder = Desc
		}
	case 25:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:142
		{
			yyVAL.autoIncrement = false
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:145
		{
			yyVAL.autoIncrement = true
		}
	case 27:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:150
		{
			yyVAL.null = true
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:153
		{
			yyVAL.null = false
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:156
		{
			yyVAL.null = true
		}
	case 30:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:161
		{
			yyVAL.unique = false
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:164
		{
			yyVAL.unique = true
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:169
		{
			yyVAL.indexedColumnDefList = []IndexDef{yyDollar[1].indexedColumnDef}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:172
		{
			yyVAL.indexedColumnDefList = append(yyDollar[1].indexedColumnDefList, yyDollar[3].indexedColumnDef)
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:177
		{
			yyVAL.indexedColumnDef = IndexDef{
				Column:    yyDollar[1].identifier,
				SortOrder: yyDollar[2].sortOrder,
			}
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:185
		{
			yylex.(*lexer).result = SelectStmt{Columns: yyDollar[2].columnList, Table: yyDollar[4].identifier}
		}
	case 36:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:190
		{
			yylex.(*lexer).result = CreateTableStmt{Table: yyDollar[3].identifier, Columns: yyDollar[5].columnDefList}
		}
	case 37:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:195
		{
			yylex.(*lexer).result = CreateIndexStmt{
				Index:          yyDollar[4].identifier,
				Table:          yyDollar[6].identifier,
				Unique:         yyDollar[2].unique,
				IndexedColumns: yyDollar[8].indexedColumnDefList,
			}
		}
	}
	goto yystack /* stack new state and value */
}
