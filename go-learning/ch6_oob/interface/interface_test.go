package interface_test

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "go programmer hello world"
}

type RustProgrammer struct {
}

func (r *RustProgrammer) WriteHelloWorld() string {
	return "rust programmer hello world"
}

func helloWorld(p Programmer) {
	fmt.Println(p.WriteHelloWorld())
}

func TestClient(t *testing.T) {
	var pg Programmer
	pg = new(GoProgrammer)
	t.Log(pg.WriteHelloWorld())
	helloWorld(pg)

	pg = &RustProgrammer{}
	t.Log(pg.WriteHelloWorld())
	helloWorld(pg)
}
