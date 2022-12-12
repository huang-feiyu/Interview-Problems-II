package day15

var cnt int // reversed count

func kthLargest(root *TreeNode, k int) int {
	cnt = k
	var result int
	countSearch(root, &result)
	return result
}

func countSearch(curr *TreeNode, res *int) {
	if curr == nil {
		return
	}
	countSearch(curr.Right, res)
	if cnt == 0 {
		return
	}
	cnt--
	if cnt == 0 {
		*res = curr.Val
	} else {
		countSearch(curr.Left, res)
	}
}
