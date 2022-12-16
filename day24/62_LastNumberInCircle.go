package day24

func lastRemaining(n int, m int) int {
	res := 0
	for i := 2; i <= n; i++ {
		res = (m + res) % i
	}
	return res
}
