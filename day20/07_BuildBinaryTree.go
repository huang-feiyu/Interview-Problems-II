package day20

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) {
		panic("length of preorder MUST equal to inorder")
	}
	in2idx := make(map[int]int)
	for i, num := range inorder {
		in2idx[num] = i
	}
	var build func(int, int, int) *TreeNode
	build = func(root, left, right int) *TreeNode {
		if left > right {
			return nil
		}
		node := &TreeNode{preorder[root], nil, nil}
		idx := in2idx[preorder[root]]
		node.Left = build(root+1, left, idx-1)
		node.Right = build(((idx-1)-left+1)+root+1, idx+1, right)
		return node
	}
	return build(0, 0, len(preorder)-1)
}
