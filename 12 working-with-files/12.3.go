package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Привет, эта программа создаёт файл только для чтения")

	lock, err := os.Create("lock.txt")
	if err := os.Chmod("lock.txt", 0444); err != nil {
		fmt.Println(err)
	}
	if err != nil {
		panic(err)
	}
	defer lock.Close()

	fmt.Println("Как и говорил - только чтение.")
}
