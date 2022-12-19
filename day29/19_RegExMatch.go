package day29

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	if len(s) == 0 {
		if len(p) > 1 && p[1] == '*' {
			return isMatch("", p[2:])
		} else {
			return false
		}
	}

	if s[0] == p[0] || p[0] == '.' {
		if len(p) > 1 && p[1] == '*' {
			return isMatch(s[1:], p[:]) ||
				isMatch(s[1:], p[2:]) ||
				isMatch(s[:], p[2:])
		} else {
			return isMatch(s[1:], p[1:])
		}
	}

	if len(p) > 1 && p[1] == '*' {
		return isMatch(s, p[2:])
	}

	return false
}
