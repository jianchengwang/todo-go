package ch8_channel

import (
	"fmt"
	"testing"
	"time"
)

type Mutex struct {
	ch chan struct{}
}

// 初始化
func NewMutex() *Mutex {
	mu := &Mutex{ch: make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}

// 请求锁
func (m *Mutex) Lock() {
	<-m.ch
}

// 释放锁
func (m *Mutex) UnLock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}

// 尝试获取锁
func (m *Mutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}
	return false
}

// 加入一个超时的设置
func (m *Mutex) LockTimeOut(timeOut time.Duration) bool {
	timer := time.NewTicker(timeOut)
	select {
	case <-m.ch:
		timer.Stop()
		return true
	case <-timer.C:
	}
	return false
}

func (m *Mutex) IsLocked() bool {
	return len(m.ch) == 0
}

func TestChannelLock(t *testing.T) {
	m := NewMutex()
	ok := m.TryLock()
	fmt.Printf("locked v %v\n", ok)
	ok = m.TryLock()
	fmt.Printf("locked %v\n", ok)
}
