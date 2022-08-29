package main

import (
	"fmt"
)

func main() {
	fmt.Println("Эта программа позволяет определить, какие из введенных 10 натуральных чисел четные, а какие нечетные", "\n")
	var num [10]int
	x := 0
	y := 0

	for i := 0; i < len(num); i++ {
		fmt.Print("Введите ", i+1, "-е число: ")
		fmt.Scan(&num[i])
	}

	for i, _ := range num {
		if num[i]%2 == 0 {
			x += 1
		} else {
			y += 1
		}
	}
	fmt.Println("\n", "Кол-во четных чисел:", x)
	fmt.Println("Кол-во нечетных чисел:", y)
}
