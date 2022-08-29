package main

import "fmt"

func main() {
	fmt.Println("Эта прога позволит вывести все числа от 0 до вашего числа")
	fmt.Println("Введите ваше число")
	var userNum int
	fmt.Scan(&userNum)

	for startNum := 0; startNum < userNum; {
		startNum = startNum + 1
		fmt.Println(startNum)
	}
}
