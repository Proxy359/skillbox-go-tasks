package main

import "fmt"

func main() {

	fmt.Print("Эта программа позволяет выяснить, получится ли оплатить товар", "\n", "тремя монетами без сдачи.", "\n", "\n",
		"Введите стоимость товара:", "\n")
	var prise int
	fmt.Scan(&prise)

	fmt.Print("Введите номинал первой монеты:", "\n")
	var firstCoin int
	fmt.Scan(&firstCoin)

	fmt.Print("Введите номинал второй монеты:", "\n")
	var secondCoin int
	fmt.Scan(&secondCoin)

	fmt.Print("Введите номинал второй монеты:", "\n")
	var therdCoin int
	fmt.Scan(&therdCoin)

	var evenAmount bool = false

	if prise == firstCoin || prise == firstCoin+secondCoin ||
		prise == firstCoin+secondCoin+therdCoin ||
		prise == secondCoin || prise == secondCoin+therdCoin ||
		prise == therdCoin {
		evenAmount = true
	}

	if evenAmount {
		fmt.Print("\n", "Этими монетами можно расплатиться без сдачи.")
	} else if prise < firstCoin+secondCoin+therdCoin {
		fmt.Print("\n", "Увы, без сдачи тут не получится.")
	} else {
		fmt.Print("\n", "Имеющихся монет недостаточно, чтобы расплатиться")
	}
}