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
	Name    string
	Type    string // as given in the CREATE TABLE
	Null    bool
	Default interface{}
	Collate string
	RowID   bool
	// todo: Affinity string // based on Type
}

type SchemaIndex struct {
	// Name is empty for the primary key in a WITHOUT ROWID table
	Name    string
	Columns []IndexColumn
}

type IndexColumn struct {
	Column    string
	Collate   string
	SortOrder sql.SortOrder
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
			Name:    c.Name,
			Type:    c.Type,
			Null:    c.Null,
			Default: c.Default,
			Collate: c.Collate,
			RowID:   false,
		}
		if c.PrimaryKey {
			col.RowID = (!ct.WithoutRowid) && isRowid(false, c.Type, c.PrimaryKeyDir)
			col.Null = !ct.WithoutRowid && c.Null // w/o rowid forces not null

			name := fmt.Sprintf("sqlite_autoindex_%s_%d", st.Table, autoindex)
			if ct.WithoutRowid {
				name = ""
			}
			if (ct.WithoutRowid || !col.RowID) && st.addIndex(
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
					col.RowID = true
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
		Name:    ci.Index,
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
				st.Indexes[i].Name = ""
			}
			return false
		}
	}
	st.Indexes = append(st.Indexes, SchemaIndex{
		Name:    name,
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
	for i, col := range st.Columns {
		if col.Name == name {
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
