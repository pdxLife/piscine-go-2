package piscine

func IsPrintable(str string) bool {
	a := []rune(str)
	for i := range a {
		if a[i] < 31 {
			return false
		}

	}
	return true
}
