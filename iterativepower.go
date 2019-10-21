package piscine

func IterativePower(nb int, power int) int {
	if power > 0 {
		result := 1
		for l := 1; l <= power; l++ {
			result = result * nb
		}
		return result
	} else {
		return 0
	}
}
