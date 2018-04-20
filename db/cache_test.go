package db

import (
	"testing"
)

func TestCache(t *testing.T) {
	c := newBtreeCache(40)
	for i := 0; i < 90; i++ {
		c.set(i, &tableLeaf{})
	}
	if have, want := len(c.elem), 10; have != want {
		t.Errorf("have %d, want %d", have, want)
	}
	c.clear()
	if have, want := len(c.elem), 0; have != want {
		t.Errorf("have %d, want %d", have, want)
	}
}
