package main

import "fmt"

func main() {

	fmt.Print("Эта программа позволяет выяснить, есть ли среди 3 введенных", "\n", "чисел положительное", "\n", "\n",
		"Введите первое число:", "\n")
	var a int
	fmt.Scan(&a)

	fmt.Print("Введите второе число:", "\n")
	var b int
	fmt.Scan(&b)

	fmt.Print("Введите третье число:", "\n", "\n")
	var c int
	fmt.Scan(&c)

	var naturalNumber bool = false

	if a >= 0 || b >= 0 || c >= 0 {
		naturalNumber = true
	}

	if naturalNumber {
		fmt.Print("Среди указанных чисел есть положительные.")
	} else {
		fmt.Print("Среди указанных чисел нет положительных.")
	}
}
