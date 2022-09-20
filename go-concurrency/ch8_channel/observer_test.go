package ch8_channel

import (
	"fmt"
	"testing"
)

func TestObserver(t *testing.T) {
	cha1 := make(chan struct{}, 1)

	over := make(chan struct{}, 1)
	listens := []chan struct{}{make(chan struct{}, 1), make(chan struct{}, 1), make(chan struct{}, 1)}
	fanOut(cha1, listens, true)

	cha1 <- struct{}{}
	cha1 <- struct{}{}
	cha1 <- struct{}{}

	close(cha1)
	go func() {
		for i := 0; i < len(listens); i++ {
			item := i
			go func(item int) {
				for {
					_, ok := <-listens[item]
					if !ok {
						over <- struct{}{}
						close(over)
						return
					}
					fmt.Printf("第%d监听者收到信息\n", item)
				}

			}(item)
		}
	}()

	<-over
}

func fanOut1(ch1 <-chan struct{}, listens []chan struct{}, isSync bool) {
	go func() {
		defer func() {
			for i := 0; i < len(listens); i++ {
				close(listens[i])
			}
		}()

		for v := range ch1 {
			v := v
			for i := 0; i < len(listens); i++ {
				item := i
				if isSync {
					listens[item] <- v
				} else {
					go func() {
						listens[item] <- v
					}()
				}
			}
		}

	}()
}

func asStream(done <-chan struct{}, values ...interface{}) <-chan interface{} {
	s := make(chan interface{}) //创建一个unbuffered的channel
	go func() { // 启动一个goroutine，往s中塞数据
		defer close(s)             // 退出时关闭chan
		for _, v := range values { // 遍历数组
			select {
			case <-done:
				return
			case s <- v: // 将数组元素塞入到chan中
			}
		}
	}()
	return s
}

func takeN(done <-chan struct{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{}) // 创建输出流
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ { // 只读取前num个元素
			select {
			case <-done:
				return
			case takeStream <- <-valueStream: //从输入流中读取元素
			}
		}
	}()
	return takeStream
}

func mapChan(in <-chan interface{}, fn func(interface{}) interface{}) <-chan interface{} {
	out := make(chan interface{}) //创建一个输出chan
	if in == nil {                // 异常检查
		close(out)
		return out
	}

	go func() { // 启动一个goroutine,实现map的主要逻辑
		defer close(out)
		for v := range in { // 从输入chan读取数据，执行业务操作，也就是map操作
			out <- fn(v)
		}
	}()

	return out
}

func reduce(in <-chan interface{}, fn func(r, v interface{}) interface{}) interface{} {
	if in == nil { // 异常检查
		return nil
	}

	out := <-in         // 先读取第一个元素
	for v := range in { // 实现reduce的主要逻辑
		out = fn(out, v)
	}

	return out
}
