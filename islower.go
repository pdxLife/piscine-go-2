package piscine

func IsLower(str string) bool {
	a := []rune(str)
	for i := range a {
		if a[i] < 'a' || a[i] > 'z' {
			return false
		}

	}
	return true
}
