package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func createTwoFiles() {
	firstString := "контент первого файла"
	secondString := "контент второго файла"

	firstFile, err := os.Create("first.txt")
	if err != nil {
		fmt.Println("Не смогли создать файл", err)
		return
	}
	defer firstFile.Close()
	firstFile.WriteString(firstString)

	secondFile, err := os.Create("second.txt")
	if err != nil {
		fmt.Println("Не смогли создать файл", err)
		return
	}
	defer secondFile.Close()
	secondFile.WriteString(secondString)
}

func createThirdFile() {
	firstFile, err := os.Open("first.txt")
	if err != nil {
		fmt.Println("Не смогли создать файл", err)
		return
	}
	defer firstFile.Close()
	firstStat, err := firstFile.Stat()
	firstBuf := make([]byte, firstStat.Size())
	if err != nil {
		panic(err)
	}
	if _, err := io.ReadFull(firstFile, firstBuf); err != nil {
		panic(err)
	}

	secondFile, err := os.Open("second.txt")
	if err != nil {
		fmt.Println("Не смогли создать файл", err)
		return
	}
	defer secondFile.Close()
	secondtStat, err := secondFile.Stat()
	secondBuf := make([]byte, secondtStat.Size())
	if err != nil {
		panic(err)
	}
	if _, err := io.ReadFull(secondFile, secondBuf); err != nil {
		panic(err)
	}

	var a []string
	a = append(a, string(firstBuf))
	b := append(a, string(secondBuf))
	firstString := strings.Join(b, " ")

	thirdFile, err := os.Create("result .txt")
	if err != nil {
		fmt.Println("Не смогли создать файл", err)
		return
	}
	defer thirdFile.Close()
	thirdFile.WriteString(firstString)
}

func main() {
	createTwoFiles()
	createThirdFile()
}
