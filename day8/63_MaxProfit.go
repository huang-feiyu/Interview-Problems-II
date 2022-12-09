package day8

import "math"

func maxProfit(prices []int) int {
	cost := math.MaxInt
	profit := 0
	for _, price := range prices {
		cost = min(price, cost)
		profit = max(profit, price-cost)
	}
	return profit
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
