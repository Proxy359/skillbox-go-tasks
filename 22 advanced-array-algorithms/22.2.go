package main

import (
	"fmt"
)

func main() {
	fmt.Println("Эта прога показывает индекс элемента в массиве.")
	fmt.Println("Введите 12 натуральных значений массива.", "\n", "\n")

	var mass [12]int
	for i := 0; i < len(mass); i++ {
		mass[i] = 0
		fmt.Scan(&mass[i])
	}

	fmt.Println("Теперь введите число, индекс которого хотите узнать.")
	num := 0
	fmt.Scan(&num)

	min := 0
	max := len(mass) - 1
	mid := (max + min) / 2

	for i := 0; i < len(mass); i++ {
		if mid == mass[0] {
			fmt.Println(mid - 1)
			break
		} else if (num == mass[mid] && num == mass[mid-1]) || num < mass[mid] {
			mid = mid - 1
		} else if num == mass[mid] {
			fmt.Println(mid)
			break
		} else {
			mid = mid + 1
		}
	}
}
