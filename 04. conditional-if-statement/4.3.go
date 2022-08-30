package main

import "fmt"

func main() {

	fmt.Print("Банкомат.", "\n",
		"Введите сумму снятия со счёта:", "\n")

	var amountReceived int
	fmt.Scan(&amountReceived)

	limit := 100000
	minSum := 100

	if amountReceived%minSum == 0 &&
		amountReceived <= limit {
		fmt.Print("\n", "Операция успешно выполнена.", "\n", "Вы сняли ", amountReceived, " рублей.")
	}

	if amountReceived%minSum != 0 {
		fmt.Print("\n", "Ошибка операции.", "\n", "Введите сумму, кратную ", minSum, " рублям.", "\n")
	}

	if amountReceived > limit {
		fmt.Print("\n", "Ошибка операции.", "\n", "У нас столько нету =(", "\n", "Введите сумму до ", limit, " рублей.")
	}
}
