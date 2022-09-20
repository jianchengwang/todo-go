package reentrant_mutex

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// 得到id字符串
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}


// RecursiveMutexGrId 包装一个Mutex，实现可重入
type RecursiveMutexGrId struct {
	sync.Mutex
	owner int64 // 当前持有锁的goroutine id
	recursion int32 // 当前这个goroutine 重入的次数
}


