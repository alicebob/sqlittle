// +build ci

package ci

import (
	"testing"
	"time"

	"github.com/alicebob/sqlittle"
)

func TestScan(t *testing.T) {
	// we don't test (or support) the 'real' storage type
	Compare(
		t,
		`
CREATE TABLE times (i int, c text, f real, b blob);
INSERT INTO times values (strftime('%s','now'), datetime('now'), julianday('now'), datetime('now'));
`,
		`SELECT datetime(i, 'unixepoch'), datetime(c), datetime(b) FROM times`,
		func(t *testing.T, db *sqlittle.DB) [][]string {
			var rows [][]string
			cb := func(r sqlittle.Row) {
				var times [4]time.Time
				if err := r.Scan(&times[0], &times[1], &times[2]); err != nil {
					t.Fatal(err)
				}
				f := "2006-01-02 15:04:05"
				rows = append(rows, []string{
					times[0].UTC().Format(f),
					times[1].UTC().Format(f),
					times[2].UTC().Format(f),
				})
			}
			if err := db.Select("times", cb, "i", "c", "b"); err != nil {
				t.Fatal(err)
			}
			return rows
		},
	)
}
