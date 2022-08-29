package main

import "fmt"

func main() {

	fmt.Print("Эта программа позволяет выяснить, есть ли среди 3 введенных", "\n", "чисел повторяющиеся", "\n", "\n",
		"Введите первое число:", "\n")
	var a int
	fmt.Scan(&a)

	fmt.Print("Введите второе число:", "\n")
	var b int
	fmt.Scan(&b)

	fmt.Print("Введите третье число:", "\n")
	var c int
	fmt.Scan(&c)

	var doubleNumber bool = false

	if a == b || b == c || c == a {
		doubleNumber = true
	}

	if doubleNumber {
		fmt.Print("\n", "Среди указанных чисел есть повторяющиеся.")
	} else {
		fmt.Print("\n", "Среди указанных чисел нет повторяющихся.")
	}
}
