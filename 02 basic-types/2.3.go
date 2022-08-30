package main

import "fmt"

func main() {
	shiftDuration := 480
	orderDuration := 2
	turnaroundTime := 4

	fmt.Println("Расчёт количества клиентов, обслуживаемых за смену.")
	fmt.Println()
	fmt.Println("Введите длительность смены в минутах:", shiftDuration)
	fmt.Println("Сколько минут клиент делает заказ:", orderDuration)
	fmt.Println("Сколько минут кассир собирает заказ:", turnaroundTime)
	fmt.Println()
	fmt.Println("-----Считаем-----")
	fmt.Println("За смену длиной", shiftDuration, "минут кассир успеет обслужить", shiftDuration/(orderDuration+turnaroundTime), "клиентов.")
}
