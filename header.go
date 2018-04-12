// main .sqlite header

package sqlittle

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/bits"
)

type header struct {
	// The database page size in bytes.
	PageSize int
	// Updated when anything changes (only for non-WAL files).
	ChangeCounter uint32
	// Updated when any table definition changes
	SchemaCookie uint32
	SchemaFormat uint32
	TextEncoding uint32
	// Journal or wal
	Mode int
}

// the file header, as described in "1.2. The Database Header"
// It checks a few constants; see also supported()
func parseHeader(b [headerSize]byte) (header, error) {
	hs := struct {
		Magic                [16]byte
		PageSize             uint16
		_                    uint8 // WriteVersion
		ReadVersion          uint8
		ReservedSpace        uint8
		MaxFraction          uint8
		MinFraction          uint8
		LeafFraction         uint8
		ChangeCounter        uint32
		_                    uint32
		_                    uint32
		_                    uint32
		SchemaCookie         uint32
		SchemaFormat         uint32
		_                    uint32
		_                    uint32
		TextEncoding         uint32
		_                    uint32
		_                    uint32
		_                    uint32
		ReservedForExpansion [20]byte
		_                    uint32
		_                    uint32
	}{}
	if err := binary.Read(bytes.NewBuffer(b[:]), binary.BigEndian, &hs); err != nil {
		return header{}, err
	}

	h := header{}

	if string(hs.Magic[:]) != headerMagic {
		return h, ErrInvalidMagic
	}

	{
		s := uint(hs.PageSize)
		if s == 1 {
			s = 1 << 16
		}
		if !validPageSize(s) {
			return header{}, ErrInvalidPageSize
		}
		h.PageSize = int(s)
	}

	switch hs.ReadVersion {
	case 1:
		h.Mode = ModeJournal
	case 2:
		h.Mode = ModeWal
	default:
		return h, ErrIncompatible
	}

	if int(hs.ReservedSpace) != 0 {
		return h, ErrReservedSpace
	}

	if hs.MaxFraction != 64 ||
		hs.MinFraction != 32 ||
		hs.LeafFraction != 32 {
		return h, ErrIncompatible
	}

	h.ChangeCounter = hs.ChangeCounter

	h.SchemaCookie = hs.SchemaCookie

	h.SchemaFormat = hs.SchemaFormat

	h.TextEncoding = hs.TextEncoding

	for _, v := range hs.ReservedForExpansion {
		if v != 0 {
			return h, ErrIncompatible
		}
	}

	return h, nil
}

func supported(h header) error {
	// 1,2,3,4 are the only valid values.
	switch h.SchemaFormat {
	case 1:
		// Version 1 ignores 'DESC' on indexes.
	case 2, 3, 4:
	default:
		return fmt.Errorf("incompatible schema format: %d", h.SchemaFormat)
	}

	switch h.TextEncoding {
	case 1:
		// UTF8. It's the only thing we currently support
	case 2, 3:
		// UTF16le and UTF16be
		return fmt.Errorf("unsupported text encoding: %d", h.TextEncoding)
	default:
		return fmt.Errorf("unknown text encoding: %d", h.TextEncoding)
	}

	return nil
}

func getHeader(l pager) (header, error) {
	p, err := l.page(1)
	if err != nil {
		return header{}, err
	}
	buf := [headerSize]byte{}
	copy(buf[:], p)
	return parseHeader(buf)
}

func validPageSize(s uint) bool {
	return s >= 512 && s <= 1<<16 && bits.OnesCount(s) == 1
}
