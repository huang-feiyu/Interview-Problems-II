package day29

func nthUglyNumber(n int) int {
	dp := make([]int, n) // dp[i]: the i+1_th ugly number
	// 2, 3, 5
	a, b, c := 0, 0, 0

	dp[0] = 1
	for i := 1; i < n; i++ {
		n1, n2, n3 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = min(n1, min(n2, n3))
		if dp[i] == n1 {
			a++
		}
		if dp[i] == n2 {
			b++
		}
		if dp[i] == n3 {
			c++
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
