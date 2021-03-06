package LinkedList

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

