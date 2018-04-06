package sql

import (
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
			sql: "create table foo",
			want: []token{
				token{CREATE, "create"},
				token{TABLE, "table"},
				token{tBare, "foo"},
			},
		},
		{
			sql: "create table foo (col1, col2, col3)",
			want: []token{
				token{CREATE, "create"},
				token{TABLE, "table"},
				token{tBare, "foo"},
				token{'(', "("},
				token{tBare, "col1"},
				token{',', ","},
				token{tBare, "col2"},
				token{',', ","},
				token{tBare, "col3"},
				token{')', ")"},
			},
		},
		{
			// *
			sql: "select * from foo",
			want: []token{
				token{SELECT, "select"},
				token{'*', "*"},
				token{FROM, "from"},
				token{tBare, "foo"},
			},
		},
		{
			// fancy whitespace
			sql: "  \tselect\n*\nfrom   foo ",
			want: []token{
				token{SELECT, "select"},
				token{'*', "*"},
				token{FROM, "from"},
				token{tBare, "foo"},
			},
		},
		{
			sql: "1 -12 +34",
			want: []token{
				token{tSignedNumber, "1"},
				token{tSignedNumber, "-12"},
				token{tSignedNumber, "+34"},
			},
		},
	} {
		ts, err := tokenize(c.sql)
		if have, want := err, c.err; !reflect.DeepEqual(have, want) {
			t.Errorf("case %d: have %#v, want %#v", n, have, want)
			continue
		}
		if have, want := ts, (c.want); !reflect.DeepEqual(have, want) {
			t.Errorf("case %d: have %#v, want %#v", n, have, want)
		}
	}

}
