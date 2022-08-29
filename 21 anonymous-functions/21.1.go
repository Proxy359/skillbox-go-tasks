package main

import (
	"fmt"
)

func first(x int16, y uint8, z float32) float32 {
	var S float32
	S = float32(x) - 3/z + float32((y-1)<<2)
	return S
}

func main() {
	fmt.Println("Эта прога вычисляет уравнение")
	var x int16 = 4
	var y uint8 = 5
	var z float32 = 3.222
	fmt.Println(first(x, y, z))
}
