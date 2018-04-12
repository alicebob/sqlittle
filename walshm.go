// Support for the shared memory wal index.
// These are the -shm files.

package sqlittle

import (
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"unsafe"

	mmap "launchpad.net/gommap"
)

const (
	shmBlock = 32768
)

var (
	enc binary.ByteOrder
)

func init() {
	// shm is in "local endianness". Can't simply cast a [4]byte to an uint32, so...
	i := 0x1
	bs := (*[strconv.IntSize]byte)(unsafe.Pointer(&i))
	if bs[0] == 0 {
		enc = binary.BigEndian
	} else {
		enc = binary.LittleEndian
	}
}

type shm struct {
	f  *os.File
	mm mmap.MMap
}

func openShm(file string) (*shm, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err // FIXME
	}
	mm, err := mmap.Map(f.Fd(), mmap.PROT_READ, mmap.MAP_SHARED)
	if err != nil {
		f.Close()
		return nil, fmt.Errorf("wal mmap error: %s", err)
	}

	return &shm{
		f:  f,
		mm: mm,
	}, nil
}

func (s *shm) Remap() error {
	fmt.Printf("remap. Old: %d\n", len(s.mm))
	mm, err := mmap.Map(s.f.Fd(), mmap.PROT_READ, mmap.MAP_SHARED)
	if err != nil {
		mm.UnsafeUnmap()
		return fmt.Errorf("wal mmap error: %s", err)
	}
	s.mm = mm
	fmt.Printf("remap. new: %d\n", len(s.mm))
	return nil
}

func (s *shm) Close() {
	s.mm.UnsafeUnmap()
	s.f.Close()
}

// read the n-th uint32 from shm
func (s shm) uint(n int) uint32 {
	return enc.Uint32(s.mm[n*4:])
}

func (s shm) Valid(w walHeader) bool {
	if len(s.mm) < 1 || len(s.mm) != shmBlock {
		// TODO: multiples of shmBlock
		return false
	}
	if iVersion := s.uint(0); iVersion != walFileFormat {
		return false
	}

	// salts are always stored bigendian.
	if salt1 := binary.BigEndian.Uint32(s.mm[8*4 : 9*4]); salt1 != w.salt1 {
		return false
	}
	if salt2 := binary.BigEndian.Uint32(s.mm[9*4 : 10*4]); salt2 != w.salt2 {
		return false
	}
	return true
}

// number of valid frames
func (s *shm) MxFrame() uint32 {
	return s.uint(4)
}

// gives (1-based) frame number for page `page`, upto mxFrame
func (s *shm) Frame(page int, mxFrame uint32) uint32 {
	mx := int(mxFrame)
	nPgno := 4062 // 4096 for later hash blocks
	if mx > nPgno {
		panic("more than 4062 pages is not handled. fixme")
		return 0 // FIXME
	}
	aPgnoOffset := (4096 - nPgno) // 0 for later hash blocks // in 4bytes
	aHashOffset := 4 * 4096       // in bytes
	max := 0
	/*
		// simple scanning algo
		for i := 0; i < mx; i++ {
			switch s.uint(aPgnoOffset + i) {
			case 0:
				return uint32(max)
			case uint32(page):
				max = i
			}
		}
	*/
	// hash algo
	h := (page * 383) % 8192
	// TODO: more byte offset checks
	for {
		frame := int(enc.Uint16(s.mm[aHashOffset+(2*h):]))
		if frame == 0 {
			break
		}
		if frame > mx {
			continue
		}
		p := int(enc.Uint32(s.mm[4*(aPgnoOffset+frame-1):]))
		if p == page && frame > max {
			max = frame
		}
		h = (h + 1) % 8192
	}
	return uint32(max)
}
