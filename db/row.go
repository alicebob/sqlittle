package db

import (
	"fmt"
	"strconv"

	"github.com/alicebob/sqlittle"
)

// A row with values as stored in the database. Use Row.Scan() to process these
// values.
//
// Values are allowed to point to bytes in the database and hence are
// only valid during a DB transaction.
type Row sqlittle.Record

// Scan converts a row with database values to the Go values you want.
// Supported Go types:
//  - string
//  - int64
//  - int32
//  - int
//  - bool
//  - nil (skips the column)
//
// Conversions are usually stricter than in SQLite:
//  - string to number does not accept trailing letters such as in "123test"
//  - string to bool needs to convert to a number cleanly
//  - numbers are stored as either int64 or float64, and are converted
//    with the normal Go conversions.
//
// Values are a copy of the database bytes; they stay valid even after closing
// the database.
func (r Row) Scan(args ...interface{}) error {
	var err error
	for i, v := range args {
		switch vt := v.(type) {
		case nil:
			// skip
		case *string:
			if len(r) <= i {
				*vt = "" // Or error?
				continue
			}
			switch rv := r[i].(type) {
			case nil:
				*vt = ""
			case int64:
				*vt = strconv.FormatInt(rv, 10)
			case float64:
				*vt = strconv.FormatFloat(rv, 'g', -1, 64)
			case string:
				*vt = rv
			case []byte:
				*vt = string(rv)
			}
		case *int64:
			if len(r) <= i {
				*vt = 0 // Or error?
				continue
			}
			switch rv := r[i].(type) {
			case nil:
				*vt = 0
			case int64:
				*vt = rv
			case float64:
				*vt = int64(rv)
			case string:
				if *vt, err = stringToInt64(rv); err != nil {
					return fmt.Errorf("invalid number: %q", rv)
				}
			case []byte:
				if *vt, err = stringToInt64(string(rv)); err != nil {
					return fmt.Errorf("invalid number: %q", rv)
				}
			}
		case *int32:
			if len(r) <= i {
				*vt = 0 // Or error?
				continue
			}
			switch rv := r[i].(type) {
			case nil:
				*vt = 0
			case int64:
				*vt = int32(rv)
			case float64:
				*vt = int32(rv)
			case string:
				v, err := stringToInt64(rv)
				if err != nil {
					return fmt.Errorf("invalid number: %q", rv)
				}
				*vt = int32(v)
			case []byte:
				v, err := stringToInt64(string(rv))
				if err != nil {
					return fmt.Errorf("invalid number: %q", rv)
				}
				*vt = int32(v)
			}
		case *int:
			if len(r) <= i {
				*vt = 0 // Or error?
				continue
			}
			switch rv := r[i].(type) {
			case nil:
				*vt = 0
			case int64:
				*vt = int(rv)
			case float64:
				*vt = int(rv)
			case string:
				n, err := stringToInt64(rv)
				if err != nil {
					return fmt.Errorf("invalid number: %q", rv)
				}
				*vt = int(n)
			case []byte:
				n, err := stringToInt64(string(rv))
				if err != nil {
					return fmt.Errorf("invalid number: %q", rv)
				}
				*vt = int(n)
			}
		case *bool:
			if len(r) <= i {
				*vt = false // Or error?
				continue
			}
			switch rv := r[i].(type) {
			case nil:
				*vt = false
			case int64:
				*vt = int(rv) != 0
			case float64:
				*vt = int(rv) != 0
			case string:
				n, err := stringToInt64(rv)
				if err != nil {
					return fmt.Errorf("invalid boolean: %q", rv)
				}
				*vt = n != 0
			case []byte:
				n, err := stringToInt64(string(rv))
				if err != nil {
					return fmt.Errorf("invalid boolean: %q", rv)
				}
				*vt = n != 0
			}
		default:
			return fmt.Errorf("unsupported Scan() type: %T", v)
		}
	}
	return nil
}

func stringToInt64(s string) (int64, error) {
	if v, err := strconv.ParseInt(s, 10, 64); err == nil {
		return v, nil
	}
	f, err := strconv.ParseFloat(s, 64)
	return int64(f), err
}

// make a new record from columns from the old record
func reRecord(r sqlittle.Record, indexes []int) sqlittle.Record {
	n := make(sqlittle.Record, len(indexes))
	for i := range n {
		n[i] = r[indexes[i]]
	}
	return n
}
