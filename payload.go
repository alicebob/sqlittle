package sqlittle

import (
	"encoding/binary"
)

// Payload represents the payload part of a cell. If overflow is non-zero the
// Payload field will be truncated. Use addOverflow() to get a full payload.
type Payload struct {
	Length   int64
	Payload  []byte
	Overflow int
}

// overflow is stored on different pages. Load whatever is needed to complete
// the payload data.
func addOverflow(db *database, pl Payload) ([]byte, error) {
	to := pl.Payload
	overflow := pl.Overflow
	for {
		if overflow == 0 {
			return to[:pl.Length], nil
		}
		buf, err := db.page(overflow)
		if err != nil {
			return nil, err
		}
		next, buf := int(binary.BigEndian.Uint32(buf[:4])), buf[4:]
		to = append(to, buf...)
		overflow = next
	}
}
