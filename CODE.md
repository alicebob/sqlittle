## code structure

This document explains the layers of sqlittle.

This is how the files work together, with the lowest level on the
bottom:

    +---------------------------------+
    | (SQL interface)                 |
    +---------------------------------+
    | (plan builder)                  |
    +---------------------------------+
    | (plan executer)                 |
    +----------------+----------------+
    | low.go         | schema.go      |
    +----------------+----------------+
    | database.go                     |
    +---------------------------------+
    | btree.go                        |
    +---------------------------------+
    | pager (pager.go, pager_unix.go) |
    +---------------------------------+

Things in parenthesis are imaginary future.

    
### pager

An `.sqlite` file consist of a sequence of pages all of the same size. Page
size is between 512 and 64K.  The `Pager` reads pages from the .sqlite database
file. It also knows how to lock that file. It knows nothing about what's in the
pages.

The Go interface is `Pager{}`, which is implemented on `pager_unix.go`. There
is an alternative in-memory implementation in the Fuzz Test. Feel free to
create a pager_windows.go if you need windows support.

### btree

Both table data and indexes are stored in binary trees, which are stored in
pages. The btree code knows how to interpret the bytes in pages. The btree code
has very low level routines to iterate and search in tables and indexes.

### database

`Database` is the main struct, which connects the pager and the btree code. It
also knows where to find the `sqlite_master` table, which stores all table
definitions. Database also deals with the caching of pages.

### low

The routines in `low.go` are the public low level routines. They mostly wrap
the iteration routines from the btree code into something a bit more friendly. 

### schema

The table and indexed definitions are stored by SQLite as `CREATE TABLE ...`
and `CREATE INDEX ...` statements in the database file. `schema.go` uses the
SQL parser from the sql/ subdir to parse those statements, and interprets the
result the same way SQLite does.

With the result you could test whether a table matches what you think it does
when you use the low level scan routines. It could also be used to build more
flexible query code.

### future

indeed
