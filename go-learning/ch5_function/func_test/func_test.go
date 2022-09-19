package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	// 初始化随机数的资源库, 如果不执行这行, 不管运行多少次都返回同样的值
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10), rand.Intn(20)
}

func timeSpend(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second + 2)
	return op
}

func TestFn(t *testing.T) {
	_, b := returnMultiValues()
	t.Log(b, rand.Intn(20))

	tsSF := timeSpend(slowFun)
	t.Log(tsSF(10))
}

func Sum(ops ...int) int {
	ret := 0
	for _, opt := range ops {
		ret += opt
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3))
	t.Log(Sum(1, 2, 3, 4))
}

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("Clear resources")
	}()

	t.Log("Started")
	panic("Fatal error") // defer仍会执行
	fmt.Println("End")
}
