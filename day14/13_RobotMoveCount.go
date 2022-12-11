package day14

var directions = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
var visited [][]bool

func movingCount(m int, n int, k int) int {
	visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	return backtrack(0, 0, m, n, k)
}

func backtrack(x, y, m, n, k int) int {
	if x < 0 || y < 0 || x >= m || y >= n || visited[x][y] || sumNum(x)+sumNum(y) > k {
		return 0
	}
	visited[x][y] = true
	sum := 1
	for _, direction := range directions {
		nx := x + direction[0]
		ny := y + direction[1]
		sum += backtrack(nx, ny, m, n, k)
	}
	return sum
}

func sumNum(a int) int {
	sum := 0
	for a/10 > 0 {
		sum += a % 10
		a /= 10
	}
	sum += a % 10
	return sum
}
