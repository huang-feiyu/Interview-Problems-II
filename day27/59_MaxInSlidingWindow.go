package day27

func maxSlidingWindow(nums []int, k int) []int {
	if k == 0 {
		return nil
	}
	var res []int //
	var queue []int

	pushQ := func(x int) {
		for len(queue) != 0 && x > queue[len(queue)-1] {
			// if x > back, pop it
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, x)
	}

	popQ := func(x int) {
		// pop if x is the head
		if len(queue) != 0 && x == queue[0] {
			queue = queue[1:]
		}
	}

	for i := 0; i < k; i++ {
		pushQ(nums[i])
	}
	res = append(res, queue[0])

	for i := k; i < len(nums); i++ {
		popQ(nums[i-k])
		pushQ(nums[i])
		res = append(res, queue[0])
	}
	return res
}
