package day22

func singleNumber(nums []int) int {
	numSet := make(map[int]int) // 3 states: 1, 2, 3 => frequency
	resSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num]++
		if numSet[num] == 1 {
			resSet[num] = true
		} else {
			delete(resSet, num)
		}
	}
	if len(resSet) != 1 {
		panic("Input wrong")
	}
	for k := range resSet {
		return k
	}
	return 0
}
