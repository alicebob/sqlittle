package db

import (
	"fmt"
	"strconv"

	"github.com/alicebob/sqlittle"
)

type Row sqlittle.Record

func (r Row) Scan(args ...interface{}) error {
	for i, v := range args {
		switch vt := v.(type) {
		case *string:
			if len(r) < i {
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
		case *int:
			if len(r) < i {
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
					return err
				}
				*vt = int(n)
			case []byte:
				n, err := stringToInt64(string(rv))
				if err != nil {
					return err
				}
				*vt = int(n)
			}
		}
	}
	return nil
}

func stringToInt64(s string) (int64, error) {
	if v, err := strconv.ParseInt(s, 10, 64); err == nil {
		return v, nil
	}
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return int64(f), nil
	}
	return 0, fmt.Errorf("invalid number: %q", s)
}
