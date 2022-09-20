## go_once

Once 可以用来执行且仅仅执行一次动作，常常用于单例对象的延迟初始化场景

### 单例对象初始化，非延迟

定义package级别的变量

```go

package abc

import time

var startTime = time.Now()
```

或者在init函数执行

```go

package abc

var startTime time.Time

func init() {
  startTime = time.Now()
}

```

又或者在main函数开始执行的时候，执行要给初始化函数

```go

package abc

var startTime time.Tim

func initApp() {
    startTime = time.Now()
}
func main() {
  initApp()
}
```

### 延迟加锁

```go

package main

import (
	"net"
	"sync"
	"time"
)

// 使用互斥锁保证线程(goroutine)安全
var connMu sync.Mutex
var conn net.Conn

func getConn() net.Conn {
	connMu.Lock()
	defer connMu.Unlock()

	// 返回已创建好的连接
	if conn != nil {
		return conn
	}

	// 创建连接
	conn, _ = net.DialTimeout("tcp", "baidu.com:80", 10*time.Second)
	return conn
}

// 使用连接
func main() {
	conn := getConn()
	if conn == nil {
		panic("conn is nil")
	}
}
```

### Once使用场景

Once 的使用场景sync.Once 只暴露了一个方法 Do，你可以多次调用 Do 方法，但是只有第一次调用 Do 方法时 f 参数才会执行，这里的 f 是一个无参数无返回值的函数。

Once 常常用来初始化单例资源，或者并发访问只需初始化一次的共享资源，或者在测试的时候初始化一次测试资源
