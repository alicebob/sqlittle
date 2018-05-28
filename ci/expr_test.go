// +build ci

package ci

import (
	"os"
	"testing"

	"github.com/alicebob/sqlittle"
)

func TestExprCol(t *testing.T) {
	// index with expression column
	if os.Getenv("TRAVIS") != "" {
		t.Skip("travis has old sqlite3 version")
	}
	Compare(
		t,
		`
CREATE TABLE expr (name varchar(255));
CREATE INDEX expr_name ON expr (substr(name, 0, 10));
INSERT INTO expr values ("aap"), ("foo"), ("qqq"), ("longestnameever");
`,
		`SELECT name FROM expr ORDER BY name`,
		func(t *testing.T, db *sqlittle.DB) [][]string {
			var rows [][]string
			cb := func(r sqlittle.Row) {
				rows = append(rows, r.ScanStrings())
			}
			if err := db.IndexedSelect("expr", "expr_name", cb, "name"); err != nil {
				t.Fatal(err)
			}
			return rows
		},
	)
}

func TestExprWhere(t *testing.T) {
	// index with WHERE expression
	Compare(
		t,
		`
CREATE TABLE expr (name varchar(255));
CREATE INDEX expr_where ON expr (name) WHERE name > "foo";
INSERT INTO expr values ("aap"), ("foo"), ("qqq"), ("longestnameever");
`,
		`SELECT name FROM expr WHERE name > "foo" ORDER BY name`,
		func(t *testing.T, db *sqlittle.DB) [][]string {
			var rows [][]string
			cb := func(r sqlittle.Row) {
				rows = append(rows, r.ScanStrings())
			}
			if err := db.IndexedSelect("expr", "expr_where", cb, "name"); err != nil {
				t.Fatal(err)
			}
			return rows
		},
	)
}
