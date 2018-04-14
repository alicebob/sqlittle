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
const CREATE = 57351
const DEFAULT = 57352
const DELETE = 57353
const DESC = 57354
const FOREIGN = 57355
const FROM = 57356
const INDEX = 57357
const KEY = 57358
const NO = 57359
const NOT = 57360
const NULL = 57361
const ON = 57362
const PRIMARY = 57363
const REFERENCES = 57364
const RESTRICT = 57365
const ROWID = 57366
const SELECT = 57367
const SET = 57368
const TABLE = 57369
const UNIQUE = 57370
const UPDATE = 57371
const WITHOUT = 57372
const tBare = 57373
const tLiteral = 57374
const tIdentifier = 57375
const tSignedNumber = 57376

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ACTION",
	"ASC",
	"AUTOINCREMENT",
	"CASCADE",
	"COLLATE",
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

const yyLast = 135

var yyAct = [...]int{

	11, 109, 9, 85, 62, 68, 74, 61, 67, 12,
	42, 13, 96, 101, 34, 19, 10, 72, 94, 22,
	106, 24, 70, 94, 27, 95, 52, 50, 32, 33,
	27, 79, 79, 93, 84, 79, 78, 80, 77, 39,
	29, 40, 64, 65, 49, 63, 23, 38, 60, 18,
	63, 69, 57, 69, 54, 36, 12, 66, 13, 64,
	65, 71, 37, 56, 107, 12, 26, 13, 73, 38,
	17, 69, 14, 16, 98, 86, 59, 36, 105, 83,
	69, 47, 108, 48, 37, 90, 89, 92, 91, 116,
	111, 46, 45, 58, 43, 6, 35, 97, 115, 99,
	113, 44, 28, 8, 86, 103, 112, 53, 51, 110,
	114, 5, 20, 75, 82, 88, 117, 102, 104, 30,
	76, 21, 41, 87, 81, 55, 15, 31, 25, 7,
	100, 4, 3, 2, 1,
}
var yyPact = [...]int{

	86, -1000, -1000, -1000, -1000, -22, 45, 35, -1000, -1000,
	-1000, -1000, -1000, -1000, 25, 97, -1000, -22, 25, 10,
	25, -1000, -1000, 25, 82, 5, -1000, 25, 25, 34,
	4, 73, 8, -9, -1000, -1000, 92, -10, 91, 56,
	33, 73, -1000, 77, -1000, -1000, 57, 25, 11, 16,
	25, -14, 25, -19, -1000, -1000, 44, -1000, 108, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 1, 0, -1000, 106,
	25, -3, 25, -1000, 109, -1000, -1000, -1000, 16, 25,
	-1000, 108, 28, -4, -1000, -12, -1000, -1000, -1000, -25,
	-1000, -1000, -1000, -1000, 25, 52, -1000, -1000, 25, -23,
	-1000, 25, 58, -17, -1000, 53, -1000, 83, 83, -1000,
	79, -1000, -1000, 112, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 134, 133, 132, 131, 0, 4, 7, 2, 103,
	3, 130, 129, 128, 66, 8, 5, 127, 126, 125,
	124, 6, 123, 10, 122, 14, 119, 1, 118, 117,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 6, 6, 5, 5, 7, 8,
	10, 10, 11, 9, 9, 12, 12, 23, 23, 23,
	23, 23, 23, 23, 24, 24, 24, 25, 25, 25,
	26, 26, 26, 22, 22, 13, 13, 14, 17, 17,
	17, 17, 20, 20, 21, 21, 21, 19, 19, 18,
	18, 15, 15, 16, 27, 27, 27, 27, 27, 28,
	28, 29, 29, 2, 3, 4,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 3, 1, 1, 1, 3, 4, 1, 1,
	2, 2, 2, 2, 0, 1, 2, 5, 4, 9,
	0, 2, 3, 0, 1, 1, 3, 3, 0, 1,
	4, 6, 0, 2, 0, 1, 1, 0, 2, 0,
	1, 1, 3, 3, 2, 2, 1, 1, 2, 3,
	3, 0, 2, 4, 8, 9,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, 25, 9, -12, -9, -8,
	38, -5, 31, 33, 27, -18, 28, 35, 14, -5,
	15, -9, -5, 36, -5, -13, -14, -5, 20, 35,
	-26, -17, -5, -5, -25, -14, 21, 28, 13, 35,
	37, -24, -23, 21, 28, 19, 18, 8, 10, 36,
	36, 16, 36, 16, -25, -19, 30, -23, 16, 19,
	-5, -7, -6, 34, 31, 32, -7, -15, -16, -5,
	36, -15, 36, 24, -21, 5, 12, 37, 35, 35,
	37, -20, 8, -15, 37, -10, -8, -22, 6, -7,
	-16, -21, -6, 37, 35, 37, 37, -8, 22, -5,
	-11, 36, -29, -10, -28, 20, 37, 11, 29, -27,
	26, 7, 23, 17, -27, 19, 10, 4,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 0, 49, 0, 15, 13,
	14, 9, 6, 7, 0, 0, 50, 0, 0, 0,
	0, 16, 63, 0, 0, 30, 35, 38, 0, 0,
	0, 24, 39, 0, 31, 36, 0, 0, 0, 0,
	47, 37, 25, 0, 18, 19, 0, 0, 0, 0,
	0, 0, 0, 0, 32, 64, 0, 26, 44, 20,
	21, 22, 23, 8, 4, 5, 0, 0, 51, 42,
	0, 0, 0, 48, 33, 45, 46, 40, 0, 0,
	65, 44, 0, 0, 28, 0, 10, 17, 34, 0,
	52, 53, 43, 27, 0, 0, 41, 11, 0, 0,
	61, 0, 29, 0, 62, 0, 12, 0, 0, 59,
	0, 56, 57, 0, 60, 54, 55, 58,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	36, 37, 38, 3, 35,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34,
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
		//line parser.go.y:96
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:99
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:104
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:107
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:112
		{
			yyVAL.signedNumber = yyDollar[1].signedNumber
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:117
		{
			yyVAL.columnName = yyDollar[1].identifier
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:122
		{
			yyVAL.columnNameList = []string{yyDollar[1].columnName}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:125
		{
			yyVAL.columnNameList = append(yyDollar[1].columnNameList, yyDollar[3].columnName)
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:130
		{
			yyVAL.columnNameList = yyDollar[2].columnNameList
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:135
		{
			yyVAL.columnName = yyDollar[1].columnName
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:138
		{
			yyVAL.columnName = "*"
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:143
		{
			yyVAL.columnNameList = []string{yyDollar[1].columnName}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:146
		{
			yyVAL.columnNameList = append(yyDollar[1].columnNameList, yyDollar[3].columnName)
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:152
		{
			yyVAL.columnConstraint = ccPrimaryKey{yyDollar[3].sortOrder, yyDollar[4].bool}
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:155
		{
			yyVAL.columnConstraint = ccUnique(true)
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:158
		{
			yyVAL.columnConstraint = ccNull(true)
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:161
		{
			yyVAL.columnConstraint = ccNull(false)
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:164
		{
			yyVAL.columnConstraint = ccCollate(yyDollar[2].identifier)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:167
		{
			yyVAL.columnConstraint = ccDefault(yyDollar[2].signedNumber)
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:170
		{
			yyVAL.columnConstraint = ccDefault(yyDollar[2].literal)
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:175
		{
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:177
		{
			yyVAL.columnConstraintList = []columnConstraint{yyDollar[1].columnConstraint}
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:180
		{
			yyVAL.columnConstraintList = append(yyDollar[1].columnConstraintList, yyDollar[2].columnConstraint)
		}
	case 27:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:185
		{
			yyVAL.tableConstraint = TablePrimaryKey{yyDollar[4].indexedColumnList}
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:188
		{
			yyVAL.tableConstraint = TableUnique{yyDollar[3].indexedColumnList}
		}
	case 29:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:191
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
		//line parser.go.y:201
		{
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:202
		{
			yyVAL.tableConstraintList = []TableConstraint{yyDollar[2].tableConstraint}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:205
		{
			yyVAL.tableConstraintList = append(yyDollar[1].tableConstraintList, yyDollar[3].tableConstraint)
		}
	case 33:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:211
		{
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:212
		{
			yyVAL.bool = true
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:217
		{
			yyVAL.columnDefList = []ColumnDef{yyDollar[1].columnDef}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:220
		{
			yyVAL.columnDefList = append(yyDollar[1].columnDefList, yyDollar[3].columnDef)
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:225
		{
			yyVAL.columnDef = makeColumnDef(yyDollar[1].identifier, yyDollar[2].name, yyDollar[3].columnConstraintList)
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:230
		{
			yyVAL.name = ""
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:233
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:236
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 41:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:239
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:244
		{
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:245
		{
			yyVAL.collate = yyDollar[2].literal
		}
	case 44:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:250
		{
			yyVAL.sortOrder = Asc
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:253
		{
			yyVAL.sortOrder = Asc
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:256
		{
			yyVAL.sortOrder = Desc
		}
	case 47:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:261
		{
			yyVAL.withoutRowid = false
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:264
		{
			yyVAL.withoutRowid = true
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:269
		{
			yyVAL.unique = false
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:272
		{
			yyVAL.unique = true
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:277
		{
			yyVAL.indexedColumnList = []IndexedColumn{yyDollar[1].indexedColumn}
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:280
		{
			yyVAL.indexedColumnList = append(yyDollar[1].indexedColumnList, yyDollar[3].indexedColumn)
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:285
		{
			yyVAL.indexedColumn = IndexedColumn{
				Column:    yyDollar[1].identifier,
				Collate:   yyDollar[2].collate,
				SortOrder: yyDollar[3].sortOrder,
			}
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:294
		{
			yyVAL.triggerAction = ActionSetNull
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:297
		{
			yyVAL.triggerAction = ActionSetDefault
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:300
		{
			yyVAL.triggerAction = ActionCascade
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:303
		{
			yyVAL.triggerAction = ActionRestrict
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:306
		{
			yyVAL.triggerAction = ActionNoAction
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:311
		{
			yyVAL.trigger = TriggerOnDelete(yyDollar[3].triggerAction)
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:314
		{
			yyVAL.trigger = TriggerOnUpdate(yyDollar[3].triggerAction)
		}
	case 61:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:319
		{
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:320
		{
			yyVAL.triggerList = append(yyDollar[1].triggerList, yyDollar[2].trigger)
		}
	case 63:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:325
		{
			yylex.(*lexer).result = SelectStmt{Columns: yyDollar[2].columnNameList, Table: yyDollar[4].identifier}
		}
	case 64:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:330
		{
			yylex.(*lexer).result = CreateTableStmt{
				Table:        yyDollar[3].identifier,
				Columns:      yyDollar[5].columnDefList,
				Constraints:  yyDollar[6].tableConstraintList,
				WithoutRowid: yyDollar[8].withoutRowid,
			}
		}
	case 65:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:340
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
