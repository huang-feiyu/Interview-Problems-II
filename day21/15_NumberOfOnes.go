package day21

func hammingWeight(num uint32) int {
	var res uint32
	for num != 0 {
		res += 1 & num
		num >>= 1
	}
	return int(res)
}
