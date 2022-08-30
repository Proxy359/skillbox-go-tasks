package main

import "fmt"

func main() {

	fmt.Print("Эта программа позволяет выяснить, набрали ли вы достаточно баллов по ЕГЭ,", "\n", "чтобы поступить в ВУЗ.", "\n", "\n",
		"Введите результат первого экзамена:", "\n")
	var resultFirst int
	fmt.Scan(&resultFirst)

	var resultSecond int
	fmt.Print("Введите результат второго экзамена:", "\n")
	fmt.Scan(&resultSecond)

	var resultThird int
	fmt.Print("Введите результат третьего экзамена:", "\n")
	fmt.Scan(&resultThird)

	totalPoints := resultFirst + resultSecond + resultThird
	lowerBarrier := 275

	if totalPoints >= lowerBarrier {
		fmt.Print("\n", "Поздравляю, вы поступили")
	} else {
		fmt.Print("\n", "Вы не поступили")
	}
}
