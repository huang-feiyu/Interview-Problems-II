package day11

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{0, head}
	prev := dummy
	curr := head
	for curr != nil {
		if curr.Val == val {
			next := curr.Next
			prev.Next = next
			curr = next
		} else {
			curr = curr.Next
			prev = prev.Next
		}
	}
	return dummy.Next
}
