package main

import (
	"fmt"
)

func main() {
	fmt.Println("Этот код позволит поменять значения двух переменных.")
	fmt.Println()

	fmt.Println("Введите желаемое значение а")
	var a int
	fmt.Scan(&a)

	fmt.Println("Введите желаемое значение b")
	var b int
	fmt.Scan(&b)
	fmt.Println("Получается, что было:", "\n", "a=", a, "\n", "b=", b)
	fmt.Println()

	a = a + b
	b = a - b
	a = a - b
	fmt.Println("Но произошла магия и теперь:", "\n", "a=", a, "\n", "b=", b)
}
