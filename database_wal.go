package sqlittle

// OpenWal opens a .sqlite file in wal format.
// Use database.Close() when done.
func OpenWal(f string) (*Database, error) {
	l, err := newFilePager(f)
	if err != nil {
		return nil, err
	}
	return newWalDB(l, f)
}

func newWalDB(l pager, f string) (*Database, error) {
	w, err := openWalPager(l, f)
	if err != nil {
		return nil, err
	}

	d := &Database{
		// dirty: true, // TODO
		l: w,
	}
	return d, nil // d.resolveDirty()
}
