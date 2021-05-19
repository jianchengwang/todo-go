package main

import (
	"fmt"
)

//type ListNode struct {
//	Val int
//	Next *ListNode
//}

/**
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		carry = (n1 + n2 + carry) / 10
		current = current.Next
	}
	return head.Next
}

func buildNodeList(values ...int) *ListNode {
	head := &ListNode{Val: 0}
	current := head
	for _,value := range values {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}
	return head.Next
}

func printNode(node *ListNode) {
	if node != nil {
		fmt.Println(node.Val)
		for node.Next!=nil {
			node = node.Next
			fmt.Println(node.Val)
		}
	}
}

func main() {
	l1 := buildNodeList(2, 4, 3)
	l2 := buildNodeList(5, 6, 4)
	fmt.Println("l1:")
	printNode(l1)
	fmt.Println("l2:")
	printNode(l2)
	l3 := addTwoNumbers(l1, l2)
	fmt.Println("l3:")
	printNode(l3)

	l1 = buildNodeList(9, 9, 9, 9, 9)
	l2 = buildNodeList(1)
	fmt.Println("l1:")
	printNode(l1)
	fmt.Println("l2:")
	printNode(l2)
	l3 = addTwoNumbers(l1, l2)
	fmt.Println("l3:")
	printNode(l3)
}
