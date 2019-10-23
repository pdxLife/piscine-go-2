package piscine

func ToUpper(s string) string {
	a := []rune(s)

	for i := range a {
		if a[i] >= 97 && a[i] <= 122 {
			a[i] -= 32
		}
	}
	s = string(a)
	return s

}
