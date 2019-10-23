package piscine

func Compare(a, b string) int {
	var x int
	if a > b {
		x = 1
	}

	if a == b {
		x = 0
	}

	if a < b {
		x = -1
	}
	return x
}
