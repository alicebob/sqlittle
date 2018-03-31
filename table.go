package sqlit

import (
	"encoding/binary"
	"errors"
)

// Payload represents the payload part of a cell. If overflow is non-zero the
// Payload field will be truncated. Use addOverflow() to get a full payload.
type Payload struct {
	Length   int64
	Payload  []byte
	Overflow int
}

// Iterate callback. Gets rowid and (possibly truncated) payload. Return true when done
type IterCB func(rowid int64, pl Payload) (bool, error)
type TableBtree interface {
	// Iter goes over every record
	Iter(*database, IterCB) (bool, error)
	// Scan starting from a key
	IterMin(*database, int64, IterCB) (bool, error)
	// Rows counts the number of rows. For debugging.
	Rows(*database) (int, error)
}

// IndexIterCB gets the (possibly truncated) payload
type IndexIterCB func(pl Payload) (bool, error)
type IndexBtree interface {
	// Iter goes over every record
	Iter(*database, IndexIterCB) (bool, error)
	// Rows counts the number of rows. For debugging.
	Rows(*database) (int, error)
}

type leafTableBtree struct {
	cells [][]byte
}

type interiorTableBtree struct {
	cells     [][]byte
	rightmost int
}

type leafIndexBtree struct {
	cells [][]byte
}

type interiorIndexBtree struct {
	cells     [][]byte
	rightmost int
}

func newTableBtree(b []byte, isFileHeader bool) (TableBtree, error) {
	hb := b
	if isFileHeader {
		hb = b[headerSize:]
	}
	cells := int(binary.BigEndian.Uint16(hb[3:5]))
	/*
		contentOffset := int(binary.BigEndian.Uint16(hb[5:7]))
		if contentOffset == 0 {
			contentOffset = 65536
		}
	*/
	switch typ := int(hb[0]); typ {
	case 0x0d:
		return newLeafTableBtree(cells, hb[8:], b)
	case 0x05:
		rightmostPointer := int(binary.BigEndian.Uint32(hb[8:12]))
		return newInteriorTableBtree(cells, hb[12:], b, rightmostPointer)
	case 0x0a, 0x02:
		return nil, errors.New("found an index, expected a table")
	default:
		return nil, errors.New("unsupported")
	}
}

func newIndexBtree(b []byte) (IndexBtree, error) {
	cells := int(binary.BigEndian.Uint16(b[3:5]))
	switch typ := int(b[0]); typ {
	case 0x0d, 0x05:
		return nil, errors.New("found a table, expected an index")
	case 0x0a:
		return newLeafIndex(cells, b[8:], b)
	case 0x02:
		rightmostPointer := int(binary.BigEndian.Uint32(b[8:12]))
		return newInteriorIndex(cells, b[12:], b, rightmostPointer)
	default:
		return nil, errors.New("unsupported")
	}
}

func newLeafTableBtree(
	count int,
	pointers []byte,
	content []byte,
) (*leafTableBtree, error) {
	cells, err := parseCellpointers(count, pointers, content)
	return &leafTableBtree{
		cells: cells,
	}, err
}

func (l *leafTableBtree) Rows(*database) (int, error) {
	return len(l.cells), nil
}

func (l *leafTableBtree) Iter(_ *database, cb IterCB) (bool, error) {
	for _, c := range l.cells {
		rowid, pl := parseCellLeaf(c)

		if done, err := cb(rowid, pl); done || err != nil {
			return done, err
		}
	}
	return false, nil
}

func (l *leafTableBtree) IterMin(db *database, rowid int64, cb IterCB) (bool, error) {
	return l.Iter(
		db,
		func(key int64, pl Payload) (bool, error) {
			if key < rowid {
				return false, nil
			}
			return cb(key, pl)
		},
	)
}

func newInteriorTableBtree(
	count int,
	pointers []byte,
	content []byte,
	rightmost int,
) (*interiorTableBtree, error) {
	cells, err := parseCellpointers(count, pointers, content)
	return &interiorTableBtree{
		cells:     cells,
		rightmost: rightmost,
	}, err
}

type interiorIterCB func(page int) (bool, error)

func (l *interiorTableBtree) cellIter(db *database, cb interiorIterCB) (bool, error) {
	for _, c := range l.cells {
		left, _ := parseInteriorLeaf(c)
		if done, err := cb(left); done || err != nil {
			return done, err
		}
	}
	return cb(l.rightmost)
}

func (l *interiorTableBtree) cellIterMin(db *database, rowid int64, cb interiorIterCB) (bool, error) {
	// Loop over all pages, skipping pages which have rows too low.
	// This could be implemented with a nice binary search.
	for _, c := range l.cells {
		left, key := parseInteriorLeaf(c)
		if key < rowid {
			continue
		}
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

func (l *interiorTableBtree) IterMin(db *database, rowid int64, cb IterCB) (bool, error) {
	// we go over the keys, skipping pages with a low max key.
	// This could be implemented with a binary search in the page.
	return l.cellIterMin(db, rowid, func(pageID int) (bool, error) {
		buf, err := db.page(pageID)
		if err != nil {
			return false, err
		}
		page, err := newTableBtree(buf, false)
		if err != nil {
			return false, err
		}
		return page.IterMin(db, rowid, cb)
	})
}

func newLeafIndex(
	count int,
	pointers []byte,
	content []byte,
) (*leafIndexBtree, error) {
	cells, err := parseCellpointers(count, pointers, content)
	return &leafIndexBtree{
		cells: cells,
	}, err
}

func (l *leafIndexBtree) Iter(db *database, cb IndexIterCB) (bool, error) {
	for _, c := range l.cells {
		pl := parseLeafIndex(c)
		if done, err := cb(pl); done || err != nil {
			return done, err
		}
	}
	return false, nil
}

func (l *leafIndexBtree) Rows(*database) (int, error) {
	return len(l.cells), nil
}

func newInteriorIndex(
	count int,
	pointers []byte,
	content []byte,
	rightmost int,
) (*interiorIndexBtree, error) {
	cells, err := parseCellpointers(count, pointers, content)
	return &interiorIndexBtree{
		cells:     cells,
		rightmost: rightmost,
	}, err
}

func (l *interiorIndexBtree) Iter(db *database, cb IndexIterCB) (bool, error) {
	for _, c := range l.cells {
		left, pl := parseInteriorIndex(c)
		page, err := db.openIndex(left)
		if err != nil {
			return false, err
		}
		if done, err := page.Iter(db, cb); done || err != nil {
			return done, err
		}

		// the btree node also has a record
		if done, err := cb(pl); done || err != nil {
			return done, err
		}
	}

	page, err := db.openIndex(l.rightmost)
	if err != nil {
		return false, err
	}
	return page.Iter(db, cb)
}

func (l *interiorIndexBtree) Rows(db *database) (int, error) {
	total := 0
	for _, c := range l.cells {
		left, _ := parseInteriorIndex(c)
		page, err := db.openIndex(left)
		if err != nil {
			return 0, err
		}
		n, err := page.Rows(db)
		if err != nil {
			return 0, err
		}
		total += n
		total += 1 // the btree node has a record, too
	}

	page, err := db.openIndex(l.rightmost)
	if err != nil {
		return 0, err
	}
	n, err := page.Rows(db)
	return total + n, err
}

// parse cell content
// returns total header length, payload
func parseCellLeaf(c []byte) (int64, Payload) {
	l, n := readVarint(c)
	c = c[n:]
	rowid, n := readVarint(c)
	c = c[n:]
	overflow := 0
	if int64(len(c)) != l {
		c, overflow = c[:len(c)-4], int(binary.BigEndian.Uint32(c[len(c)-4:]))
	}
	return rowid, Payload{l, c, overflow}
}

func parseInteriorLeaf(c []byte) (int, int64) {
	left := int(binary.BigEndian.Uint32(c[:4]))
	key, _ := readVarint(c[4:])
	return left, key
}

// returns: payload
func parseLeafIndex(c []byte) Payload {
	l, n := readVarint(c)
	c = c[n:]
	overflow := 0
	if int64(len(c)) != l {
		c, overflow = c[:len(c)-4], int(binary.BigEndian.Uint32(c[len(c)-4:]))
	}
	return Payload{l, c, overflow}
}

// returns: left page, payload length, payload we have, overflow pageid
func parseInteriorIndex(c []byte) (int, Payload) {
	left := int(binary.BigEndian.Uint32(c[:4]))
	c = c[4:]
	l, n := readVarint(c)
	c = c[n:]
	overflow := 0
	if int64(len(c)) != l {
		c, overflow = c[:len(c)-4], int(binary.BigEndian.Uint32(c[len(c)-4:]))
	}
	return int(left), Payload{l, c, overflow}
}

// The readVarint() from encoding/binary is little endian :(
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

// Parse the list of pointers to cells into a slice of byte slices.
// This format is used in all four page types.
// N is the nr of cells, pointers point to the start of the cells, until end of
// the page, content points to the whole page (because cell pointers use page
// offsets).
func parseCellpointers(
	n int,
	pointers []byte,
	content []byte,
) ([][]byte, error) {
	if len(pointers) < n*2 {
		return nil, errors.New("invalid cell pointer array")
	}
	cs := make([][]byte, n)
	end := len(content)
	// cell pointers go [p1, p2, p3], contents goes [c3, c2, c1]
	// SQLite docs aren't too clear about this, though.
	for i := range cs {
		start := int(binary.BigEndian.Uint16(pointers[2*i : 2*i+2]))
		if start > len(content) || start > end {
			return nil, errors.New("invalid cell pointer")
		}
		cs[i] = content[start:end]
		end = start
	}
	return cs, nil
}

// Parse an index cell. Last element of the row is the rowid
func parseIndexRow(pl []byte) (int64, Row, error) {
	row, err := parseRecord(pl)
	if err != nil {
		return 0, nil, err
	}
	if len(row) == 0 {
		return 0, nil, errors.New("no fields in index")
	}
	rowid, ok := row[len(row)-1].(int64)
	if !ok {
		return 0, nil, errors.New("invalid rowid pointer in index")
	}
	row = row[:len(row)-1]
	return rowid, row, nil
}

func addOverflow(db *database, pl Payload) ([]byte, error) {
	to := pl.Payload
	overflow := pl.Overflow
	for {
		if overflow == 0 {
			return to[:pl.Length], nil
		}
		buf, err := db.page(overflow)
		if err != nil {
			return nil, err
		}
		next, buf := int(binary.BigEndian.Uint32(buf[:4])), buf[4:]
		to = append(to, buf...)
		overflow = next
	}
}
