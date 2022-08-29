package main

import (
	"fmt"
)

func first(a int) string {
	if a%2 == 0 {
		return "true"
	} else {
		return "false"
	}
}

func main() {
	fmt.Println("Эта программа запрашивает у вас число")
	fmt.Println("И выводит 'true', если оно четное")
	fmt.Println("И 'false', если оно нечетное", "\n")

	fmt.Println("Введите ваше натуральное число")
	num := 0
	fmt.Scan(&num)
	fmt.Println(first(num))
}
