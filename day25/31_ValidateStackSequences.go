package day25

func validateStackSequences(pushed []int, popped []int) bool {
	n := len(pushed)
	p, q := 0, 0 // push, pop pointers

	var stack []int

	for p = 0; p < n; p++ {
		stack = append(stack, pushed[p])
		for len(stack) != 0 && stack[len(stack)-1] == popped[q] {
			stack = stack[:len(stack)-1]
			q++
		}
	}
	return len(stack) == 0
}
