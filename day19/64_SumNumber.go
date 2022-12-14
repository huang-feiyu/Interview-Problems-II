package day19

func sumNums(n int) int {
	res := 0
	var sum func(int) bool
	sum = func(n int) bool {
		res += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return res
}
