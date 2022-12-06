package day1

import "math"

type MinStack struct {
	cStack []int // container stack
	mStack []int // helper min stack, store elements non-strict desc
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (s *MinStack) Push(x int) {
	min := s.Min()
	s.cStack = append(s.cStack, x)
	if min >= x {
		s.mStack = append(s.mStack, x)
	}
}

func (s *MinStack) Pop() {
	if len(s.mStack) == 0 {
		return
	}
	min := s.Min()
	val := s.Top()
	s.cStack = s.cStack[:len(s.cStack)-1]
	if min == val {
		s.mStack = s.mStack[:len(s.mStack)-1]
	}
}

func (s *MinStack) Top() int {
	if len(s.mStack) == 0 {
		return math.MaxInt
	}
	return s.cStack[len(s.cStack)-1]
}

func (s *MinStack) Min() int {
	if len(s.mStack) > 0 {
		return s.mStack[len(s.mStack)-1] // top, aka min
	}
	return math.MaxInt // represent as invalid
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
