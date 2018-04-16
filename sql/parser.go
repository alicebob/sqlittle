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
const WITHOUT = 57373
const tBare = 57374
const tLiteral = 57375
const tIdentifier = 57376
const tSignedNumber = 57377

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
	"WITHOUT",
	"tBare",
	"tLiteral",
	"tIdentifier",
	"tSignedNumber",
	"','",
	"'('",
	"')'",
	"'*'",
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

const yyLast = 137

var yyAct = [...]int{

	11, 113, 9, 95, 62, 67, 68, 49, 75, 12,
	40, 13, 61, 97, 34, 19, 10, 105, 99, 22,
	110, 24, 86, 99, 27, 100, 84, 71, 32, 33,
	27, 80, 80, 98, 94, 48, 80, 53, 81, 79,
	37, 78, 38, 47, 64, 65, 60, 63, 23, 69,
	57, 29, 54, 18, 12, 36, 13, 111, 63, 56,
	66, 52, 73, 64, 65, 14, 16, 74, 26, 50,
	120, 102, 69, 109, 17, 112, 51, 85, 12, 119,
	13, 69, 28, 45, 6, 69, 46, 90, 92, 96,
	93, 91, 89, 59, 44, 43, 115, 41, 35, 8,
	5, 72, 101, 103, 42, 70, 20, 117, 96, 107,
	58, 76, 36, 116, 118, 83, 114, 21, 88, 77,
	121, 106, 108, 30, 39, 87, 82, 55, 15, 31,
	25, 7, 104, 4, 3, 2, 1,
}
var yyPact = [...]int{

	74, -1000, -1000, -1000, -1000, -23, 37, 38, -1000, -1000,
	-1000, -1000, -1000, -1000, 22, 90, -1000, -23, 22, 11,
	22, -1000, -1000, 22, 61, 15, -1000, 22, 22, 46,
	4, 75, 6, -2, 47, -1000, 22, 103, 28, 75,
	-1000, 93, -1000, -1000, 73, 22, 12, 23, 22, -1000,
	88, -10, 84, -1000, 47, -1000, 42, -1000, 106, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 3, 0, -1000, 107,
	-11, 22, -15, -1000, -1000, 112, -1000, -1000, -1000, 23,
	22, -1000, 106, 31, 22, -4, 22, -1000, -1000, -25,
	-1000, -1000, -1000, -5, -1000, -13, -1000, -1000, -1000, 22,
	48, -1000, 22, -20, -1000, 22, 52, -18, -1000, 45,
	-1000, 89, 89, -1000, 59, -1000, -1000, 116, -1000, -1000,
	-1000, -1000,
}
var yyPgo = [...]int{

	0, 136, 135, 134, 133, 0, 4, 12, 2, 99,
	3, 132, 131, 130, 68, 5, 6, 129, 14, 128,
	127, 126, 8, 125, 10, 124, 7, 123, 1, 122,
	121,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 6, 6, 5, 5, 7, 8,
	10, 10, 11, 9, 9, 12, 12, 24, 24, 24,
	24, 24, 24, 24, 25, 25, 25, 26, 26, 26,
	18, 18, 27, 27, 27, 23, 23, 13, 13, 14,
	17, 17, 17, 17, 21, 21, 22, 22, 22, 20,
	20, 19, 19, 15, 15, 16, 28, 28, 28, 28,
	28, 29, 29, 30, 30, 2, 3, 4,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 3, 1, 1, 1, 3, 4, 1, 1,
	2, 2, 2, 2, 0, 1, 2, 5, 4, 9,
	0, 2, 0, 3, 4, 0, 1, 1, 3, 3,
	0, 1, 4, 6, 0, 2, 0, 1, 1, 0,
	2, 0, 1, 1, 3, 3, 2, 2, 1, 1,
	2, 3, 3, 0, 2, 4, 8, 9,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, 26, 10, -12, -9, -8,
	39, -5, 32, 34, 28, -19, 29, 36, 15, -5,
	16, -9, -5, 37, -5, -13, -14, -5, 21, 36,
	-27, -17, -5, -5, -18, -14, 9, 36, 38, -25,
	-24, 22, 29, 20, 19, 8, 11, 37, 37, -26,
	22, 29, 14, -5, -18, -20, 31, -24, 17, 20,
	-5, -7, -6, 35, 32, 33, -7, -15, -16, -5,
	17, 37, 17, -26, 25, -22, 5, 13, 38, 36,
	36, 38, -21, 8, 37, -15, 37, -23, 6, -7,
	-16, -22, -6, -15, 38, -10, -8, 38, 38, 36,
	38, -8, 23, -5, -11, 37, -30, -10, -29, 21,
	38, 12, 30, -28, 27, 7, 24, 18, -28, 20,
	11, 4,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 0, 51, 0, 15, 13,
	14, 9, 6, 7, 0, 0, 52, 0, 0, 0,
	0, 16, 65, 0, 0, 32, 37, 40, 0, 30,
	0, 24, 41, 0, 0, 38, 0, 30, 49, 39,
	25, 0, 18, 19, 0, 0, 0, 0, 0, 33,
	0, 0, 0, 31, 0, 66, 0, 26, 46, 20,
	21, 22, 23, 8, 4, 5, 0, 0, 53, 44,
	0, 0, 0, 34, 50, 35, 47, 48, 42, 0,
	0, 67, 46, 0, 0, 0, 0, 17, 36, 0,
	54, 55, 45, 0, 28, 0, 10, 43, 27, 0,
	0, 11, 0, 0, 63, 0, 29, 0, 64, 0,
	12, 0, 0, 61, 0, 58, 59, 0, 62, 56,
	57, 60,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	37, 38, 39, 3, 36,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35,
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
		//line parser.go.y:97
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:100
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:105
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:108
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:113
		{
			yyVAL.signedNumber = yyDollar[1].signedNumber
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:118
		{
			yyVAL.columnName = yyDollar[1].identifier
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:123
		{
			yyVAL.columnNameList = []string{yyDollar[1].columnName}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:126
		{
			yyVAL.columnNameList = append(yyDollar[1].columnNameList, yyDollar[3].columnName)
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:131
		{
			yyVAL.columnNameList = yyDollar[2].columnNameList
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:136
		{
			yyVAL.columnName = yyDollar[1].columnName
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:139
		{
			yyVAL.columnName = "*"
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:144
		{
			yyVAL.columnNameList = []string{yyDollar[1].columnName}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:147
		{
			yyVAL.columnNameList = append(yyDollar[1].columnNameList, yyDollar[3].columnName)
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:153
		{
			yyVAL.columnConstraint = ccPrimaryKey{yyDollar[3].sortOrder, yyDollar[4].bool}
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:156
		{
			yyVAL.columnConstraint = ccUnique(true)
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:159
		{
			yyVAL.columnConstraint = ccNull(true)
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:162
		{
			yyVAL.columnConstraint = ccNull(false)
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:165
		{
			yyVAL.columnConstraint = ccCollate(yyDollar[2].identifier)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:168
		{
			yyVAL.columnConstraint = ccDefault(yyDollar[2].signedNumber)
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:171
		{
			yyVAL.columnConstraint = ccDefault(yyDollar[2].literal)
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:176
		{
			yyVAL.columnConstraintList = nil
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:179
		{
			yyVAL.columnConstraintList = []columnConstraint{yyDollar[1].columnConstraint}
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:182
		{
			yyVAL.columnConstraintList = append(yyDollar[1].columnConstraintList, yyDollar[2].columnConstraint)
		}
	case 27:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:187
		{
			yyVAL.tableConstraint = TablePrimaryKey{yyDollar[4].indexedColumnList}
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:190
		{
			yyVAL.tableConstraint = TableUnique{yyDollar[3].indexedColumnList}
		}
	case 29:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:193
		{
			yyVAL.tableConstraint = TableForeignKey{
				Columns:        yyDollar[4].columnNameList,
				ForeignTable:   yyDollar[7].identifier,
				ForeignColumns: yyDollar[8].columnNameList,
				Triggers:       yyDollar[9].triggerList,
			}
		}
	case 30:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:203
		{
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:204
		{
		}
	case 32:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:208
		{
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:209
		{
			yyVAL.tableConstraintList = []TableConstraint{yyDollar[3].tableConstraint}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:212
		{
			yyVAL.tableConstraintList = append(yyDollar[1].tableConstraintList, yyDollar[4].tableConstraint)
		}
	case 35:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:218
		{
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:219
		{
			yyVAL.bool = true
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:224
		{
			yyVAL.columnDefList = []ColumnDef{yyDollar[1].columnDef}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:227
		{
			yyVAL.columnDefList = append(yyDollar[1].columnDefList, yyDollar[3].columnDef)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:232
		{
			yyVAL.columnDef = makeColumnDef(yyDollar[1].identifier, yyDollar[2].name, yyDollar[3].columnConstraintList)
		}
	case 40:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:237
		{
			yyVAL.name = ""
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:240
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:243
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 43:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:246
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 44:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:251
		{
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:252
		{
			yyVAL.collate = yyDollar[2].literal
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:257
		{
			yyVAL.sortOrder = Asc
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:260
		{
			yyVAL.sortOrder = Asc
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:263
		{
			yyVAL.sortOrder = Desc
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:268
		{
			yyVAL.withoutRowid = false
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:271
		{
			yyVAL.withoutRowid = true
		}
	case 51:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:276
		{
			yyVAL.unique = false
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:279
		{
			yyVAL.unique = true
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:284
		{
			yyVAL.indexedColumnList = []IndexedColumn{yyDollar[1].indexedColumn}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:287
		{
			yyVAL.indexedColumnList = append(yyDollar[1].indexedColumnList, yyDollar[3].indexedColumn)
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:292
		{
			yyVAL.indexedColumn = IndexedColumn{
				Column:    yyDollar[1].identifier,
				Collate:   yyDollar[2].collate,
				SortOrder: yyDollar[3].sortOrder,
			}
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:301
		{
			yyVAL.triggerAction = ActionSetNull
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:304
		{
			yyVAL.triggerAction = ActionSetDefault
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:307
		{
			yyVAL.triggerAction = ActionCascade
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:310
		{
			yyVAL.triggerAction = ActionRestrict
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:313
		{
			yyVAL.triggerAction = ActionNoAction
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:318
		{
			yyVAL.trigger = TriggerOnDelete(yyDollar[3].triggerAction)
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:321
		{
			yyVAL.trigger = TriggerOnUpdate(yyDollar[3].triggerAction)
		}
	case 63:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:326
		{
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:327
		{
			yyVAL.triggerList = append(yyDollar[1].triggerList, yyDollar[2].trigger)
		}
	case 65:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:332
		{
			yylex.(*lexer).result = SelectStmt{Columns: yyDollar[2].columnNameList, Table: yyDollar[4].identifier}
		}
	case 66:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:337
		{
			yylex.(*lexer).result = CreateTableStmt{
				Table:        yyDollar[3].identifier,
				Columns:      yyDollar[5].columnDefList,
				Constraints:  yyDollar[6].tableConstraintList,
				WithoutRowid: yyDollar[8].withoutRowid,
			}
		}
	case 67:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:347
		{
			yylex.(*lexer).result = CreateIndexStmt{
				Index:          yyDollar[4].identifier,
				Table:          yyDollar[6].identifier,
				Unique:         yyDollar[2].unique,
				IndexedColumns: yyDollar[8].indexedColumnList,
			}
		}
	}
	goto yystack /* stack new state and value */
}
