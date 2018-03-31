package sqlit

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
	"strings"
)

// Row is a list of: nil, int64, float64, string, []byte
// TODO: rename to Record
type Row []interface{}

func parseRecord(r []byte) ([]interface{}, error) {
	var res []interface{}
	hSize, n := readVarint(r)
	header, body := r[n:hSize], r[hSize:]
	for len(header) > 0 {
		c, n := readVarint(header)
		header = header[n:]
		switch c {
		case 0:
			// NULL
			res = append(res, nil)
		case 1:
			// 8-bit twos-complement integer.
			res = append(res, int64(int8(body[0])))
			body = body[1:]
		case 2:
			// Value is a big-endian 16-bit twos-complement integer.
			res = append(res, int64(int16(binary.BigEndian.Uint16(body[:2]))))
			body = body[2:]
		case 3:
			// Value is a big-endian 24-bit twos-complement integer.
			n := int64(uint64(body[0])<<16 | uint64(body[1])<<8 | uint64(body[2]))
			if n&(1<<23) != 0 {
				n -= (1 << 24)
			}
			res = append(res, n)
			body = body[3:]
		case 4:
			// Value is a big-endian 32-bit twos-complement integer.
			res = append(res, int64(int32(binary.BigEndian.Uint32(body[:4]))))
			body = body[4:]
		case 5:
			// Value is a big-endian 48-bit twos-complement integer.
			n := int64(uint64(body[0])<<40 | uint64(body[1])<<32 | uint64(body[2])<<24 |
				uint64(body[3])<<16 | uint64(body[4])<<8 | uint64(body[5]))
			if n&(1<<47) != 0 {
				n -= (1 << 48)
			}
			res = append(res, n)
			body = body[6:]
		case 6:
			// Value is a big-endian 64-bit twos-complement integer.
			res = append(res, int64(binary.BigEndian.Uint64(body[:8])))
			body = body[8:]
		case 7:
			// Value is a big-endian IEEE 754-2008 64-bit floating point number.
			res = append(res, math.Float64frombits(binary.BigEndian.Uint64(body[:8])))
			body = body[8:]
		case 8:
			// Value is the integer 0. (Only available for schema format 4 and higher.)
			res = append(res, int64(0))
		case 9:
			// Value is the integer 1. (Only available for schema format 4 and higher.)
			res = append(res, int64(1))
		case 10, 11:
			// internal types. Should not happen.
			return nil, errors.New("unexpected record type found")
		default:
			if c&1 == 0 {
				// even, blob
				l := (c - 12) / 2
				p := body[:l]
				body = body[l:]
				res = append(res, p)
			} else {
				// odd, string
				// TODO: deal with encoding
				l := (c - 13) / 2
				p := body[:l]
				body = body[l:]
				res = append(res, string(p))
			}
		}
	}
	return res, nil
}

// cmp two rows, according to the 'Record Sort Order' docs.
// Note: strings are always compared binary; no collating functions are used.
// Can return an error on impossible column pair comparison.
func cmp(a, b Row) (int, error) {
	for i, ac := range a {
		if len(b)-1 < i {
			return 1, nil
		}
		if c, err := columnCmp(ac, b[i]); c != 0 || err != nil {
			return c, err
		}
	}
	if len(b) > len(a) {
		return -1, nil
	}
	return 0, nil
}

// returns -1 when a is smaller, 0 when a and b are equal, and else 1
// Column can be: nil, int64, float64, string, []byte
// Becuase float columns might be stored as int that's conversion is supported.
var errCmp = errors.New("impossible comparison")

func columnCmp(a, b interface{}) (int, error) {
	switch at := a.(type) {
	case nil:
		switch b.(type) {
		case nil:
			return 0, nil
		case int64, float64, string, []byte:
			return -1, nil
		default:
			panic("impossible cmp type")
		}
	case int64:
		switch bt := b.(type) {
		case nil:
			return 1, nil
		case int64:
			return cmpInt64(at, bt), nil
		case float64:
			return cmpFloat64(float64(at), bt), nil
		case string, []byte:
			return 0, errCmp
		default:
			panic("impossible cmp type")
		}
	case float64:
		switch bt := b.(type) {
		case nil:
			return 1, nil
		case int64:
			return cmpFloat64(at, float64(bt)), nil
		case float64:
			return cmpFloat64(at, bt), nil
		case string, []byte:
			return 0, errCmp
		default:
			panic("impossible cmp type")
		}
	case string:
		switch bt := b.(type) {
		case nil:
			return 1, nil
		case int64, float64:
			return 0, errCmp
		case string:
			return strings.Compare(at, bt), nil
		case []byte:
			return 0, errCmp
		default:
			panic("impossible cmp type")
		}
	case []byte:
		switch bt := b.(type) {
		case nil:
			return 1, nil
		case int64, float64, string:
			return 0, errCmp
		case []byte:
			return bytes.Compare(at, bt), nil
		default:
			panic("impossible cmp type")
		}
	default:
		panic("impossible cmp type")
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
