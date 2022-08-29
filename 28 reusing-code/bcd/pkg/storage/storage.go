package storage

import "test/pkg/student"

type StudentsStorage struct {
	ns map[int]*student.Student
}

func (ss *StudentsStorage) Get() map[int]*student.Student {
	return ss.ns
}

func (ss *StudentsStorage) Put(i int, s *student.Student) {
	if ss.ns == nil {
		ss.ns = make(map[int]*student.Student, 0)
	}
	ss.ns[i] = s
}
