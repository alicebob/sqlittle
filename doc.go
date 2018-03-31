// Package SQLittle provides read-only access to SQLite (version 3) database
// files.
//
// Both tables and index files are supported. The current interface is very low
// level, and you'll have to know the table structure.
// Can't be used on files which are being written to (WAL files are ignored).
package sqlittle
