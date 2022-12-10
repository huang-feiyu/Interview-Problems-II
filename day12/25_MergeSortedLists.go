package day12

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	left := l1
	right := l2
	curr := dummy

	for left != nil && right != nil {
		if left.Val > right.Val {
			curr.Next = right
			right = right.Next
		} else {
			curr.Next = left
			left = left.Next
		}
		curr = curr.Next
	}

	if left != nil {
		curr.Next = left
	}
	if right != nil {
		curr.Next = right
	}

	return dummy.Next
}
