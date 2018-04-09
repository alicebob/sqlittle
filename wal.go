package sqlittle

import (
	"bytes"
	"encoding/binary"
	"errors"
)

const (
	walHeaderSize = 32
	walMagicLE    = 0x377f0682
	walMagicBE    = 0x377f0683
	walFileFormat = 3007000
)

var (
	ErrInvalidWal = errors.New("invalid wal file")
)

type walHeader struct {
	Magic         uint32
	FileFormat    uint32
	PageSize      uint32
	CheckpointSeq uint32
	Salt1         uint32
	Salt2         uint32
	Checksum1     uint32
	Checksum2     uint32
}

func parseWalHeader(b [walHeaderSize]byte) (walHeader, error) {
	wal := walHeader{}
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
		return wal, ErrInvalidWal
	}

	if wal.FileFormat != walFileFormat {
		return wal, ErrInvalidWal
	}

	if !validPageSize(uint(wal.PageSize)) {
		return wal, ErrInvalidWal
	}

	s0, s1 := walChecksum(enc, 0, 0, b[:24])
	if wal.Checksum1 != s0 || wal.Checksum2 != s1 {
		return wal, ErrInvalidWal
	}

	return wal, nil
}

func walChecksum(enc binary.ByteOrder, s0, s1 uint32, b []byte) (uint32, uint32) {
	for len(b) >= 8 {
		s0 += enc.Uint32(b) + s1
		b = b[4:]
		s1 += enc.Uint32(b) + s0
		b = b[4:]
	}
	return s0, s1
}
