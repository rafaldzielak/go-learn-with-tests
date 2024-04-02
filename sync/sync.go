package sync

import "sync"

type Counter struct {
	value int
	// A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
	mu sync.Mutex
}

func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func NewCounter() *Counter {
	return &Counter{}
}
