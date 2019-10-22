package piscine

func RecursiveFactorial(nb int, power int) int {
	if nb > 20 || nb < -20 || power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	} else {
		return 1
	}
	return nb * RecursiveFactorial(nb-1)
}
