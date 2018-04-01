// Btree page types.
// 'Table' are data tables. InteriorTable pages have no data, and
// points to other pages. InteriorLeaf pages have data and don't point to other
// pages.
// 'Index' tables have index keys. Both the internal and leaf pages contain
// keys.

package sqlittle

import (
	"encoding/binary"
	"errors"
)

// Iterate callback. Gets rowid and (possibly truncated) payload. Return true when done
type iterCB func(rowid int64, pl cellPayload) (bool, error)
type tableBtree interface {
	// Iter goes over every record
	Iter(*Database, iterCB) (bool, error)
	// Scan starting from a key
	IterMin(*Database, int64, iterCB) (bool, error)
	// Count counts the number of records. For debugging.
	Count(*Database) (int, error)
}

// IndexIterCB gets the (possibly truncated) payload
type indexIterCB func(pl cellPayload) (bool, error)
type indexIterMinCB func(rowid int64, row Record) (bool, error)
type indexBtree interface {
	// Iter goes over every record
	Iter(*Database, indexIterCB) (bool, error)
	// Scan starting from an index value
	IterMin(*Database, Record, indexIterMinCB) (bool, error)
	// Count counts the number of records. For debugging.
	Count(*Database) (int, error)
}

type tableLeafCell struct {
	left    int64 // rowID
	payload cellPayload
}
type tableLeaf struct {
	cells []tableLeafCell
}

type tableInteriorCell struct {
	left int
	key  int64
}
type tableInterior struct {
	cells     []tableInteriorCell
	rightmost int
}

type indexLeaf struct {
	cells []cellPayload
}

type indexInteriorCell struct {
	left    int // pageID
	payload cellPayload
}
type indexInterior struct {
	cells     []indexInteriorCell
	rightmost int
}

func newTableBtree(b []byte, isFileHeader bool) (tableBtree, error) {
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

func newIndexBtree(b []byte) (indexBtree, error) {
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
) (*tableLeaf, error) {
	cells, err := parseCellpointers(count, pointers, content)
	if err != nil {
		return nil, err
	}
	leafs := make([]tableLeafCell, len(cells))
	for i, c := range cells {
		leafs[i], err = parseTableLeaf(c)
		if err != nil {
			return nil, err
		}
	}
	return &tableLeaf{
		cells: leafs,
	}, nil
}

func (l *tableLeaf) Count(*Database) (int, error) {
	return len(l.cells), nil
}

func (l *tableLeaf) Iter(_ *Database, cb iterCB) (bool, error) {
	for _, c := range l.cells {
		if done, err := cb(c.left, c.payload); done || err != nil {
			return done, err
		}
	}
	return false, nil
}

func (l *tableLeaf) IterMin(db *Database, rowid int64, cb iterCB) (bool, error) {
	return l.Iter(
		db,
		func(key int64, pl cellPayload) (bool, error) {
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
) (*tableInterior, error) {
	cells, err := parseCellpointers(count, pointers, content)
	if err != nil {
		return nil, err
	}
	cs := make([]tableInteriorCell, len(cells))
	for i, c := range cells {
		cs[i], err = parseTableInterior(c)
		if err != nil {
			return nil, err
		}
	}
	return &tableInterior{
		cells:     cs,
		rightmost: rightmost,
	}, nil
}

type interiorIterCB func(page int) (bool, error)

func (l *tableInterior) cellIter(db *Database, cb interiorIterCB) (bool, error) {
	for _, c := range l.cells {
		if done, err := cb(c.left); done || err != nil {
			return done, err
		}
	}
	return cb(l.rightmost)
}

func (l *tableInterior) cellIterMin(db *Database, rowid int64, cb interiorIterCB) (bool, error) {
	// Loop over all pages, skipping pages which have rows too low.
	// This could be implemented with a nice binary search.
	for _, c := range l.cells {
		if c.key < rowid {
			continue
		}
		if done, err := cb(c.left); done || err != nil {
			return done, err
		}
	}
	return cb(l.rightmost)
}

func (l *tableInterior) Count(db *Database) (int, error) {
	total := 0
	l.cellIter(db, func(p int) (bool, error) {
		page, err := db.openTable(p)
		if err != nil {
			return false, err
		}
		n, err := page.Count(db)
		if err != nil {
			return false, err
		}
		total += n
		return false, nil
	})
	return total, nil
}

func (l *tableInterior) Iter(db *Database, cb iterCB) (bool, error) {
	return l.cellIter(db, func(p int) (bool, error) {
		page, err := db.openTable(p)
		if err != nil {
			return false, err
		}
		if done, err := page.Iter(db, cb); done || err != nil {
			return done, err
		}
		return false, nil
	})
}

func (l *tableInterior) IterMin(db *Database, rowid int64, cb iterCB) (bool, error) {
	// we go over the keys, skipping pages with a low max key.
	// This could be implemented with a binary search in the page.
	return l.cellIterMin(db, rowid, func(pageID int) (bool, error) {
		page, err := db.openTable(pageID)
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
) (*indexLeaf, error) {
	cells, err := parseCellpointers(count, pointers, content)
	if err != nil {
		return nil, err
	}
	cs := make([]cellPayload, len(cells))
	for i, c := range cells {
		cs[i], err = parseIndexLeaf(c)
		if err != nil {
			return nil, err
		}
	}
	return &indexLeaf{
		cells: cs,
	}, nil
}

func (l *indexLeaf) Iter(db *Database, cb indexIterCB) (bool, error) {
	for _, pl := range l.cells {
		if done, err := cb(pl); done || err != nil {
			return done, err
		}
	}
	return false, nil
}

func (l *indexLeaf) IterMin(db *Database, min Record, cb indexIterMinCB) (bool, error) {
	for _, pl := range l.cells {
		cmpRes, rec, err := lazyCmp(db, pl, min)
		if err != nil {
			return false, err
		}
		if cmpRes < 0 {
			continue
		}

		if rec == nil {
			full, err := addOverflow(db, pl)
			if err != nil {
				return false, err
			}
			if rec, err = parseRecord(full); err != nil {
				return false, err
			}
		}

		rowid, rec, err := chompRowid(rec)
		if err != nil {
			return false, err
		}
		if done, err := cb(rowid, rec); done || err != nil {
			return done, err
		}
	}
	return false, nil
}

func (l *indexLeaf) Count(*Database) (int, error) {
	return len(l.cells), nil
}

func newInteriorIndex(
	count int,
	pointers []byte,
	content []byte,
	rightmost int,
) (*indexInterior, error) {
	cells, err := parseCellpointers(count, pointers, content)
	if err != nil {
		return nil, err
	}
	cs := make([]indexInteriorCell, len(cells))
	for i, c := range cells {
		cs[i], err = parseIndexInterior(c)
		if err != nil {
			return nil, err
		}
	}
	return &indexInterior{
		cells:     cs,
		rightmost: rightmost,
	}, nil
}

func (l *indexInterior) Iter(db *Database, cb indexIterCB) (bool, error) {
	for _, c := range l.cells {
		page, err := db.openIndex(c.left)
		if err != nil {
			return false, err
		}
		if done, err := page.Iter(db, cb); done || err != nil {
			return done, err
		}

		// the btree node also has a record
		if done, err := cb(c.payload); done || err != nil {
			return done, err
		}
	}

	page, err := db.openIndex(l.rightmost)
	if err != nil {
		return false, err
	}
	return page.Iter(db, cb)
}

func (l *indexInterior) IterMin(db *Database, min Record, cb indexIterMinCB) (bool, error) {
	for _, c := range l.cells {
		cmpRes, rec, err := lazyCmp(db, c.payload, min)
		if err != nil {
			return false, err
		}
		// on equal we still need to check left on non-unique indexes.
		if cmpRes < 0 {
			continue
		}

		page, err := db.openIndex(c.left)
		if err != nil {
			return false, err
		}
		if done, err := page.IterMin(db, min, cb); done || err != nil {
			return done, err
		}

		if rec == nil {
			full, err := addOverflow(db, c.payload)
			if err != nil {
				return false, err
			}
			if rec, err = parseRecord(full); err != nil {
				return false, err
			}
		}
		rowid, rec, err := chompRowid(rec)
		if err != nil {
			return false, err
		}
		if done, err := cb(rowid, rec); done || err != nil {
			return done, err
		}
	}
	page, err := db.openIndex(l.rightmost)
	if err != nil {
		return false, err
	}
	return page.IterMin(db, min, cb)
}

func (l *indexInterior) Count(db *Database) (int, error) {
	total := 0
	for _, c := range l.cells {
		page, err := db.openIndex(c.left)
		if err != nil {
			return 0, err
		}
		n, err := page.Count(db)
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
	n, err := page.Count(db)
	return total + n, err
}

// shared code for parsing payload from cells
func parsePayload(l int64, c []byte) (cellPayload, error) {
	overflow := 0
	if int64(len(c)) != l {
		if len(c) < 4 {
			return cellPayload{}, ErrCorrupted
		}
		c, overflow = c[:len(c)-4], int(binary.BigEndian.Uint32(c[len(c)-4:]))
	}
	return cellPayload{l, c, overflow}, nil
}

func parseTableLeaf(c []byte) (tableLeafCell, error) {
	l, n := readVarint(c)
	if n < 0 {
		return tableLeafCell{}, ErrCorrupted
	}
	c = c[n:]
	rowid, n := readVarint(c)
	if n < 0 {
		return tableLeafCell{}, ErrCorrupted
	}
	pl, err := parsePayload(l, c[n:])
	return tableLeafCell{
		left:    rowid,
		payload: pl,
	}, err
}

func parseTableInterior(c []byte) (tableInteriorCell, error) {
	if len(c) < 4 {
		return tableInteriorCell{}, ErrCorrupted
	}
	left := int(binary.BigEndian.Uint32(c[:4]))
	key, n := readVarint(c[4:])
	if n < 0 {
		return tableInteriorCell{}, ErrCorrupted
	}
	return tableInteriorCell{
		left: left,
		key:  key,
	}, nil
}

func parseIndexLeaf(c []byte) (cellPayload, error) {
	l, n := readVarint(c)
	if n < 0 {
		return cellPayload{}, ErrCorrupted
	}
	return parsePayload(l, c[n:])
}

func parseIndexInterior(c []byte) (indexInteriorCell, error) {
	if len(c) < 4 {
		return indexInteriorCell{}, ErrCorrupted
	}
	left := int(binary.BigEndian.Uint32(c[:4]))
	c = c[4:]
	l, n := readVarint(c)
	if n < 0 {
		return indexInteriorCell{}, ErrCorrupted
	}
	pl, err := parsePayload(l, c[n:])
	return indexInteriorCell{
		left:    int(left),
		payload: pl,
	}, err
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

// Given a (non-expanded) payload, runs cmp() against it.
// The returned Record may be nil if the non-expanded payload was enough to
// determine the result.
// (that's TODO).
func lazyCmp(db *Database, pl cellPayload, against Record) (int, Record, error) {
	// The TODO idea is that this function doesn't load the overflow page if it
	// can determine the outcome without loading the overflow.
	full, err := addOverflow(db, pl)
	if err != nil {
		return 0, nil, err
	}
	irec, err := parseRecord(full)
	if err != nil {
		return 0, nil, err
	}

	r, err := cmp(irec, against)
	return r, irec, err
}
