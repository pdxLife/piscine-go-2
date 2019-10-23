package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	r := 0
	var neue int
	arr := []rune{}

	if n == 0 {
		arr = append(arr, rune(n))
	}

	for n != 0 {
		neue = n % 10
		n = n / 10
		arr = append(arr, rune(neue))
	}
	for range arr {
		r++
	}
	for i1 := 0; i1 < r-1; i1++ {
		for i2 := i1; i2 < r; i2++ {
			if arr[i1] > arr[i2] {
				arr[i1], arr[i2] = arr[i2], arr[i1]
			}
		}
	}
	for l := 0; l < r; l++ {
		z01.PrintRune(arr[l] + 48)
	}
}
