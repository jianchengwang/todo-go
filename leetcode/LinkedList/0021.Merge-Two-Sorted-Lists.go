package main

import "fmt"

//type ListNode struct {
//     Val int
//     Next *ListNode
//}

func buildNodeList1(values ...int) *ListNode {
	head := &ListNode{Val: 0}
	current := head
	for _,value := range values {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}
	return head.Next
}

func printNode1(node *ListNode) {
	if node != nil {
		fmt.Println(node.Val)
		for node.Next!=nil {
			node = node.Next
			fmt.Println(node.Val)
		}
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := &ListNode{Val: 0}
	current := head
	for l1 != nil || l2 != nil {
		if l1 == nil {
			current.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		} else if l2 == nil {
			current.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			if l1.Val < l2.Val {
				current.Next = &ListNode{Val: l1.Val}
				l1 = l1.Next
			} else {
				current.Next = &ListNode{Val: l2.Val}
				l2 = l2.Next
			}
		}
		current = current.Next

	}
	return head.Next
}

func main()  {
	l1 := buildNodeList1(1, 1, 5)
	l2 := buildNodeList1(2, 4, 6)
	fmt.Println("l1:")
	printNode1(l1)
	fmt.Println("l2:")
	printNode1(l2)
	l3 := mergeTwoLists(l1, l2)
	fmt.Println("l3:")
	printNode1(l3)
}


