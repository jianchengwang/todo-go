package encap

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := &Employee{Name: "Mike", Age: 30}
	e2 := new(Employee) // 返回指针
	e2.Id = "2"
	e2.Name = "Rose"
	e.Age = 12
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("&e is %T", &e)
	t.Logf("e1 is %T", e1)
	t.Logf("e2 is %T", e2)
}

//func (e Employee) String() string {
//	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID;%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
//}
//
//func TestStructOperations(t *testing.T) {
//	e := Employee{"0", "Bob", 20}
//	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
//	t.Log(e.String())
//
//	//e1 := &Employee{"0", "Bob", 20}
//	//t.Log(e1.String())
//}

func (e *Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID;%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())

	//e1 := &Employee{"0", "Bob", 20}
	//fmt.Printf("Address is %x\n", unsafe.Pointer(&e1.Name))
	//t.Log(e1.String())
}
