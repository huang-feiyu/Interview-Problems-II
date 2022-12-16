package day22

func majorityElement(nums []int) int {
	length := len(nums)
	num2freq := make(map[int]int)
	for _, num := range nums {
		num2freq[num]++
		if freq, ok := num2freq[num]; ok && freq > length/2 {
			return num
		}
	}
	return 0
}
