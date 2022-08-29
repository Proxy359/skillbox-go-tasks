package main

import "fmt"

func main() {
	fmt.Println("Включение устройства")
	fmt.Println("Loading...")
	fmt.Println("Loading...")
	fmt.Println("Loading...")
	fmt.Println("...")
	fmt.Println("Устройство по упрощению жизни маршруточного водилы готово к работе.")
	fmt.Println()
	fmt.Println()

	fmt.Println("Ваша первая остановка - улица Начальная.")
	var passengersVal float32 = 0

	var passengersCame float32
	fmt.Println("Вошло пассажиров:")
	fmt.Scan(&passengersCame)

	passengersVal = passengersCame + passengersVal

	fmt.Println("Отъезжаем, в салоне ", passengersVal, " пассажиров.")
	fmt.Println()
	fmt.Println()

	var passengersTotal float32 = 0
	passengersTotal += passengersCame

	fmt.Println("След. остановка - улица Печальная.")

	var passengersLeft float32
	fmt.Println("Вышло пассажиров:")
	fmt.Scan(&passengersLeft)

	fmt.Println("Вошло пассажиров:")
	fmt.Scan(&passengersCame)

	passengersVal = -passengersLeft + passengersCame + passengersVal

	fmt.Print("Отъезжаем, в салоне ", passengersVal,
		" пассажиров.", "\n", "\n", "\n")

	passengersTotal += passengersCame

	fmt.Println("След. остановка - улица Конечная.")
	fmt.Println()

	fmt.Print("Ну всё, время подсчитывать барыши.", "\n")
	var revenue float32 = passengersTotal * 20
	var salary float32 = revenue / 4
	fifthPart := salary / 5

	fmt.Print("Из полученных тобою ", revenue, " рублей:", "\n",
		salary, " можешь смело убирать себе в карман", "\n",
		fifthPart, " рублей отдай мужику на заправке, ремонтнику и налоговой.", "\n", "А теперь обнови приложение и езжай по новому кругу.")
}
