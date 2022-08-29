package main

import "fmt"

func main() {

	fmt.Println("Введите стоимость товара")
	var prise int
	fmt.Scan(&prise)

	fmt.Println("На какую скидку вы рассчитываете?")
	var discount int
	fmt.Scan(&discount)

	if discount <= prise/10*3 &&
		discount <= 2000 {
		fmt.Println("Вы правы, ваша скидка действительно", discount, "рублей")
		//мы заскамили покупателя, ведь он потенциально мог претендовать на большую скидку
	} else if prise/10*3 > 2000 {
		fmt.Println("Увы, но скидка будет лишь 2000 рублей")
	} else {
		fmt.Println("Увы, но скидка будет лишь", prise/10*3, "рублей")
	}
}
