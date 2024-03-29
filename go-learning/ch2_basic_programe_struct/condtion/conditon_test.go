package condtion_test

import (
	"errors"
	"testing"
)

func TestIfMultiSec(t *testing.T) {
	if v, err := sumFun(); err != nil {
		t.Log("err")
	} else {
		t.Log("v==", v)
	}
}

func sumFun() (int, error) {
	return 1, errors.New("1")
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unKnow")
		}
	}
}
