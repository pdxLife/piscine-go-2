package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	name := os.Args
	for j := range name {
		if j > 0 {
			a := []rune(name[j])
			for i := range name[j] {
				z01.PrintRune(a[i])
			}
			z01.PrintRune(10)
		}
	}
}
