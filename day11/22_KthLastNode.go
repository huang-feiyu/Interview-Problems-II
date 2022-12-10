package day11

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow := head
	fast := head // forward k steps

	for i := 0; i < k; i++ {
		if fast == nil {
			// no enough length
			return nil
		}
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
