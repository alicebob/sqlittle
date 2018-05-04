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

const yyLast = 145

var yyAct = [...]int{

	10, 121, 9, 96, 61, 99, 66, 74, 48, 103,
	98, 118, 100, 67, 18, 105, 113, 85, 21, 110,
	23, 83, 103, 26, 104, 70, 47, 31, 32, 26,
	79, 79, 102, 95, 46, 79, 52, 80, 78, 36,
	77, 37, 62, 22, 101, 59, 63, 64, 68, 17,
	62, 105, 28, 62, 11, 55, 12, 91, 60, 65,
	73, 35, 72, 63, 64, 39, 13, 15, 108, 117,
	119, 68, 128, 27, 25, 58, 33, 84, 6, 16,
	68, 19, 44, 51, 68, 45, 127, 93, 97, 92,
	94, 88, 11, 89, 12, 120, 123, 49, 43, 42,
	71, 5, 40, 34, 56, 50, 107, 106, 69, 111,
	41, 109, 125, 53, 57, 8, 97, 115, 75, 35,
	124, 87, 126, 122, 82, 129, 76, 90, 114, 116,
	29, 38, 20, 86, 81, 54, 14, 30, 24, 7,
	112, 4, 3, 2, 1,
}
var yyPact = [...]int{

	67, -1000, -1000, -1000, -1000, 13, 30, 33, -1000, -1000,
	-1000, -1000, -1000, 13, 62, -1000, 13, 13, -4, 13,
	-1000, -1000, 13, 46, 6, -1000, 13, 13, 51, -7,
	73, -13, -21, 68, -1000, 13, 109, 15, 73, -1000,
	93, -1000, -1000, 49, 13, 5, 8, 13, -1000, 87,
	-22, 79, -1000, 68, -1000, 27, -1000, 112, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -8, -11, -1000, 115, -26,
	13, -30, -1000, -1000, 114, -1000, -1000, -1000, 8, 13,
	18, 112, 22, 13, -15, 13, -1000, -1000, -38, -1000,
	-1000, -3, -1000, -1000, -16, -1000, -24, -1000, -1000, 7,
	-1000, -3, -1000, 13, 38, -3, -29, -1000, 13, 7,
	-1000, -31, -1000, 13, 42, -37, -1000, 57, -1000, 88,
	88, -1000, 60, -1000, -1000, 121, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 144, 143, 142, 141, 0, 4, 12, 2, 115,
	3, 140, 139, 138, 74, 6, 13, 137, 76, 136,
	135, 134, 7, 133, 65, 131, 8, 130, 1, 129,
	128, 127, 5,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 6, 6, 5, 5, 7, 8,
	10, 10, 11, 9, 12, 12, 24, 24, 24, 24,
	24, 24, 24, 25, 25, 25, 26, 26, 26, 18,
	18, 27, 27, 27, 23, 23, 13, 13, 14, 17,
	17, 17, 17, 21, 21, 22, 22, 22, 20, 20,
	19, 19, 15, 15, 16, 28, 28, 28, 28, 28,
	29, 29, 30, 30, 31, 31, 32, 32, 32, 2,
	3, 4,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 3, 1, 1, 3, 4, 1, 1, 2,
	2, 2, 2, 0, 1, 2, 5, 4, 9, 0,
	2, 0, 3, 4, 0, 1, 1, 3, 3, 0,
	1, 4, 6, 0, 2, 0, 1, 1, 0, 2,
	0, 1, 1, 3, 3, 2, 2, 1, 1, 2,
	3, 3, 0, 2, 0, 2, 1, 3, 3, 4,
	8, 10,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, 34, 11, -12, -9, -8,
	-5, 41, 43, 36, -19, 37, 46, 16, -5, 19,
	-9, -5, 47, -5, -13, -14, -5, 27, 46, -27,
	-17, -5, -5, -18, -14, 10, 46, 48, -25, -24,
	29, 37, 26, 25, 9, 12, 47, 47, -26, 29,
	37, 15, -5, -18, -20, 40, -24, 21, 26, -5,
	-7, -6, 45, 41, 42, -7, -15, -16, -5, 21,
	47, 21, -26, 33, -22, 6, 14, 48, 46, 46,
	48, -21, 9, 47, -15, 47, -23, 7, -7, -16,
	-31, 39, -22, -6, -15, 48, -10, -8, 48, -32,
	-7, 47, 48, 46, 48, 44, -32, -8, 30, -32,
	48, -5, -11, 47, -30, -10, -29, 27, 48, 13,
	38, -28, 35, 8, 32, 24, -28, 26, 12, 4,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 0, 50, 0, 14, 13,
	9, 6, 7, 0, 0, 51, 0, 0, 0, 0,
	15, 69, 0, 0, 31, 36, 39, 0, 29, 0,
	23, 40, 0, 0, 37, 0, 29, 48, 38, 24,
	0, 17, 18, 0, 0, 0, 0, 0, 32, 0,
	0, 0, 30, 0, 70, 0, 25, 45, 19, 20,
	21, 22, 8, 4, 5, 0, 0, 52, 43, 0,
	0, 0, 33, 49, 34, 46, 47, 41, 0, 0,
	64, 45, 0, 0, 0, 0, 16, 35, 0, 53,
	71, 0, 54, 44, 0, 27, 0, 10, 42, 65,
	66, 0, 26, 0, 0, 0, 0, 11, 0, 67,
	68, 0, 62, 0, 28, 0, 63, 0, 12, 0,
	0, 60, 0, 57, 58, 0, 61, 55, 56, 59,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	47, 48, 3, 3, 46,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45,
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
		//line parser.go.y:110
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:113
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:118
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:121
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:126
		{
			yyVAL.signedNumber = yyDollar[1].signedNumber
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:131
		{
			yyVAL.columnName = yyDollar[1].identifier
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:136
		{
			yyVAL.columnNameList = []string{yyDollar[1].columnName}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:139
		{
			yyVAL.columnNameList = append(yyDollar[1].columnNameList, yyDollar[3].columnName)
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:144
		{
			yyVAL.columnNameList = yyDollar[2].columnNameList
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:149
		{
			yyVAL.columnName = yyDollar[1].columnName
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:154
		{
			yyVAL.columnNameList = []string{yyDollar[1].columnName}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:157
		{
			yyVAL.columnNameList = append(yyDollar[1].columnNameList, yyDollar[3].columnName)
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:163
		{
			yyVAL.columnConstraint = ccPrimaryKey{yyDollar[3].sortOrder, yyDollar[4].bool}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:166
		{
			yyVAL.columnConstraint = ccUnique(true)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:169
		{
			yyVAL.columnConstraint = ccNull(true)
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:172
		{
			yyVAL.columnConstraint = ccNull(false)
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:175
		{
			yyVAL.columnConstraint = ccCollate(yyDollar[2].identifier)
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:178
		{
			yyVAL.columnConstraint = ccDefault(yyDollar[2].signedNumber)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:181
		{
			yyVAL.columnConstraint = ccDefault(yyDollar[2].literal)
		}
	case 23:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:186
		{
			yyVAL.columnConstraintList = nil
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:189
		{
			yyVAL.columnConstraintList = []columnConstraint{yyDollar[1].columnConstraint}
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:192
		{
			yyVAL.columnConstraintList = append(yyDollar[1].columnConstraintList, yyDollar[2].columnConstraint)
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:197
		{
			yyVAL.tableConstraint = TablePrimaryKey{yyDollar[4].indexedColumnList}
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:200
		{
			yyVAL.tableConstraint = TableUnique{yyDollar[3].indexedColumnList}
		}
	case 28:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:203
		{
			yyVAL.tableConstraint = TableForeignKey{
				Columns:        yyDollar[4].columnNameList,
				ForeignTable:   yyDollar[7].identifier,
				ForeignColumns: yyDollar[8].columnNameList,
				Triggers:       yyDollar[9].triggerList,
			}
		}
	case 29:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:213
		{
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:214
		{
		}
	case 31:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:218
		{
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:219
		{
			yyVAL.tableConstraintList = []TableConstraint{yyDollar[3].tableConstraint}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:222
		{
			yyVAL.tableConstraintList = append(yyDollar[1].tableConstraintList, yyDollar[4].tableConstraint)
		}
	case 34:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:228
		{
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:229
		{
			yyVAL.bool = true
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:234
		{
			yyVAL.columnDefList = []ColumnDef{yyDollar[1].columnDef}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:237
		{
			yyVAL.columnDefList = append(yyDollar[1].columnDefList, yyDollar[3].columnDef)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:242
		{
			yyVAL.columnDef = makeColumnDef(yyDollar[1].identifier, yyDollar[2].name, yyDollar[3].columnConstraintList)
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:247
		{
			yyVAL.name = ""
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:250
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:253
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 42:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:256
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 43:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:261
		{
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:262
		{
			yyVAL.collate = yyDollar[2].literal
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:267
		{
			yyVAL.sortOrder = Asc
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:270
		{
			yyVAL.sortOrder = Asc
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:273
		{
			yyVAL.sortOrder = Desc
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:278
		{
			yyVAL.withoutRowid = false
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:281
		{
			yyVAL.withoutRowid = true
		}
	case 50:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:286
		{
			yyVAL.unique = false
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:289
		{
			yyVAL.unique = true
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:294
		{
			yyVAL.indexedColumnList = []IndexedColumn{yyDollar[1].indexedColumn}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:297
		{
			yyVAL.indexedColumnList = append(yyDollar[1].indexedColumnList, yyDollar[3].indexedColumn)
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:302
		{
			yyVAL.indexedColumn = IndexedColumn{
				Column:    yyDollar[1].identifier,
				Collate:   yyDollar[2].collate,
				SortOrder: yyDollar[3].sortOrder,
			}
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:311
		{
			yyVAL.triggerAction = ActionSetNull
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:314
		{
			yyVAL.triggerAction = ActionSetDefault
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:317
		{
			yyVAL.triggerAction = ActionCascade
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:320
		{
			yyVAL.triggerAction = ActionRestrict
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:323
		{
			yyVAL.triggerAction = ActionNoAction
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:328
		{
			yyVAL.trigger = TriggerOnDelete(yyDollar[3].triggerAction)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:331
		{
			yyVAL.trigger = TriggerOnUpdate(yyDollar[3].triggerAction)
		}
	case 62:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:336
		{
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:337
		{
			yyVAL.triggerList = append(yyDollar[1].triggerList, yyDollar[2].trigger)
		}
	case 64:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:342
		{
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:343
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:348
		{
			yyVAL.expr = yyDollar[1].signedNumber
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:351
		{
			yyVAL.expr = ExBinaryOp{yyDollar[2].identifier, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:354
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:359
		{
			yylex.(*lexer).result = SelectStmt{Columns: yyDollar[2].columnNameList, Table: yyDollar[4].identifier}
		}
	case 70:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:364
		{
			yylex.(*lexer).result = CreateTableStmt{
				Table:        yyDollar[3].identifier,
				Columns:      yyDollar[5].columnDefList,
				Constraints:  yyDollar[6].tableConstraintList,
				WithoutRowid: yyDollar[8].withoutRowid,
			}
		}
	case 71:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line parser.go.y:374
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
