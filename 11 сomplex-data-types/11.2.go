package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Исходная строка:")

	var a string = "a10 10 20b 20 30c30 30 dd"
	fmt.Println(a)

	ans := ""

	for strings.Contains(a, " ") {
		blankIndex := strings.Index(a, " ")
		word := a[:blankIndex]

		i, err := strconv.Atoi(word)
		if err != nil {
			i = 1
		}

		if i%10 == 0 {
			ans = ans + word + " "
		}
		a = a[blankIndex+1:]
	}

	i, err := strconv.Atoi(a)
	if err != nil {
		i = 1
	}

	if i%10 == 0 {
		ans = ans + a + " "
	}
	ans = strings.Trim(ans, " ")
	fmt.Println(ans)
}
