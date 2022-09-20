package wrong_usage

import (
	"fmt"
	"sync"
	"testing"
)

// Lock/Unlock 不是成对出现
func TestLockUnLockNotComeInPairs(t *testing.T) {
	var mut sync.Mutex
	defer mut.Unlock()
	t.Log("hello world!")
}

// Copy 已使用的 Mutex，Mutex是一个有状态state的对象，我们期望的都是一个零值的Mutex
// go vet copy.go 检查死锁
type Counter struct {
	sync.Mutex
	Count int
}

func foo(c Counter)  {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}

func TestCopyMutex(t *testing.T) {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo(c) // 复制锁
}

// 重入，mutex是不可重入的锁，
// 因为 Mutex 的实现中没有记录哪个 goroutine 拥有这把锁。
// 理论上，任何 goroutine 都可以随意地 Unlock 这把锁，所以没办法计算重入条件
// 所以不要把锁作为参数传递
func foo1(l sync.Locker)  {
	fmt.Println("in foo")
	l.Lock()
	bar1(l)
	l.Unlock()
}

func bar1(l sync.Locker)  {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}

func TestReentrant (t *testing.T) {
	l := &sync.Mutex{}
	foo1(l)
}

