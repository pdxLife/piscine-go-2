package piscine

func IsUpper(str string) bool {
	a := []rune(str)
	for i := range a {
		if a[i] < 'A' || a[i] > 'Z' {
			return false
		}

	}
	return true
}
