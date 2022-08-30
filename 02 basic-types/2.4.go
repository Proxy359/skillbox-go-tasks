package main

import "fmt"

func main() {
	totalPrice := 4000000
	numberOfDriveway := 10
	apartmentsInDriveway := 40

	fmt.Println("Расчёт стоимости ремонта для каждой квартиры.")
	fmt.Println()
	fmt.Println("Сумма, указанная в квитанции:", totalPrice, "рублей.")
	fmt.Println("Подъездов в доме:", numberOfDriveway)
	fmt.Println("Квартир в каждом подъезде:", apartmentsInDriveway)
	fmt.Println()
	fmt.Println("-----Считаем-----")
	fmt.Println("Каждая квартира должна платить по", totalPrice/(numberOfDriveway*apartmentsInDriveway), "рублей.")
}
