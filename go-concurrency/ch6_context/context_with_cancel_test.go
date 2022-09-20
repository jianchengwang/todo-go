package ch6_context

import (
	"context"
	"testing"
	"time"
)

func TestContextWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer func() {
			t.Log("goroutine exit")
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}
