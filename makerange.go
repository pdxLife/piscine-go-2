package piscine

func MakeRange(min, max int) []int {
	if max-min <= 0 {
		var a []int
		return a
	}
	a := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		a[i] = i + min
	}
	return a
}
