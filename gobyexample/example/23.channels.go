package main

import (
	"fmt"
	"time"
)

func main()  {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	time.Sleep(time.Second)

	msg := <-messages
	fmt.Println(msg)

}
