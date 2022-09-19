package util_all_done

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(i int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The result is from %d", i)
}

func AllResponse() string {
	numberOfRunner := 10
	ch := make(chan string, 10)
	for i := 0; i < numberOfRunner; i++ {
		go func(i int) {
			ch <- runTask(i)
		}(i)
	}
	finalRet := ""
	for j := 0; j < numberOfRunner; j++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}

func TestAllResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}
