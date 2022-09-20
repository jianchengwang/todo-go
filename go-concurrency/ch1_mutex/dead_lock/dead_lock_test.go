package dead_lock

import (
	"sync"
	"testing"
	"time"
)

func TestDeadLock(t *testing.T) {
	// 派出所证明
	var psCertificate sync.Mutex
	// 物业证明
	var propertyCertificate sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2)

	// 派出所处理goroutine
	go func() {
		defer wg.Done() // 派出所处理完成

		psCertificate.Lock()
		defer psCertificate.Unlock()

		// 检查材料
		time.Sleep(time.Second * 5)
		// 请求物业证明
		propertyCertificate.Lock()
		propertyCertificate.Unlock()
	}()

	// 物业处理goroutine
	go func() {
		defer wg.Done() // 物业处理完成
		propertyCertificate.Lock()
		defer propertyCertificate.Unlock()

		// 检查材料
		time.Sleep(time.Second * 5)
		// 请求派出所证明
		psCertificate.Lock()
		psCertificate.Unlock()
	}()

	wg.Wait()
	t.Log("done")
}
