//line parser.go.y:2
package sql

import __yyfmt__ "fmt"

//line parser.go.y:2
//line parser.go.y:5
type yySymType struct {
	yys           int
	identifier    string
	signedNumber  string
	expr          interface{}
	columnList    []string
	columnName    string
	columnDefList []ColumnDef
	columnDef     ColumnDef
	name          string
	primaryKey    PrimaryKey
	autoIncrement bool
	unique        bool
	null          bool
}

const SELECT = 57346
const FROM = 57347
const CREATE = 57348
const TABLE = 57349
const PRIMARY = 57350
const KEY = 57351
const ASC = 57352
const DESC = 57353
const AUTOINCREMENT = 57354
const NOT = 57355
const NULL = 57356
const UNIQUE = 57357
const tBare = 57358
const tSignedNumber = 57359

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"SELECT",
	"FROM",
	"CREATE",
	"TABLE",
	"PRIMARY",
	"KEY",
	"ASC",
	"DESC",
	"AUTOINCREMENT",
	"NOT",
	"NULL",
	"UNIQUE",
	"tBare",
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

const yyLast = 47

var yyAct = [...]int{

	32, 8, 40, 21, 39, 22, 19, 45, 28, 17,
	13, 33, 10, 14, 9, 16, 10, 42, 43, 20,
	35, 36, 24, 20, 12, 7, 37, 38, 25, 30,
	31, 27, 11, 4, 29, 5, 41, 34, 15, 26,
	23, 44, 18, 6, 3, 2, 1,
}
var yyPact = [...]int{

	29, -1000, -1000, -1000, -4, 25, 5, -1000, -1000, -1000,
	-1000, 0, -4, 0, -11, -1000, -1000, 0, -16, -1000,
	0, 0, -1000, 23, -12, -1000, 17, 21, -6, 7,
	-1000, 16, -17, -1000, 2, 4, -1000, -1000, -1000, -1000,
	-6, -1000, -1000, -1000, -14, -1000,
}
var yyPgo = [...]int{

	0, 46, 45, 44, 1, 0, 43, 25, 42, 6,
	40, 39, 37, 36, 34,
}
var yyR1 = [...]int{

	0, 1, 1, 4, 5, 7, 7, 6, 6, 2,
	8, 8, 9, 10, 10, 10, 10, 11, 11, 11,
	11, 14, 14, 12, 12, 12, 13, 13, 3,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 3, 4,
	1, 3, 6, 0, 1, 4, 6, 0, 2, 3,
	3, 0, 1, 0, 2, 1, 0, 1, 6,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, 4, 6, -6, -7, -4, 18,
	16, 7, 19, 5, -4, -7, -4, 20, -8, -9,
	-4, 19, 21, -10, -4, -9, -11, 8, 20, -14,
	12, 9, -5, 17, -12, 13, 14, 10, 11, 21,
	19, -13, 15, 14, -5, 21,
}
var yyDef = [...]int{

	0, -2, 1, 2, 0, 0, 0, 7, 5, 6,
	3, 0, 0, 0, 0, 8, 9, 0, 0, 10,
	13, 0, 28, 17, 14, 11, 21, 0, 0, 23,
	22, 18, 0, 4, 26, 0, 25, 19, 20, 15,
	0, 12, 27, 24, 0, 16,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	20, 21, 18, 3, 19,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17,
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

	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:46
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:51
		{
			yyVAL.signedNumber = yyDollar[1].signedNumber
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:56
		{
			yyVAL.columnName = yyDollar[1].identifier
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:59
		{
			yyVAL.columnName = "*"
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:64
		{
			yyVAL.columnList = []string{yyDollar[1].columnName}
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:67
		{
			yyVAL.columnList = append(yyDollar[1].columnList, yyDollar[3].columnName)
		}
	case 9:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:72
		{
			yylex.(*Lexer).result = SelectStmt{Columns: yyDollar[2].columnList, Table: yyDollar[4].identifier}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:77
		{
			yyVAL.columnDefList = []ColumnDef{yyDollar[1].columnDef}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:80
		{
			yyVAL.columnDefList = append(yyDollar[1].columnDefList, yyDollar[3].columnDef)
		}
	case 12:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:85
		{
			yyVAL.columnDef = ColumnDef{
				Name:          yyDollar[1].identifier,
				Type:          yyDollar[2].name,
				PrimaryKey:    yyDollar[3].primaryKey,
				AutoIncrement: yyDollar[4].autoIncrement,
				Null:          yyDollar[5].null,
				Unique:        yyDollar[6].unique,
			}
		}
	case 13:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:97
		{
			yyVAL.name = ""
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:100
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:103
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 16:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:106
		{
			yyVAL.name = yyDollar[1].identifier
		}
	case 17:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:111
		{
			yyVAL.primaryKey = PKNone
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:114
		{
			yyVAL.primaryKey = PKAsc
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:117
		{
			yyVAL.primaryKey = PKAsc
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:120
		{
			yyVAL.primaryKey = PKDesc
		}
	case 21:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:125
		{
			yyVAL.autoIncrement = false
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:128
		{
			yyVAL.autoIncrement = true
		}
	case 23:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:133
		{
			yyVAL.null = true
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:136
		{
			yyVAL.null = false
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:139
		{
			yyVAL.null = true
		}
	case 26:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:144
		{
			yyVAL.unique = false
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:147
		{
			yyVAL.unique = true
		}
	case 28:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:152
		{
			yylex.(*Lexer).result = CreateTableStmt{Table: yyDollar[3].identifier, Columns: yyDollar[5].columnDefList}
		}
	}
	goto yystack /* stack new state and value */
}
