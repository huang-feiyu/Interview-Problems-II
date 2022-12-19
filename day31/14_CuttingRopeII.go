package day31

func cuttingRope(n int) int {
	if n < 3 {
		return 1
	}
	if n == 3 {
		return 2
	}
	res := 1
	for n > 4 {
		res *= 3
		res %= 1000000007
		n -= 3
	}
	return (res * n) % 1000000007
}
