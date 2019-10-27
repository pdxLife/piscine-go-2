package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	name := os.Args
	size := 0
	for range name {
		size++
	}
	for j := size - 1; j > 0; j-- {
		a := []rune(name[j])
		for i := range name[j] {
			z01.PrintRune(a[i])
		}
		z01.PrintRune(10)
	}
}
