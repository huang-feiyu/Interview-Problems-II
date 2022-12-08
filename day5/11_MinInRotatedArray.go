package day5

func minArray(numbers []int) int {
	l, r := 0, len(numbers)-1
	// find first number with i makes: nums[i] > nums[i+1]
	for l < r {
		mid := (l + r) / 2
		if numbers[mid] > numbers[r] {
			l = mid + 1
		} else if numbers[mid] < numbers[r] {
			r = mid
		} else {
			r--
		}
	}

	return numbers[l]
}
