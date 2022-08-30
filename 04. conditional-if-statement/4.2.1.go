package main

import "fmt"

func main() {

	fmt.Print("Эта программа позволяет выяснить, сколько из 3 чисел больше 5,", "\n", "\n",
		"Введите первое число:", "\n")

	var resultFirst int
	fmt.Scan(&resultFirst)

	var resultSecond int
	fmt.Print("Введите второе число:", "\n")
	fmt.Scan(&resultSecond)

	var resultThird int
	fmt.Print("Введите третье число:", "\n")
	fmt.Scan(&resultThird)

	limit := 5
	overLimit := 0

	if resultFirst <= limit &&
		resultSecond <= limit &&
		resultThird <= limit {
		fmt.Print("\n", "Чисел больше 5 не обнаружено:", "\n", "\n")
	} else {
		if resultFirst > limit {
			overLimit += 1
		}

		if resultSecond > limit {
			overLimit += 1
		}

		if resultThird > limit {
			overLimit += 1
		}

		fmt.Print("\n", "Обнаружено чисел больше 5:", "\n", overLimit)
	}
}
