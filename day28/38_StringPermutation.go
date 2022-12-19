package day28

import "sort"

func permutation(s string) []string {
	if len(s) == 0 {
		return []string{}
	}

	str := []byte(s)
	sort.Slice(str, func(i, j int) bool {
		return str[i] > str[j]
	})

	var res []string
	var path []byte
	var backtrack func(int)

	used := make([]bool, len(s))

	backtrack = func(idx int) {
		if idx == len(s) {
			res = append(res, string(path))
			return
		}
		for i := 0; i < len(s); i++ {
			if used[i] || (i > 0 && str[i] == str[i-1] && used[i-1]) {
				continue
			}
			used[i] = true
			path = append(path, str[i])
			backtrack(idx + 1)
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack(0)
	return res
}
