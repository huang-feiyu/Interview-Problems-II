package day12

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeA, nodeB := headA, headB

	for nodeA != nodeB {
		if nodeA == nil {
			nodeA = headB
		} else {
			nodeA = nodeA.Next
		}
		if nodeB == nil {
			nodeB = headA
		} else {
			nodeB = nodeB.Next
		}
	}
	return nodeA
}
