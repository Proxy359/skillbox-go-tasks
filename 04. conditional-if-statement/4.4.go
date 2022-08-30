package main

import "fmt"

func main() {

	fmt.Print("Эта программа позволяет выяснить, сработает ли акция на ваше посещение ресторана.", "\n", "\n",
		"Введите день недели, когда собираетесь посетить ресторан:", "\n")
	var day int
	fmt.Scan(&day)

	var peopleNum int
	fmt.Print("Введите количество гостей:", "\n")
	fmt.Scan(&peopleNum)

	var checkAmount int
	fmt.Print("Введите планируемую сумму чека:", "\n")
	fmt.Scan(&checkAmount)

	totalDiscount := 0

	if day == 1 {
		mondayDiscount := checkAmount / 10
		fmt.Print("\n", "За посещение в понедельник ваша скидка составляет ", mondayDiscount, " рублей.")
		totalDiscount += -mondayDiscount

	} else if day == 5 &&
		checkAmount > 10000 {
		fridayDiscount := checkAmount / 20
		fmt.Print("\n", "За посещение в пятницу и счет более 10к", "\n", "ваша скидка составляет ", fridayDiscount, " рублей")
		totalDiscount += -fridayDiscount
	}

	if peopleNum > 5 {
		addCost := checkAmount / 10
		totalDiscount += addCost
		fmt.Print("\n", "За посещение ресторана компанией из более чем 5 человек", "\n",
			"к сумме заказов добавляется ", addCost, " рублей.")
	}

	fmt.Print("\n", "\n", "Итого ваш счет составит: ", checkAmount+totalDiscount)
}
