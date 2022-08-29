package main

import (
	"fmt"
)

func first(mass [10]int) ([5]int, [5]int) {
	var a [5]int
	var b [5]int
	x := 0
	y := 0
	for i := 0; i < len(mass); i++ {
		if mass[i]%2 == 0 {
			a[x] = mass[i]
			x += 1
		} else {
			b[y] = mass[i]
			y += 1
		}
	}
	return a, b
}

func main() {
	mass := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(first(mass))
}
