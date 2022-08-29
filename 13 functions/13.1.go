package main

import (
	"fmt"
)

func second(a, b int) {
	y := 0
	if a == b && a%2 == 0 {
		fmt.Println(a)

	} else if a < b {
		for x := a; x <= b; x++ {
			if x%2 == 0 {
				y += x
			}
		}
		fmt.Println(y)

	} else {
		for x := b; x <= a; x++ {
			if x%2 == 0 {
				y += x
			}
		}
		fmt.Println(y)
	}
}

func main() {
	fmt.Println("Привет, эта программа выводит сумму четных числел в указанном тобой диапазоне")
	var firstNum, secondNum int
	fmt.Println("Введи первое натуральное число")
	fmt.Scan(&firstNum)
	fmt.Println("Введи второе натуральное число")
	fmt.Scan(&secondNum)
	second(firstNum, secondNum)
}
