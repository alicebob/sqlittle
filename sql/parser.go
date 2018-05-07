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
	where                Expression
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
	-1, 77,
	50, 6,
	-2, 77,
	-1, 78,
	50, 7,
	-2, 78,
}

const yyPrivate = 57344

const yyLast = 182

var yyAct = [...]int{

	73, 147, 9, 125, 88, 71, 10, 68, 74, 48,
	127, 61, 75, 132, 18, 144, 69, 10, 21, 132,
	23, 133, 72, 26, 95, 17, 131, 31, 32, 26,
	130, 95, 129, 124, 139, 108, 52, 77, 76, 78,
	106, 62, 82, 80, 81, 59, 79, 102, 95, 99,
	96, 84, 101, 100, 60, 67, 122, 94, 16, 93,
	65, 66, 47, 86, 62, 46, 63, 64, 22, 36,
	28, 37, 91, 92, 62, 82, 80, 81, 99, 65,
	66, 101, 100, 55, 62, 103, 63, 64, 35, 91,
	92, 114, 107, 104, 105, 11, 87, 12, 145, 39,
	149, 136, 115, 111, 51, 117, 118, 119, 121, 10,
	116, 126, 112, 33, 123, 25, 151, 143, 49, 11,
	128, 12, 27, 146, 150, 154, 50, 148, 6, 44,
	13, 15, 45, 10, 58, 135, 134, 137, 56, 153,
	10, 19, 126, 141, 34, 43, 42, 85, 152, 40,
	53, 5, 83, 57, 8, 89, 35, 41, 110, 98,
	155, 120, 70, 90, 113, 140, 142, 29, 38, 109,
	97, 20, 54, 14, 30, 24, 7, 138, 4, 3,
	2, 1,
}
var yyPact = [...]int{

	117, -1000, -1000, -1000, -1000, 54, 94, 9, -1000, -1000,
	-1000, -1000, -1000, 54, 122, -1000, 54, 54, 18, 54,
	-1000, -1000, 54, 95, 21, -1000, 54, 54, 78, 20,
	120, 15, 12, 89, -1000, 54, 146, 43, 120, -1000,
	132, -1000, -1000, 108, 54, 19, 39, -4, -1000, 131,
	1, 126, -1000, 89, -1000, 63, -1000, 149, -1000, -1000,
	-1000, -1000, -1000, 39, 39, -1000, -1000, 8, -1, -1000,
	150, 34, -1000, -3, -1000, -1000, -1000, -1000, -1000, -4,
	29, 29, -1000, -10, -4, -15, -1000, -1000, 151, -1000,
	-1000, -1000, -1000, -1000, 39, -4, 52, 149, 38, -4,
	-4, -4, -4, 5, -1000, -1000, -4, -18, 54, -1000,
	-1000, -41, -1000, -1000, -4, -1000, -1000, 34, 34, 34,
	-19, 34, -1000, -25, -1000, -30, -1000, -1000, 34, -1000,
	-4, -1000, 54, 71, 34, -1000, 54, -16, -1000, 54,
	90, -36, -1000, 85, -1000, 92, 92, -1000, 113, -1000,
	-1000, 156, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 181, 180, 179, 178, 0, 11, 8, 12, 2,
	154, 3, 177, 176, 175, 115, 7, 16, 174, 113,
	173, 172, 170, 4, 169, 99, 168, 9, 167, 1,
	166, 165, 164, 162, 5, 161,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 6, 6, 5, 5, 7, 7,
	7, 8, 8, 8, 9, 11, 11, 12, 10, 13,
	13, 25, 25, 25, 25, 25, 25, 25, 26, 26,
	26, 27, 27, 27, 19, 19, 28, 28, 28, 24,
	24, 14, 14, 15, 18, 18, 18, 18, 22, 22,
	23, 23, 23, 21, 21, 20, 20, 16, 16, 33,
	17, 29, 29, 29, 29, 29, 30, 30, 31, 31,
	32, 32, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 35, 35, 35, 2, 3, 4,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	2, 1, 2, 2, 1, 1, 3, 3, 1, 1,
	3, 4, 1, 1, 2, 2, 2, 2, 0, 1,
	2, 5, 4, 9, 0, 2, 0, 3, 4, 0,
	1, 1, 3, 3, 0, 1, 4, 6, 0, 2,
	0, 1, 1, 0, 2, 0, 1, 1, 3, 1,
	3, 2, 2, 1, 1, 2, 3, 3, 0, 2,
	0, 2, 1, 4, 1, 1, 1, 1, 1, 3,
	3, 3, 3, 0, 1, 3, 4, 8, 10,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, 34, 11, -13, -10, -9,
	-5, 41, 43, 36, -20, 37, 49, 16, -5, 19,
	-10, -5, 50, -5, -14, -15, -5, 27, 49, -28,
	-18, -5, -5, -19, -15, 10, 49, 51, -26, -25,
	29, 37, 26, 25, 9, 12, 50, 50, -27, 29,
	37, 15, -5, -19, -21, 40, -25, 21, 26, -5,
	-7, -6, 45, 47, 48, 41, 42, -7, -16, -17,
	-33, -34, 26, -5, -7, -8, 42, 41, 43, 50,
	47, 48, 46, 21, 50, 21, -27, 33, -23, 6,
	14, -7, -7, 51, 49, 49, 51, -22, 9, 44,
	48, 47, 50, -34, -8, -8, 50, -16, 50, -24,
	7, -7, -17, -32, 39, -23, -6, -34, -34, -34,
	-35, -34, 51, -16, 51, -11, -9, 51, -34, 51,
	49, 51, 49, 51, -34, -9, 30, -5, -12, 50,
	-31, -11, -30, 27, 51, 13, 38, -29, 35, 8,
	32, 24, -29, 26, 12, 4,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 0, 55, 0, 19, 18,
	14, 6, 7, 0, 0, 56, 0, 0, 0, 0,
	20, 86, 0, 0, 36, 41, 44, 0, 34, 0,
	28, 45, 0, 0, 42, 0, 34, 53, 43, 29,
	0, 22, 23, 0, 0, 0, 0, 0, 37, 0,
	0, 0, 35, 0, 87, 0, 30, 50, 24, 25,
	26, 27, 8, 0, 0, 4, 5, 0, 0, 57,
	48, 59, 72, 0, 74, 75, 76, -2, -2, 0,
	0, 0, 11, 0, 0, 0, 38, 54, 39, 51,
	52, 9, 10, 46, 0, 0, 70, 50, 0, 0,
	0, 0, 83, 0, 12, 13, 0, 0, 0, 21,
	40, 0, 58, 88, 0, 60, 49, 79, 80, 81,
	0, 84, 82, 0, 32, 0, 15, 47, 71, 73,
	0, 31, 0, 0, 85, 16, 0, 0, 68, 0,
	33, 0, 69, 0, 17, 0, 0, 66, 0, 63,
	64, 0, 67, 61, 62, 65,
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
		//line parser.go.y:117
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:120
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:125
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:128
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:133
		{
			yyVAL.signedNumber = yyDollar[1].signedNumber
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:136
		{
			yyVAL.signedNumber = -yyDollar[2].signedNumber
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:139
		{
			yyVAL.signedNumber = yyDollar[2].signedNumber
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:144
		{
			yyVAL.float = yyDollar[1].float
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:147
		{
			yyVAL.float = -yyDollar[2].float
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:150
		{
			yyVAL.float = yyDollar[2].float
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:155
		{
			yyVAL.columnName = yyDollar[1].identifier
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:160
		{
			yyVAL.columnNameList = []string{yyDollar[1].columnName}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:163
		{
			yyVAL.columnNameList = append(yyDollar[1].columnNameList, yyDollar[3].columnName)
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:168
		{
			yyVAL.columnNameList = yyDollar[2].columnNameList
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:173
		{
			yyVAL.columnName = yyDollar[1].columnName
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:178
		{
			yyVAL.columnNameList = []string{yyDollar[1].columnName}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:181
		{
			yyVAL.columnNameList = append(yyDollar[1].columnNameList, yyDollar[3].columnName)
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:187
		{
			yyVAL.columnConstraint = ccPrimaryKey{yyDollar[3].sortOrder, yyDollar[4].bool}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:190
		{
			yyVAL.columnConstraint = ccUnique(true)
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:193
		{
			yyVAL.columnConstraint = ccNull(true)
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:196
		{
			yyVAL.columnConstraint = ccNull(false)
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:199
		{
			yyVAL.columnConstraint = ccCollate(yyDollar[2].identifier)
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:202
		{
			yyVAL.columnConstraint = ccDefault(yyDollar[2].signedNumber)
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:205
		{
			yyVAL.columnConstraint = ccDefault(yyDollar[2].literal)
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:210
		{
			yyVAL.columnConstraintList = nil
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:213
		{
			yyVAL.columnConstraintList = []columnConstraint{yyDollar[1].columnConstraint}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:216
		{
			yyVAL.columnConstraintList = append(yyDollar[1].columnConstraintList, yyDollar[2].columnConstraint)
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:221
		{
			yyVAL.tableConstraint = TablePrimaryKey{yyDollar[4].indexedColumnList}
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:224
		{
			yyVAL.tableConstraint = TableUnique{yyDollar[3].indexedColumnList}
		}
	case 33:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:227
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
		//line parser.go.y:237
		{
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:238
		{
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:242
		{
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:243
		{
			yyVAL.tableConstraintList = []TableConstraint{yyDollar[3].tableConstraint}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:246
		{
			yyVAL.tableConstraintList = append(yyDollar[1].tableConstraintList, yyDollar[4].tableConstraint)
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:252
		{
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:253
		{
			yyVAL.bool = true
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:258
		{
			yyVAL.columnDefList = []ColumnDef{yyDollar[1].columnDef}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:261
		{
			yyVAL.columnDefList = append(yyDollar[1].columnDefList, yyDollar[3].columnDef)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:266
		{
			yyVAL.columnDef = makeColumnDef(yyDollar[1].identifier, yyDollar[2].name, yyDollar[3].columnConstraintList)
		}
	case 44:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:271
		{
			yyVAL.name = ""
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:274
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 46:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:277
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 47:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:280
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:285
		{
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:286
		{
			yyVAL.collate = yyDollar[2].literal
		}
	case 50:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:291
		{
			yyVAL.sortOrder = Asc
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:294
		{
			yyVAL.sortOrder = Asc
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:297
		{
			yyVAL.sortOrder = Desc
		}
	case 53:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:302
		{
			yyVAL.withoutRowid = false
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:305
		{
			yyVAL.withoutRowid = true
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:310
		{
			yyVAL.unique = false
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:313
		{
			yyVAL.unique = true
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:318
		{
			yyVAL.indexedColumnList = []IndexedColumn{yyDollar[1].indexedColumn}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:321
		{
			yyVAL.indexedColumnList = append(yyDollar[1].indexedColumnList, yyDollar[3].indexedColumn)
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:326
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:331
		{
			yyVAL.indexedColumn = newIndexColumn(yyDollar[1].expr, yyDollar[2].collate, yyDollar[3].sortOrder)
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:336
		{
			yyVAL.triggerAction = ActionSetNull
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:339
		{
			yyVAL.triggerAction = ActionSetDefault
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:342
		{
			yyVAL.triggerAction = ActionCascade
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:345
		{
			yyVAL.triggerAction = ActionRestrict
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:348
		{
			yyVAL.triggerAction = ActionNoAction
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:353
		{
			yyVAL.trigger = TriggerOnDelete(yyDollar[3].triggerAction)
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:356
		{
			yyVAL.trigger = TriggerOnUpdate(yyDollar[3].triggerAction)
		}
	case 68:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:361
		{
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:362
		{
			yyVAL.triggerList = append(yyDollar[1].triggerList, yyDollar[2].trigger)
		}
	case 70:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:367
		{
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:368
		{
			yyVAL.where = yyDollar[2].expr
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:373
		{
			yyVAL.expr = nil
		}
	case 73:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:376
		{
			yyVAL.expr = ExFunction{yyDollar[1].identifier, yyDollar[3].exprList}
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:379
		{
			yyVAL.expr = yyDollar[1].signedNumber
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:382
		{
			yyVAL.expr = yyDollar[1].float
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:385
		{
			yyVAL.expr = yyDollar[1].identifier
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:388
		{
			yyVAL.expr = ExColumn(yyDollar[1].identifier)
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:391
		{
			yyVAL.expr = ExColumn(yyDollar[1].identifier)
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:394
		{
			yyVAL.expr = ExBinaryOp{yyDollar[2].identifier, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:397
		{
			yyVAL.expr = ExBinaryOp{"+", yyDollar[1].expr, yyDollar[3].expr}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:400
		{
			yyVAL.expr = ExBinaryOp{"-", yyDollar[1].expr, yyDollar[3].expr}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:403
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 83:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:408
		{
			yyVAL.exprList = nil
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:411
		{
			yyVAL.exprList = []Expression{yyDollar[1].expr}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:414
		{
			yyVAL.exprList = append(yyDollar[1].exprList, yyDollar[3].expr)
		}
	case 86:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:419
		{
			yylex.(*lexer).result = SelectStmt{Columns: yyDollar[2].columnNameList, Table: yyDollar[4].identifier}
		}
	case 87:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:424
		{
			yylex.(*lexer).result = CreateTableStmt{
				Table:        yyDollar[3].identifier,
				Columns:      yyDollar[5].columnDefList,
				Constraints:  yyDollar[6].tableConstraintList,
				WithoutRowid: yyDollar[8].withoutRowid,
			}
		}
	case 88:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line parser.go.y:434
		{
			yylex.(*lexer).result = CreateIndexStmt{
				Index:          yyDollar[4].identifier,
				Table:          yyDollar[6].identifier,
				Unique:         yyDollar[2].unique,
				IndexedColumns: yyDollar[8].indexedColumnList,
				Where:          yyDollar[10].where,
			}
		}
	}
	goto yystack /* stack new state and value */
}
