package main

import "fmt"

func main() {

	num1 := 0
	num2 := 0
	num3 := 0

	for {
		num1 = num1 + 1
		num2 = num2 + 1
		num3 = num3 + 1

		if num1 == 10 {
			fmt.Println(num1)
			num1 = 10
		}

		if num2 == 100 {
			fmt.Println(num2)
			num2 = 100
		}

		if num3 == 1000 {
			fmt.Println(num3)
			break
		}
	}
}
