package main

import (
	"fmt"
)

func first(a [10]int) [10]int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}

func main() {
	a := [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(first(a))
}
