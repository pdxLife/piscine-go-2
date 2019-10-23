package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	size := 0
	for range toFind {
		size++
	}
	b := -1
	count := 0
	for i := range s {
		if count == size {
			break
		} else if s[i] == toFind[count] {
			if count == 0 {
				b = i
			}
			count++
		} else {
			count = 0
			b = -1
		}
	}
	if count != size {
		return -1
	}
	if b == -1 {
		return -1
	}
	return b

}
