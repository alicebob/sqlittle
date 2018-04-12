package sqlittle

func Fuzz(data []byte) int {
	if err := fuzz(data); err != nil {
		return 0
	}
	return 1
}

func fuzz(data []byte) error {
	p, err := newBytePager(data)
	if err != nil {
		return err
	}
	db, err := newJournalDB(p, "")
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

		if _, err := table.Rowid(42); err != nil {
			return err
		}

		if _, err := table.Def(); err != nil {
			return err
		}
	}

	indexes, err := db.Indexes()
	if err != nil {
		return err
	}
	for _, in := range indexes {
		index, err := db.Index(in)
		if err != nil {
			return err
		}

		if err := index.Scan(
			func(rowid int64, rec Record) bool {
				return false
			},
		); err != nil {
			return err
		}

		if err := index.ScanMin(
			Record{"q"},
			func(rowid int64, rec Record) bool {
				return false
			},
		); err != nil {
			return err
		}

		if _, err := index.Def(); err != nil {
			return err
		}
	}
	return nil
}

type bytePager struct {
	b      []byte
	header header
}

func newBytePager(b []byte) (*bytePager, error) {
	buf := [headerSize]byte{}
	copy(buf[:], b)
	h, err := parseHeader(buf)
	return &bytePager{
		b:      b,
		header: h,
	}, err
}

func (b *bytePager) page(n int) ([]byte, error) {
	x := b.header.PageSize * (n - 1)
	y := x + b.header.PageSize
	if x < 0 || y > len(b.b) {
		return nil, ErrCorrupted
	}
	return (b.b)[x:y], nil
}

func (b *bytePager) RLock() error                     { return nil }
func (b *bytePager) RUnlock() error                   { return nil }
func (b *bytePager) CheckReservedLock() (bool, error) { return false, nil }
func (b *bytePager) Close() error                     { return nil }
