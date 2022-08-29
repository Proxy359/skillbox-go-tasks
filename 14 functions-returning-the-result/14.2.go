package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNum(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func second() {
	x := randomNum(10)
	y := randomNum(10)
	x = 2*x + 10
	y = -3*y - 5
	fmt.Print("x: ", x, "y: ", y)
}

func main() {
	fmt.Println("Эта программа генерирует рандомные точки координат и перемножает их по формуле.")
	fmt.Print("Результат вычислений первой точки: ")
	second()
	fmt.Print("\n", "Результат вычислений второй точки: ")
	second()
	fmt.Print("\n", "Результат вычислений третьей точки: ")
	second()
}
