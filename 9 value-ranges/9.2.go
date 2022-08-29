package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Сейчас мы подберём наиболее ёмкий тип для резултата умножения ваших двух чисел")

	fmt.Println("Введите первое число")
	var firstNum int16
	fmt.Scan(&firstNum)

	fmt.Println("Введите второе число")
	var secondNum int16
	fmt.Scan(&secondNum)

	sum := int64(firstNum) * int64(secondNum)

	if sum > 0 {
		if sum <= math.MaxUint8 {
			fmt.Println("Uint8")
		} else if sum <= math.MaxUint16 {
			fmt.Println("Uint16")
		} else {
			fmt.Println("Uint32")
		}
	} else {
		if sum > math.MinInt8 {
			fmt.Println("int8")
		} else if sum > math.MinInt16 {
			fmt.Println("int16")
		} else {
			fmt.Println("int32")
		}
	}
}
