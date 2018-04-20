package db

import (
	"strings"

	"github.com/alicebob/sqlittle"
)

// for non-rowid tables only:
// given an index gives back the indexes in a row which form the primary key.
func pkColumns(schema *sqlittle.SchemaTable, ind *sqlittle.SchemaIndex) []int {
	if !schema.WithoutRowid {
		panic("can't call pkColumns on an rowid table")
	}

	var res []int
	for _, c := range schema.PK {
		if in := ind.Column(c.Column); in < 0 {
			ind.Columns = append(ind.Columns, c)
			res = append(res, len(ind.Columns)-1)
		} else {
			res = append(res, in)
		}
	}
	return res
}

// given a non-rowid table, gives the order columns are stored on disk
func columnStoreOrder(schema *sqlittle.SchemaTable) []int {
	if !schema.WithoutRowid {
		panic("can't call columnStoreOrder on an rowid table")
	}

	// all PK columns come first, then all other columns, in order
	var cols = make([]string, 0, len(schema.Columns))
	for _, c := range schema.PK {
		cols = append(cols, strings.ToLower(c.Column))
	}
loop:
	for _, c := range schema.Columns {
		n := strings.ToLower(c.Column)
		for _, oc := range cols {
			if oc == n {
				continue loop
			}
		}
		cols = append(cols, n)
	}

	res := make([]int, len(cols))
loop2:
	for i, c := range schema.Columns {
		n := strings.ToLower(c.Column)
		for j, oc := range cols {
			if oc == n {
				res[i] = j
				continue loop2
			}
		}
	}
	return res
}
