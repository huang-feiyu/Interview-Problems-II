package day25

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	l, r := 0, n-1 // column
	t, b := 0, m-1 // row

	res := make([]int, m*n)
	x := 0 // counter for result

	for true {
		// left to right
		for i := l; i <= r; i++ {
			res[x] = matrix[t][i]
			x++
		}
		t++
		if t > b {
			break
		}

		// top to bottom
		for i := t; i <= b; i++ {
			res[x] = matrix[i][r]
			x++
		}
		r--
		if l > r {
			break
		}

		// right to left
		for i := r; i >= l; i-- {
			res[x] = matrix[b][i]
			x++
		}
		b--
		if t > b {
			break
		}

		// bottom to up
		for i := b; i >= t; i-- {
			res[x] = matrix[i][l]
			x++
		}
		l++
		if l > r {
			break
		}
	}
	return res
}
