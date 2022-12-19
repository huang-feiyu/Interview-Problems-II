package day29

func dicesProbability(n int) []float64 {
	dp := make([][]float64, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]float64, 6*n+1)
	}

	for j := 1; j <= 6; j++ {
		dp[1][j] = 1.0 / 6
	}
	for i := 2; i <= n; i++ {
		for j := i; j <= 6*n; j++ {
			for k := 1; k <= 6; k++ {
				if j <= k {
					continue
				}
				dp[i][j] += dp[i-1][j-k] * (1.0 / 6)
			}
		}
	}

	return dp[n][n:]
}
