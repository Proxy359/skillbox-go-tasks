package main

import "fmt"

func main() {

	fmt.Print("Эта программа позволяет выяснить, какие из 3 % ставок выгоднее.", "\n", "\n",
		"Введите первую % ставку:", "\n")
	var firstInterest int
	fmt.Scan(&firstInterest)

	fmt.Print("Введите вторую % ставку:", "\n")
	var secondInterest int
	fmt.Scan(&secondInterest)

	fmt.Print("Введите третью % ставку:", "\n")
	var thirdInterest int
	fmt.Scan(&thirdInterest)

	if firstInterest == secondInterest &&
		firstInterest == thirdInterest &&
		secondInterest == thirdInterest {
		fmt.Print("\n", "Все вышеуказанные ставки равны")
	} else if firstInterest <= secondInterest &&
		firstInterest <= thirdInterest {
		fmt.Print("\n", "Наиболее выгодные ставки это ", secondInterest, " и ", thirdInterest)
	} else if secondInterest <= firstInterest &&
		secondInterest <= thirdInterest {
		fmt.Print("\n", "Наиболее выгодные ставки это ", firstInterest, " и ", thirdInterest)
	} else {
		fmt.Print("\n", "Наиболее выгодные ставки это ", firstInterest, " и ", secondInterest)
	}
}