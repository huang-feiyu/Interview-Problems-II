package day4

func missingNumber(nums []int) int {
	// find the first number makes: i < nums[i] => i is what we want
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) / 2
		if nums[mid] <= mid {
			// no less than in fact
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
