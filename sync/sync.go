package main

import "sync"

type Counter struct {
	mu    sync.Mutex //mutual exclusion lock
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock() // gouroutine wywolujaca funckje Inc() otrzyma lock na Counter jesli jest pierwsza.
	// Wszystkie pozostale beda musialy poczekac az zostanie odblokowany
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
