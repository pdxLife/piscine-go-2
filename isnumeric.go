package piscine

func IsNumeric(str string) bool {
	a := []rune(str)
	for i := range a {
		if a[i] > '9' || a[i] < '0' {
			return false
		}

	}
	return true
}
