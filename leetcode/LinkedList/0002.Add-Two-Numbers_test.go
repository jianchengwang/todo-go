package LinkedList

import (
	"fmt"
	"testing"
)

/**
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/
func TestAddTwoNumbers(t *testing.T) {
	l1 := BuildNodeList(2, 4, 3)
	l2 := BuildNodeList(5, 6, 4)
	fmt.Println("l1:")
	PrintNode(l1)
	fmt.Println("l2:")
	PrintNode(l2)
	l3 := AddTwoNumbers(l1, l2)
	fmt.Println("l3:")
	PrintNode(l3)

	l1 = BuildNodeList(9, 9, 9, 9, 9)
	l2 = BuildNodeList(1)
	fmt.Println("l1:")
	PrintNode(l1)
	fmt.Println("l2:")
	PrintNode(l2)
	l3 = AddTwoNumbers(l1, l2)
	fmt.Println("l3:")
	PrintNode(l3)
}
