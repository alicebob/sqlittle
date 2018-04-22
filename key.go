package sqlittle

import (
	sdb "github.com/alicebob/sqlittle/db"
)

// Key is used to find a record.
type Key []interface{}

// TODO: deal with DESC columns, collate, and accept more basic datatypes.
func asDbKey(k Key) sdb.Key {
	return sdb.Key(k)
}
