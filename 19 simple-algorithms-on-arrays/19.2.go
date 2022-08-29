package main

import (
	"fmt"
)

func main() {
	mass := [6]int{1, 2, 5, 7, -1, 0}
	BubbleSort(mass)
}

func BubbleSort(mass [6]int) {
	for y := 0; y < 6; y++ {
		for i := 0; i < 5; i++ {
			if mass[i] > mass[i+1] {
				x := mass[i]
				mass[i] = mass[i+1]
				mass[i+1] = x
			}
		}
	}
	fmt.Print(mass)
}
