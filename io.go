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

func (db *database) pageMaster() (TableBtree, error) {
	buf, err := db.page(1)
	if err != nil {
		return nil, err
	}
	return newTableBtree(buf, true)
}

// n starts a 1, sqlite style
func (db *database) page(id int) ([]byte, error) {
	if id < 1 {
		return nil, errors.New("invalid page number")
	}
	buf := make([]byte, db.header.PageSize)
	n, err := db.f.ReadAt(buf[:], int64(id-1)*int64(db.header.PageSize))
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

type table struct {
	name string
	root TableBtree
	// TODO: point to indices, &c.
}

// master records are defined as:
// CREATE TABLE sqlite_master(
//     type text,
//     name text,
//     tbl_name text,
//     rootpage integer,
//     sql text
// );
type Master struct {
	typ, name, tblName string
	rootPage           int
	sql                string
}

func (db *database) master() ([]Master, error) {
	master, err := db.pageMaster()
	if err != nil {
		return nil, err
	}

	var tables []Master
	_, err = master.Iter(db, func(rowid int64, c []byte) (bool, error) {
		e, err := parseRecord(c)
		if err != nil {
			return false, err
		}
		m := Master{}
		if s, ok := e[0].(string); !ok {
			return false, errors.New("wrong column type for sqlite_master")
		} else {
			m.typ = s
		}
		if s, ok := e[1].(string); !ok {
			return false, errors.New("wrong column type for sqlite_master")
		} else {
			m.name = s
		}
		if s, ok := e[2].(string); !ok {
			return false, errors.New("wrong column type for sqlite_master")
		} else {
			m.tblName = s
		}
		if n, ok := e[3].(int64); !ok {
			return false, errors.New("wrong column type for sqlite_master")
		} else {
			m.rootPage = int(n)
		}
		if s, ok := e[4].(string); !ok {
			return false, errors.New("wrong column type for sqlite_master")
		} else {
			m.sql = s
		}
		tables = append(tables, m)
		return false, nil
	})
	return tables, err
}

func (db *database) Table(name string) (*table, error) {
	tables, err := db.master()
	if err != nil {
		return nil, err
	}
	for _, t := range tables {
		if t.typ == "table" && t.name == name {
			root, err := db.openRoot(t.rootPage)
			if err != nil {
				return nil, err
			}
			return &table{
				name: t.name,
				root: root,
			}, nil
		}
	}
	return nil, nil
}

func (db *database) openRoot(page int) (TableBtree, error) {
	buf, err := db.page(page)
	if err != nil {
		return nil, err
	}
	return newTableBtree(buf, false)
}

func (db *database) addOverflow(length int64, page int, to []byte) ([]byte, error) {
	buf, err := db.page(page)
	if err != nil {
		return nil, err
	}
	next, buf := int(binary.BigEndian.Uint32(buf[:4])), buf[4:]
	to = append(to, buf...)
	if next != 0 {
		return db.addOverflow(length, next, to)
	}
	return to[:length], nil
}
