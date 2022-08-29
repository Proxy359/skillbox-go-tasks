package main

import (
	"fmt"
	"strconv"
)

func first(ans int) chan int {
	firstChan := make(chan int)
	go func() {
		a := ans * ans
		firstChan <- a
	}()
	return firstChan
}

func second(a chan int) chan int {
	secondChan := make(chan int)
	aa := <-a
	fmt.Println(aa)
	go func() {
		secondChan <- aa * 2
	}()
	return secondChan
}

func main() {
	ans := ""

	for {
		fmt.Println("Введите число")
		fmt.Scan(&ans)

		if ans == "стоп" {
			break
		}

		intAns, _ := strconv.Atoi(ans)

		fc := first(intAns)
		sc := second(fc)
		fmt.Println(<-sc)

	}
}
