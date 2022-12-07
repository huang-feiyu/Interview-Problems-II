package day2

// Normal version
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	curr := head
	next := head.Next
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// Recursive version
func reverseListR(head *ListNode) *ListNode {
	var prev *ListNode
	return reverse(prev, head)
}

func reverse(prev, head *ListNode) *ListNode {
	// encounter the end of source list
	if head == nil {
		return prev
	}
	newHead := reverse(head, head.Next)
	head.Next = prev
	return newHead
}
