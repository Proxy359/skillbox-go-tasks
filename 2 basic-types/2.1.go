package main

import "fmt"

func main() {
	lapNumber := 4
	engineBuff := 254
	wheelsBuff := 93
	rudderBuff := 49
	windDebuff := -21
	rainDebuff := -17
	driverSpeed := engineBuff + wheelsBuff + rudderBuff + windDebuff + rainDebuff

	fmt.Println("===================")
	fmt.Println("Супер гонки. Круг", lapNumber)
	fmt.Println("===================")
	fmt.Println()
	fmt.Print("Шумахер (", driverSpeed, ")")
	fmt.Println()
	fmt.Println("===================")
	fmt.Println("Водитель: Шумахер")
	fmt.Println("Скорость:", driverSpeed)
	fmt.Println("-------------------")
	fmt.Println("Оснащение")
	fmt.Println("Двигатель: +", engineBuff)
	fmt.Println("Колеса: +", wheelsBuff)
	fmt.Println("Руль: +", rudderBuff)
	fmt.Println("-------------------")
	fmt.Println("Действия плохой погоды")
	fmt.Println("Ветер:", windDebuff)
	fmt.Println("Дождь:", rainDebuff)
}
