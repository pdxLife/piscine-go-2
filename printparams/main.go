package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	name := os.Args
	for j := range name {
		if j > 0 {
			a := []rune(name[j]) // создаешь слайс из рун с элементами j которые сейчас находятся в цикле 
			for i := range name[j] { //цикл создается, 
				z01.PrintRune(a[i])
			}
			z01.PrintRune(10)
		}
	}
}
