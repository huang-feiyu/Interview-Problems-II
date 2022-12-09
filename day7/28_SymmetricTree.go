package day7

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isEqual(root.Left, root.Right)
}

func isEqual(left, right *TreeNode) bool {
	// both nil
	if left == nil && right == nil {
		return true
	}
	// only one nil
	if left == nil || right == nil {
		return false
	}

	// in-order
	return left.Val == right.Val && isEqual(left.Left, right.Right) && isEqual(left.Right, right.Left)
}
