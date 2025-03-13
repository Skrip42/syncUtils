package syncutils

import (
	"sync"
)

type Cond struct {
	awaitStack []func()
	mu         *sync.Mutex
}

func NewCond() *Cond {
	return &Cond{
		awaitStack: make([]func(), 0, 8),
		mu:         &sync.Mutex{},
	}
}

func (c *Cond) Wait() <-chan struct{} {
	ch := make(chan struct{})
	c.mu.Lock()
	c.awaitStack = append(
		c.awaitStack,
		func() {
			close(ch)
		},
	)
	c.mu.Unlock()
	return ch
}

func (c *Cond) Signal() {
	c.mu.Lock()
	if len(c.awaitStack) > 0 {
		c.awaitStack[len(c.awaitStack)-1]()
		c.awaitStack = c.awaitStack[:len(c.awaitStack)-1]
	}
	c.mu.Unlock()
}

func (c *Cond) Broadcast() {
	c.mu.Lock()
	for i := range c.awaitStack {
		c.awaitStack[i]()
	}
	c.awaitStack = c.awaitStack[:0]
	c.mu.Unlock()
}
