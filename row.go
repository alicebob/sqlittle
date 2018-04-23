package sqlittle

import (
	"fmt"
	"strconv"

	sdb "github.com/alicebob/sqlittle/db"
)

// A row with values as stored in the database. Use Row.Scan() to process these
// values.
//
// Values are allowed to point to bytes in the database and hence are
// only valid during a DB transaction.
type Row sdb.Record

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
			*vt = r.scanString(i)
		case *int64:
			*vt, err = r.scanInt64(i)
			if err != nil {
				return err
			}
		case *int32:
			n, err := r.scanInt64(i)
			if err != nil {
				return err
			}
			*vt = int32(n)
		case *int:
			n, err := r.scanInt64(i)
			if err != nil {
				return err
			}
			*vt = int(n)
		case *bool:
			n, err := r.scanInt64(i)
			if err != nil {
				return fmt.Errorf("invalid boolean: %q", r[i])
			}
			*vt = n != 0
		default:
			return fmt.Errorf("unsupported Scan() type: %T", v)
		}
	}
	return nil
}

func (r Row) scanString(i int) string {
	if len(r) <= i {
		return "" // Or error?
	}
	switch rv := r[i].(type) {
	case nil:
		return ""
	case int64:
		return strconv.FormatInt(rv, 10)
	case float64:
		return strconv.FormatFloat(rv, 'g', -1, 64)
	case string:
		return rv
	case []byte:
		return string(rv)
	default:
		panic("impossible")
	}
}

func (r Row) scanInt64(i int) (int64, error) {
	if len(r) <= i {
		return 0, nil // Or error?
	}
	switch rv := r[i].(type) {
	case nil:
		return 0, nil
	case int64:
		return rv, nil
	case float64:
		return int64(rv), nil
	case string:
		vt, err := stringToInt64(rv)
		if err != nil {
			return 0, fmt.Errorf("invalid number: %q", rv)
		}
		return vt, nil
	case []byte:
		vt, err := stringToInt64(string(rv))
		if err != nil {
			return 0, fmt.Errorf("invalid number: %q", rv)
		}
		return vt, nil
	default:
		panic("impossible")
	}
}

// ScanString is a shortcut for row.Scan(&string)
func (r Row) ScanString(args ...interface{}) (string, error) {
	var s1 string
	return s1, r.Scan(&s1)
}

// ScanStringString is a shortcut for row.Scan(&string, &string)
func (r Row) ScanStringString(args ...interface{}) (string, string, error) {
	var s1, s2 string
	return s1, s2, r.Scan(&s1, &s2)
}

// ScanStrings is a shortcut to scan all columns as string
//
// Since everything can be converted to strings nothing can possibly go wrong.
func (r Row) ScanStrings(args ...interface{}) []string {
	s := make([]string, len(r))
	for i := range s {
		s[i] = r.scanString(i)
	}
	return s
}

func stringToInt64(s string) (int64, error) {
	if v, err := strconv.ParseInt(s, 10, 64); err == nil {
		return v, nil
	}
	f, err := strconv.ParseFloat(s, 64)
	return int64(f), err
}

// make a key from columns from the record
func asKey(r sdb.Record, indexes []int) sdb.Key {
	n := make(sdb.Key, len(indexes))
	for i := range n {
		n[i] = r[indexes[i]]
	}
	return n
}
