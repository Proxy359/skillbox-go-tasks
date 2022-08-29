package main

import "fmt"

func main() {
	fmt.Println("Эта программа выводит ёлочку на экран")

	fmt.Println("Укажите высоту ёлочки")
	var height int
	fmt.Scan(&height)

	length := height*2 - 1
	low := height
	high := height

	a := " "
	b := "*"
	aa := ""

	for y := 1; y <= height; y++ {
		for x := 1; x <= length; x++ {

			if x < low || x > high {
				aa = aa + a
			} else {
				aa = aa + b
			}
		}
		fmt.Println(aa)
		low = low - 1
		high = high + 1
		aa = ""
	}
}
