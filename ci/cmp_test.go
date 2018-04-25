// +build ci

package ci

import (
	"testing"

	"github.com/alicebob/sqlittle"
)

func TestSelect(t *testing.T) {
	Compare(
		t,
		`
CREATE TABLE foo (a, b, c);
INSERT INTO foo values ("aa", "bb", "cc");
INSERT INTO foo values ("aa", "bb", "cc");
INSERT INTO foo values ("aa2", "bb2", "cc2");
INSERT INTO foo values ("aa2", "bb2", "cc2");
INSERT INTO foo values ("aa3", "bb3", "cc3");
`,
		`SELECT a, c FROM foo`,
		func(t *testing.T, db *sqlittle.DB) [][]string {
			var rows [][]string
			cb := func(r sqlittle.Row) {
				rows = append(rows, r.ScanStrings())
			}
			if err := db.Select("foo", cb, "a", "c"); err != nil {
				t.Fatal(err)
			}
			return rows
		},
	)
}

func TestSelectRowid(t *testing.T) {
	Compare(
		t,
		`
CREATE TABLE foo (a, b, c);
INSERT INTO foo values ("aa", "bb", "cc");
INSERT INTO foo values ("a2", "bb", "cc");
INSERT INTO foo values ("a3", "bb", "cc");
INSERT INTO foo values ("a4", "bb", "cc");
INSERT INTO foo values ("a5", "bb", "cc");
`,
		`SELECT a FROM foo WHERE rowid=3`,
		func(t *testing.T, db *sqlittle.DB) [][]string {
			row, err := db.SelectRowid("foo", int64(3), "a")
			if err != nil {
				t.Fatal(err)
			}
			s, _ := row.ScanString()
			return [][]string{{s}}
		},
	)
}

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

func TestSelectIndexEqDesc(t *testing.T) {
	Compare(
		t,
		`
CREATE TABLE foo (a, b);
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
CREATE INDEX foo_desc ON foo (a DESC, b DESC);
`,
		`SELECT a, b FROM foo WHERE a='a3' ORDER BY a DESC, b DESC`,
		func(t *testing.T, db *sqlittle.DB) [][]string {
			var rows [][]string
			cb := func(r sqlittle.Row) {
				rows = append(rows, r.ScanStrings())
			}
			err := db.IndexedSelectEq("foo", "foo_desc", sqlittle.Key{"a3"}, cb, "a", "b")
			if err != nil {
				t.Fatal(err)
			}
			return rows
		},
	)
}
