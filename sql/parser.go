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

const yyLast = 101

var yyAct = [...]int{

	59, 65, 70, 58, 64, 40, 66, 75, 33, 87,
	88, 75, 9, 80, 75, 74, 76, 73, 25, 37,
	18, 38, 67, 9, 21, 50, 23, 48, 47, 26,
	22, 17, 28, 31, 32, 26, 11, 11, 12, 12,
	10, 61, 62, 60, 60, 54, 51, 34, 69, 35,
	53, 63, 57, 16, 78, 68, 36, 61, 62, 41,
	56, 11, 35, 12, 44, 43, 42, 45, 8, 36,
	46, 13, 79, 71, 72, 27, 82, 84, 83, 86,
	85, 15, 55, 49, 19, 20, 5, 29, 6, 39,
	81, 77, 52, 14, 30, 24, 7, 4, 3, 2,
	1,
}
var yyPact = [...]int{

	82, -1000, -1000, -1000, -1000, 14, 64, 26, -1000, -1000,
	-1000, -1000, -1000, 15, 76, -1000, 14, 15, 2, 15,
	-1000, -1000, 15, 66, 5, -1000, 15, 15, 39, -8,
	49, 0, -1, -1000, -1000, 72, -3, 52, 31, 49,
	-1000, 71, -1000, -1000, 44, 15, 19, 18, 15, -6,
	15, -1000, -1000, 28, -1000, 61, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -12, -13, -1000, 36, 15, -16, -1000,
	62, -1000, -1000, -1000, 18, 15, -1000, 61, 35, -20,
	-1000, -1000, -1000, -19, -1000, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 100, 99, 98, 97, 6, 0, 3, 96, 68,
	95, 18, 4, 1, 94, 93, 92, 91, 2, 90,
	5, 89, 8, 87,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 6, 6, 5, 5, 7, 9,
	9, 8, 8, 20, 20, 20, 20, 20, 20, 20,
	21, 21, 21, 22, 22, 23, 23, 23, 19, 19,
	10, 10, 11, 14, 14, 14, 14, 17, 17, 18,
	18, 18, 16, 16, 15, 15, 12, 12, 13, 2,
	3, 4,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 4, 1, 1, 2, 2, 2, 2,
	0, 1, 2, 5, 4, 0, 2, 3, 0, 1,
	1, 3, 3, 0, 1, 4, 6, 0, 2, 0,
	1, 1, 0, 2, 0, 1, 1, 3, 3, 4,
	8, 9,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, 4, 6, -8, -9, -5,
	26, 22, 24, 7, -15, 17, 27, 5, -5, 8,
	-9, -5, 28, -5, -10, -11, -5, 9, 27, -23,
	-14, -5, -5, -22, -11, 10, 17, 27, 29, -21,
	-20, 10, 17, 16, 15, 18, 21, 28, 28, 11,
	28, -22, -16, 19, -20, 11, 16, -5, -7, -6,
	25, 22, 23, -7, -12, -13, -5, 28, -12, 20,
	-18, 12, 13, 29, 27, 27, 29, -17, 18, -12,
	29, -19, 14, -7, -13, -18, -6, 29, 29,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 0, 44, 0, 11, 9,
	10, 6, 7, 0, 0, 45, 0, 0, 0, 0,
	12, 49, 0, 0, 25, 30, 33, 0, 0, 0,
	20, 34, 0, 26, 31, 0, 0, 0, 42, 32,
	21, 0, 14, 15, 0, 0, 0, 0, 0, 0,
	0, 27, 50, 0, 22, 39, 16, 17, 18, 19,
	8, 4, 5, 0, 0, 46, 37, 0, 0, 43,
	28, 40, 41, 35, 0, 0, 51, 39, 0, 0,
	24, 13, 29, 0, 47, 48, 38, 23, 36,
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
		//line parser.go.y:65
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:68
		{
			yyVAL.literal = yyDollar[1].identifier
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:73
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:76
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:81
		{
			yyVAL.signedNumber = yyDollar[1].signedNumber
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:86
		{
			yyVAL.columnName = yyDollar[1].identifier
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:89
		{
			yyVAL.columnName = "*"
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:94
		{
			yyVAL.columnList = []string{yyDollar[1].columnName}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:97
		{
			yyVAL.columnList = append(yyDollar[1].columnList, yyDollar[3].columnName)
		}
	case 13:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:102
		{
			yyVAL.columnConstraint = ccPrimaryKey{yyDollar[3].sortOrder, yyDollar[4].bool}
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:105
		{
			yyVAL.columnConstraint = ccUnique(true)
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:108
		{
			yyVAL.columnConstraint = ccNull(true)
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:111
		{
			yyVAL.columnConstraint = ccNull(false)
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:114
		{
			yyVAL.columnConstraint = ccCollate(yyDollar[2].identifier)
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:117
		{
			yyVAL.columnConstraint = ccDefault(yyDollar[2].signedNumber)
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:120
		{
			yyVAL.columnConstraint = ccDefault(yyDollar[2].literal)
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:125
		{
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:127
		{
			yyVAL.columnConstraintList = []columnConstraint{yyDollar[1].columnConstraint}
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:130
		{
			yyVAL.columnConstraintList = append(yyDollar[1].columnConstraintList, yyDollar[2].columnConstraint)
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:135
		{
			yyVAL.tableConstraint = TablePrimaryKey{yyDollar[4].indexedColumnList}
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:138
		{
			yyVAL.tableConstraint = TableUnique{yyDollar[3].indexedColumnList}
		}
	case 25:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:143
		{
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:144
		{
			yyVAL.tableConstraintList = []TableConstraint{yyDollar[2].tableConstraint}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:147
		{
			yyVAL.tableConstraintList = append(yyDollar[1].tableConstraintList, yyDollar[3].tableConstraint)
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:153
		{
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:154
		{
			yyVAL.bool = true
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:159
		{
			yyVAL.columnDefList = []ColumnDef{yyDollar[1].columnDef}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:162
		{
			yyVAL.columnDefList = append(yyDollar[1].columnDefList, yyDollar[3].columnDef)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:167
		{
			yyVAL.columnDef = makeColumnDef(yyDollar[1].identifier, yyDollar[2].name, yyDollar[3].columnConstraintList)
		}
	case 33:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:172
		{
			yyVAL.name = ""
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:175
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:178
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 36:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:181
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:186
		{
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:187
		{
			yyVAL.collate = yyDollar[2].literal
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:192
		{
			yyVAL.sortOrder = Asc
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:195
		{
			yyVAL.sortOrder = Asc
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:198
		{
			yyVAL.sortOrder = Desc
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:203
		{
			yyVAL.withoutRowid = false
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:206
		{
			yyVAL.withoutRowid = true
		}
	case 44:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:211
		{
			yyVAL.unique = false
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:214
		{
			yyVAL.unique = true
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:219
		{
			yyVAL.indexedColumnList = []IndexedColumn{yyDollar[1].indexedColumn}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:222
		{
			yyVAL.indexedColumnList = append(yyDollar[1].indexedColumnList, yyDollar[3].indexedColumn)
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:227
		{
			yyVAL.indexedColumn = IndexedColumn{
				Column:    yyDollar[1].identifier,
				Collate:   yyDollar[2].collate,
				SortOrder: yyDollar[3].sortOrder,
			}
		}
	case 49:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:236
		{
			yylex.(*lexer).result = SelectStmt{Columns: yyDollar[2].columnList, Table: yyDollar[4].identifier}
		}
	case 50:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:241
		{
			yylex.(*lexer).result = CreateTableStmt{
				Table:        yyDollar[3].identifier,
				Columns:      yyDollar[5].columnDefList,
				Constraints:  yyDollar[6].tableConstraintList,
				WithoutRowid: yyDollar[8].withoutRowid,
			}
		}
	case 51:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:251
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
