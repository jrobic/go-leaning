package sync_mutex

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}
