package sqlit

import (
	"encoding/binary"
	"errors"
)

// Iterate callback. Return true when done.
type IterCB func(rowid int64, row []byte) bool

type TableBtree interface {
	Iter(IterCB)
	Rows() int
}

// Knuth B*-Tree, leaf
type leafTableBtree struct {
	cellCount    int
	cellPointers []byte
	_content     []byte
}

func newTableBtree(b []byte, isFileHeader bool) (*leafTableBtree, error) {
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
	case 13:
		return newLeafTableBtree(cells, hb[8:8+2*cells], b), nil
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

func (l *leafTableBtree) Rows() int {
	i := 0
	l.Iter(func(int64, []byte) bool {
		i++
		return false
	})
	return i
}

func (l *leafTableBtree) Iter(cb IterCB) {
	end := len(l._content)
	// cell pointers go: [p1, p2, p3], contents goes [c3, c2, c1]
	// sqlite docs aren't too clear about this, though.
	for i := 0; i < l.cellCount; i++ {
		start := int(binary.BigEndian.Uint16(l.cellPointers[2*i : 2*i+2]))
		c := l._content[start:end]
		rowid, content := parseCellTableLeaf(c)
		if cb(rowid, content) {
			return
		}
		end = start
	}
}

// parse cell content
// TODO: overflow
func parseCellTableLeaf(c []byte) (int64, []byte) {
	l, c := readVarint(c)
	rowid, c := readVarint(c)
	if int64(len(c)) != l {
		panic("overflow!")
	}
	return rowid, c
}

// readVarint from encoding/binary is little endian :(
func readVarint(b []byte) (int64, []byte) {
	var n uint64
	for i := 0; ; i++ {
		if i >= len(b) {
			return int64(n), nil
		}
		c := b[i]
		if i == 8 {
			n = (n << 8) | uint64(c)
			return int64(n), b[i+1:]
		}
		n = (n << 7) | uint64(c&0x7f)
		if c < 0x80 {
			return int64(n), b[i+1:]
		}
	}
}
