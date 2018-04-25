package sqlittle

import (
	"errors"
	"reflect"
	"testing"

	sdb "github.com/alicebob/sqlittle/db"
	"github.com/alicebob/sqlittle/sql"
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
	test(Key{nil}, cols, sdb.Key{sdb.KeyCol{V: nil}}, nil)
	test(Key{int64(1)}, cols, sdb.Key{sdb.KeyCol{V: int64(1)}}, nil)
	test(Key{3.14}, cols, sdb.Key{sdb.KeyCol{V: 3.14}}, nil)
	test(Key{"foo"}, cols, sdb.Key{sdb.KeyCol{V: "foo"}}, nil)
	test(Key{[]byte("foo")}, cols, sdb.Key{sdb.KeyCol{V: []byte("foo")}}, nil)

	// simple conversions
	test(Key{1}, cols, sdb.Key{sdb.KeyCol{V: int64(1)}}, nil)
	test(Key{int32(42)}, cols, sdb.Key{sdb.KeyCol{V: int64(42)}}, nil)
	test(Key{uint(42)}, cols, sdb.Key{sdb.KeyCol{V: int64(42)}}, nil)
	test(Key{uint32(42)}, cols, sdb.Key{sdb.KeyCol{V: int64(42)}}, nil)
	test(Key{float32(300)}, cols, sdb.Key{sdb.KeyCol{V: float64(300)}}, nil)
	test(Key{true}, cols, sdb.Key{sdb.KeyCol{V: int64(1)}}, nil)
	test(Key{false}, cols, sdb.Key{sdb.KeyCol{V: int64(0)}}, nil)

	twoCols := []sdb.IndexColumn{{Column: "test"}, {Column: "test2", SortOrder: sql.Desc}}
	test(
		Key{int64(1), "foo"},
		twoCols,
		sdb.Key{
			sdb.KeyCol{V: int64(1)},
			sdb.KeyCol{V: "foo", Desc: true},
		},
		nil,
	)
	test(Key{1, 2, 3}, twoCols, nil, errors.New("too many columns in Key"))
}
