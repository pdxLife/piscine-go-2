package piscine

func Swap(a *int, b *int) {
	tempa := *a
	tempb := *b
	*a = tempb
	*b = tempa
}
