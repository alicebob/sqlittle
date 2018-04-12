package sqlittle

import (
	"errors"
	"reflect"
	"testing"
)

func TestHeader(t *testing.T) {
	// This tests both parseHeader() and supported()

	// hexdump -v -e '/1 "0x%02x, "' -n 100 test/single.sqlite
	base := [headerSize]byte{
		0x53, 0x51, 0x4c, 0x69, 0x74, 0x65, 0x20, 0x66,
		0x6f, 0x72, 0x6d, 0x61, 0x74, 0x20, 0x33, 0x00,
		0x10, 0x00, 0x01, 0x01, 0x00, 0x40, 0x20, 0x20,
		0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x04,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04,
		0x00, 0x2e, 0x1c, 0xb0,
	}
	type change func([headerSize]byte) [headerSize]byte
	type cas struct {
		change    change
		want      header
		err       error
		supported error
	}
	for n, c := range []cas{
		// All fine
		{
			change: func(h [headerSize]byte) [headerSize]byte {
				return h
			},
			want: header{
				PageSize:      4096,
				ChangeCounter: 4,
				SchemaCookie:  1,
				SchemaFormat:  4,
				TextEncoding:  1,
				Mode:          ModeJournal,
			},
			supported: nil,
		},

		// Magic number
		{
			// invalid magic numner
			change: func(h [headerSize]byte) [headerSize]byte {
				h[0] = 's'
				return h
			},
			err: ErrInvalidMagic,
		},

		// PageSize
		{
			// page size 4
			change: func(h [headerSize]byte) [headerSize]byte {
				h[16], h[17] = 0, 4
				return h
			},
			err: errors.New("invalid page size"),
		},
		{
			// page size not a power of two
			change: func(h [headerSize]byte) [headerSize]byte {
				h[17] = 0x12
				return h
			},
			err: ErrInvalidPageSize,
		},
		{
			// page size 0xffff
			change: func(h [headerSize]byte) [headerSize]byte {
				h[16], h[17] = 0xFF, 0xFF
				return h
			},
			err: ErrInvalidPageSize,
		},
		{
			// page size 1 is special case, according to the docs
			change: func(h [headerSize]byte) [headerSize]byte {
				h[16], h[17] = 0, 1
				return h
			},
			want: header{
				PageSize:      0x010000,
				ChangeCounter: 4,
				SchemaCookie:  1,
				SchemaFormat:  4,
				TextEncoding:  1,
				Mode:          ModeJournal,
			},
		},

		// read version
		{
			// read version > 2
			change: func(h [headerSize]byte) [headerSize]byte {
				h[19] = 3
				return h
			},
			err: ErrIncompatible,
		},

		// reserved space
		{
			// test #7
			change: func(h [headerSize]byte) [headerSize]byte {
				h[20] = 0x10
				return h
			},
			err: ErrReservedSpace,
		},

		// constants
		{
			// maximum fraction
			change: func(h [headerSize]byte) [headerSize]byte {
				h[21] = 123
				return h
			},
			err: ErrIncompatible,
		},
		{
			// minimum fraction
			change: func(h [headerSize]byte) [headerSize]byte {
				h[22] = 123
				return h
			},
			err: ErrIncompatible,
		},
		{
			// leaf fraction
			change: func(h [headerSize]byte) [headerSize]byte {
				h[23] = 123
				return h
			},
			err: ErrIncompatible,
		},

		// Schema format numner
		{
			// we do support version 1
			change: func(h [headerSize]byte) [headerSize]byte {
				h[44+3] = 1
				return h
			},
			want: header{
				PageSize:      4096,
				ChangeCounter: 4,
				SchemaCookie:  1,
				SchemaFormat:  1,
				TextEncoding:  1,
				Mode:          ModeJournal,
			},
			supported: nil,
		},
		{
			// invalid value
			change: func(h [headerSize]byte) [headerSize]byte {
				h[44+3] = 0
				return h
			},
			want: header{
				PageSize:      4096,
				ChangeCounter: 4,
				SchemaCookie:  1,
				SchemaFormat:  0,
				TextEncoding:  1,
				Mode:          ModeJournal,
			},
			supported: errors.New("incompatible schema format: 0"),
		},
		{
			// invalid value
			change: func(h [headerSize]byte) [headerSize]byte {
				h[44+3] = 5
				return h
			},
			want: header{
				PageSize:      4096,
				ChangeCounter: 4,
				SchemaCookie:  1,
				SchemaFormat:  5,
				TextEncoding:  1,
				Mode:          ModeJournal,
			},
			supported: errors.New("incompatible schema format: 5"),
		},

		// Text Encoding
		{
			// invalid value
			change: func(h [headerSize]byte) [headerSize]byte {
				h[56+3] = 0
				return h
			},
			want: header{
				PageSize:      4096,
				ChangeCounter: 4,
				SchemaCookie:  1,
				SchemaFormat:  4,
				TextEncoding:  0,
				Mode:          ModeJournal,
			},
			supported: errors.New("unknown text encoding: 0"),
		},
		{
			// invalid value
			change: func(h [headerSize]byte) [headerSize]byte {
				h[56+3] = 2
				return h
			},
			want: header{
				PageSize:      4096,
				ChangeCounter: 4,
				SchemaCookie:  1,
				SchemaFormat:  4,
				TextEncoding:  2,
				Mode:          ModeJournal,
			},
			supported: errors.New("unsupported text encoding: 2"),
		},

		// empty
		{
			// 'Reserved for expansion'. Must be 0s.
			change: func(h [headerSize]byte) [headerSize]byte {
				h[78] = 1
				return h
			},
			err: ErrIncompatible,
		},
	} {
		h, err := parseHeader(c.change(base))
		if have, want := err, c.err; !reflect.DeepEqual(have, want) {
			t.Fatalf("case %d: have %v, want %v", n, have, want)
		}
		if c.err != nil {
			continue
		}
		if have, want := h, c.want; have != want {
			t.Errorf("case %d: have %#v, want %#v", n, have, want)
			continue
		}
		if have, want := supported(h), c.supported; !reflect.DeepEqual(have, want) {
			t.Errorf("case %d: have %#v, want %#v", n, have, want)
			continue
		}
	}
}
