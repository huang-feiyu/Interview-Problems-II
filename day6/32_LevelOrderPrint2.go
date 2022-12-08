package day6

func levelOrder2(root *TreeNode) [][]int {
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) != 0 {
		size := len(queue)
		path := make([]int, 0)
		// pop one level & push next level
		for i := 0; i < size; i++ {
			head := queue[0]
			queue = queue[1:]
			path = append(path, head.Val)
			if head.Left != nil {
				queue = append(queue, head.Left)
			}
			if head.Right != nil {
				queue = append(queue, head.Right)
			}
		}
		res = append(res, path)
	}
	return res
}
