package sqlittle

func Fuzz(data []byte) int {
	if err := fuzz(data); err != nil {
		return 0
	}
	return 1
}

func fuzz(data []byte) error {
	p := bytePager(data)
	db, err := newDatabase(&p)
	if err != nil {
		return err
	}
	tables, err := db.Tables()
	if err != nil {
		return err
	}
	for _, t := range tables {
		table, err := db.Table(t)
		if err != nil {
			return err
		}
		if err := table.Scan(
			func(rowid int64, rec Record) bool {
				return false
			},
		); err != nil {
			return err
		}
	}
	return nil
}

type bytePager []byte

func (b *bytePager) header() ([headerSize]byte, error) {
	if len(*b) < headerSize {
		return [headerSize]byte{}, ErrCorrupted
	}
	var bb [headerSize]byte
	copy(bb[:], (*b)[:headerSize])
	return bb, nil
}

func (b *bytePager) page(n int, pagesize int) ([]byte, error) {
	x := pagesize * (n - 1)
	y := x + pagesize
	if x < 0 || y > len(*b) {
		return nil, ErrCorrupted
	}
	return (*b)[x:y], nil
}

func (b *bytePager) RLock() error   { return nil }
func (b *bytePager) RUnlock() error { return nil }
func (b *bytePager) Close() error   { return nil }
