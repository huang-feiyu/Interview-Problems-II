package day26

import (
	"math"
	"strings"
)

func strToInt(s string) int {
	s = strings.TrimSpace(s)
	res := 0
	sign := 1
	length := len(s)
	bndry := math.MaxInt32 / 10
	p := 0
	if length == 0 {
		return 0
	}
	if s[p] == '-' {
		sign = -1
		p++
	} else if s[p] == '+' {
		p++
	}

	for j := p; j < length; j++ {
		if s[j] < '0' || s[j] > '9' {
			break
		}
		if res > bndry || res == bndry && s[j] > '7' {
			if sign == -1 {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
		res = res*10 + int(s[j]-'0')
	}
	return sign * res
}
