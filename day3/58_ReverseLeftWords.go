package day3

func reverseLeftWords(s string, n int) string {
	var res []byte
	for i := range s {
		res = append(res, s[(i+n)%len(s)])
	}
	return string(res)
}
