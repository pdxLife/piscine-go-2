package piscine

func IsAlpha(str string) bool {
	a := []rune(str)
	for i := range a {
		if a[i] < '0' || (a[i] > '9' && a[i] < 'A') || (a[i] > 'Z' && a[i] < 'a') || a[i] > 'z' {
			return false
		}

	}
	return true
}
