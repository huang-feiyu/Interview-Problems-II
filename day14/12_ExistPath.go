package day14

var direcs = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

func exist(board [][]byte, word string) bool {
	// for each char in board => backtracking
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if backtracking(i, j, 0, word, board) {
				return true
			}
		}
	}
	return false
}

func backtracking(x, y, p int, word string, board [][]byte) bool {
	if !(x >= 0 && x < len(board) && y >= 0 && y < len(board[0])) || board[x][y] != word[p] {
		return false
	}

	// !important: avoid out of boundary => len(word)-1
	if p == len(word)-1 {
		return true
	}

	tmp := board[x][y]
	board[x][y] = 0 // mark as visited
	for _, direction := range direcs {
		nx := x + direction[0]
		ny := y + direction[1]
		if backtracking(nx, ny, p+1, word, board) {
			return true
		}
	}
	board[x][y] = tmp

	return false
}
