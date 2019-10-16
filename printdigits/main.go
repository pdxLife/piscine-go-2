package main

import "github.com/01-edu/z01"

func main() {
	var aRune rune = 48
	for aRune <= 57 {

		z01.PrintRune(aRune)
		aRune++
	}
	z01.PrintRune(10)
}
