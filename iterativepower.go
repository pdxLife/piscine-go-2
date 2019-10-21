package piscine

func IterativePower(nb int, power int) int {
	if nb > 20 || nb < -20 || power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	} else {
		result := 1
		for l := 1; l <= power; l++ {
			result = result * nb
		}
		return result
	}
}
