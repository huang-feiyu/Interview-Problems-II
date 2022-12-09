package day10

import (
	"strconv"
	"strings"
)

func translateNum(num int) int {
	numStr := strconv.Itoa(num)
	n := len(numStr)
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1) // dp[i] => number of strings translated from numStr[0, i]
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		subStr := numStr[i-2 : i]
		if strings.Compare(subStr, "10") >= 0 && strings.Compare(subStr, "25") <= 0 {
			dp[i] = dp[i-2] + dp[i-1]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n]
}
