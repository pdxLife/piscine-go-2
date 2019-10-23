package piscine

func NRune(s string, n int) rune {
	size := 0
	for range s {
		size++
	}
	if n > size || n < 1 {
		return '\x00'
	}
	return []rune(s)[n-1]

}
