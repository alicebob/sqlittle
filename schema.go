// SchemaTable describes a table and all indexes on that table.
// Both indexes from the `CREATE TABLE` and from any relevant `CREATE INDEX`-es
// are processed.
// It knows the SQLite conventions how tables and indexes are used, such as the
// names for internal indexes, when a column is stored in the rowid, &c.

package sqlittle

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/alicebob/sqlittle/sql"
)

/*
type Schema struct {
	Tables map[string]SchemaTable
	// Indexes map[string]Index
}
*/

type SchemaTable struct {
	Table        string
	WithoutRowid bool
	Columns      []TableColumn
	Indexes      []SchemaIndex
}

type TableColumn struct {
	Column  string
	Type    string // as given in the CREATE TABLE
	Null    bool
	Default interface{}
	Collate string
	Rowid   bool
	// todo: Affinity string // based on Type
}

type SchemaIndex struct {
	// Index is empty for the primary key in a WITHOUT ROWID table
	Index     string
	Columns   []IndexColumn
	PKColumns []int // indexes in Columns. Only filled for non-rowid tables
}

type IndexColumn struct {
	Column    string
	Collate   string
	SortOrder sql.SortOrder
	Rowid     bool
}

func newSchema(table string, master []sqliteMaster) (*SchemaTable, error) {
	var createSQL string
	for _, m := range master {
		if m.typ == "table" && m.name == table {
			createSQL = m.sql
			break
		}
	}
	if createSQL == "" {
		return nil, fmt.Errorf("no such table: %q", table)
	}

	t, err := sql.Parse(createSQL)
	if err != nil {
		return nil, err
	}
	ct, ok := t.(sql.CreateTableStmt)
	if !ok {
		return nil, errors.New("unsupported CREATE TABLE statement")
	}

	st := newCreateTable(ct)

	for _, m := range master {
		if m.typ == "index" && m.tblName == table && m.sql != "" {
			// silently ignore indexes we don't understand
			if t, err := sql.Parse(m.sql); err == nil {
				if ci, ok := t.(sql.CreateIndexStmt); ok {
					st.addCreateIndex(ci)
				}
			}
		}
	}

	st.addRefColumns()

	return st, nil
}

// transform a `create table` statement into a SchemaTable, which knows which
// indexes are used
func newCreateTable(ct sql.CreateTableStmt) *SchemaTable {
	st := &SchemaTable{
		Table:        ct.Table,
		WithoutRowid: ct.WithoutRowid,
	}
	autoindex := 1
	for _, c := range ct.Columns {
		col := TableColumn{
			Column:  c.Name,
			Type:    c.Type,
			Null:    c.Null,
			Default: c.Default,
			Collate: c.Collate,
			Rowid:   false,
		}
		if c.PrimaryKey {
			col.Rowid = (!ct.WithoutRowid) && isRowid(false, c.Type, c.PrimaryKeyDir)
			col.Null = !ct.WithoutRowid && c.Null // w/o rowid forces not null

			name := fmt.Sprintf("sqlite_autoindex_%s_%d", st.Table, autoindex)
			if ct.WithoutRowid {
				name = ""
			}
			if (ct.WithoutRowid || !col.Rowid) && st.addIndex(
				name,
				[]IndexColumn{
					{
						Column:    c.Name,
						SortOrder: c.PrimaryKeyDir,
					},
				},
			) {
				autoindex++
			}
		}
		if c.Unique {
			if st.addIndex(
				fmt.Sprintf("sqlite_autoindex_%s_%d", st.Table, autoindex),
				[]IndexColumn{
					{
						Column:    c.Name,
						SortOrder: sql.Asc,
					},
				},
			) {
				autoindex++
			}
		}
		st.Columns = append(st.Columns, col)
	}
constraint:
	for _, c := range ct.Constraints {
		switch c := c.(type) {
		case sql.TablePrimaryKey:
			if !ct.WithoutRowid && len(c.IndexedColumns) == 1 {
				// is this column an alias for the rowid?
				col := st.column(c.IndexedColumns[0].Column)
				if isRowid(true, col.Type, c.IndexedColumns[0].SortOrder) {
					col.Rowid = true
					continue constraint
				}
			}
			name := fmt.Sprintf("sqlite_autoindex_%s_%d", st.Table, autoindex)
			if ct.WithoutRowid {
				for _, co := range c.IndexedColumns {
					st.column(co.Column).Null = false
				}
				name = ""
			}
			if st.addIndexed(name, c.IndexedColumns) {
				autoindex++
			}
		case sql.TableUnique:
			if st.addIndexed(
				fmt.Sprintf("sqlite_autoindex_%s_%d", st.Table, autoindex),
				c.IndexedColumns,
			) {
				autoindex++
			}
		}
	}

	return st
}

// add `CREATE INDEX` statement to a table
// Does not check for duplicate indexes.
func (st *SchemaTable) addCreateIndex(ci sql.CreateIndexStmt) {
	var cs []IndexColumn
	for _, col := range ci.IndexedColumns {
		cs = append(cs, IndexColumn{
			Column:    col.Column,
			Collate:   col.Collate,
			SortOrder: col.SortOrder,
		})
	}
	st.Indexes = append(st.Indexes, SchemaIndex{
		Index:   ci.Index,
		Columns: cs,
	})
}

// add an index. This is a noop if an equivalent index already exists. Returns
// whether the indexed got added.
func (st *SchemaTable) addIndex(name string, cols []IndexColumn) bool {
	for i, ind := range st.Indexes {
		if reflect.DeepEqual(ind.Columns, cols) {
			if name == "" {
				// special case if we found the 'WITHOUT ROWID' PK later on
				st.Indexes[i].Index = ""
			}
			return false
		}
	}
	st.Indexes = append(st.Indexes, SchemaIndex{
		Index:   name,
		Columns: cols,
	})
	return true
}

// helper for addIndex when you already have an []IndexedColumn.
func (st *SchemaTable) addIndexed(name string, cols []sql.IndexedColumn) bool {
	var cs []IndexColumn
	for _, col := range cols {
		cs = append(cs, IndexColumn{
			Column:    col.Column,
			Collate:   col.Collate,
			SortOrder: col.SortOrder,
		})
	}
	return st.addIndex(name, cs)
}

// Returns the index of the named column, or -1.
func (st *SchemaTable) Column(name string) int {
	u := strings.ToUpper(name)
	for i, col := range st.Columns {
		if strings.ToUpper(col.Column) == u {
			return i
		}
	}
	return -1
}

func (st *SchemaTable) column(name string) *TableColumn {
	n := st.Column(name)
	if n < 0 {
		return nil // you're asking for non-exising columns and for trouble
	}
	return &st.Columns[n]
}

// NamedIndex returns the index with the name (case insensitive)
func (st *SchemaTable) NamedIndex(name string) *SchemaIndex {
	u := strings.ToUpper(name)
	for i, ind := range st.Indexes {
		if strings.ToUpper(ind.Index) == u {
			return &st.Indexes[i]
		}
	}
	return nil
}

// addRefColumns adds all columns sqlite adds to indexes.
//  - for rowid tables that's a rowid column in every index
//  - for non-rowid tables that's all primary key columns which are not already
//  present in the index.
//
// This can only be done once the primary key is known.
func (st *SchemaTable) addRefColumns() {
	if st.WithoutRowid {
		if pk := st.NamedIndex(""); pk != nil {
			// pk can't be nil, really

			// add all missing PK columns to the indexes
			for i, ind := range st.Indexes {
				for _, c := range pk.Columns {
					if in := ind.Column(c.Column); in < 0 {
						ind.Columns = append(ind.Columns, c)
						ind.PKColumns = append(ind.PKColumns, len(ind.Columns)-1)
					} else {
						ind.PKColumns = append(ind.PKColumns, in)
					}
				}
				st.Indexes[i] = ind
			}

			// PK gets all the column from the table
			// Effectively it gets the column-store-order of the whole table
			// (which can be different from the column order, if you primary
			// key doesn't use the columns in the same order)
			for _, c := range st.Columns {
				found := false
				for _, pkc := range pk.Columns {
					if strings.ToUpper(c.Column) == strings.ToUpper(pkc.Column) {
						found = true
						break
					}
				}
				if !found {
					pk.Columns = append(pk.Columns, IndexColumn{Column: c.Column})
				}
			}
		}
	} else {
		for i, s := range st.Indexes {
			s.Columns = append(s.Columns, IndexColumn{Rowid: true})
			st.Indexes[i] = s
		}
	}
}

// Returns the index of the named column, or -1.
func (si *SchemaIndex) Column(name string) int {
	u := strings.ToUpper(name)
	for i, col := range si.Columns {
		if strings.ToUpper(col.Column) == u {
			return i
		}
	}
	return -1
}

// A primary key can be an alias for the rowid iff:
//  - this is not a `WITHOUT ROWID` table (not tested here)
//  - it's a single column of type 'INTEGER'
//  - ASC and DESC are fine for table constraints:
//     CREATE TABLE foo (a integer, primary key (a DESC))
//    but in a column statement it only works with a ASC:
//     CREATE TABLE foo (a integer primary key)
//    invalid:
//     ~~CREATE TABLE foo (a integer primary key DESC)~~
// If the row is an alias for the rowid it won't be stored in the datatable;
// all values will be null.
// See https://sqlite.org/lang_createtable.html#rowid
func isRowid(tableConstraint bool, typ string, dir sql.SortOrder) bool {
	if strings.ToUpper(typ) != "INTEGER" {
		return false
	}
	return tableConstraint || dir == sql.Asc
}
