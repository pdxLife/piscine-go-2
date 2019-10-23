package piscine

func LastRune(s string) rune {
	size := 0
	for range s {
		size++
	}
	return []rune(s)[size-1]

}
