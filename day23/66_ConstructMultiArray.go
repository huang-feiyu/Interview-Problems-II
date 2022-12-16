package day22

func constructArr(a []int) []int {
	n := len(a)
	if n == 0 {
		return nil
	}
	dp1 := make([]int, n) // left product
	dp2 := make([]int, n) // right product

	dp1[0] = 1
	dp2[n-1] = 1
	for i := 1; i < n; i++ {
		dp1[i] = dp1[i-1] * a[i-1]
	}
	for j := n - 2; j >= 0; j-- {
		dp2[j] = dp2[j+1] * a[j+1]
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = dp1[i] * dp2[i]
	}
	return res
}
