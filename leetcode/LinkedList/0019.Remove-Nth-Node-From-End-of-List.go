package LinkedList

//type ListNode struct {
//	Val int
//	Next *ListNode
//}

/**
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
 */

// 解法一
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	preSlow, slow, fast := dummyHead, head, head
	for fast != nil {
		if n <= 0 {
			preSlow = slow
			slow = slow.Next
		}
		n--
		fast = fast.Next
	}
	preSlow.Next = slow.Next
	return dummyHead.Next
}


