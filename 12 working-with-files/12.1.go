package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Привет, эта программа документирует твой ввод в консоль в новом файле")
	fmt.Println("Пиши что хочешь, а чтобы выйти, напиши 'exit'")

	documentation, err := os.Create("documentation.txt")
	writer := bufio.NewWriter(documentation)

	if err != nil {
		panic(err)
	}
	defer documentation.Close()

	x := 0

	for {

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}

		x = x + 1
		a := strconv.Itoa(x)

		time := time.Now()
		t := time.String()
		t = t[:19]

		writer.WriteString(a)
		writer.WriteString(" ")
		writer.WriteString(t)
		writer.WriteString(" ")
		writer.WriteString(input)
		writer.WriteString("\n")
		writer.Flush()
	}
	fmt.Println("Славно провели время, возвращайся, ещё попишем")
}
