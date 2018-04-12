package sqlittle

import (
	"io/ioutil"
	"testing"
)

func TestWalShm(t *testing.T) {
	file := "./test/wal_crashed.sqlite"
	wal, err := openWal(file + "-wal")
	if err != nil {
		t.Fatal(err)
	}
	defer wal.Close()

	b, err := ioutil.ReadFile(file + "-shm")
	if err != nil {
		t.Fatal(err)
	}

	s := &shm{mm: b}
	if have, want := s.Valid(wal.header), true; have != want {
		t.Fatalf("have %t, want %t", have, want)
	}

	if have, want := s.MxFrame(), uint32(8); have != want {
		t.Fatalf("have %v, want %v", have, want)
	}

	// check page->frameID mapping
	pm := map[int]uint32{
		0: 0, // not there
		1: 3, 2: 4, 3: 5, 4: 6, 5: 7, 6: 8,
		7: 0, 42: 0, // not there
	}
	for page, want := range pm {
		frame := s.Frame(page, s.MxFrame())
		if frame != want {
			t.Errorf("page %d: have %v, want %v", page, frame, want)
		}
	}
}
