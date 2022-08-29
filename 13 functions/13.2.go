package main

import (
	"fmt"
)

func first(a, b int) {
	a = a * b
	b = a / b
	a = a / b
	fmt.Println(a, b)
}

func main() {
	fmt.Println("Эта программа меняет значения переменных через вызов вспомогат. функции")

	fmt.Println("Введите первое число")
	var x1 int
	fmt.Scan(&x1)

	fmt.Println("Введите второе число")
	var x2 int
	fmt.Scan(&x2)

	fmt.Println("Так у нас было раньше")
	fmt.Println(x1, x2)
	fmt.Println("Так стало после вызова фунцкии")
	first(x1, x2)
}
