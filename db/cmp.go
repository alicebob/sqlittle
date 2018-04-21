// docs:
// https://sqlite.org/datatype3.html chapter "4. Comparison Expressions"

package db

import (
	"strings"
)

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
