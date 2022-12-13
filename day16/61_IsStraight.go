package day16

func isStraight(nums []int) bool {
	max := 0
	min := 14
	numSet := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if ok := numSet[nums[i]]; ok {
			return false
		}
		numSet[nums[i]] = true
		max = maxNum(max, nums[i])
		min = minNum(min, nums[i])
	}
	return max-min < 5
}

func maxNum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minNum(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
