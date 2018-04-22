Pure Go SQLite3 file reader

# what

SQLittle can read SQLite3 tables and indexes. It can iterate over tables, and
can efficiently search using indexes.  SQLittle will deal with all SQLite
storage quirks, but otherwise it doesn't try to be smart; if you want to use
an index you have to give the name of the index.

There is no support for SQL, and if you want to do the most efficient joins
possible you'll have to use the low level code.

Based on https://sqlite.org/fileformat2.html and some SQLite source code reading.


# why

This whole thing is mostly for fun. The normal SQLite libraries are perfectly great, and
there is no real need for this. However, since this library is pure Go
cross-compilation is much easier. Given the constraints a valid use-case would
for example be storing app configuration in read-only sqlite files.

# example

simple SELECT over the whole table:

    db, _ := Open("./testdata/music.sqlite")
    defer db.Close()

    cb := func(r Row) {
        var (
            name   string
            length int
        )
        _ = r.Scan(&name, &length)
        fmt.Printf("%s: %d seconds\n", name, length)
    }
    // iterate in rowid order:
    db.Select("tracks", cb, "name", "length")

    // iterate by length, using an index
    db.IndexedSelect("tracks", "tracks_length", cb, "name, "length")


Select a primary key:

    db, _ := Open("./testdata/music.sqlite")
    defer db.Close()

    cb := func(r Row) {
        name, _ := r.ScanString()
        fmt.Printf("%s\n", name)
    }
    db.PKSelect("tracks", Key{int64(4)}, cb, "name")


# docs

https://godoc.org/github.com/alicebob/sqlittle for the go doc and examples.

See [LOWLEVEL.md](LOWLEVEL.md) about the low level reader.
See [CODE.md](CODE.md) for an overview how the code is structured.


# features

- table scan in row order, or table scan in index order, simple searches with
  use of (partial) indexes
- works on both rowid and non-rowid tables
- behaves nicely on corrupted database files (no panics)
- files can be used concurrently with sqlite (compatible locks)
- detects corrupt journal files


# constraints

- read-only
- only supports UTF8 strings
- only supports binary string comparisons
- no joins
- WAL files are not supported
- indexes are used for sorting; no on-the-fly sorting


# locks

SQLittle has a read-lock on the file during the whole execution of a Select function. It's safe to change the database using SQLite while the file is opened in SQLittle.


# status

The current level of abstraction is likely the final one (that is: deal
with reading single tables; don't even try joins or SQL or query planning), but
the API might still change.

TODOs:
- deal with DESC indexes
- deal with collate functions somehow
- the table and index definitions SQL parser is not finished enough
- add some more databases found in the wild to sqlittle-ci
- add a helper to find indexes. That would be especially useful for the
  `sqlite_autoindex_...` indexes
- optimize loading when all requested columns are available in the index
- IndexedSelectCmd()

# &c.

[Travis](https://travis-ci.org/alicebob/sqlittle)

https://github.com/alicebob/sqlittle-ci tests sqlite and sqlittle interaction

`make fuzz` uses [go-fuzz](https://github.com/dvyukov/go-fuzz)

https://github.com/cznic/sqlite2go/ for another approach to pure Go SQLite
