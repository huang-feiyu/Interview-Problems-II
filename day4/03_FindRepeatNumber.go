package day4

func findRepeatNumber(nums []int) int {
	numset := make(map[int]bool) // check if encounter before
	for _, num := range nums {
		if ok := numset[num]; ok {
			return num
		}
		numset[num] = true
	}
	return -1
}
