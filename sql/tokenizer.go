package sql

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	keywords = map[string]int{
		"ASC":           ASC,
		"AUTOINCREMENT": AUTOINCREMENT,
		"COLLATE":       COLLATE,
		"CREATE":        CREATE,
		"DESC":          DESC,
		"FROM":          FROM,
		"INDEX":         INDEX,
		"KEY":           KEY,
		"NOT":           NOT,
		"NULL":          NULL,
		"ON":            ON,
		"PRIMARY":       PRIMARY,
		"ROWID":         ROWID,
		"SELECT":        SELECT,
		"TABLE":         TABLE,
		"UNIQUE":        UNIQUE,
		"WITHOUT":       WITHOUT,
	}
)

type token struct {
	typ int
	s   string
}

func tokenize(s string) ([]token, error) {
	var res []token
	for i := 0; ; {
		if i >= len(s) {
			return res, nil
		}
		c, l := utf8.DecodeRuneInString(s[i:])
		switch {
		case unicode.IsSpace(c):
			// ignore
		case unicode.IsLetter(c) || c == '_':
			bt, bl := readBareword(s[i:])
			tnr := tBare
			if n, ok := keywords[strings.ToUpper(bt)]; ok {
				tnr = n
			}
			res = append(res, token{tnr, bt})
			i += bl - 1
		case unicode.IsDigit(c) || c == '+' || c == '-':
			d, l := readSignedNumber(s[i:])
			res = append(res, token{tSignedNumber, d})
			i += l - 1
		case c == '(' || c == ')' || c == ',' || c == '*':
			res = append(res, token{int(c), string(c)})
		case c == '\'':
			bt, bl := readSingleQuoted(s[i+1:])
			if bl == -1 {
				return res, errors.New("no terminating ' found")
			}
			res = append(res, token{tLiteral, bt})
			i += bl
		case c == '"' || c == '`' || c == '[':
			close := c
			if close == '[' {
				close = ']'
			}
			bt, bl := readQuoted(close, s[i+1:])
			if bl == -1 {
				return res, fmt.Errorf("no terminating %q found", close)
			}
			res = append(res, token{tIdentifier, bt})
			i += bl
		default:
			return nil, fmt.Errorf("unexpected char at pos:%d: %q", i, c)
		}
		i += l
	}
}

func readBareword(s string) (string, int) {
	for i, r := range s {
		switch {
		case unicode.IsLetter(r):
		case i > 0 && unicode.IsDigit(r):
		case r == '_':
		default:
			return s[:i], i
		}
	}
	return s, len(s)
}

func readSignedNumber(s string) (string, int) {
	// TODO: decimals, scientific notation
	for i, r := range s {
		switch {
		case i == 0 && r == '+' || r == '-':
		case unicode.IsDigit(r):
		default:
			return s[:i], i
		}
	}
	return s, len(s)
}

// parse a 'bareword'. Opening ' is already gone. No escape sequences.
func readSingleQuoted(s string) (string, int) {
	for i, r := range s {
		switch r {
		case '\'':
			return s[:i], i + 1
		default:
		}
	}
	return "", -1
}

// parse a quoted string until `close`. Opening char is already gone. No escape sequences.
func readQuoted(close rune, s string) (string, int) {
	for i, r := range s {
		switch r {
		case close:
			return s[:i], i + 1
		default:
		}
	}
	return "", -1
}
