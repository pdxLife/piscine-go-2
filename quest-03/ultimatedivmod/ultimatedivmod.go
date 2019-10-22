package piscine

func UltimateDivMod(a *int, b *int) {
	x := *a / *b
	z := *a % *b
	*a = x
	*b = z

}
