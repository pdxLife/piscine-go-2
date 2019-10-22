package piscine

func BasicAtoi(s string) int {
	aS := []rune(s)
	n := len(s)
	a := 0
	for i := 0; i < n; i++ {
		if aS[i] < '0' || aS[i] > '5' {
			return a
		} else {
			a *= 10
			a += int(aS[i]) - '0'
		}
	}
	return a
}
