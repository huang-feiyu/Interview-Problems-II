package day10

func lengthOfLongestSubstring(s string) int {
	chars := []byte(s)
	letter2idx := make(map[byte]int) // char and idx in current path

	left := 0
	res := 0
	for i := 0; i < len(chars); i++ {
		c := chars[i]
		// there is the char in previous => move start idx forward
		if idx, ok := letter2idx[c]; ok {
			left = max(left, idx+1)
		}
		letter2idx[c] = i
		res = max(res, i-left+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
