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
const tBare = 57363
const tLiteral = 57364
const tIdentifier = 57365
const tSignedNumber = 57366

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

const yyLast = 71

var yyAct = [...]int{

	9, 53, 55, 50, 37, 60, 59, 61, 58, 28,
	25, 29, 67, 44, 18, 43, 22, 11, 21, 12,
	23, 10, 11, 26, 12, 17, 51, 31, 32, 26,
	45, 35, 48, 64, 13, 8, 56, 57, 19, 33,
	47, 46, 38, 49, 15, 54, 16, 41, 40, 39,
	42, 27, 20, 5, 36, 6, 63, 62, 34, 14,
	30, 54, 66, 65, 52, 24, 7, 4, 3, 2,
	1,
}
var yyPact = [...]int{

	49, -1000, -1000, -1000, -1000, -4, 27, 20, -1000, -1000,
	-1000, -1000, -1000, 1, 30, -1000, -4, 1, -11, 1,
	-1000, -1000, 1, 42, -17, -1000, 1, 1, 1, 12,
	32, -12, -14, -1000, -1000, 10, 32, -1000, 29, -1000,
	-1000, 16, 1, 2, 1, -1000, -1000, 24, -1000, -1000,
	-20, -1000, -21, -1000, 24, 19, -1000, -1000, -1000, 2,
	1, -1000, -1000, -1000, -1000, -16, -1000, -1000,
}
var yyPgo = [...]int{

	0, 70, 69, 68, 67, 0, 3, 66, 35, 65,
	10, 64, 1, 60, 59, 58, 2, 56, 4, 54,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 5, 5, 6, 8, 8, 7,
	7, 18, 18, 18, 18, 18, 19, 19, 19, 17,
	17, 9, 9, 10, 13, 13, 13, 13, 16, 16,
	16, 15, 15, 14, 14, 11, 11, 12, 2, 3,
	4,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 4, 1, 1, 2, 2, 0, 1, 2, 0,
	1, 1, 3, 3, 0, 1, 4, 6, 0, 1,
	1, 0, 2, 0, 1, 1, 3, 2, 4, 7,
	9,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, 4, 6, -7, -8, -5,
	25, 21, 23, 7, -14, 17, 26, 5, -5, 8,
	-8, -5, 27, -5, -9, -10, -5, 9, 26, 28,
	-13, -5, -5, -10, -15, 19, -19, -18, 10, 17,
	16, 15, 18, 27, 27, 20, -18, 11, 16, -5,
	-6, 24, -11, -12, -5, -16, 12, 13, 28, 26,
	26, 28, -16, -17, 14, -6, -12, 28,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 0, 33, 0, 9, 7,
	8, 4, 5, 0, 0, 34, 0, 0, 0, 0,
	10, 38, 0, 0, 0, 21, 24, 0, 0, 31,
	16, 25, 0, 22, 39, 0, 23, 17, 0, 12,
	13, 0, 0, 0, 0, 32, 18, 28, 14, 15,
	0, 6, 0, 35, 28, 19, 29, 30, 26, 0,
	0, 40, 37, 11, 20, 0, 36, 27,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	27, 28, 25, 3, 26,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24,
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
		//line parser.go.y:56
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:59
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:64
		{
			yyVAL.signedNumber = yyDollar[1].signedNumber
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:69
		{
			yyVAL.columnName = yyDollar[1].identifier
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:72
		{
			yyVAL.columnName = "*"
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:77
		{
			yyVAL.columnList = []string{yyDollar[1].columnName}
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:80
		{
			yyVAL.columnList = append(yyDollar[1].columnList, yyDollar[3].columnName)
		}
	case 11:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:85
		{
			yyVAL.ifaceList = []interface{}{primaryKey(yyDollar[3].sortOrder), yyDollar[4].iface}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:88
		{
			yyVAL.ifaceList = []interface{}{unique(true)}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:91
		{
			yyVAL.ifaceList = []interface{}{null(true)}
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:94
		{
			yyVAL.ifaceList = []interface{}{null(false)}
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:97
		{
			yyVAL.ifaceList = []interface{}{collate(yyDollar[2].identifier)}
		}
	case 16:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:102
		{
			yyVAL.ifaceList = nil
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:105
		{
			yyVAL.ifaceList = yyDollar[1].ifaceList
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:108
		{
			yyVAL.ifaceList = append(yyDollar[1].ifaceList, yyDollar[2].ifaceList...)
		}
	case 19:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:113
		{
			yyVAL.iface = autoincrement(false)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:116
		{
			yyVAL.iface = autoincrement(true)
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:121
		{
			yyVAL.columnDefList = []ColumnDef{yyDollar[1].columnDef}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:124
		{
			yyVAL.columnDefList = append(yyDollar[1].columnDefList, yyDollar[3].columnDef)
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:129
		{
			yyVAL.columnDef = makeDef(yyDollar[1].identifier, yyDollar[2].name, yyDollar[3].ifaceList)
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:134
		{
			yyVAL.name = ""
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:137
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:140
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 27:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:143
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:148
		{
			yyVAL.sortOrder = Asc
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:151
		{
			yyVAL.sortOrder = Asc
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:154
		{
			yyVAL.sortOrder = Desc
		}
	case 31:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:159
		{
			yyVAL.withoutRowid = false
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:162
		{
			yyVAL.withoutRowid = true
		}
	case 33:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:167
		{
			yyVAL.unique = false
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:170
		{
			yyVAL.unique = true
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:175
		{
			yyVAL.indexedColumnDefList = []IndexDef{yyDollar[1].indexedColumnDef}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:178
		{
			yyVAL.indexedColumnDefList = append(yyDollar[1].indexedColumnDefList, yyDollar[3].indexedColumnDef)
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:183
		{
			yyVAL.indexedColumnDef = IndexDef{
				Column:    yyDollar[1].identifier,
				SortOrder: yyDollar[2].sortOrder,
			}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:191
		{
			yylex.(*lexer).result = SelectStmt{Columns: yyDollar[2].columnList, Table: yyDollar[4].identifier}
		}
	case 39:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:196
		{
			yylex.(*lexer).result = CreateTableStmt{
				Table:        yyDollar[3].identifier,
				Columns:      yyDollar[5].columnDefList,
				WithoutRowid: yyDollar[7].withoutRowid,
			}
		}
	case 40:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:205
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
