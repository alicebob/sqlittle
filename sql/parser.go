//line parser.go.y:2
package sql

import __yyfmt__ "fmt"

//line parser.go.y:2
//line parser.go.y:5
type yySymType struct {
	yys           int
	identifier    string
	expr          interface{}
	columnList    []string
	columnName    string
	columnDefList []ColumnDef
	columnDef     ColumnDef
	null          bool
}

const SELECT = 57346
const FROM = 57347
const CREATE = 57348
const TABLE = 57349
const NOT = 57350
const NULL = 57351
const tBare = 57352

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"SELECT",
	"FROM",
	"CREATE",
	"TABLE",
	"NOT",
	"NULL",
	"tBare",
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

const yyLast = 33

var yyAct = [...]int{

	23, 8, 19, 21, 13, 22, 17, 7, 25, 26,
	24, 12, 10, 14, 29, 16, 10, 9, 11, 20,
	15, 25, 26, 20, 27, 28, 4, 18, 5, 6,
	3, 2, 1,
}
var yyPact = [...]int{

	22, -1000, -1000, -1000, 6, 11, -1, -1000, -1000, -1000,
	-1000, 2, 6, 2, -7, -1000, -1000, 2, -9, -1000,
	0, 2, -1000, -1000, 13, 5, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 32, 31, 30, 1, 29, 7, 27, 2, 0,
}
var yyR1 = [...]int{

	0, 1, 1, 4, 6, 6, 5, 5, 2, 7,
	7, 8, 8, 9, 9, 9, 3,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 3, 4, 1,
	3, 2, 3, 0, 2, 1, 6,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, 4, 6, -5, -6, -4, 11,
	10, 7, 12, 5, -4, -6, -4, 13, -7, -8,
	-4, 12, 14, -9, 10, 8, 9, -8, -9, 9,
}
var yyDef = [...]int{

	0, -2, 1, 2, 0, 0, 0, 6, 4, 5,
	3, 0, 0, 0, 0, 7, 8, 0, 0, 9,
	13, 0, 16, 11, 13, 0, 15, 10, 12, 14,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	13, 14, 11, 3, 12,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10,
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
		//line parser.go.y:35
		{
			yyVAL.identifier = yyDollar[1].identifier
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:40
		{
			yyVAL.columnName = yyDollar[1].identifier
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:43
		{
			yyVAL.columnName = "*"
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:48
		{
			yyVAL.columnList = []string{yyDollar[1].columnName}
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:51
		{
			yyVAL.columnList = append(yyDollar[1].columnList, yyDollar[3].columnName)
		}
	case 8:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:56
		{
			yylex.(*Lexer).result = SelectStmt{Columns: yyDollar[2].columnList, Table: yyDollar[4].identifier}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:61
		{
			yyVAL.columnDefList = []ColumnDef{yyDollar[1].columnDef}
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:64
		{
			yyVAL.columnDefList = append(yyDollar[1].columnDefList, yyDollar[3].columnDef)
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:69
		{
			yyVAL.columnDef = ColumnDef{Name: yyDollar[1].identifier, Null: yyDollar[2].null}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:72
		{
			yyVAL.columnDef = ColumnDef{Name: yyDollar[1].identifier, Type: yyDollar[2].identifier, Null: yyDollar[3].null}
		}
	case 13:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:77
		{
			yyVAL.null = true
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:80
		{
			yyVAL.null = false
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:83
		{
			yyVAL.null = true
		}
	case 16:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:88
		{
			yylex.(*Lexer).result = CreateTableStmt{Table: yyDollar[3].identifier, Columns: yyDollar[5].columnDefList}
		}
	}
	goto yystack /* stack new state and value */
}
