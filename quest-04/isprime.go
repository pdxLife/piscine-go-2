package piscine

func IsPrime(s int) bool {
	if s < 2 {
		return false
	}
	var a bool = true
	for i := 2; i < s; i++ {
		if s%i == 0 {
			a = false
		}
	}
	return a
}
