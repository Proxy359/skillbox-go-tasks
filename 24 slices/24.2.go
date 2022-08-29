package main

import (
	"fmt"
)

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	f := func(a [10]int) [10]int {
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(a)-1; j++ {
				if a[j] < a[j+1] {
					a[j], a[j+1] = a[j+1], a[j]
				}
			}
		}
		return a
	}
	fmt.Println(f(a))
}
