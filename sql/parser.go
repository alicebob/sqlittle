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
	105, 118, 100, 67, 18, 110, 98, 103, 21, 104,
	23, 113, 79, 26, 102, 85, 83, 31, 32, 26,
	79, 79, 95, 80, 70, 78, 52, 77, 36, 62,
	37, 47, 101, 46, 22, 59, 28, 17, 68, 63,
	64, 62, 35, 62, 11, 105, 12, 55, 60, 65,
	63, 64, 72, 91, 39, 13, 15, 119, 73, 108,
	117, 68, 51, 33, 27, 25, 6, 84, 16, 58,
	68, 123, 128, 11, 68, 12, 49, 93, 97, 92,
	94, 88, 120, 89, 50, 71, 127, 125, 44, 5,
	69, 45, 57, 56, 34, 124, 107, 106, 122, 111,
	53, 109, 19, 8, 43, 42, 97, 115, 40, 75,
	35, 82, 126, 87, 129, 90, 41, 76, 114, 116,
	20, 29, 38, 86, 81, 54, 14, 30, 24, 7,
	112, 4, 3, 2, 1,
}
var yyPact = [...]int{

	65, -1000, -1000, -1000, -1000, 13, 29, 31, -1000, -1000,
	-1000, -1000, -1000, 13, 93, -1000, 13, 13, -4, 13,
	-1000, -1000, 13, 47, -1, -1000, 13, 13, 42, -9,
	89, -5, -7, 57, -1000, 13, 110, 17, 89, -1000,
	81, -1000, -1000, 53, 13, 8, 6, 13, -1000, 79,
	-14, 74, -1000, 57, -1000, 35, -1000, 113, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -12, -16, -1000, 112, -22,
	13, -23, -1000, -1000, 116, -1000, -1000, -1000, 6, 13,
	24, 113, 19, 13, -17, 13, -1000, -1000, -33, -1000,
	-1000, -6, -1000, -1000, -25, -1000, -30, -1000, -1000, 11,
	-1000, -6, -1000, 13, 39, -6, -34, -1000, 13, 11,
	-1000, -27, -1000, 13, 43, -38, -1000, 54, -1000, 73,
	73, -1000, 70, -1000, -1000, 120, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 144, 143, 142, 141, 0, 4, 12, 2, 113,
	3, 140, 139, 138, 75, 6, 13, 137, 73, 136,
	135, 134, 7, 133, 64, 132, 8, 131, 1, 129,
	128, 125, 5,
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
	-5, 41, 43, 36, -19, 37, 47, 16, -5, 19,
	-9, -5, 48, -5, -13, -14, -5, 27, 47, -27,
	-17, -5, -5, -18, -14, 10, 47, 49, -25, -24,
	29, 37, 26, 25, 9, 12, 48, 48, -26, 29,
	37, 15, -5, -18, -20, 40, -24, 21, 26, -5,
	-7, -6, 45, 41, 42, -7, -15, -16, -5, 21,
	48, 21, -26, 33, -22, 6, 14, 49, 47, 47,
	49, -21, 9, 48, -15, 48, -23, 7, -7, -16,
	-31, 39, -22, -6, -15, 49, -10, -8, 49, -32,
	-7, 48, 49, 47, 49, 44, -32, -8, 30, -32,
	49, -5, -11, 48, -30, -10, -29, 27, 49, 13,
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
	48, 49, 3, 3, 47,
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
		//line parser.go.y:111
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:114
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:119
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:122
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:127
		{
			yyVAL.signedNumber = yyDollar[1].signedNumber
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:132
		{
			yyVAL.columnName = yyDollar[1].identifier
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:137
		{
			yyVAL.columnNameList = []string{yyDollar[1].columnName}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:140
		{
			yyVAL.columnNameList = append(yyDollar[1].columnNameList, yyDollar[3].columnName)
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:145
		{
			yyVAL.columnNameList = yyDollar[2].columnNameList
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:150
		{
			yyVAL.columnName = yyDollar[1].columnName
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:155
		{
			yyVAL.columnNameList = []string{yyDollar[1].columnName}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:158
		{
			yyVAL.columnNameList = append(yyDollar[1].columnNameList, yyDollar[3].columnName)
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:164
		{
			yyVAL.columnConstraint = ccPrimaryKey{yyDollar[3].sortOrder, yyDollar[4].bool}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:167
		{
			yyVAL.columnConstraint = ccUnique(true)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:170
		{
			yyVAL.columnConstraint = ccNull(true)
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:173
		{
			yyVAL.columnConstraint = ccNull(false)
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:176
		{
			yyVAL.columnConstraint = ccCollate(yyDollar[2].identifier)
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:179
		{
			yyVAL.columnConstraint = ccDefault(yyDollar[2].signedNumber)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:182
		{
			yyVAL.columnConstraint = ccDefault(yyDollar[2].literal)
		}
	case 23:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:187
		{
			yyVAL.columnConstraintList = nil
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:190
		{
			yyVAL.columnConstraintList = []columnConstraint{yyDollar[1].columnConstraint}
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:193
		{
			yyVAL.columnConstraintList = append(yyDollar[1].columnConstraintList, yyDollar[2].columnConstraint)
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:198
		{
			yyVAL.tableConstraint = TablePrimaryKey{yyDollar[4].indexedColumnList}
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:201
		{
			yyVAL.tableConstraint = TableUnique{yyDollar[3].indexedColumnList}
		}
	case 28:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:204
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
		//line parser.go.y:214
		{
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:215
		{
		}
	case 31:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:219
		{
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:220
		{
			yyVAL.tableConstraintList = []TableConstraint{yyDollar[3].tableConstraint}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:223
		{
			yyVAL.tableConstraintList = append(yyDollar[1].tableConstraintList, yyDollar[4].tableConstraint)
		}
	case 34:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:229
		{
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:230
		{
			yyVAL.bool = true
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:235
		{
			yyVAL.columnDefList = []ColumnDef{yyDollar[1].columnDef}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:238
		{
			yyVAL.columnDefList = append(yyDollar[1].columnDefList, yyDollar[3].columnDef)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:243
		{
			yyVAL.columnDef = makeColumnDef(yyDollar[1].identifier, yyDollar[2].name, yyDollar[3].columnConstraintList)
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:248
		{
			yyVAL.name = ""
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:251
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:254
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 42:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:257
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 43:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:262
		{
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:263
		{
			yyVAL.collate = yyDollar[2].literal
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:268
		{
			yyVAL.sortOrder = Asc
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:271
		{
			yyVAL.sortOrder = Asc
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:274
		{
			yyVAL.sortOrder = Desc
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:279
		{
			yyVAL.withoutRowid = false
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:282
		{
			yyVAL.withoutRowid = true
		}
	case 50:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:287
		{
			yyVAL.unique = false
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:290
		{
			yyVAL.unique = true
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:295
		{
			yyVAL.indexedColumnList = []IndexedColumn{yyDollar[1].indexedColumn}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:298
		{
			yyVAL.indexedColumnList = append(yyDollar[1].indexedColumnList, yyDollar[3].indexedColumn)
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:303
		{
			yyVAL.indexedColumn = IndexedColumn{
				Column:    yyDollar[1].identifier,
				Collate:   yyDollar[2].collate,
				SortOrder: yyDollar[3].sortOrder,
			}
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:312
		{
			yyVAL.triggerAction = ActionSetNull
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:315
		{
			yyVAL.triggerAction = ActionSetDefault
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:318
		{
			yyVAL.triggerAction = ActionCascade
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:321
		{
			yyVAL.triggerAction = ActionRestrict
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:324
		{
			yyVAL.triggerAction = ActionNoAction
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:329
		{
			yyVAL.trigger = TriggerOnDelete(yyDollar[3].triggerAction)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:332
		{
			yyVAL.trigger = TriggerOnUpdate(yyDollar[3].triggerAction)
		}
	case 62:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:337
		{
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:338
		{
			yyVAL.triggerList = append(yyDollar[1].triggerList, yyDollar[2].trigger)
		}
	case 64:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:343
		{
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:344
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:349
		{
			yyVAL.expr = yyDollar[1].signedNumber
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:352
		{
			yyVAL.expr = ExBinaryOp{yyDollar[2].identifier, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:355
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:360
		{
			yylex.(*lexer).result = SelectStmt{Columns: yyDollar[2].columnNameList, Table: yyDollar[4].identifier}
		}
	case 70:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:365
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
		//line parser.go.y:375
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
