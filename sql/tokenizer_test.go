package sql

import (
	"errors"
	"reflect"
	"testing"
)

func TestTokens(t *testing.T) {
	type cas struct {
		sql  string
		want []token
		err  error
	}
	for n, c := range []cas{
		{
			sql: "foo foo_bar FoObAr foo1 _foo café",
			want: []token{
				stoken(tBare, "foo"),
				stoken(tBare, "foo_bar"),
				stoken(tBare, "FoObAr"),
				stoken(tBare, "foo1"),
				stoken(tBare, "_foo"),
				stoken(tBare, "café"),
			},
		},
		{
			sql: "1 -12 +34",
			want: []token{
				ntoken(tSignedNumber, 1),
				ntoken(tSignedNumber, -12),
				ntoken(tSignedNumber, +34),
			},
		},
		{
			sql: "create table foo",
			want: []token{
				stoken(CREATE, "create"),
				stoken(TABLE, "table"),
				stoken(tBare, "foo"),
			},
		},
		{
			sql: "create table foo (col1, col2, col3)",
			want: []token{
				stoken(CREATE, "create"),
				stoken(TABLE, "table"),
				stoken(tBare, "foo"),
				stoken('(', "("),
				stoken(tBare, "col1"),
				stoken(',', ","),
				stoken(tBare, "col2"),
				stoken(',', ","),
				stoken(tBare, "col3"),
				stoken(')', ")"),
			},
		},
		{
			// *
			sql: "select * from foo",
			want: []token{
				stoken(SELECT, "select"),
				stoken(tOperator, "*"),
				stoken(FROM, "from"),
				stoken(tBare, "foo"),
			},
		},
		{
			// fancy whitespace
			sql: "  \tselect\n*\nfrom   foo ",
			want: []token{
				stoken(SELECT, "select"),
				stoken(tOperator, "*"),
				stoken(FROM, "from"),
				stoken(tBare, "foo"),
			},
		},
		{
			sql: "from FROM 'from' ''",
			want: []token{
				stoken(FROM, "from"),
				stoken(FROM, "FROM"),
				stoken(tLiteral, "from"),
				stoken(tLiteral, ""),
			},
		},
		{
			sql: "bare \"id 1\" [id 2] `id 3` 'lit 1'",
			want: []token{
				stoken(tBare, "bare"),
				stoken(tIdentifier, "id 1"),
				stoken(tIdentifier, "id 2"),
				stoken(tIdentifier, "id 3"),
				stoken(tLiteral, "lit 1"),
			},
		},
		{
			sql: "|| | * > >=",
			want: []token{
				stoken(tOperator, "||"),
				stoken(tOperator, "|"),
				stoken(tOperator, "*"),
				stoken(tOperator, ">"),
				stoken(tOperator, ">="),
			},
		},
		{
			sql: "foo 'bar",
			err: errors.New("no terminating ' found"),
		},
	} {
		ts, err := tokenize(c.sql)
		if have, want := err, c.err; !reflect.DeepEqual(have, want) {
			t.Errorf("case %d: have %#v, want %#v", n, have, want)
			continue
		}
		if c.err != nil {
			continue
		}
		if have, want := ts, (c.want); !reflect.DeepEqual(have, want) {
			t.Errorf("case %d: have %#v, want %#v", n, have, want)
		}
	}

}
