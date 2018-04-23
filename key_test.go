package sqlittle

import (
	"errors"
	"reflect"
	"testing"

	sdb "github.com/alicebob/sqlittle/db"
)

func TestKeys(t *testing.T) {
	test := func(k Key, cols []sdb.IndexColumn, want sdb.Key, wantErr error) {
		t.Helper()

		have, err := asDbKey(k, cols)
		if !reflect.DeepEqual(err, wantErr) {
			t.Errorf("err: have %v, want %v", err, wantErr)
			return
		}
		if wantErr != nil {
			return
		}
		if !reflect.DeepEqual(have, want) {
			t.Errorf("have %v, want %v", have, want)
		}
	}

	cols := []sdb.IndexColumn{{Column: "test"}}
	// basic types
	test(Key{nil}, cols, sdb.Key{nil}, nil)
	test(Key{int64(1)}, cols, sdb.Key{int64(1)}, nil)
	test(Key{3.14}, cols, sdb.Key{3.14}, nil)
	test(Key{"foo"}, cols, sdb.Key{"foo"}, nil)
	test(Key{[]byte("foo")}, cols, sdb.Key{[]byte("foo")}, nil)

	// simple conversions
	test(Key{1}, cols, sdb.Key{int64(1)}, nil)
	test(Key{int32(42)}, cols, sdb.Key{int64(42)}, nil)
	test(Key{uint(42)}, cols, sdb.Key{int64(42)}, nil)
	test(Key{uint32(42)}, cols, sdb.Key{int64(42)}, nil)
	test(Key{float32(300)}, cols, sdb.Key{float64(300)}, nil)
	test(Key{true}, cols, sdb.Key{int64(1)}, nil)
	test(Key{false}, cols, sdb.Key{int64(0)}, nil)

	twoCols := []sdb.IndexColumn{{Column: "test"}, {Column: "test2"}}
	test(Key{int64(1), "foo"}, twoCols, sdb.Key{int64(1), "foo"}, nil)
	test(Key{1, 2, 3}, twoCols, nil, errors.New("too many columns in Key"))
}
