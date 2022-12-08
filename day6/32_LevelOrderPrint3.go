package day6

func levelOrder(root *TreeNode) [][]int {
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
		if len(res)%2 == 1 {
			// reverse
			for i := 0; i < len(path)/2; i++ {
				path[i], path[len(path)-i-1] = path[len(path)-i-1], path[i]
			}
		}
		res = append(res, path)
	}
	return res
}
