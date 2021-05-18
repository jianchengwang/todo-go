package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check1(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	filePath1 := "D:\\tmp\\test.txt"
	filePath2 := "D:\\tmp\\test2.txt"

	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile(filePath1, d1, 0644)
	check1(err)

	f, err := os.Create(filePath2)
	check1(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check1(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check1(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check1(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}