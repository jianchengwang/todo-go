package recursive_mutex

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// TokenRecursiveMutex 包装一个Mutex，实现可重入
type TokenRecursiveMutex struct {
	sync.Mutex
	token     int64 // 当前持有锁的token
	recursion int32 // 当前这个goroutine 重入的次数
}

func (m *TokenRecursiveMutex) Lock(token int64) {
	// 如果当前持有锁的token就是调用的token，说明是重入
	if atomic.LoadInt64(&m.token) == token {
		m.recursion++
		return
	}
	m.Mutex.Lock() // 传入的token不一致，说明不是递归调用
	// 抢到锁之后记录这个token
	atomic.StoreInt64(&m.token, token)
	m.recursion = 1
}

func (m *TokenRecursiveMutex) UnLock(token int64) {
	if atomic.LoadInt64(&m.token) != token { // 释放其他token持有的锁，panic
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.token, token))
	}
	m.recursion--         // 当前持有这个锁的token释放
	if m.recursion != 0 { // 还没有回退到最初的递归调用
		return
	}
	atomic.StoreInt64(&m.token, 0) // 没有递归调用，释放锁
	m.Mutex.Unlock()
}
