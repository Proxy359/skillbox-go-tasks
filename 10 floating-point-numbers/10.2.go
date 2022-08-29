package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Эта программа позволит рассчитать сумму, которую вы хотите инвестировать")

	fmt.Println("Введите сумму, которую хотите инвестировать:")
	var investment float64
	fmt.Scan(&investment)

	fmt.Println("Введите срок, на который вы хотите инвестировать вышеуказанную сумму:")
	var timeline float64
	fmt.Scan(&timeline)
	timeline = timeline * 12

	fmt.Println("Введите проценткую ставку:")
	var interestRate float64
	fmt.Scan(&interestRate)

	userSum := investment
	bankSum := 0.0

	for x := 1.0; x <= timeline; x++ {
		userSum = userSum*interestRate/100 + userSum
		bankSum = bankSum + userSum
		userSum = math.Floor((userSum * 100)) / 100
		bankSum = bankSum - userSum
	}

	fmt.Println("Вы получите", userSum, "рублей")
	fmt.Println("Банк получит", bankSum, "рублей")
}
