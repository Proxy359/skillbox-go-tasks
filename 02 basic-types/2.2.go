package main

import "fmt"

func main() {
	initialValue := 6400
	deliveryFee := 350
	discount := 700
	finalCost := initialValue + deliveryFee - discount

	fmt.Println("Калькулятор стоимости товара со скидкой.")
	fmt.Println()
	fmt.Println("Стоимость товара:", initialValue)
	fmt.Println("Стоимость доставки:", deliveryFee)
	fmt.Println("Размер скидки:", discount)
	fmt.Println("---------")
	fmt.Println("Итого:", finalCost)
}
