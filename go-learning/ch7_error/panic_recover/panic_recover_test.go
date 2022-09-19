package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		fmt.Println("Clear Resources")
		if err := recover(); err != nil {
			fmt.Println("Recover from", err)
		}
	}()

	fmt.Println("Start")
	panic(errors.New("Something Wrong"))
	//os.Exit(-1)
	fmt.Println("End")
}
