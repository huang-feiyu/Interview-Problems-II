package day24

func cuttingRope(n int) int {
	max3 := func(a, b, c int) int {
		if a >= b && a >= c {
			return a
		}
		if b >= a && b >= c {
			return b
		}
		return c
	}

	dp := make([]int, n+1) // dp[i]: the maximum product of broken i
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 0; j <= i; j++ {
			dp[i] = max3(dp[i], (i-j)*j, dp[i-j]*j)
		}
	}
	return dp[n]
}
