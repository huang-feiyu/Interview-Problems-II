package day18

import "math"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBalanced(root.Left) {
		return false
	}
	if !isBalanced(root.Right) {
		return false
	}
	lDepth := getDepth(root.Left)
	rDepth := getDepth(root.Right)
	return math.Abs(float64(lDepth-rDepth)) <= 1
}

func getDepth(curr *TreeNode) int {
	if curr == nil {
		return 0
	}
	return max(getDepth(curr.Left), getDepth(curr.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
