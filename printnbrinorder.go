package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	m := [10]int{}
	if n == 0 {
		m[n]++
	}
	for {
		if n == 0 {
			break
		}
		m[n%10]++
		n /= 10
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < m[i]; j++ {
			z01.PrintRune(rune(48 + i))
		}
	}
}
