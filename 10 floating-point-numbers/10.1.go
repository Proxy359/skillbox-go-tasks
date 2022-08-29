package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Давайте вычислим 'ex' посредством разложения в ряд Тейлора")
	fmt.Println("введите 'x' для которого необходимо рассчитать sin")
	var x float64
	fmt.Scan(&x)

	var epsilion float64
	fmt.Println("уточните, до какого знака после запятой необходимо вычислить значение функции?")
	fmt.Scan(&epsilion)
	epsilion = 1 / (math.Pow(10, epsilion))

	fact := 1.0
	result := 0.0
	prevResult := 0.0
	k := 0.0

	for {
		if k > 0 {
			fact = fact * k
		}
		result = result + math.Pow(x, k)/fact
		if math.Abs(result-prevResult) < epsilion {
			fmt.Println("/n", result)
			break
		}
		prevResult = result
		k = k + 1
	}
}
