package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func dataConsumer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			ret, ok := <-ch
			fmt.Println(ret)
			if ok {
				fmt.Println("ok")
			} else {
				fmt.Println("close")
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int, 10)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataConsumer(ch, &wg)
	wg.Add(1)
	dataConsumer(ch, &wg)
	wg.Wait()
	fmt.Println("Done")
}
