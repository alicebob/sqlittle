package sqlittle

import (
	"errors"
	"golang.org/x/exp/mmap"
)

var (
	ErrFileTruncated = errors.New("file truncated")
)

type loader interface {
	// read the file header bytes. Page size is unknown yet.
	header() ([headerSize]byte, error)
	// load a page from storage. The loader is allowed to use a cache.
	page(n int, pagesize int) ([]byte, error)
	// as it says
	Close() error
}

type mmapLoader mmap.ReaderAt

func newMMapLoader(f string) (*mmapLoader, error) {
	r, err := mmap.Open(f)
	return (*mmapLoader)(r), err
}

func (mm *mmapLoader) header() ([headerSize]byte, error) {
	buf := [headerSize]byte{}
	n, err := (*mmap.ReaderAt)(mm).ReadAt(buf[:], 0)
	if n != headerSize {
		return buf, ErrFileTruncated
	}
	return buf, err
}

func (mm *mmapLoader) page(id int, pagesize int) ([]byte, error) {
	buf := make([]byte, pagesize)
	// pages start counting at 1
	n, err := (*mmap.ReaderAt)(mm).ReadAt(buf[:], int64(id-1)*int64(pagesize))
	if err != nil {
		return buf, err
	}
	if n != len(buf) {
		return buf, ErrFileTruncated
	}
	return buf, nil
}

func (mm *mmapLoader) Close() error {
	return (*mmap.ReaderAt)(mm).Close()
}
