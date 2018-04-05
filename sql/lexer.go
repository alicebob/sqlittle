package sql

import (
	"errors"
)

type Lexer struct {
	tokens []token
	result interface{}
	err    error
}

func (l *Lexer) Lex(lval *yySymType) int {
	if len(l.tokens) == 0 {
		return 0
	}
	tok := l.tokens[0]
	l.tokens = l.tokens[1:]

	lval.identifier = tok.s
	return tok.typ
}

func (l *Lexer) Error(e string) {
	l.err = errors.New(e)
}

func Parse(sql string) (interface{}, error) {
	ts, err := tokenize(sql)
	if err != nil {
		return nil, err
	}
	l := &Lexer{tokens: ts}
	yyParse(l)
	return l.result, l.err
}
