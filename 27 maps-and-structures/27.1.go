package main

import (
	"bufio"
	"fmt"
	"os"
)

type Student struct {
	name  string
	age   int
	grade int
}

type Students struct {
	ns map[int]*Student
}

func NewStudent(name string, age, grade int) *Student {
	return &Student{name: name, age: age, grade: grade}
}

func (s *Students) Put(i int, student *Student) {
	if s.ns == nil {
		s.ns = make(map[int]*Student, 0)
	}
	s.ns[i] = student
}

func (s *Students) Get() map[int]*Student {
	return s.ns
}

func main() {
	var studentName string
	var age, grade int
	scanner := bufio.NewScanner(os.Stdin)
	m := &Students{}
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
		student := NewStudent(studentName, age, grade)
		i++
		m.Put(i, student)
	}

	for _, val := range m.Get() {
		fmt.Println(val.name)
	}
}
