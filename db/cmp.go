// docs:
// https://sqlite.org/datatype3.html chapter "4. Comparison Expressions"

package db

import (
	"bytes"
	"strings"
)

// CmpPrefix can be used as a Cmp value
func CmpPrefix(w string) func(string) int {
	lw := len(w)
	return func(s string) int {
		if len(s) > lw {
			s = s[:lw]
		}
		return Compare(w, s)
	}
}

// example collate function
func CollateRtrim(w string) func(string) int {
	tw := strings.TrimRight(w, " \t\n\r")
	return func(s string) int {
		return Compare(tw, strings.TrimRight(s, " \t\n\r"))
	}
}

// NewCmpDesc reverses the comparison logic.
func NewCmpDesc(a interface{}) func(interface{}) int {
	return func(b interface{}) int {
		return -Compare(a, b)
	}
}

// nil, int64, float64, string, []byte

// Compare record values, with ordering according to SQLite's type sort order:
//    nil < {int64|float64} < string < []byte
// Return value follows memcmp and strings.Compare:
//    The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
//
// In addition to the types mentioned, `a` may be a function on b.
func Compare(a, b interface{}) int {
	switch at := a.(type) {
	case nil:
		switch b.(type) {
		case nil:
			return 0
		case int64, float64, string, []byte:
			return -1
		default:
			panic("impossible cmp type")
		}
	case int64:
		switch bt := b.(type) {
		case nil:
			return 1
		case int64:
			return cmpInt64(at, bt)
		case float64:
			return cmpFloat64(float64(at), bt)
		case string, []byte:
			return -1
		default:
			panic("impossible cmp type")
		}
	case float64:
		switch bt := b.(type) {
		case nil:
			return 1
		case int64:
			return cmpFloat64(at, float64(bt))
		case float64:
			return cmpFloat64(at, bt)
		case string, []byte:
			return -1
		default:
			panic("impossible cmp type")
		}
	case string:
		switch bt := b.(type) {
		case nil, int64, float64:
			return 1
		case string:
			return strings.Compare(at, bt)
		case []byte:
			return -1
		default:
			panic("impossible cmp type")
		}
	case []byte:
		switch bt := b.(type) {
		case nil, int64, float64, string:
			return 1
		case []byte:
			return bytes.Compare(at, bt)
		default:
			panic("impossible cmp type")
		}

	// speciality functions
	case func(string) int:
		switch bt := b.(type) {
		case nil, int64, float64:
			return 1
		case string:
			return at(bt)
		case []byte:
			return -1
		default:
			panic("impossible cmp type")
		}
	case func(interface{}) int:
		return at(b)

	default:
		panic("impossible cmp type for a")
	}
}

func cmpInt64(a, b int64) int {
	switch {
	case a < b:
		return -1
	case a == b:
		return 0
	default:
		return 1
	}
}

func cmpFloat64(a, b float64) int {
	switch {
	case a < b:
		return -1
	case a == b:
		return 0
	default:
		return 1
	}
}

// Cmp returns:
//  -1 if the given value is smaller than what we're looking for
//  0 if the given value is what we're looking for
//  1 is the given value is larger than what we're looking for
type Cmp func(interface{}) int

// CmpDesc reverses any Cmp function. For use in DESC indexes
func CmpDesc(c Cmp) Cmp {
	return func(r interface{}) int {
		return -c(r)
	}
}

func CmpString(s string) Cmp {
	return func(r interface{}) int {
		switch r := r.(type) {
		case string:
			return strings.Compare(r, s)
		default:
			panic("cmp string fixme!")
		}
	}
}

func CmpInt64(n int64) Cmp {
	return func(r interface{}) int {
		switch r := r.(type) {
		case int64:
			return cmpInt64(r, n) // TODO: fix cmpInt64!
		default:
			panic("cmp int64 fixme!")
		}
	}
}
