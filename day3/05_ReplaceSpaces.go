package day3

func replaceSpace(s string) string {
	str := []byte{'%', '2', '0'}
	chars := []byte(s)
	var res []byte

	for i := range s {
		if chars[i] == ' ' {
			res = append(res, str...)
		} else {
			res = append(res, chars[i])
		}
	}
	return string(res)
}
