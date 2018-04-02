// table/index page cache

package sqlittle

import (
	"sync"
)

type tableCache struct {
	limit int
	elem  map[int]tableBtree
	mu    sync.RWMutex
}

func newTableCache(limit int) *tableCache {
	return &tableCache{
		limit: limit,
		elem:  make(map[int]tableBtree, limit),
	}
}

func (t *tableCache) get(p int) tableBtree {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.elem[p]
}

func (t *tableCache) set(p int, tab tableBtree) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if len(t.elem) >= t.limit {
		// cache full? Simply drop the whole thing.
		t.elem = make(map[int]tableBtree, t.limit)
	}
	t.elem[p] = tab
}
