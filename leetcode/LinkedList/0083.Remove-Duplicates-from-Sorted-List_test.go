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
func TestDeleteDuplicates(t *testing.T) {
	l1 := BuildNodeList(1, 1, 5, 8, 7, 7, 9)
	fmt.Println("l1:")
	PrintNode(l1)
	l2 := DeleteDuplicates(l1)
	fmt.Println("l2:")
	PrintNode(l2)
}
