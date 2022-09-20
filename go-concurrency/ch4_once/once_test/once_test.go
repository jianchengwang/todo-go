package once_test

import (
	"net"
	"sync"
	"testing"
	"time"
)

// 使用互斥锁保证goroutine安全
var connMu sync.Mutex
var conn net.Conn

func getMutexConn() net.Conn {
	connMu.Lock()
	defer connMu.Unlock()

	// 返回已经创建好的连接
	if conn != nil {
		return conn
	}

	// 创建连接
	conn, _ = net.DialTimeout("tcp", "baidu.com:80", 10*time.Second)
	return conn
}

func TestGetMutexConn(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			t.Log(getMutexConn())
		}()
	}
	wg.Wait()
}

func TestOnce(t *testing.T) {
	var once sync.Once

	// 第一个初始化函数
	f1 := func() {
		t.Log("in f1")
	}
	once.Do(f1) // 打印f1

	// 第二个函数
	var addr = "baidu.com"
	once.Do(func() {
		t.Log(addr)
	}) // 无输出
}
