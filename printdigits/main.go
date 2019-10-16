package main

import "github.com/01-edu/z01"

func main() {
	var fuck rune = 48
	for fuck <= 57 {

		z01.PrintRune(fuck)
		fuck++
	}
	z01.PrintRune(10)
}
