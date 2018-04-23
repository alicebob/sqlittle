package sqlittle

import (
	"fmt"

	sdb "github.com/alicebob/sqlittle/db"
)

// Key is used to find a record.
//
// It accepts most Go datatypes, but they will be converted to the set SQLite
// supports:
// nil, int64, float64, string, []byte
type Key []interface{}

// TODO: deal with DESC columns, collate, and accept more basic datatypes.
// asDbKey translates a Key to a db.Key. Applies DESC and collate, and changes
// values to the few datatypes db.Key accepts.
func asDbKey(k Key) (sdb.Key, error) {
	var dbk sdb.Key
	for _, kv := range k {
		switch kv := kv.(type) {
		case nil:
			dbk = append(dbk, nil)
		case int64:
			dbk = append(dbk, kv)
		case float64:
			dbk = append(dbk, kv)
		case string:
			dbk = append(dbk, kv)
		case []byte:
			dbk = append(dbk, kv)

		case int:
			dbk = append(dbk, int64(kv))
		case uint:
			dbk = append(dbk, int64(kv))
		case int32:
			dbk = append(dbk, int64(kv))
		case uint32:
			dbk = append(dbk, int64(kv))
		case float32:
			dbk = append(dbk, float64(kv))
		case bool:
			v := int64(0)
			if kv {
				v = 1
			}
			dbk = append(dbk, v)

		default:
			return nil, fmt.Errorf("unknown Key datatype: %T", kv)
		}
	}
	return dbk, nil
}
