package main

import (
	"fmt"
)

const (
	a = 1
	b = 2
	c = 3
)

func first(n int) int {
	y := n + a + c
	return y
}

func second(n int) int {
	y := n + a + b
	return y
}

func third(n int) int {
	y := n + b + c
	return y
}

func main() {
	var n int
	n = 1
	fmt.Println(first(second(third(n))))
}
