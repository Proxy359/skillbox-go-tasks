package main

import "fmt"

func main() {
	fmt.Println("Переполнение числа типа uint8 и uint16 в диапазоне от 0 до uint32 будет:")
	const uint8 = 256
	const uint16 = 65536
	const uint32 = 4294967296
	x8 := 0
	x16 := 0

	for x := 1; x <= uint32; x++ {
		if x%uint8 == 0 {
			x8 = x8 + 1
		}
		if x%uint16 == 0 {
			x16 = x16 + 1
		}
	}
	fmt.Println("uint8 переполнился", x8-1, "раз")
	fmt.Println("uint16 переполнился", x16-1, "раз")
}
