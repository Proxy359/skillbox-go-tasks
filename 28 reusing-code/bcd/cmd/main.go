package main

import (
	"bufio"
	"fmt"
	"os"
	"test/pkg/storage"
	"test/pkg/student"
)

func main() {
	var studentName string
	var age, grade int
	scanner := bufio.NewScanner(os.Stdin)
	m := &storage.StudentsStorage{}
	i := 0
	for {
		fmt.Print("Введите имя ", i+1, "-го студента:")
		if !scanner.Scan() {
			break
		}

		studentName = scanner.Text()
		fmt.Print("Введите возраст ", i+1, "-го студента:")
		fmt.Scanf("%d", &age)
		fmt.Print("Введите оценку ", i+1, "-го студента:")
		fmt.Scanf("%d", &grade)
		student := student.NewStudent(studentName, age, grade)
		i++
		m.Put(i, student)
	}

	for _, val := range m.Get() {
		fmt.Println(val.Name)
	}
}
