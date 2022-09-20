package main

import (
	"fmt"
	"sync"
)

// go run -race counter_unsafe.go // 运行检查并发问题
// go tool compile -race -S counter_unsafe.go // 编译增加data race检测

func main() {
	var count = 0
	// 使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i:=0; i<10; i++ {
		go func() {
			defer wg.Done()
			for j:=0; j<100000; j++ {
				count++
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
