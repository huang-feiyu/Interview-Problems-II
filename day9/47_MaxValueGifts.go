package day9

func maxValue(grid [][]int) int {
	dp := make([][]int, len(grid)) // dp[i][j]
	for i := range grid {
		dp[i] = make([]int, len(grid[i]))
	}
	sum := 0
	for j := range grid[0] {
		sum += grid[0][j]
		dp[0][j] = sum
	}
	sum = 0
	for i := range grid {
		sum += grid[i][0]
		dp[i][0] = sum
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = max(dp[i][j-1], dp[i-1][j]) + grid[i][j]
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
