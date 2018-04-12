package sqlittle

import (
	"errors"
	"fmt"
)

type formatWal struct {
	l          pager
	wal        *wal
	shm        *shm
	btreeCache *btreeCache // table and index page cache
	mxFrame    uint32
}

// Open a wal formatted DB. `file` is the name of the sqlite file.
func newFormatWal(l pager, file string) (*formatWal, error) {
	// TODO: try an EXCLUSIVE lock and rebuild shm file

	// WAL database readers keep an rlock while it's open
	fmt.Printf("go rlock\n")
	if err := l.RLock(); err != nil {
		return nil, err
	}
	fmt.Printf("done rlock\n")

	wal, err := openWal(file + "-wal")
	if err != nil {
		return nil, err // FIXME
	}

	shm, err := openShm(file + "-shm")
	if err != nil {
		return nil, err // FIXME
	}

	return &formatWal{
		l:          l,
		wal:        wal,
		shm:        shm,
		btreeCache: newBtreeCache(CachePages),
	}, nil
}

// Page returns a tableBtree or indexBtree
func (fw *formatWal) Page(page int) (interface{}, error) {
	if fw.mxFrame == 0 {
		return nil, errors.New("need a lock")
	}
	if p := fw.btreeCache.get(page); p != nil {
		return p, nil
	}
	fmt.Printf("load page %d\n", page)

	buf, err := fw.page(page)
	if err != nil {
		return nil, err
	}
	p, err := newBtree(buf, page == 1)
	if err == nil && fw.btreeCache != nil {
		fw.btreeCache.set(page, p)
	}
	fmt.Printf("return page %d\n", page)
	return p, err
}

func (p *formatWal) Close() error {
	p.wal.Close()
	p.shm.Close()
	return p.l.RUnlock()
}

func (p *formatWal) page(n int) ([]byte, error) {
	fmt.Printf("load page %d in %d\n", n, p.mxFrame)
	if frame := p.shm.Frame(n, p.mxFrame); frame > 0 {
		return p.wal.frame(int(frame))
	}
	return p.l.page(n)
}

func (p *formatWal) RLock() error {
	if err := p.wal.Remap(); err != nil {
		return err
	}
	if err := p.shm.Remap(); err != nil {
		return err
	}

	// TODO: actual lock

	p.mxFrame = p.shm.MxFrame()

	// TODO: this is not always needed
	p.btreeCache.clear()

	return nil
}

func (p *formatWal) RUnlock() error {
	return nil
}
