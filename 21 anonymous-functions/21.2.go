package main

import "fmt"

func first(F func(int, int) int) int {
	num1 := 3
	num2 := 5
	summ := num1 + num2
	ans := F(summ, num2)
	defer fmt.Println(ans)
	fmt.Println("Ответ будет:")
	return ans
}

func main() {
	first(func(x, y int) int { return (x * y) << 2 })
	first(func(x, y int) int { return x + y })
	first(func(x, y int) int { return x - y })
}
