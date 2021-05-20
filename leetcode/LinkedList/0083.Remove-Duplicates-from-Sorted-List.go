package LinkedList

func DeleteDuplicates(head *ListNode) *ListNode {
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