package day30

func printNumbers(n int) []int {
	max := maxNumber(n)
	res := make([]int, max+1)
	for i := 1; i <= max; i++ {
		res[i] = i
	}
	return res[1:]
}

func maxNumber(n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res *= 10
	}
	res--
	return res
}
