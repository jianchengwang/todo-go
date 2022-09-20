package main

import (
	"fmt"
	"sync"
)

func main()  {
	ch := make(chan struct {})
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				ch <- struct{}{}
			}
		}()
	}

	go func() {
		wg.Wait() // 等待所以goroutine完成
 		close(ch) // 关闭ch通道
	}()

	count := 0
	for range ch { // 如果ch通道读取完，ch是关闭状态，则for循环结束
		count++
	}
	fmt.Println(count)
}
