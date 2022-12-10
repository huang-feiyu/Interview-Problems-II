package day13

func exchange(nums []int) []int {
	res := make([]int, len(nums))
	left := 0              // [0, left) => odd
	right := len(nums) - 1 // (right, n) => even

	for _, num := range nums {
		if num%2 == 1 {
			res[left] = num
			left++
		} else {
			res[right] = num
			right--
		}
	}
	return res
}
