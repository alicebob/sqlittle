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
	statement            interface{}
	columnNameList       []string
	columnName           string
	columnDefList        []ColumnDef
	columnDef            ColumnDef
	indexedColumnList    []IndexedColumn
	indexedColumn        IndexedColumn
	name                 string
	withoutRowid         bool
	unique               bool
	bool                 bool
	collate              string
	sortOrder            SortOrder
	columnConstraint     columnConstraint
	columnConstraintList []columnConstraint
	tableConstraint      TableConstraint
	tableConstraintList  []TableConstraint
	triggerAction        TriggerAction
	trigger              Trigger
	triggerList          []Trigger
	expr                 Expression
	exprList             []Expression
	float                float64
}

const ACTION = 57346
const AND = 57347
const ASC = 57348
const AUTOINCREMENT = 57349
const CASCADE = 57350
const COLLATE = 57351
const CONSTRAINT = 57352
const CREATE = 57353
const DEFAULT = 57354
const DELETE = 57355
const DESC = 57356
const FOREIGN = 57357
const FROM = 57358
const GLOB = 57359
const IN = 57360
const INDEX = 57361
const IS = 57362
const KEY = 57363
const LIKE = 57364
const MATCH = 57365
const NO = 57366
const NOT = 57367
const NULL = 57368
const ON = 57369
const OR = 57370
const PRIMARY = 57371
const REFERENCES = 57372
const REGEXP = 57373
const RESTRICT = 57374
const ROWID = 57375
const SELECT = 57376
const SET = 57377
const TABLE = 57378
const UNIQUE = 57379
const UPDATE = 57380
const WHERE = 57381
const WITHOUT = 57382
const tBare = 57383
const tLiteral = 57384
const tIdentifier = 57385
const tOperator = 57386
const tSignedNumber = 57387
const tFloat = 57388

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ACTION",
	"AND",
	"ASC",
	"AUTOINCREMENT",
	"CASCADE",
	"COLLATE",
	"CONSTRAINT",
	"CREATE",
	"DEFAULT",
	"DELETE",
	"DESC",
	"FOREIGN",
	"FROM",
	"GLOB",
	"IN",
	"INDEX",
	"IS",
	"KEY",
	"LIKE",
	"MATCH",
	"NO",
	"NOT",
	"NULL",
	"ON",
	"OR",
	"PRIMARY",
	"REFERENCES",
	"REGEXP",
	"RESTRICT",
	"ROWID",
	"SELECT",
	"SET",
	"TABLE",
	"UNIQUE",
	"UPDATE",
	"WHERE",
	"WITHOUT",
	"tBare",
	"tLiteral",
	"tIdentifier",
	"tOperator",
	"tSignedNumber",
	"tFloat",
	"'-'",
	"'+'",
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
	-1, 109,
	50, 6,
	-2, 76,
	-1, 110,
	50, 7,
	-2, 77,
}

const yyPrivate = 57344

const yyLast = 184

var yyAct = [...]int{

	105, 146, 9, 103, 100, 68, 10, 107, 61, 48,
	76, 102, 137, 69, 18, 106, 121, 10, 21, 116,
	23, 143, 104, 26, 135, 28, 134, 31, 32, 26,
	116, 83, 117, 115, 89, 87, 52, 109, 108, 110,
	72, 62, 114, 112, 113, 59, 111, 118, 70, 47,
	120, 119, 46, 83, 132, 99, 83, 82, 84, 81,
	22, 60, 67, 74, 65, 66, 55, 36, 62, 37,
	63, 64, 17, 70, 62, 114, 112, 113, 88, 79,
	80, 35, 118, 144, 70, 120, 119, 11, 70, 12,
	10, 95, 101, 98, 39, 97, 96, 93, 92, 62,
	75, 63, 64, 65, 66, 16, 126, 6, 145, 33,
	148, 142, 11, 27, 12, 122, 58, 10, 73, 125,
	123, 124, 127, 128, 129, 131, 150, 133, 79, 80,
	5, 51, 44, 56, 149, 45, 25, 147, 10, 138,
	101, 153, 140, 13, 15, 49, 53, 151, 43, 42,
	19, 71, 40, 50, 57, 152, 8, 35, 77, 86,
	41, 91, 154, 130, 94, 34, 78, 139, 141, 29,
	38, 90, 85, 20, 54, 14, 30, 24, 7, 136,
	4, 3, 2, 1,
}
var yyPact = [...]int{

	96, -1000, -1000, -1000, -1000, 46, 107, 56, -1000, -1000,
	-1000, -1000, -1000, 46, 131, -1000, 46, 46, 10, 46,
	-1000, -1000, 46, 86, -24, -1000, 46, 46, 71, 18,
	123, 2, -1, 116, -1000, 46, 147, 26, 123, -1000,
	133, -1000, -1000, 90, 46, 23, 54, 46, -1000, 130,
	-10, 97, -1000, 116, -1000, 67, -1000, 152, -1000, -1000,
	-1000, -1000, -1000, 54, 54, -1000, -1000, 8, 7, -1000,
	150, -15, 46, -16, -1000, -1000, 154, -1000, -1000, -1000,
	-1000, -1000, 54, 46, 52, 152, 62, 46, 4, 46,
	-1000, -1000, -40, -1000, -1000, -4, -1000, -1000, -18, -1000,
	-19, -1000, -1000, 38, -1000, -34, -1000, -1000, -1000, -1000,
	-1000, -4, 29, 29, -1000, -1000, 46, 76, -4, -4,
	-4, -4, 3, -1000, -1000, -1000, 46, 38, 38, 38,
	-25, 38, -1000, -38, -1000, -4, -1000, 46, 38, 84,
	-30, -1000, 70, -1000, 102, 102, -1000, 129, -1000, -1000,
	158, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 183, 182, 181, 180, 0, 8, 15, 7, 2,
	156, 4, 179, 178, 177, 136, 5, 13, 176, 109,
	175, 174, 172, 10, 171, 94, 170, 9, 169, 1,
	168, 167, 164, 3, 163,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 6, 6, 5, 5, 7, 7,
	7, 8, 8, 8, 9, 11, 11, 12, 10, 13,
	13, 25, 25, 25, 25, 25, 25, 25, 26, 26,
	26, 27, 27, 27, 19, 19, 28, 28, 28, 24,
	24, 14, 14, 15, 18, 18, 18, 18, 22, 22,
	23, 23, 23, 21, 21, 20, 20, 16, 16, 17,
	29, 29, 29, 29, 29, 30, 30, 31, 31, 32,
	32, 33, 33, 33, 33, 33, 33, 33, 33, 33,
	33, 33, 34, 34, 34, 2, 3, 4,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	2, 1, 2, 2, 1, 1, 3, 3, 1, 1,
	3, 4, 1, 1, 2, 2, 2, 2, 0, 1,
	2, 5, 4, 9, 0, 2, 0, 3, 4, 0,
	1, 1, 3, 3, 0, 1, 4, 6, 0, 2,
	0, 1, 1, 0, 2, 0, 1, 1, 3, 3,
	2, 2, 1, 1, 2, 3, 3, 0, 2, 0,
	2, 1, 4, 1, 1, 1, 1, 1, 3, 3,
	3, 3, 0, 1, 3, 4, 8, 10,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, 34, 11, -13, -10, -9,
	-5, 41, 43, 36, -20, 37, 49, 16, -5, 19,
	-10, -5, 50, -5, -14, -15, -5, 27, 49, -28,
	-18, -5, -5, -19, -15, 10, 49, 51, -26, -25,
	29, 37, 26, 25, 9, 12, 50, 50, -27, 29,
	37, 15, -5, -19, -21, 40, -25, 21, 26, -5,
	-7, -6, 45, 47, 48, 41, 42, -7, -16, -17,
	-5, 21, 50, 21, -27, 33, -23, 6, 14, -7,
	-7, 51, 49, 49, 51, -22, 9, 50, -16, 50,
	-24, 7, -7, -17, -32, 39, -23, -6, -16, 51,
	-11, -9, 51, -33, 26, -5, -7, -8, 42, 41,
	43, 50, 47, 48, 46, 51, 49, 51, 44, 48,
	47, 50, -33, -8, -8, -9, 30, -33, -33, -33,
	-34, -33, 51, -5, 51, 49, -12, 50, -33, -31,
	-11, -30, 27, 51, 13, 38, -29, 35, 8, 32,
	24, -29, 26, 12, 4,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 0, 55, 0, 19, 18,
	14, 6, 7, 0, 0, 56, 0, 0, 0, 0,
	20, 85, 0, 0, 36, 41, 44, 0, 34, 0,
	28, 45, 0, 0, 42, 0, 34, 53, 43, 29,
	0, 22, 23, 0, 0, 0, 0, 0, 37, 0,
	0, 0, 35, 0, 86, 0, 30, 50, 24, 25,
	26, 27, 8, 0, 0, 4, 5, 0, 0, 57,
	48, 0, 0, 0, 38, 54, 39, 51, 52, 9,
	10, 46, 0, 0, 69, 50, 0, 0, 0, 0,
	21, 40, 0, 58, 87, 0, 59, 49, 0, 32,
	0, 15, 47, 70, 71, 0, 73, 74, 75, -2,
	-2, 0, 0, 0, 11, 31, 0, 0, 0, 0,
	0, 82, 0, 12, 13, 16, 0, 78, 79, 80,
	0, 83, 81, 0, 72, 0, 67, 0, 84, 33,
	0, 68, 0, 17, 0, 0, 65, 0, 62, 63,
	0, 66, 60, 61, 64,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	50, 51, 3, 48, 49, 47,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46,
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
		//line parser.go.y:114
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:117
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:122
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:125
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:130
		{
			yyVAL.signedNumber = yyDollar[1].signedNumber
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:133
		{
			yyVAL.signedNumber = -yyDollar[2].signedNumber
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:136
		{
			yyVAL.signedNumber = yyDollar[2].signedNumber
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:141
		{
			yyVAL.float = yyDollar[1].float
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:144
		{
			yyVAL.float = -yyDollar[2].float
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:147
		{
			yyVAL.float = yyDollar[2].float
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:152
		{
			yyVAL.columnName = yyDollar[1].identifier
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:157
		{
			yyVAL.columnNameList = []string{yyDollar[1].columnName}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:160
		{
			yyVAL.columnNameList = append(yyDollar[1].columnNameList, yyDollar[3].columnName)
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:165
		{
			yyVAL.columnNameList = yyDollar[2].columnNameList
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:170
		{
			yyVAL.columnName = yyDollar[1].columnName
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:175
		{
			yyVAL.columnNameList = []string{yyDollar[1].columnName}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:178
		{
			yyVAL.columnNameList = append(yyDollar[1].columnNameList, yyDollar[3].columnName)
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:184
		{
			yyVAL.columnConstraint = ccPrimaryKey{yyDollar[3].sortOrder, yyDollar[4].bool}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:187
		{
			yyVAL.columnConstraint = ccUnique(true)
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:190
		{
			yyVAL.columnConstraint = ccNull(true)
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:193
		{
			yyVAL.columnConstraint = ccNull(false)
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:196
		{
			yyVAL.columnConstraint = ccCollate(yyDollar[2].identifier)
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:199
		{
			yyVAL.columnConstraint = ccDefault(yyDollar[2].signedNumber)
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:202
		{
			yyVAL.columnConstraint = ccDefault(yyDollar[2].literal)
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:207
		{
			yyVAL.columnConstraintList = nil
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:210
		{
			yyVAL.columnConstraintList = []columnConstraint{yyDollar[1].columnConstraint}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:213
		{
			yyVAL.columnConstraintList = append(yyDollar[1].columnConstraintList, yyDollar[2].columnConstraint)
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:218
		{
			yyVAL.tableConstraint = TablePrimaryKey{yyDollar[4].indexedColumnList}
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:221
		{
			yyVAL.tableConstraint = TableUnique{yyDollar[3].indexedColumnList}
		}
	case 33:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:224
		{
			yyVAL.tableConstraint = TableForeignKey{
				Columns:        yyDollar[4].columnNameList,
				ForeignTable:   yyDollar[7].identifier,
				ForeignColumns: yyDollar[8].columnNameList,
				Triggers:       yyDollar[9].triggerList,
			}
		}
	case 34:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:234
		{
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:235
		{
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:239
		{
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:240
		{
			yyVAL.tableConstraintList = []TableConstraint{yyDollar[3].tableConstraint}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:243
		{
			yyVAL.tableConstraintList = append(yyDollar[1].tableConstraintList, yyDollar[4].tableConstraint)
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:249
		{
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:250
		{
			yyVAL.bool = true
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:255
		{
			yyVAL.columnDefList = []ColumnDef{yyDollar[1].columnDef}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:258
		{
			yyVAL.columnDefList = append(yyDollar[1].columnDefList, yyDollar[3].columnDef)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:263
		{
			yyVAL.columnDef = makeColumnDef(yyDollar[1].identifier, yyDollar[2].name, yyDollar[3].columnConstraintList)
		}
	case 44:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:268
		{
			yyVAL.name = ""
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:271
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 46:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:274
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 47:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:277
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:282
		{
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:283
		{
			yyVAL.collate = yyDollar[2].literal
		}
	case 50:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:288
		{
			yyVAL.sortOrder = Asc
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:291
		{
			yyVAL.sortOrder = Asc
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:294
		{
			yyVAL.sortOrder = Desc
		}
	case 53:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:299
		{
			yyVAL.withoutRowid = false
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:302
		{
			yyVAL.withoutRowid = true
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:307
		{
			yyVAL.unique = false
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:310
		{
			yyVAL.unique = true
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:315
		{
			yyVAL.indexedColumnList = []IndexedColumn{yyDollar[1].indexedColumn}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:318
		{
			yyVAL.indexedColumnList = append(yyDollar[1].indexedColumnList, yyDollar[3].indexedColumn)
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:323
		{
			yyVAL.indexedColumn = IndexedColumn{
				Column:    yyDollar[1].identifier,
				Collate:   yyDollar[2].collate,
				SortOrder: yyDollar[3].sortOrder,
			}
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:332
		{
			yyVAL.triggerAction = ActionSetNull
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:335
		{
			yyVAL.triggerAction = ActionSetDefault
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:338
		{
			yyVAL.triggerAction = ActionCascade
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:341
		{
			yyVAL.triggerAction = ActionRestrict
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:344
		{
			yyVAL.triggerAction = ActionNoAction
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:349
		{
			yyVAL.trigger = TriggerOnDelete(yyDollar[3].triggerAction)
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:352
		{
			yyVAL.trigger = TriggerOnUpdate(yyDollar[3].triggerAction)
		}
	case 67:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:357
		{
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:358
		{
			yyVAL.triggerList = append(yyDollar[1].triggerList, yyDollar[2].trigger)
		}
	case 69:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:363
		{
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:364
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:369
		{
			yyVAL.expr = nil
		}
	case 72:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:372
		{
			yyVAL.expr = ExFunction{yyDollar[1].identifier, yyDollar[3].exprList}
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:375
		{
			yyVAL.expr = yyDollar[1].signedNumber
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:378
		{
			yyVAL.expr = yyDollar[1].float
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:381
		{
			yyVAL.expr = yyDollar[1].identifier
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:384
		{
			yyVAL.expr = yyDollar[1].identifier
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:387
		{
			yyVAL.expr = ExColumn(yyDollar[1].identifier)
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:390
		{
			yyVAL.expr = ExBinaryOp{yyDollar[2].identifier, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:393
		{
			yyVAL.expr = ExBinaryOp{"+", yyDollar[1].expr, yyDollar[3].expr}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:396
		{
			yyVAL.expr = ExBinaryOp{"-", yyDollar[1].expr, yyDollar[3].expr}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:399
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 82:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:404
		{
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:405
		{
			yyVAL.exprList = []Expression{yyDollar[1].expr}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:408
		{
			yyVAL.exprList = append(yyDollar[1].exprList, yyDollar[3].expr)
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:414
		{
			yylex.(*lexer).result = SelectStmt{Columns: yyDollar[2].columnNameList, Table: yyDollar[4].identifier}
		}
	case 86:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:419
		{
			yylex.(*lexer).result = CreateTableStmt{
				Table:        yyDollar[3].identifier,
				Columns:      yyDollar[5].columnDefList,
				Constraints:  yyDollar[6].tableConstraintList,
				WithoutRowid: yyDollar[8].withoutRowid,
			}
		}
	case 87:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line parser.go.y:429
		{
			yylex.(*lexer).result = CreateIndexStmt{
				Index:          yyDollar[4].identifier,
				Table:          yyDollar[6].identifier,
				Unique:         yyDollar[2].unique,
				IndexedColumns: yyDollar[8].indexedColumnList,
				Where:          yyDollar[10].expr,
			}
		}
	}
	goto yystack /* stack new state and value */
}
