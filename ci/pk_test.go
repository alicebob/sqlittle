// +build ci

package ci

import (
	"testing"

	"github.com/alicebob/sqlittle"
)

func TestPKSelect(t *testing.T) {
	Compare(
		t,
		`
CREATE TABLE foo (a NOT NULL PRIMARY KEY, b, c);
INSERT INTO foo values ("aa", "bb1", "1");
INSERT INTO foo values ("a2", "bb2", "2");
INSERT INTO foo values ("a3", "bb3", "3");
INSERT INTO foo values ("a4", "bb4", "4");
INSERT INTO foo values ("a5", "bb5", "5");
`,
		`SELECT b, c FROM foo WHERE a='a3'`,
		func(t *testing.T, db *sqlittle.DB) [][]string {
			var rows [][]string
			cb := func(r sqlittle.Row) {
				rows = append(rows, r.ScanStrings())
			}
			err := db.PKSelect("foo", sqlittle.Key{"a3"}, cb, "b", "c")
			if err != nil {
				t.Fatal(err)
			}
			return rows
		},
	)
}

func TestPKDesc(t *testing.T) {
	Compare(
		t,
		`
CREATE TABLE foo (a not null, b not null, primary key (a, b DESC));
INSERT INTO foo values ("a1", 1);
INSERT INTO foo values ("a1", 2);
INSERT INTO foo values ("a1", 3);
INSERT INTO foo values ("a2", 1);
INSERT INTO foo values ("a2", 2);
INSERT INTO foo values ("a2", 3);
INSERT INTO foo values ("a3", 1);
INSERT INTO foo values ("a3", 2);
INSERT INTO foo values ("a3", 3);
INSERT INTO foo values ("a4", 1);
INSERT INTO foo values ("a4", 2);
INSERT INTO foo values ("a4", 3);
INSERT INTO foo values ("a5", 1);
INSERT INTO foo values ("a5", 2);
INSERT INTO foo values ("a5", 3);
`,
		`SELECT a, b FROM foo WHERE a='a2' and b==2`,
		func(t *testing.T, db *sqlittle.DB) [][]string {
			var rows [][]string
			cb := func(r sqlittle.Row) {
				rows = append(rows, r.ScanStrings())
			}
			err := db.PKSelect("foo", sqlittle.Key{"a2", 2}, cb, "a", "b")
			if err != nil {
				t.Fatal(err)
			}
			return rows
		},
	)
}

func TestPKNonrowidDesc(t *testing.T) {
	Compare(
		t,
		`
CREATE TABLE foo (a not null, b not null, primary key (a, b DESC)) WITHOUT ROWID;
INSERT INTO foo values ("a1", 1);
INSERT INTO foo values ("a1", 2);
INSERT INTO foo values ("a1", 3);
INSERT INTO foo values ("a2", 1);
INSERT INTO foo values ("a2", 2);
INSERT INTO foo values ("a2", 3);
INSERT INTO foo values ("a3", 1);
INSERT INTO foo values ("a3", 2);
INSERT INTO foo values ("a3", 3);
INSERT INTO foo values ("a4", 1);
INSERT INTO foo values ("a4", 2);
INSERT INTO foo values ("a4", 3);
INSERT INTO foo values ("a5", 1);
INSERT INTO foo values ("a5", 2);
INSERT INTO foo values ("a5", 3);
`,
		`SELECT a, b FROM foo WHERE a='a2' and b==2`,
		func(t *testing.T, db *sqlittle.DB) [][]string {
			var rows [][]string
			cb := func(r sqlittle.Row) {
				rows = append(rows, r.ScanStrings())
			}
			err := db.PKSelect("foo", sqlittle.Key{"a2", 2}, cb, "a", "b")
			if err != nil {
				t.Fatal(err)
			}
			return rows
		},
	)
}
