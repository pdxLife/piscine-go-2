package piscine

func ToUpper(s string) string {
	a := []rune(s)

	for i := range a {
		if a[i] >= 65 && a[i] <= 90 {
			a[i] += 32
		}
	}
	s = string(a)
	return s

}
