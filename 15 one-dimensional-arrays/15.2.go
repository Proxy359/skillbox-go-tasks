package main

import (
	"fmt"
)

func first(a [10]int) (x [10]int) {
	for i := 0; i < 10; i++ {
		x[i] = a[9-i]
	}
	return x
}

func main() {
	fmt.Println("Эта прога меняет значения массива местами")
	var a [10]int
	for i := 0; i < 10; i++ {
		fmt.Print("Введите ", i+1, "-е число: ")
		fmt.Scan(&a[i])
	}
	fmt.Println(a)
	fmt.Println(first(a))
}
