package rwmutex_test

import (
	"sync"
	"testing"
	"time"
)

type Counter struct {
	mu    sync.RWMutex
	count uint64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Count() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func TestRWMutex(t *testing.T) {
	var counter Counter
	for i := 0; i < 10; i++ {
		go func() {
			for {
				t.Log(counter.Count()) // 计数器读操作
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for {
		counter.Incr() // 计数器写操作
		time.Sleep(time.Second)
	}
}
