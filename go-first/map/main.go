package main

import "fmt"

func main() {
	//s1 := make([]int, 1)
	//s2 := make([]int, 2)
	//f1 := func() {}
	//f2 := func() {}
	//m1 := make(map[int]string)
	//m2 := make(map[int]string)
	//println(s1 == s2) // 错误：invalid operation: s1 == s2 (slice can only be compared to nil)
	//println(f1 == f2) // 错误：invalid operation: f1 == f2 (func can only be compared to nil)
	//println(m1 == m2) // 错误：invalid operation: m1 == m2 (map can only be compared to nil)

	//var m map[string]int // 一个map[string]int类型的变量
	//m["key"] = 1         // 发生运行时异常：panic: assignment to entry in nil map

	//m := map[string]int{}
	//m["key"] = 1
	//fmt.Println(m)

	//m1 := map[int][]string{
	//	1: []string{"val1_1", "val1_2"},
	//	3: {"val3_1", "val3_2", "val3_3"},
	//	7: {"val7_1"},
	//}
	//fmt.Println(m1)
	//
	//type Position struct {
	//	x float64
	//	y float64
	//}

	//m2 := map[Position]string{
	//	Position{29.935523, 52.568915}:  "school",
	//	Position{25.352594, 113.304361}: "shopping-mall",
	//	Position{73.224455, 111.804306}: "hospital",
	//}
	//m2 := map[Position]string{
	//	{29.935523, 52.568915}:  "school",
	//	{25.352594, 113.304361}: "shopping-mall",
	//	{73.224455, 111.804306}: "hospital",
	//}
	//fmt.Println(m2)

	//m1 := make(map[int]string)    // 未指定初始容量
	//m2 := make(map[int]string, 8) // 指定初始容量为8

	//m := map[string]int{
	//	"key1": 1,
	//	"key2": 2,
	//}
	//
	//fmt.Println(len(m)) // 2
	//m["key3"] = 3
	//fmt.Println(len(m)) // 3

	//m := make(map[string]int)
	//_, ok := m["key1"]
	//if !ok {
	//	// "key1"不在map中
	//	fmt.Println("key1不在map中")
	//} else {
	//	// "key1"在map中，v将被赋予"key1"键对应的value
	//	fmt.Println("key1在map中，v将被赋予key1键对应的value")
	//}

	//m := map[string]int{
	//	"key1": 1,
	//	"key2": 2,
	//}
	//
	//fmt.Println(m)    // map[key1:1 key2:2]
	//delete(m, "key2") // 删除"key2"
	//fmt.Println(m)    // map[key1:1
	m := map[int]int{1: 11, 2: 12, 3: 13}
	fmt.Printf("{ ")
	for k, v := range m {
		fmt.Printf("[%d, %d] ", k, v)
	}
	fmt.Printf("}\n")
	fmt.Println(m)\
}
