package ch7_atomic

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type Config struct {
	NodeName string
	Addr     string
	Count    int32
}

func LoadNewConfig() Config {
	return Config{
		NodeName: "北京",
		Addr:     "10.77.95.27",
		Count:    rand.Int31(),
	}
}

func TestAtomicValue(t *testing.T) {
	var config atomic.Value
	config.Store(LoadNewConfig())
	var cond = sync.NewCond(&sync.Mutex{})

	// 设置新的config
	go func() {
		for {
			time.Sleep(time.Duration(5+rand.Int63n(5)) * time.Second)
			config.Store(LoadNewConfig())
			cond.Broadcast()
		}
	}()

	go func() {
		for {
			cond.L.Lock()
			cond.Wait()                 // 等待变更信息
			c := config.Load().(Config) // 读取新的配置
			t.Logf("new config:%+v\n", c)
			cond.L.Unlock()
		}
	}()

	select {}
}
