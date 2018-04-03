package sqlittle

type pager interface {
	// read the file header bytes. Page size is unknown yet.
	header() ([headerSize]byte, error)
	// load a page from storage.
	page(n int, pagesize int) ([]byte, error)
	// as it says
	Close() error
	// read lock
	RLock() error
	// unlock read lock
	RUnlock() error
}
