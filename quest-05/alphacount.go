package piscine

func AlphaCount(str string) int {
	i := 0
	for _, l := range str {
		if l >= 'a' && l <= 'z' || l >= 'A' && l <= 'Z' {
			i++
		}
	}
	return i
}
