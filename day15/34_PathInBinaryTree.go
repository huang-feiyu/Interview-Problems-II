package day15

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var path []int
var res [][]int

func pathSum(root *TreeNode, target int) [][]int {
	res = [][]int{}
	path = []int{}
	inorderSearch(root, target)
	return res
}

func inorderSearch(curr *TreeNode, target int) {
	if curr == nil {
		return
	}
	target -= curr.Val
	path = append(path, curr.Val)
	if target == 0 && curr.Left == nil && curr.Right == nil {
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		res = append(res, pathCopy)
	}
	inorderSearch(curr.Left, target)
	inorderSearch(curr.Right, target)
	path = path[:len(path)-1]
}
