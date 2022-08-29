package main

import "fmt"

func main() {
	a := "*"
	b := " "
	aa := ""
	fmt.Println("Шахматная доска")
	fmt.Println("Введите высоту")
	var height int
	fmt.Scan(&height)

	fmt.Println("Введите ширину")
	var width int
	fmt.Scan(&width)

	for x := 1; x <= height; x++ {
		for i := 1; i <= width; i++ {

			if x%2 == 1 {
				if i%2 == 1 {
					aa = aa + a
				} else {
					aa = aa + b
				}
			} else {
				if i%2 == 1 {
					aa = aa + b
				} else {
					aa = aa + a
				}
			}
		}
		fmt.Println(aa)
		aa = ""
	}
}
