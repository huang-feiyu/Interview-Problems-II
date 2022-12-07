package day2

type ListNode struct {
	Val  int
	Next *ListNode
}

// Recursive version
func reversePrintR(head *ListNode) []int {
	var res []int
	subPrint(head, &res)
	return res
}

func subPrint(head *ListNode, path *[]int) {
	if head == nil {
		return
	}
	subPrint(head.Next, path) // add child's values first
	*path = append(*path, head.Val)
}

// Iterative version
func reversePrint(head *ListNode) []int {
	var stack []int
	curr := head

	// store all in stack first
	for curr != nil {
		stack = append(stack, curr.Val)
		curr = curr.Next
	}

	// reverse the stack in order =>
	for i := 0; i < len(stack)/2; i++ {
		stack[i], stack[len(stack)-i-1] = stack[len(stack)-i-1], stack[i]
	}
	return stack
}
