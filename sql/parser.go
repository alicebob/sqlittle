//line parser.go.y:2
package sql

import __yyfmt__ "fmt"

//line parser.go.y:2
//line parser.go.y:5
type yySymType struct {
	yys                  int
	literal              string
	identifier           string
	signedNumber         int64
	expr                 interface{}
	columnList           []string
	columnName           string
	columnDefList        []ColumnDef
	columnDef            ColumnDef
	indexedColumnDefList []IndexDef
	indexedColumnDef     IndexDef
	name                 string
	withoutRowid         bool
	unique               bool
	sortOrder            SortOrder
	iface                interface{}
	ifaceList            []interface{}
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
const COLLATE = 57360
const WITHOUT = 57361
const ROWID = 57362
const DEFAULT = 57363
const tBare = 57364
const tLiteral = 57365
const tIdentifier = 57366
const tSignedNumber = 57367

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
	"COLLATE",
	"WITHOUT",
	"ROWID",
	"DEFAULT",
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

const yyLast = 77

var yyAct = [...]int{

	9, 58, 60, 51, 37, 65, 64, 66, 63, 28,
	72, 29, 45, 25, 18, 44, 22, 11, 21, 12,
	23, 10, 17, 26, 11, 53, 12, 31, 32, 26,
	54, 55, 46, 53, 35, 49, 69, 13, 8, 61,
	62, 47, 33, 50, 16, 38, 59, 15, 56, 48,
	41, 40, 39, 42, 27, 20, 43, 19, 5, 36,
	6, 68, 67, 34, 14, 30, 59, 71, 70, 57,
	24, 7, 52, 4, 3, 2, 1,
}
var yyPact = [...]int{

	54, -1000, -1000, -1000, -1000, -5, 30, 17, -1000, -1000,
	-1000, -1000, -1000, 2, 49, -1000, -5, 2, -12, 2,
	-1000, -1000, 2, 45, -18, -1000, 2, 2, 2, 15,
	35, -13, -16, -1000, -1000, 12, 35, -1000, 38, -1000,
	-1000, 19, 2, 8, 0, 2, -1000, -1000, 27, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -21, -22, -1000, 27,
	22, -1000, -1000, -1000, 0, 2, -1000, -1000, -1000, -1000,
	-19, -1000, -1000,
}
var yyPgo = [...]int{

	0, 76, 75, 74, 73, 0, 72, 3, 71, 38,
	70, 13, 69, 1, 65, 64, 63, 2, 61, 4,
	59,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 6, 6, 5, 5, 7, 9,
	9, 8, 8, 19, 19, 19, 19, 19, 19, 19,
	20, 20, 20, 18, 18, 10, 10, 11, 14, 14,
	14, 14, 17, 17, 17, 16, 16, 15, 15, 12,
	12, 13, 2, 3, 4,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 4, 1, 1, 2, 2, 2, 2,
	0, 1, 2, 0, 1, 1, 3, 3, 0, 1,
	4, 6, 0, 1, 1, 0, 2, 0, 1, 1,
	3, 2, 4, 7, 9,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, 4, 6, -8, -9, -5,
	26, 22, 24, 7, -15, 17, 27, 5, -5, 8,
	-9, -5, 28, -5, -10, -11, -5, 9, 27, 29,
	-14, -5, -5, -11, -16, 19, -20, -19, 10, 17,
	16, 15, 18, 21, 28, 28, 20, -19, 11, 16,
	-5, -7, -6, 25, 22, 23, -7, -12, -13, -5,
	-17, 12, 13, 29, 27, 27, 29, -17, -18, 14,
	-7, -13, 29,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 0, 37, 0, 11, 9,
	10, 6, 7, 0, 0, 38, 0, 0, 0, 0,
	12, 42, 0, 0, 0, 25, 28, 0, 0, 35,
	20, 29, 0, 26, 43, 0, 27, 21, 0, 14,
	15, 0, 0, 0, 0, 0, 36, 22, 32, 16,
	17, 18, 19, 8, 4, 5, 0, 0, 39, 32,
	23, 33, 34, 30, 0, 0, 44, 41, 13, 24,
	0, 40, 31,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	28, 29, 26, 3, 27,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25,
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
		//line parser.go.y:58
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:61
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:66
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:69
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:74
		{
			yyVAL.signedNumber = yyDollar[1].signedNumber
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:79
		{
			yyVAL.columnName = yyDollar[1].identifier
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:82
		{
			yyVAL.columnName = "*"
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:87
		{
			yyVAL.columnList = []string{yyDollar[1].columnName}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:90
		{
			yyVAL.columnList = append(yyDollar[1].columnList, yyDollar[3].columnName)
		}
	case 13:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:95
		{
			yyVAL.ifaceList = []interface{}{primaryKey(yyDollar[3].sortOrder), yyDollar[4].iface}
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:98
		{
			yyVAL.ifaceList = []interface{}{unique(true)}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:101
		{
			yyVAL.ifaceList = []interface{}{null(true)}
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:104
		{
			yyVAL.ifaceList = []interface{}{null(false)}
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:107
		{
			yyVAL.ifaceList = []interface{}{collate(yyDollar[2].identifier)}
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:110
		{
			yyVAL.ifaceList = []interface{}{defaul(yyDollar[2].signedNumber)}
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:113
		{
			yyVAL.ifaceList = []interface{}{defaul(yyDollar[2].literal)}
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:118
		{
			yyVAL.ifaceList = nil
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:121
		{
			yyVAL.ifaceList = yyDollar[1].ifaceList
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:124
		{
			yyVAL.ifaceList = append(yyDollar[1].ifaceList, yyDollar[2].ifaceList...)
		}
	case 23:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:129
		{
			yyVAL.iface = autoincrement(false)
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:132
		{
			yyVAL.iface = autoincrement(true)
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:137
		{
			yyVAL.columnDefList = []ColumnDef{yyDollar[1].columnDef}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:140
		{
			yyVAL.columnDefList = append(yyDollar[1].columnDefList, yyDollar[3].columnDef)
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:145
		{
			yyVAL.columnDef = makeDef(yyDollar[1].identifier, yyDollar[2].name, yyDollar[3].ifaceList)
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:150
		{
			yyVAL.name = ""
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:153
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:156
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 31:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:159
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 32:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:164
		{
			yyVAL.sortOrder = Asc
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:167
		{
			yyVAL.sortOrder = Asc
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:170
		{
			yyVAL.sortOrder = Desc
		}
	case 35:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:175
		{
			yyVAL.withoutRowid = false
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:178
		{
			yyVAL.withoutRowid = true
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:183
		{
			yyVAL.unique = false
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:186
		{
			yyVAL.unique = true
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:191
		{
			yyVAL.indexedColumnDefList = []IndexDef{yyDollar[1].indexedColumnDef}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:194
		{
			yyVAL.indexedColumnDefList = append(yyDollar[1].indexedColumnDefList, yyDollar[3].indexedColumnDef)
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:199
		{
			yyVAL.indexedColumnDef = IndexDef{
				Column:    yyDollar[1].identifier,
				SortOrder: yyDollar[2].sortOrder,
			}
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:207
		{
			yylex.(*lexer).result = SelectStmt{Columns: yyDollar[2].columnList, Table: yyDollar[4].identifier}
		}
	case 43:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:212
		{
			yylex.(*lexer).result = CreateTableStmt{
				Table:        yyDollar[3].identifier,
				Columns:      yyDollar[5].columnDefList,
				WithoutRowid: yyDollar[7].withoutRowid,
			}
		}
	case 44:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:221
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
