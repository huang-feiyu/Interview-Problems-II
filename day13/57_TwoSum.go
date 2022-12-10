package day13

func twoSum(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	for left < right {
		s := nums[left] + nums[right]
		if s < target {
			left++
		} else if s > target {
			right--
		} else {
			return []int{nums[left], nums[right]}
		}
	}
	return nil
}
