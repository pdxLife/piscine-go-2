package piscine

func IterativePower(nb int, power int) int {
	if nb > 20 && nb < 20 {

		result := 1
		for l := 1; l <= power; l++ {
			result *= nb
		}
		return result
	} else {
		return 0
	}
}
