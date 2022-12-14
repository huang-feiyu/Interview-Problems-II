package day19

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	pathToP := make([]*TreeNode, 0)
	pathToQ := make([]*TreeNode, 0)
	if !findNode(root, p, &pathToP) {
		panic("Not exist p")
	}
	if !findNode(root, q, &pathToQ) {
		panic("Not exist q")
	}
	res := root
	for len(pathToP) > 0 && len(pathToQ) > 0 && pathToP[0] == pathToQ[0] {
		res = pathToP[0]
		pathToP = pathToP[1:]
		pathToQ = pathToQ[1:]
	}
	return res
}

func findNode(root, node *TreeNode, queue *[]*TreeNode) bool {
	if root == nil {
		return false
	}
	*queue = append(*queue, root)
	if root == node {
		return true
	}
	if findNode(root.Left, node, queue) {
		return true
	}
	if findNode(root.Right, node, queue) {
		return true
	}
	*queue = (*queue)[:len(*queue)-1]
	return false
}
