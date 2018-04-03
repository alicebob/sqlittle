Pure Go low-level SQLite3 file reader

# what

This is (for now) a set of low level routines to read SQLite files. Both 
tables and indices can be read, but there is no support for SQL.

Based on https://sqlite.org/fileformat2.html and empirical evidence.


# why

This whole thing is mostly for fun. The normal SQLite libraries are perfectly great, and
there is no real need for this. However, since this library is pure Go
cross-compilation is much easier. Given the constraints a valid use-case would
for example be storing app configuration in read-only sqlite files.


# docs

https://godoc.org/github.com/alicebob/sqlittle for the go doc and examples.


# features

- low level interface to access tables and indices. Full table/index
  scan and basic search are supported
- behaves nicely on corrupted database files (no panics)
- files can be used concurrently with sqlite (compatible locks)


# constraints

- read-only
- only supports UTF8 strings
- only supports binary string comparisons
- no joins/sorting/ranges
- can only read clean files. No corrupted transaction journals can be present
- does not work with WAL journal mode files (WAL is not the default journal mode)


# low level interface

See [godoc](https://godoc.org/github.com/alicebob/sqlittle) for all available
methods and examples, but the gist of a table scan is:

	db, _ := OpenFile("test/single.sqlite")
	defer db.Close()
	table, _ := db.Table("hello")
	table.Scan(func(rowid int64, rec Record) bool {
			fmt.Printf("row %d: %s\n", rowid, rec[0].(string))
			return false // we want all the rows
    })


## low level locks

If you somehow know that no-one will change the .sqlite file you don't have to
use locks. Otherwise sandwich your logic between database.RLock() and
database.RUnlock() calls. Any *Table or *Index pointer you have is invalid
after database.RUnlock().


# low level sqlite gotchas

The low level routines don't change any fields, they simply pass on how data is
stored in the database by SQLite. Notably that includes:
- float64 columns might be stored as int64
- after an alter table which adds columns a row might miss those new columns
- "integer primary key" columns will be always be stored as `nil` in a table,
  and the rowid should be used as the value
- string indexes are compared with a simple binary comparison, no collating
  functions are used. If a column uses any other collating function for strings
  you can't use the index.


# low level todos

- ~~remove all panics on wrong input~~
- ~~fail on non-utf8 encoding~~
- ~~check all constant header fields~~
- proper ~~page loading abstraction~~/~~page cache~~/index cache
- ~~locks~~
- deal with the reserved region
- refuse to open files with a non-committed/failed non-wal journal
- refuse to open wal journal files
- parse embedded table and index definitions and make them available
