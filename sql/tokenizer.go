package sql

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	keywords = map[string]int{
		"SELECT":        SELECT,
		"FROM":          FROM,
		"CREATE":        CREATE,
		"TABLE":         TABLE,
		"NOT":           NOT,
		"NULL":          NULL,
		"UNIQUE":        UNIQUE,
		"PRIMARY":       PRIMARY,
		"KEY":           KEY,
		"ASC":           ASC,
		"DESC":          DESC,
		"AUTOINCREMENT": AUTOINCREMENT,
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
		case unicode.IsLetter(c):
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
