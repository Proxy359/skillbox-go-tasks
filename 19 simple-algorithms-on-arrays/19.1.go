package main

import (
	"fmt"
)

func first(a [4]int, b [5]int) [9]int {
	var c [9]int
	var x, y int = 0, 0
	for i := 0; i < 9; i++ {
		if i <= 3 {
			c[i] = a[x]
			x += 1
		} else {
			c[i] = b[y]
			y += 1
		}
	}
	return c
}

func main() {
	fmt.Println("Эта прога скрепляет 2 массива")
	a := [4]int{1, 2, 3, 4}
	b := [5]int{5, 6, 7, 8, 9}
	c := first(a, b)
	fmt.Println(c)
}
