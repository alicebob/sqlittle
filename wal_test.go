package sqlittle

import (
	"encoding/binary"
	"reflect"
	"testing"
)

func TestWalHeader(t *testing.T) {
	// hexdump -v -e '/1 "0x%02x, "' -n 32 test/wal_crashed.sqlite-wal
	base := [walHeaderSize]byte{
		0x37, 0x7f, 0x06, 0x82, 0x00, 0x2d, 0xe2, 0x18,
		0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x31, 0xa6, 0x02, 0xf2, 0xe0, 0xc1, 0xc7, 0xfd,
		0xc6, 0xe9, 0xa9, 0x44, 0x7c, 0x99, 0x42, 0xc9,
	}
	type change func(*[walHeaderSize]byte)
	type cas struct {
		change change
		want   walHeader
		err    error
	}
	for n, c := range []cas{
		// All fine
		{
			change: func(h *[walHeaderSize]byte) {},
			want: walHeader{
				pageSize: 0x1000,
				// CheckpointSeq: 0x0,
				salt1:    0x31a602f2,
				salt2:    0xe0c1c7fd,
				checksum: checksum{binary.LittleEndian, 0xc6e9a944, 0x7c9942c9},
			},
		},
		// magic nr
		{
			change: func(h *[walHeaderSize]byte) {
				h[1] = 'q'
			},
			err: ErrInvalidWal,
		},
		// file format
		{
			change: func(h *[walHeaderSize]byte) {
				h[7] = 'q'
			},
			err: ErrInvalidWal,
		},
		// page size
		{
			change: func(h *[walHeaderSize]byte) {
				h[11] = 0xff
			},
			err: ErrInvalidWal,
		},
		// change salt1 to mess up the checksum
		{
			change: func(h *[walHeaderSize]byte) {
				h[16] = 0xff
			},
			err: ErrInvalidWal,
		},
		// change salt2 to mess up the checksum
		{
			change: func(h *[walHeaderSize]byte) {
				h[21] = 0xff
			},
			err: ErrInvalidWal,
		},
		// mess up the checksum
		{
			change: func(h *[walHeaderSize]byte) {
				h[25] = 0xff
			},
			err: ErrInvalidWal,
		},
	} {
		header := base
		c.change(&header)
		h, err := parseWalHeader(header)
		if have, want := err, c.err; !reflect.DeepEqual(have, want) {
			t.Fatalf("case %d: have %v, want %v", n, have, want)
		}
		if c.err != nil {
			continue
		}
		if have, want := h, c.want; have != want {
			t.Fatalf("case %d: have %#v, want %#v", n, have, want)
		}
	}
}

/*
func TestWalRead(t *testing.T) {
	file := "./test/wal_crashed.sqlite-wal"
	w, err := readWal(file)
	if err != nil {
		t.Fatal(err)
	}
	defer w.Close()

	if err := w.ReadCommits(); err != nil {
		t.Fatal(err)
	}
	if have, want := w.pages, (pageMap{
		1: 2, 2: 3, 3: 4, 4: 5, 5: 6, 6: 7,
	}); !reflect.DeepEqual(have, want) {
		t.Fatalf("have %v, want %v", have, want)
	}
}
*/
