package day5

// Binary Search version
func findNumberIn2DArrayB(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	// find all rows might contain target => leftmost
	l, r := 0, len(matrix)
	for l < r {
		mid := (l + r) / 2
		if matrix[mid][0] < target+1 {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if l == 0 || matrix[l-1][0] > target {
		return false
	}
	top := l // => end
	// find all rows might contain target => rightmost
	l, r = 0, top
	for l < r {
		mid := (l + r) / 2
		if matrix[mid][len(matrix[mid])-1] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	down := l // start
	// => [down, top)

	// find the row in the array
	for i := down; i < top; i++ {
		l, r = 0, len(matrix[i])
		for l < r {
			mid := (l + r) / 2
			if matrix[i][mid] < target {
				l = mid + 1
			} else if matrix[i][mid] > target {
				r = mid
			} else {
				return true
			}
		}
	}

	return false
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	i, j := len(matrix)-1, 0
	for i >= 0 && j < len(matrix[0]) {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}
	return false
}
