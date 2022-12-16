package day24

func findContinuousSequence(target int) [][]int {
	i, j := 1, 2
	res := make([][]int, 0)
	sum := 3

	for i < j {
		if sum == target {
			var tmp []int
			for k := i; k <= j; k++ {
				tmp = append(tmp, k)
			}
			res = append(res, tmp)
		}
		if sum >= target {
			sum -= i
			i++
		} else {
			j++
			sum += j
		}
	}
	return res
}
