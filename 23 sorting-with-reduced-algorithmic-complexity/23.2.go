package main

import (
	"fmt"
	"strings"
)

func parseTest(a [4]string, b [5]string) [4][5]int {
	var mass [4][5]int
	for i := 0; i < 4; i++ {
		fmt.Println("\n", "В фразе", a[i], "элементы находятся в след позиции:")
		for j := 0; j < 5; j++ {
			mass[i][j] = strings.LastIndex(a[i], b[j])
			fmt.Println("Буква", b[j], "имеет индекс", mass[i][j])
		}
	}
	return mass
}

func main() {
	var a [5]string
	sentences := [4]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := [5]rune{'H', 'e', 'l', 'П', 'М'}
	for i := 0; i < 5; i++ {
		a[i] = string(chars[i])
	}
	parseTest(sentences, a)
}
