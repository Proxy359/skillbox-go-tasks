package main

import "fmt"

func main() {
	fmt.Println("Времена года.")
	fmt.Println("Введите месяц:")
	var answer string
	fmt.Scan(&answer)

	switch answer {
	case "декабрь", "январь", "февраль":
		fmt.Println("зима")
	case "март", "апрель", "май":
		fmt.Println("весна")
	case "июнь", "июль", "август":
		fmt.Println("лето")
	case "сентябрь", "октябрь", "ноябрь":
		fmt.Println("осень")
	}
}
