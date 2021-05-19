package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func buildNodeList2(values ...int) *ListNode {
	head := &ListNode{Val: 0}
	current := head
	for _,value := range values {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}
	return head.Next
}

func printNode2(node *ListNode) {
	if node != nil {
		fmt.Println(node.Val)
		for node.Next!=nil {
			node = node.Next
			fmt.Println(node.Val)
		}
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	preValue := current.Val
	for current.Next != nil {
		if current.Next.Val == preValue {
			current.Next = current.Next.Next
		} else {
			preValue = current.Val
			current = current.Next
		}
	}
	return head
}

func main()  {
	l1 := buildNodeList2(1, 1, 5, 8, 7, 7, 9)
	fmt.Println("l1:")
	printNode2(l1)
	l2 := deleteDuplicates(l1)
	fmt.Println("l2:")
	printNode2(l2)

}