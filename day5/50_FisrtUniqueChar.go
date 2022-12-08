package day5

func firstUniqChar(s string) byte {
	charSet := make(map[byte]bool)
	for _, c := range []byte(s) {
		_, ok := charSet[c]
		// for all more than once chars => false
		charSet[c] = !ok
	}
	for _, c := range []byte(s) {
		if charSet[c] {
			return c
		}
	}
	return ' '
}
