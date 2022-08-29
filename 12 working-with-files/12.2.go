package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Привет, эта программа выводит, что на писали в прошлой проге")

	documentation, err := os.Open("documentation.txt")
	if err != nil {
		fmt.Println("Файл отсутсвует или пуст")
		panic(err)
	}
	defer documentation.Close()

	fmt.Println("Вот, что мы записали")
	stat, err := documentation.Stat()
	buf := make([]byte, stat.Size())
	if err != nil {
		panic(err)
	}
	if _, err := io.ReadFull(documentation, buf); err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}
