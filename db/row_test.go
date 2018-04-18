package db

import (
	"errors"
	"reflect"
	"testing"
)

func TestScanString(t *testing.T) {
	row := Row{nil, int64(42), float64(3.14), "world", []byte("hello")}
	vs := [5]string{}
	if err := row.Scan(&vs[0], &vs[1], &vs[2], &vs[3], &vs[4]); err != nil {
		t.Fatal(err)
	}
	if have, want := vs[0], ""; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
	if have, want := vs[1], "42"; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
	if have, want := vs[2], "3.14"; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
	if have, want := vs[3], "world"; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
	if have, want := vs[4], "hello"; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}

func TestScanInt(t *testing.T) {
	row := Row{nil, int64(42), float64(3.14), "3.1415", []byte("2.71828")}
	vs := [5]int{}
	if err := row.Scan(&vs[0], &vs[1], &vs[2], &vs[3], &vs[4]); err != nil {
		t.Fatal(err)
	}
	if have, want := vs[0], 0; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
	if have, want := vs[1], 42; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
	if have, want := vs[2], 3; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
	if have, want := vs[3], 3; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
	if have, want := vs[4], 2; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	var n int
	if have, want := (Row{"hi"}.Scan(&n)), errors.New(`invalid number: "hi"`); !reflect.DeepEqual(have, want) {
		t.Errorf("have %v, want %v", have, want)
	}
	if have, want := (Row{[]byte("bye")}.Scan(&n)), errors.New(`invalid number: "bye"`); !reflect.DeepEqual(have, want) {
		t.Errorf("have %v, want %v", have, want)
	}
}
