package sqlittle

import (
	"fmt"

	sdb "github.com/alicebob/sqlittle/db"
	"github.com/alicebob/sqlittle/sql"
)

// Key is used to find a record.
//
// It accepts most Go datatypes, but they will be converted to the set SQLite
// supports:
// nil, int64, float64, string, []byte
type Key []interface{}

// asDbKey translates a Key to a db.Key. Applies DESC and collate, and changes
// values to the few datatypes db.Key accepts.
// TODO: deal collate
func asDbKey(k Key, cols []sdb.IndexColumn) (sdb.Key, error) {
	var dbk sdb.Key
	for i, kv := range k {
		if i > len(cols)-1 {
			return nil, fmt.Errorf("too many columns in Key")
		}
		add := func(v interface{}) {
			dbk = append(dbk, v)
		}
		c := cols[i]
		if c.SortOrder == sql.Desc {
			oldadd := add
			add = func(v interface{}) {
				oldadd(sdb.KeyDesc(v))
			}
		}
		switch kv := kv.(type) {
		case nil:
			add(nil)
		case int64:
			add(kv)
		case float64:
			add(kv)
		case string:
			add(kv)
		case []byte:
			add(kv)

		case int:
			add(int64(kv))
		case uint:
			add(int64(kv))
		case int32:
			add(int64(kv))
		case uint32:
			add(int64(kv))
		case float32:
			add(float64(kv))
		case bool:
			v := int64(0)
			if kv {
				v = 1
			}
			add(v)

		default:
			return nil, fmt.Errorf("unknown Key datatype: %T", kv)
		}
	}
	return dbk, nil
}
