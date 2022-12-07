package day4

func search(nums []int, target int) int {
	// search the most left
	left := searchLeft(nums, target)
	right := searchLeft(nums, target+1)
	if left == len(nums) || nums[left] != target {
		return 0
	}
	return right - left
}

// find first i makes nums[i] >= target
func searchLeft(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
