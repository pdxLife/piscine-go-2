package piscine

func BasicAtoi(s string) int {
	x := 0
	for range s {
		x++
	}
	z := 1
	y := 0
	for i := x - 1; i >= 0; i-- {
		y += z * int(s[i]-'0')
		z *= 10
	}

	return y
}
