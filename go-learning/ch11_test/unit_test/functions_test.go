package unit_test

import (
	"fmt"
	"testing"
)

func square(op int) int {
	return op * op
}

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 10}
	for i, v := range inputs {
		ret := square(v)
		if ret != expected[i] {
			t.Errorf("input is %d, the expected is %d, the actual is %d", v, expected[i], ret)
		}
	}
}

func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	t.Error("Error")
	fmt.Println("End")
}

func TestFatalInCode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("Fatal")
	fmt.Println("End")
}
