package main

import "fmt"

func main() {
	fmt.Println("Введите будний день недели: пн, вт, ср, чт, пт:")
	var answer string
	fmt.Scan(&answer)

	switch answer {
	case "пн":
		fmt.Println("понедельник")
		fallthrough
	case "вт":
		fmt.Println("вторник")
		fallthrough
	case "ср":
		fmt.Println("среда")
		fallthrough
	case "чт":
		fmt.Println("четверг")
		fallthrough
	case "пт":
		fmt.Println("пятница")
	}
}
