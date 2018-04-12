// Wal or Journal backend formats

package sqlittle

type format interface {
	// load a page, is expected to use caching
	Page(int) (interface{}, error)
	RLock() error
	RUnlock() error
	Close() error
}

type formatJournal struct {
	journal    string
	l          pager
	header     *header
	btreeCache *btreeCache // table and index page cache
}

func newFJournal(l pager, journal string) (*formatJournal, error) {
	fj := &formatJournal{
		journal:    journal,
		l:          l,
		btreeCache: newBtreeCache(CachePages),
	}
	return fj, fj.checkJournal()
}

// Page returns a tableBtree or indexBtree
func (fj *formatJournal) Page(page int) (interface{}, error) {
	if p := fj.btreeCache.get(page); p != nil {
		return p, nil
	}

	buf, err := fj.l.page(page)
	if err != nil {
		return nil, err
	}
	p, err := newBtree(buf, page == 1)
	if err == nil && fj.btreeCache != nil {
		fj.btreeCache.set(page, p)
	}
	return p, err
}

func (fj *formatJournal) RLock() error {
	if err := fj.checkJournal(); err != nil {
		return err
	}
	if err := fj.checkCache(); err != nil {
		return err
	}
	return fj.l.RLock()
}

func (fj *formatJournal) RUnlock() error {
	return fj.l.RUnlock()
}

func (fj *formatJournal) Close() error {
	// return fj.l.Close()
	return nil
}

func (fj *formatJournal) checkJournal() error {
	if fj.journal != "" {
		hot, err := validJournal(fj.journal)
		if err != nil {
			return err
		}
		if hot {
			// If something is using the transaction the db will have a RESERVED
			// lock.
			locked, err := fj.l.CheckReservedLock()
			if err != nil {
				return err
			}
			if !locked {
				return ErrHotJournal
			}
		}
	}
	return nil
}

func (fj *formatJournal) checkCache() error {
	h, err := getHeader(fj.l)
	if err != nil {
		return err
	}
	if fj.header != nil && fj.header.ChangeCounter != h.ChangeCounter {
		fj.btreeCache.clear()
	}
	fj.header = &h
	return nil
}
