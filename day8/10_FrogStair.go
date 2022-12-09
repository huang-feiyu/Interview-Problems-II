package day8

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}
	mod := 1000000007
	dp := make([]int, n+1) // dp[i] => number of ways to jump to i-th stair
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		//      jump one  jump two
		dp[i] = (dp[i-1] + dp[i-2]) % mod
	}
	return dp[n] % mod
}
