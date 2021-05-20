package LinkedList

import (
	"fmt"
	"leetcode/structures"
)

type ListNode = structures.ListNode
func BuildNodeList(values ...int) *ListNode {
	head := &ListNode{Val: 0}
	current := head
	for _,value := range values {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}
	return head.Next
}

func PrintNode(node *ListNode) {
	if node != nil {
		fmt.Println(node.Val)
		for node.Next!=nil {
			node = node.Next
			fmt.Println(node.Val)
		}
	}
}
