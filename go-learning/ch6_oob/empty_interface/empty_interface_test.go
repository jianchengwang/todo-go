package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer:", i)
		return
	}
	if i, ok := p.(string); ok {
		fmt.Println("String", i)
		return
	}
	fmt.Println("UnKnow Type")
}

func DoSomethingSwitch(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("UnKnow Type")
	}
}

func TestEmpty(t *testing.T) {
	DoSomething(10)
	DoSomethingSwitch("10")
}
