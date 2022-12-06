package day1

type CQueue struct {
	aStack []int // append stack
	sStack []int // store stack
}

func Constructor9() CQueue {
	return CQueue{[]int{}, []int{}}
}

func (q *CQueue) AppendTail(value int) {
	q.aStack = append(q.aStack, value) // push stack
}

func (q *CQueue) DeleteHead() int {
	if len(q.sStack) == 0 {
		// no remaining element => push aStack to sStack
		for len(q.aStack) > 0 {
			// pop from append stack
			val := q.aStack[len(q.aStack)-1]
			q.aStack = q.aStack[:len(q.aStack)-1]
			// push to store stack
			q.sStack = append(q.sStack, val)
		}
	}

	if len(q.sStack) != 0 {
		val := q.sStack[len(q.sStack)-1]
		q.sStack = q.sStack[:len(q.sStack)-1]
		return val
	}

	// no element in both stacks => -1
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
