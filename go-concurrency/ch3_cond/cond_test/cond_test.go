package cond_test

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestCond(t *testing.T) {
	c := sync.NewCond(&sync.Mutex{})
	var ready int

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Second * time.Duration(rand.Int63n(10)))

			// 加锁更改等待条件
			c.L.Lock()
			ready++
			c.L.Unlock()

			t.Logf("运动员%d已准备就绪\n", i)

			// 广播唤醒所有等待者
			c.Broadcast()
		}(i)
	}

	c.L.Lock()
	for ready != 10 {
		c.Wait()
		t.Log("裁判员被唤醒一次")
	}
	c.L.Unlock()

	// 所有运动员准备就绪
	t.Log("所有运动员准备就绪，比赛开始，3，2，1，。。。。")
}
