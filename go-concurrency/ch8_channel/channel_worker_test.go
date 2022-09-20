package ch8_channel

import (
	"fmt"
	"testing"
	"time"
)

type Token struct {
}

func newWorker(id int, ch chan Token, nextCh chan Token) {
	for {
		token := <-ch       // 取得令牌
		fmt.Println(id + 1) // id从1开始
		time.Sleep(time.Second)
		nextCh <- token
	}
}

func TestChannelWorker(t *testing.T) {
	chArr := [4]chan struct{}{
		make(chan struct{}),
		make(chan struct{}),
		make(chan struct{}),
		make(chan struct{}),
	}

	for i := 0; i < 4; i++ {
		go func(i int) {
			for {
				<-chArr[i%4]
				fmt.Printf("i am %d\n", i)

				time.Sleep(1 * time.Second)
				chArr[(i+1)%4] <- struct{}{}
			}
		}(i)
	}

	chArr[0] <- struct{}{}
	select {}
}

func TestChannelWorkerToken(t *testing.T) {
	chs := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}
	// 创建4个worker
	for i := 0; i < 4; i++ {
		go newWorker(i, chs[i], chs[(i+1)%4])
	}

	// 首先帮令牌交给第一个worker
	chs[0] <- struct{}{}

	select {}
}
