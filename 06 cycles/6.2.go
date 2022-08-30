package main

import "fmt"

func main() {
	fmt.Println("Эта прога позволит путем сомнительного перебора вычестить сумму двух натуральных цисел")

	fmt.Println("Введите первое число")
	var userNum1 int
	fmt.Scan(&userNum1)

	fmt.Println("Введите второе число")
	var userNum2 int
	fmt.Scan(&userNum2)

	sum := userNum1 + userNum2

	for x := userNum1; x <= sum; {
		x = x + 1
		fmt.Println(x)
	}
}
