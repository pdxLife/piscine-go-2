package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	for _, a := range str {
		z01.PrintRune(a)
	}
}
