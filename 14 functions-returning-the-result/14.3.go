package main

import (
	"fmt"
)

func first(n int) int {
	n *= 3
	return n
}

func second(n int) int {
	n += 6
	return n
}

func main() {
	fmt.Println("Эта программа преобразует число двумя функциями.")
	var x int
	x = 1
	x = first(x)
	x = second(x)
	fmt.Println(x)
}
