package sqlit

import (
	"encoding/binary"
	"errors"
	"math/bits"

	"golang.org/x/exp/mmap"
)

const (
	Magic      = "SQLite format 3\x00"
	headerSize = 100
)

var (
	ErrFileZeroLength        = errors.New("file is 0 bytes")
	ErrFileTooShort          = errors.New("file is too short")
	ErrHeaderInvalidMagic    = errors.New("invalid magic number")
	ErrHeaderInvalidPageSize = errors.New("invalid page size")
	ErrFileTruncated         = errors.New("file truncated")
)

type header struct {
	Magic    string
	PageSize int
}

type database struct {
	f      *mmap.ReaderAt
	header header
}

func openFile(f string) (*database, error) {
	r, err := mmap.Open(f)
	if err != nil {
		return nil, err
	}
	if r.Len() == 0 {
		return nil, ErrFileZeroLength
	}

	buf := [headerSize]byte{}
	n, err := r.ReadAt(buf[:], 0)
	if n != headerSize {
		return nil, ErrFileTooShort
	}
	if err != nil {
		return nil, err
	}
	header, err := parseHeader(buf)
	if err != nil {
		return nil, err
	}

	db := &database{
		f:      r,
		header: header,
	}
	return db, nil
}

func (db *database) Close() error {
	return db.f.Close()
}

func (db *database) pageMaster() (*leafTableBtree, error) {
	buf, err := db.page(1)
	if err != nil {
		return nil, err
	}
	return newTableBtree(buf, true)
}

// n starts a 1, sqlite style
func (db *database) page(n int) ([]byte, error) {
	if n < 1 {
		return nil, errors.New("invalid page number")
	}
	buf := make([]byte, db.header.PageSize)
	n, err := db.f.ReadAt(buf[:], (int64(n)-1)*int64(db.header.PageSize))
	if err != nil {
		return nil, err
	}
	if n != len(buf) {
		return nil, ErrFileTruncated
	}
	return buf, nil
}

func parseHeader(b [headerSize]byte) (header, error) {
	magic := string(b[:16])
	if magic != Magic {
		return header{}, ErrHeaderInvalidMagic
	}

	pageSize := uint(binary.BigEndian.Uint16(b[16:18]))
	if pageSize == 1 {
		pageSize = 1 << 16
	}
	isPower := func(n uint) bool {
		return bits.OnesCount(n) == 1
	}
	if pageSize < 512 || pageSize > 1<<16 || !isPower(pageSize) {
		// TODO: special case for 1
		return header{}, ErrHeaderInvalidPageSize
	}

	h := header{
		Magic:    magic,
		PageSize: int(pageSize),
	}
	return h, nil
}
