package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Определение количества слов, начинающихся с большой буквы в строке:")

	var a string = "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software"
	fmt.Println(a)

	result := 0

	for strings.Contains(a, " ") {
		blankIndex := strings.Index(a, " ")
		word := a[:blankIndex]
		if word == strings.Title(word) {
			result = result + 1
		}
		a = a[blankIndex+1:]
	}
	if a == strings.Title(a) {
		result = result + 1
	}
	fmt.Println("\n", "Строка содержит", result, "слов с большой буквы.")
}
