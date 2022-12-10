package day13

import "strings"

func reverseWords(s string) string {
	str := strings.TrimSpace(s)
	var res []byte

	// for right to left
	right := len(str) - 1
	left := right
	for left >= 0 {
		for left >= 0 && str[left] != ' ' {
			left-- // find first space
		}
		res = append(append(res, str[left+1:right+1]...), ' ')
		for left >= 0 && str[left] == ' ' {
			left-- // skip all spaces between
		}
		right = left
	}
	return strings.TrimSpace(string(res))
}
