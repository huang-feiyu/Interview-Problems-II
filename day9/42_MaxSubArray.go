package day9

func maxSubArray(nums []int) int {
	max := nums[0]
	frontSum := 0
	for _, num := range nums {
		if frontSum > 0 {
			frontSum += num
		} else {
			frontSum = num
		}
		max = maxNum(max, frontSum)
	}
	return max
}

func maxNum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
