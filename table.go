package sqlit

import (
	"encoding/binary"
	"errors"
	"fmt"
)

// Iterate callback. Return true when done.
type IterCB func(rowid int64, row []byte) (bool, error)

type TableBtree interface {
	// Iter goes over every record
	Iter(*database, IterCB) (bool, error)
	// Rows counts the number of rows
	Rows(*database) (int, error)
}

type leafTableBtree struct {
	cellCount    int
	cellPointers []byte
	_content     []byte
}
type interiorTableBtree struct {
	cellCount    int
	cellPointers []byte
	_content     []byte
	rightmost    int
}

func newTableBtree(b []byte, isFileHeader bool) (TableBtree, error) {
	hb := b
	if isFileHeader {
		hb = b[headerSize:]
	}
	cells := int(binary.BigEndian.Uint16(hb[3:5]))
	contentOffset := int(binary.BigEndian.Uint16(hb[5:7]))
	if contentOffset == 0 {
		contentOffset = 65536
	}
	switch typ := int(hb[0]); typ {
	case 0x0d:
		return newLeafTableBtree(cells, hb[8:8+2*cells], b), nil
	case 0x05:
		rightmostPointer := int(binary.BigEndian.Uint32(hb[8:12]))
		return newInteriorTableBtree(cells, hb[12:8+2*cells], b, rightmostPointer), nil
	default:
		return nil, errors.New("unsupported")
	}
}

func newLeafTableBtree(cellCount int, cellPointers []byte, content []byte) *leafTableBtree {
	return &leafTableBtree{
		cellCount:    cellCount,
		cellPointers: cellPointers,
		_content:     content,
	}
}

func (l *leafTableBtree) Rows(*database) (int, error) {
	return l.cellCount, nil
}

func (l *leafTableBtree) Iter(_ *database, cb IterCB) (bool, error) {
	end := len(l._content)
	// cell pointers go [p1, p2, p3], contents goes [c3, c2, c1]
	// SQLite docs aren't too clear about this, though.
	for i := 0; i < l.cellCount; i++ {
		start := int(binary.BigEndian.Uint16(l.cellPointers[2*i : 2*i+2]))
		c := l._content[start:end]
		rowid, content := parseCellTableLeaf(c)
		if done, err := cb(rowid, content); done || err != nil {
			return done, err
		}
		end = start
	}
	return false, nil
}

func newInteriorTableBtree(cellCount int, cellPointers []byte, content []byte, rightmost int) *interiorTableBtree {
	return &interiorTableBtree{
		cellCount:    cellCount,
		cellPointers: cellPointers,
		_content:     content,
		rightmost:    rightmost,
	}
}

type interiorIterCB func(left int) (bool, error)

func (l *interiorTableBtree) cellIter(db *database, cb interiorIterCB) (bool, error) {
	end := len(l._content)
	// cell pointers go [p1, p2, p3], contents goes [c3, c2, c1]
	// SQLite docs aren't too clear about this, though.
	for i := 0; i < l.cellCount; i++ {
		start := int(binary.BigEndian.Uint16(l.cellPointers[2*i : 2*i+2]))
		c := l._content[start:end]
		left, _ := parseInteriorTableLeaf(c)
		if done, err := cb(left); done || err != nil {
			return done, err
		}
	}
	return cb(l.rightmost)
}

func (l *interiorTableBtree) Rows(db *database) (int, error) {
	total := 0
	l.cellIter(db, func(p int) (bool, error) {
		buf, err := db.page(p)
		if err != nil {
			return false, err
		}
		page, err := newTableBtree(buf, false)
		if err != nil {
			return false, err
		}
		n, err := page.Rows(db)
		if err != nil {
			return false, err
		}
		total += n
		return false, nil
	})
	return total, nil
}

func (l *interiorTableBtree) Iter(db *database, cb IterCB) (bool, error) {
	return l.cellIter(db, func(p int) (bool, error) {
		buf, err := db.page(p)
		if err != nil {
			return false, err
		}
		page, err := newTableBtree(buf, false)
		if err != nil {
			return false, err
		}
		if done, err := page.Iter(db, cb); done || err != nil {
			return done, err
		}
		return false, nil
	})
}

// parse cell content
// TODO: overflow
func parseCellTableLeaf(c []byte) (int64, []byte) {
	l, n := readVarint(c)
	c = c[n:]
	rowid, n := readVarint(c)
	c = c[n:]
	if int64(len(c)) != l {
		panic("overflow!")
	}
	return rowid, c
}

func parseInteriorTableLeaf(c []byte) (int, int64) {
	left := int(binary.BigEndian.Uint32(c[:4]))
	key, _ := readVarint(c[4:])
	return left, key
}

// readVarint from encoding/binary is little endian :(
func readVarint(b []byte) (int64, int) {
	var n uint64
	for i := 0; ; i++ {
		if i >= len(b) {
			return int64(n), i
		}
		c := b[i]
		if i == 8 {
			n = (n << 8) | uint64(c)
			return int64(n), i + 1
		}
		n = (n << 7) | uint64(c&0x7f)
		if c < 0x80 {
			return int64(n), i + 1
		}
	}
}

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
		case 2, 3, 4, 5, 6, 7, 8, 9:
			return nil, fmt.Errorf("unimplemented record type: %d. fix me!", c)
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
