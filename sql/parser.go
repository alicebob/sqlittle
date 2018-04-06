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
const tSignedNumber = 57361

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

const yyLast = 65

var yyAct = [...]int{

	9, 43, 13, 40, 24, 52, 51, 53, 50, 27,
	61, 28, 36, 17, 35, 21, 41, 20, 16, 22,
	11, 11, 25, 10, 14, 12, 30, 31, 25, 46,
	47, 58, 32, 8, 15, 14, 38, 44, 55, 56,
	48, 49, 39, 34, 26, 18, 54, 37, 57, 19,
	5, 45, 6, 44, 60, 59, 33, 29, 42, 23,
	7, 4, 3, 2, 1,
}
var yyPact = [...]int{

	46, -1000, -1000, -1000, -1000, 3, 18, 13, -1000, -1000,
	-1000, -1000, 2, 37, -1000, 3, 2, -7, 2, -1000,
	-1000, 2, 35, -12, -1000, 2, 2, 2, -1000, 33,
	-8, -10, -1000, 22, 31, -3, 2, 14, -1000, 28,
	-15, -1000, -16, -1000, 26, 7, 15, -1000, -1000, -1000,
	-1000, -3, 2, -1000, -1000, -1000, -1000, -1000, -1000, -13,
	-1000, -1000,
}
var yyPgo = [...]int{

	0, 64, 63, 62, 61, 0, 3, 60, 33, 59,
	4, 58, 1, 57, 56, 51, 2, 47, 46,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 5, 6, 8, 8, 7, 7,
	9, 9, 10, 13, 13, 13, 13, 14, 14, 14,
	14, 18, 18, 18, 17, 17, 15, 15, 15, 16,
	16, 11, 11, 12, 2, 3, 4,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 3,
	1, 3, 6, 0, 1, 4, 6, 0, 2, 3,
	3, 0, 1, 1, 0, 1, 0, 2, 1, 0,
	1, 1, 3, 2, 4, 6, 9,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, 4, 6, -7, -8, -5,
	20, 18, 7, -16, 17, 21, 5, -5, 8, -8,
	-5, 22, -5, -9, -10, -5, 9, 21, 23, -13,
	-5, -5, -10, -14, 10, 22, 22, -17, 14, 11,
	-6, 19, -11, -12, -5, -15, 15, 16, 12, 13,
	23, 21, 21, 23, -18, 12, 13, -16, 16, -6,
	-12, 23,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 0, 29, 0, 8, 6,
	7, 4, 0, 0, 30, 0, 0, 0, 0, 9,
	34, 0, 0, 0, 10, 13, 0, 0, 35, 17,
	14, 0, 11, 24, 0, 0, 0, 26, 25, 18,
	0, 5, 0, 31, 21, 29, 0, 28, 19, 20,
	15, 0, 0, 36, 33, 22, 23, 12, 27, 0,
	32, 16,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	22, 23, 20, 3, 21,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19,
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
		//line parser.go.y:59
		{
			yyVAL.signedNumber = yyDollar[1].signedNumber
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:64
		{
			yyVAL.columnName = yyDollar[1].identifier
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:67
		{
			yyVAL.columnName = "*"
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:72
		{
			yyVAL.columnList = []string{yyDollar[1].columnName}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:75
		{
			yyVAL.columnList = append(yyDollar[1].columnList, yyDollar[3].columnName)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:80
		{
			yyVAL.columnDefList = []ColumnDef{yyDollar[1].columnDef}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:83
		{
			yyVAL.columnDefList = append(yyDollar[1].columnDefList, yyDollar[3].columnDef)
		}
	case 12:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:88
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
	case 13:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:100
		{
			yyVAL.name = ""
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:103
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:106
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 16:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:109
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 17:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:114
		{
			yyVAL.primaryKey = PKNone
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:117
		{
			yyVAL.primaryKey = PKAsc
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:120
		{
			yyVAL.primaryKey = PKAsc
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:123
		{
			yyVAL.primaryKey = PKDesc
		}
	case 21:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:128
		{
			yyVAL.sortOrder = Asc
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:131
		{
			yyVAL.sortOrder = Asc
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:134
		{
			yyVAL.sortOrder = Desc
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:139
		{
			yyVAL.autoIncrement = false
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:142
		{
			yyVAL.autoIncrement = true
		}
	case 26:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:147
		{
			yyVAL.null = true
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:150
		{
			yyVAL.null = false
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:153
		{
			yyVAL.null = true
		}
	case 29:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:158
		{
			yyVAL.unique = false
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:161
		{
			yyVAL.unique = true
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:166
		{
			yyVAL.indexedColumnDefList = []IndexDef{yyDollar[1].indexedColumnDef}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:169
		{
			yyVAL.indexedColumnDefList = append(yyDollar[1].indexedColumnDefList, yyDollar[3].indexedColumnDef)
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:174
		{
			yyVAL.indexedColumnDef = IndexDef{
				Column:    yyDollar[1].identifier,
				SortOrder: yyDollar[2].sortOrder,
			}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:182
		{
			yylex.(*lexer).result = SelectStmt{Columns: yyDollar[2].columnList, Table: yyDollar[4].identifier}
		}
	case 35:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:187
		{
			yylex.(*lexer).result = CreateTableStmt{Table: yyDollar[3].identifier, Columns: yyDollar[5].columnDefList}
		}
	case 36:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:192
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
