package piscine

func TrimAtoi(s string) int {
	r := 0
	m := false
	for _, c := range s {
		if c == '-' && r == 0 {
			m = true
		}
		if c >= '0' && c <= '9' {
			r = r*10 + int(c-48)
		}
	}
	if m {
		return -r
	} else {
		return r
	}
}
