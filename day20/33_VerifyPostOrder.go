package day20

func verifyPostorder(postorder []int) bool {
	var verify func(int, int) bool
	verify = func(i, j int) bool {
		if i >= j {
			return true
		}
		p := i
		for postorder[p] < postorder[j] {
			p++
		}
		m := p
		for postorder[p] > postorder[j] {
			p++
		}
		return p == j && verify(i, m-1) && verify(m, j-1)
	}
	return verify(0, len(postorder)-1)
}
