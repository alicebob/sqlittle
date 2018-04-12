package sqlittle

type walPager struct {
	l   pager
	wal *wal
	shm *shm
}

// Open a wal formatted DB. `file` is the name of the sqlite file.
func openWalPager(l pager, file string) (*walPager, error) {
	// TODO: try an EXCLUSIVE lock and rebuild shm file

	// WAL database readers keep an rlock while it's open
	if err := l.RLock(); err != nil {
		return nil, err
	}

	wal, err := openWal(file + "-wal")
	if err != nil {
		return nil, err // FIXME
	}

	shm, err := openShm(file + "-shm")
	if err != nil {
		return nil, err // FIXME
	}

	return &walPager{
		l:   l,
		wal: wal,
		shm: shm,
	}, nil
}

func (p *walPager) Close() error {
	p.wal.Close()
	p.shm.Close()
	return p.l.RUnlock()
}

func (p *walPager) page(n int) ([]byte, error) {
	if frame := p.shm.Frame(n, p.shm.MxFrame()); frame > 0 {
		return p.wal.frame(int(frame))
	}
	return p.l.page(n)
}

func (p *walPager) RLock() error {
	if err := p.wal.Remap(); err != nil {
		return err
	}
	if err := p.shm.Remap(); err != nil {
		return err
	}
	// TODO: actual lock
	return nil
}

func (p *walPager) RUnlock() error {
	return nil
}

func (p *walPager) CheckReservedLock() (bool, error) {
	return p.l.CheckReservedLock()
}
