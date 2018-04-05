package sql

import (
	"errors"
	"reflect"
	"testing"
)

func TestSQL(t *testing.T) {

	type sqlCase struct {
		sql  string
		want interface{}
		err  error
	}

	cases := []sqlCase{
		// unknown
		{
			sql: "INSERT INTO FOO",
			err: errors.New("syntax error"),
		},

		// select
		{
			sql:  "SELECT * FROM foo",
			want: SelectStmt{Columns: []string{"*"}, Table: "foo"},
		},
		{
			sql:  "SELECT aap,noot, mies FROM foo2",
			want: SelectStmt{Columns: []string{"aap", "noot", "mies"}, Table: "foo2"},
		},

		// create table
		{
			sql:  "CREATE TABLE foo",
			want: CreatTableStmt{"foo"},
		},
		{
			sql: "CREATE nothing foo",
			err: errors.New("syntax error"),
		},
	}

	for n, c := range cases {
		stmt, err := Parse(c.sql)
		if have, want := err, c.err; !reflect.DeepEqual(have, want) {
			t.Errorf("case %d: have %#v, want %#v", n, have, want)
			continue
		}
		if c.err != nil {
			continue
		}
		if have, want := stmt, c.want; !reflect.DeepEqual(have, want) {
			t.Errorf("case %d: have %#v, want %#v", n, have, want)
		}
	}
}
