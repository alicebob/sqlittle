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
const ASC = 57347
const AUTOINCREMENT = 57348
const CASCADE = 57349
const COLLATE = 57350
const CONSTRAINT = 57351
const CREATE = 57352
const DEFAULT = 57353
const DELETE = 57354
const DESC = 57355
const FOREIGN = 57356
const FROM = 57357
const INDEX = 57358
const KEY = 57359
const NO = 57360
const NOT = 57361
const NULL = 57362
const ON = 57363
const PRIMARY = 57364
const REFERENCES = 57365
const RESTRICT = 57366
const ROWID = 57367
const SELECT = 57368
const SET = 57369
const TABLE = 57370
const UNIQUE = 57371
const UPDATE = 57372
const WHERE = 57373
const WITHOUT = 57374
const tBare = 57375
const tLiteral = 57376
const tIdentifier = 57377
const tOperator = 57378
const tSignedNumber = 57379

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ACTION",
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
	"INDEX",
	"KEY",
	"NO",
	"NOT",
	"NULL",
	"ON",
	"PRIMARY",
	"REFERENCES",
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
	62, 105, 28, 39, 62, 35, 55, 33, 60, 65,
	119, 11, 72, 12, 63, 64, 91, 51, 13, 15,
	25, 68, 16, 108, 73, 49, 117, 84, 120, 11,
	68, 12, 50, 27, 68, 123, 128, 93, 97, 92,
	94, 88, 56, 89, 53, 127, 125, 6, 71, 34,
	58, 69, 124, 57, 8, 122, 107, 106, 19, 111,
	75, 109, 44, 5, 82, 45, 97, 115, 76, 35,
	87, 20, 126, 43, 42, 129, 40, 90, 114, 116,
	29, 38, 86, 41, 81, 54, 14, 30, 24, 7,
	112, 4, 3, 2, 1,
}
var yyPact = [...]int{

	87, -1000, -1000, -1000, -1000, 28, 40, 34, -1000, -1000,
	-1000, -1000, -1000, 28, 92, -1000, 28, 28, 4, 28,
	-1000, -1000, 28, 62, 14, -1000, 28, 28, 46, 1,
	104, -5, -13, 53, -1000, 28, 110, 24, 104, -1000,
	86, -1000, -1000, 80, 28, 13, 17, 28, -1000, 84,
	-14, 81, -1000, 53, -1000, 49, -1000, 105, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 0, -3, -1000, 106, -18,
	28, -22, -1000, -1000, 114, -1000, -1000, -1000, 17, 28,
	35, 105, 31, 28, -7, 28, -1000, -1000, -30, -1000,
	-1000, 5, -1000, -1000, -8, -1000, -16, -1000, -1000, 15,
	-1000, 5, -1000, 28, 50, 5, -21, -1000, 28, 15,
	-1000, -23, -1000, 28, 55, -29, -1000, 48, -1000, 78,
	78, -1000, 75, -1000, -1000, 121, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 144, 143, 142, 141, 0, 4, 12, 2, 104,
	3, 140, 139, 138, 70, 6, 13, 137, 57, 136,
	135, 134, 7, 132, 53, 131, 8, 130, 1, 129,
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

	-1000, -1, -2, -3, -4, 26, 10, -12, -9, -8,
	-5, 33, 35, 28, -19, 29, 38, 15, -5, 16,
	-9, -5, 39, -5, -13, -14, -5, 21, 38, -27,
	-17, -5, -5, -18, -14, 9, 38, 40, -25, -24,
	22, 29, 20, 19, 8, 11, 39, 39, -26, 22,
	29, 14, -5, -18, -20, 32, -24, 17, 20, -5,
	-7, -6, 37, 33, 34, -7, -15, -16, -5, 17,
	39, 17, -26, 25, -22, 5, 13, 40, 38, 38,
	40, -21, 8, 39, -15, 39, -23, 6, -7, -16,
	-31, 31, -22, -6, -15, 40, -10, -8, 40, -32,
	-7, 39, 40, 38, 40, 36, -32, -8, 23, -32,
	40, -5, -11, 39, -30, -10, -29, 21, 40, 12,
	30, -28, 27, 7, 24, 18, -28, 20, 11, 4,
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
	39, 40, 3, 3, 38,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37,
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
		//line parser.go.y:102
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:105
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:110
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:113
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:118
		{
			yyVAL.signedNumber = yyDollar[1].signedNumber
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:123
		{
			yyVAL.columnName = yyDollar[1].identifier
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:128
		{
			yyVAL.columnNameList = []string{yyDollar[1].columnName}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:131
		{
			yyVAL.columnNameList = append(yyDollar[1].columnNameList, yyDollar[3].columnName)
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:136
		{
			yyVAL.columnNameList = yyDollar[2].columnNameList
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:141
		{
			yyVAL.columnName = yyDollar[1].columnName
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:146
		{
			yyVAL.columnNameList = []string{yyDollar[1].columnName}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:149
		{
			yyVAL.columnNameList = append(yyDollar[1].columnNameList, yyDollar[3].columnName)
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:155
		{
			yyVAL.columnConstraint = ccPrimaryKey{yyDollar[3].sortOrder, yyDollar[4].bool}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:158
		{
			yyVAL.columnConstraint = ccUnique(true)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:161
		{
			yyVAL.columnConstraint = ccNull(true)
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:164
		{
			yyVAL.columnConstraint = ccNull(false)
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:167
		{
			yyVAL.columnConstraint = ccCollate(yyDollar[2].identifier)
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:170
		{
			yyVAL.columnConstraint = ccDefault(yyDollar[2].signedNumber)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:173
		{
			yyVAL.columnConstraint = ccDefault(yyDollar[2].literal)
		}
	case 23:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:178
		{
			yyVAL.columnConstraintList = nil
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:181
		{
			yyVAL.columnConstraintList = []columnConstraint{yyDollar[1].columnConstraint}
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:184
		{
			yyVAL.columnConstraintList = append(yyDollar[1].columnConstraintList, yyDollar[2].columnConstraint)
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:189
		{
			yyVAL.tableConstraint = TablePrimaryKey{yyDollar[4].indexedColumnList}
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:192
		{
			yyVAL.tableConstraint = TableUnique{yyDollar[3].indexedColumnList}
		}
	case 28:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:195
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
		//line parser.go.y:205
		{
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:206
		{
		}
	case 31:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:210
		{
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:211
		{
			yyVAL.tableConstraintList = []TableConstraint{yyDollar[3].tableConstraint}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:214
		{
			yyVAL.tableConstraintList = append(yyDollar[1].tableConstraintList, yyDollar[4].tableConstraint)
		}
	case 34:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:220
		{
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:221
		{
			yyVAL.bool = true
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:226
		{
			yyVAL.columnDefList = []ColumnDef{yyDollar[1].columnDef}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:229
		{
			yyVAL.columnDefList = append(yyDollar[1].columnDefList, yyDollar[3].columnDef)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:234
		{
			yyVAL.columnDef = makeColumnDef(yyDollar[1].identifier, yyDollar[2].name, yyDollar[3].columnConstraintList)
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:239
		{
			yyVAL.name = ""
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:242
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:245
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 42:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:248
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 43:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:253
		{
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:254
		{
			yyVAL.collate = yyDollar[2].literal
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:259
		{
			yyVAL.sortOrder = Asc
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:262
		{
			yyVAL.sortOrder = Asc
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:265
		{
			yyVAL.sortOrder = Desc
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:270
		{
			yyVAL.withoutRowid = false
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:273
		{
			yyVAL.withoutRowid = true
		}
	case 50:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:278
		{
			yyVAL.unique = false
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:281
		{
			yyVAL.unique = true
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:286
		{
			yyVAL.indexedColumnList = []IndexedColumn{yyDollar[1].indexedColumn}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:289
		{
			yyVAL.indexedColumnList = append(yyDollar[1].indexedColumnList, yyDollar[3].indexedColumn)
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:294
		{
			yyVAL.indexedColumn = IndexedColumn{
				Column:    yyDollar[1].identifier,
				Collate:   yyDollar[2].collate,
				SortOrder: yyDollar[3].sortOrder,
			}
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:303
		{
			yyVAL.triggerAction = ActionSetNull
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:306
		{
			yyVAL.triggerAction = ActionSetDefault
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:309
		{
			yyVAL.triggerAction = ActionCascade
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:312
		{
			yyVAL.triggerAction = ActionRestrict
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:315
		{
			yyVAL.triggerAction = ActionNoAction
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:320
		{
			yyVAL.trigger = TriggerOnDelete(yyDollar[3].triggerAction)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:323
		{
			yyVAL.trigger = TriggerOnUpdate(yyDollar[3].triggerAction)
		}
	case 62:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:328
		{
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:329
		{
			yyVAL.triggerList = append(yyDollar[1].triggerList, yyDollar[2].trigger)
		}
	case 64:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:334
		{
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:335
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:340
		{
			yyVAL.expr = yyDollar[1].signedNumber
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:343
		{
			yyVAL.expr = ExBinaryOp{yyDollar[2].identifier, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:346
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:351
		{
			yylex.(*lexer).result = SelectStmt{Columns: yyDollar[2].columnNameList, Table: yyDollar[4].identifier}
		}
	case 70:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:356
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
		//line parser.go.y:366
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
