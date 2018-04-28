// +build ci

package ci

import (
	"testing"

	"github.com/alicebob/sqlittle"
)

func TestCollateRtrim(t *testing.T) {
	// collate on an index. Doesn't need any support.
	Compare(
		t,
		`
CREATE TABLE col (a text);
INSERT INTO col values ('Ondt på mig');
INSERT INTO col values ('ondt på mig   ');
INSERT INTO col values ('OnDT på MIG');
CREATE INDEX col_a_nocase ON col (a collate rtrim);
`,
		`SELECT rowid, a FROM col ORDER BY a collate 'rtrim'`,
		func(t *testing.T, db *sqlittle.DB) [][]string {
			var rows [][]string
			cb := func(r sqlittle.Row) {
				rows = append(rows, r.ScanStrings())
			}
			if err := db.IndexedSelect("col", "col_a_nocase", cb, "rowid", "a"); err != nil {
				t.Fatal(err)
			}
			return rows
		},
	)
}

func TestCollateEq(t *testing.T) {
	Compare(
		t,
		`
CREATE TABLE col (a text);
INSERT INTO col values ('Ondt på mig');
INSERT INTO col values ('ondt på mig   ');
INSERT INTO col values ('OnDT på MIG');
CREATE INDEX col_a_nocase ON col (a collate rtrim);
`,
		`SELECT rowid, a FROM col where a='ondt på mig' collate 'rtrim'`,
		func(t *testing.T, db *sqlittle.DB) [][]string {
			var rows [][]string
			cb := func(r sqlittle.Row) {
				rows = append(rows, r.ScanStrings())
			}
			if err := db.IndexedSelectEq("col", "col_a_nocase", sqlittle.Key{"ondt på mig"}, cb, "rowid", "a"); err != nil {
				t.Fatal(err)
			}
			return rows
		},
	)
}

func TestCollateIndexedEqCol(t *testing.T) {
	t.Skip("broken")
	// collate on a column should end up on the index
	Compare(
		t,
		`
CREATE TABLE col (a text collate rtrim);
INSERT INTO col values ('Ondt på mig');
INSERT INTO col values ('ondt på mig   ');
INSERT INTO col values ('OnDT på MIG');
CREATE INDEX col_a_nocase ON col (a);
`,
		`SELECT rowid, a FROM col where a='ondt på mig'`,
		func(t *testing.T, db *sqlittle.DB) [][]string {
			var rows [][]string
			cb := func(r sqlittle.Row) {
				rows = append(rows, r.ScanStrings())
			}
			if err := db.IndexedSelectEq("col", "col_a_nocase", sqlittle.Key{"ondt på mig"}, cb, "rowid", "a"); err != nil {
				t.Fatal(err)
			}
			return rows
		},
	)
}
