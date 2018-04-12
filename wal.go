package sqlittle

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"

	mmap "launchpad.net/gommap"
)

const (
	walHeaderSize      = 32
	walFrameHeaderSize = 24
	walMagicLE         = 0x377f0682
	walMagicBE         = 0x377f0683
	walFileFormat      = 3007000
)

var (
	ErrInvalidWal = errors.New("invalid wal file")
)

// wal is an open -wal file. pages is the pagenr->framenr map, upto-including
// the last valid commit found in the wal file.
type wal struct {
	f      *os.File
	mm     mmap.MMap
	header walHeader
	// pages          pageMap  // of all frames up to mxFrame
	// mxFrame        int      // last frame nr + 1. Aka the next frame to read
	// commitChecksum checksum // of everything up to mxFrame
}

type pageMap map[int]int

type walHeader struct {
	pageSize     uint32
	salt1, salt2 uint32
	checksum     checksum
}

// open WAL file. `file` needs to have -wal appended.
func openWal(file string) (*wal, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err // FIXME
	}
	wal := &wal{
		f: f,
	}

	if err := wal.Remap(); err != nil {
		f.Close()
		return nil, fmt.Errorf("wal mmap error: %s", err)
	}

	header := [walHeaderSize]byte{}
	if n := copy(header[:], wal.mm); n != walHeaderSize {
		wal.mm.UnsafeUnmap()
		f.Close()
		return nil, errors.New("wal too small")
	}
	h, err := parseWalHeader(header)
	if err != nil {
		return nil, err
	}
	wal.header = h
	return wal, nil
}

func (w *wal) Remap() error {
	mm, err := mmap.Map(w.f.Fd(), mmap.PROT_READ, mmap.MAP_SHARED)
	if err != nil {
		return fmt.Errorf("wal mmap error: %s", err)
	}
	w.mm = mm
	return nil
}

func (w *wal) Close() {
	w.mm.UnsafeUnmap()
	w.f.Close()
}

func parseWalHeader(b [walHeaderSize]byte) (walHeader, error) {
	wal := struct {
		Magic         uint32
		FileFormat    uint32
		PageSize      uint32
		CheckpointSeq uint32
		Salt1         uint32
		Salt2         uint32
		Checksum1     uint32
		Checksum2     uint32
	}{}
	if err := binary.Read(bytes.NewBuffer(b[:]), binary.BigEndian, &wal); err != nil {
		return walHeader{}, err
	}

	var enc binary.ByteOrder
	switch wal.Magic {
	case walMagicLE:
		enc = binary.LittleEndian
	case walMagicBE:
		enc = binary.BigEndian
	default:
		return walHeader{}, ErrInvalidWal
	}

	if wal.FileFormat != walFileFormat {
		return walHeader{}, ErrInvalidWal
	}

	if !validPageSize(uint(wal.PageSize)) {
		return walHeader{}, ErrInvalidWal
	}

	check := checksum{enc, 0, 0}.Add(b[:24])
	if wal.Checksum1 != check.s0 || wal.Checksum2 != check.s1 {
		return walHeader{}, ErrInvalidWal
	}
	return walHeader{
		pageSize: wal.PageSize,
		salt1:    wal.Salt1,
		salt2:    wal.Salt2,
		checksum: check,
	}, nil
}

/*
func readWal(file string) (*wal, error) {
	f, err := mmap.Open(file)
	if err != nil {
		return nil, err
	}
	header := [walHeaderSize]byte{}
	if _, err := f.ReadAt(header[:], 0); err != nil {
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}
	h, err := parseWalHeader(header)
	if err != nil {
		return nil, err
	}
	return &wal{
		f:              f,
		header:         h,
		pages:          pageMap{},
		mxFrame:        0,
		commitChecksum: h.checksum,
	}, nil
}
*/

/*
// Read all (new) valid commits and update w.pages. You can call this as often
// as you like, as long as the main wal header has not changed.
// Not calling this function will effectively give you a read transaction.
func (w *wal) ReadCommits() error {
	for {
		pl, check, err := w.readCommit(w.mxFrame, w.commitChecksum)
		if err != nil {
			return err
		}
		if len(pl) == 0 {
			return nil
		}
		for _, p := range pl {
			w.pages[p] = w.mxFrame
			w.mxFrame++
		}
		w.commitChecksum = check
	}
}
*/

// read the list of pages starting at frame `frame`. Will return at the first
// commit frame. If there is no (final) commit frame it will return an empty
// list.
func (w *wal) readCommit(frame int, check checksum) ([]int, checksum, error) {
	type walFrame struct {
		Page              uint32
		DatabasePageCount uint32
		Salt1             uint32
		Salt2             uint32
		Checksum1         uint32
		Checksum2         uint32
	}
	var pl []int
	for {
		buf, err := w.frame(frame)
		if buf == nil || err != nil {
			return nil, check, err
		}
		f := walFrame{}
		if err := binary.Read(bytes.NewBuffer(buf), binary.BigEndian, &f); err != nil {
			return nil, check, err
		}
		if f.Salt1 != w.header.salt1 || f.Salt2 != w.header.salt2 {
			return nil, check, err
		}
		check = check.Add(buf[:8])
		check = check.Add(buf[walFrameHeaderSize:])
		if f.Checksum1 != check.s0 || f.Checksum2 != check.s1 {
			return nil, check, err
		}
		pl = append(pl, int(f.Page))
		if f.DatabasePageCount != 0 {
			return pl, check, nil
		}
		frame++
	}
}

// return content bytes of frame nr (starts counting at 1)
func (w *wal) frame(i int) ([]byte, error) {
	size := walFrameHeaderSize + int64(w.header.pageSize)
	offset := walHeaderSize + int64(i-1)*size
	fmt.Printf("load frame %d [%d->%d]\n", i, offset, size)
	return w.mm[offset+walFrameHeaderSize : offset+size], nil
}

type checksum struct {
	enc    binary.ByteOrder
	s0, s1 uint32
}

func (c checksum) Add(b []byte) checksum {
	for len(b) >= 8 {
		c.s0 += c.enc.Uint32(b) + c.s1
		b = b[4:]
		c.s1 += c.enc.Uint32(b) + c.s0
		b = b[4:]
	}
	return c
}
