package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(str string) {
	arrayStr := []rune(str)

	for i := 0; i < len(arrayStr); i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	length := 0
	for char := range os.Args[1:] {
		length = char + 1
	}

	if isEven(length) == true {
		printStr("I have an odd number of arguments")
	} else {
		printStr("I have an even number of arguments")
	}
}
