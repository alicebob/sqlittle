// pager loads the main sqlite database file

package sqlittle

type pager interface {
	// load a page from storage. Starts at 1. Shouldn't cache.
	page(n int) ([]byte, error)
	// as it says
	Close() error
	// read lock
	RLock() error
	// unlock read lock
	RUnlock() error
	// true if there is any 'RESERVED' lock on this file
	CheckReservedLock() (bool, error)
}
