package day20

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n == -1 {
		return 1 / x
	}
	res := myPow(x, n/2)
	res *= res
	res *= myPow(x, n%2)
	return res
}
